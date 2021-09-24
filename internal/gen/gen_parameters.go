package gen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ogen-go/ogen"
)

func (g *Generator) parseParameter(param ogen.Parameter, path string) (Parameter, error) {
	types := map[string]ParameterLocation{
		"query":  LocationQuery,
		"header": LocationHeader,
		"path":   LocationPath,
		"cookie": LocationCookie,
	}

	locatedIn, exists := types[strings.ToLower(param.In)]
	if !exists {
		return Parameter{}, fmt.Errorf("unsupported parameter type %s", param.In)
	}

	if locatedIn == LocationPath {
		if !param.Required {
			return Parameter{}, fmt.Errorf("parameters located in 'path' must be required")
		}

		exists, err := regexp.MatchString(fmt.Sprintf("{%s}", param.Name), path)
		if err != nil {
			return Parameter{}, fmt.Errorf("match path param '%s': %w", param.Name, err)
		}

		if !exists {
			return Parameter{}, fmt.Errorf("param '%s' not found in path '%s'", param.Name, path)
		}
	}

	name := pascal(param.Name)
	schema, err := g.generateSchema(name, param.Schema)
	if err != nil {
		return Parameter{}, fmt.Errorf("schema: %w", err)
	}

	switch schema.Kind {
	case KindStruct:
		return Parameter{}, fmt.Errorf("object type not supported")
	case KindArray:
		if schema.Item.Kind != KindPrimitive {
			return Parameter{}, fmt.Errorf("only arrays with primitive types supported")
		}
	}

	style, err := paramStyle(locatedIn, param.Style)
	if err != nil {
		return Parameter{}, fmt.Errorf("style: %w", err)
	}

	return Parameter{
		Name:       name,
		In:         locatedIn,
		SourceName: param.Name,
		Schema:     schema,
		Style:      style,
		Explode:    paramExplode(locatedIn, param.Explode),
		Required:   param.Required,
	}, nil
}

// paramStyle checks parameter style field.
// https://swagger.io/docs/specification/serialization/
func paramStyle(locatedIn ParameterLocation, style string) (string, error) {
	if style == "" {
		defaultStyles := map[ParameterLocation]string{
			LocationPath:   "simple",
			LocationQuery:  "form",
			LocationHeader: "simple",
			LocationCookie: "form",
		}

		return defaultStyles[locatedIn], nil
	}

	allowedStyles := map[ParameterLocation]map[string]struct{}{
		LocationPath: {
			"simple": struct{}{},
			"label":  struct{}{},
			"matrix": struct{}{},
		},
		LocationQuery: {
			"form":           struct{}{},
			"spaceDelimited": struct{}{},
			"pipeDelimited":  struct{}{},
			"deepObject":     struct{}{},
		},
		LocationHeader: {
			"simple": struct{}{},
		},
		LocationCookie: {
			"form": struct{}{},
		},
	}

	if _, found := allowedStyles[locatedIn][style]; !found {
		return "", fmt.Errorf("unexpected style: %s", style)
	}

	return style, nil
}

// paramExplode checks parameter explode field.
// https://swagger.io/docs/specification/serialization/
func paramExplode(locatedIn ParameterLocation, explode *bool) bool {
	if explode != nil {
		return *explode
	}

	// When style is form, the default value is true.
	// For all other styles, the default value is false.
	if locatedIn == LocationQuery || locatedIn == LocationCookie {
		return true
	}

	return false
}
