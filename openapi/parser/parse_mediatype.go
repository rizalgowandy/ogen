package parser

import (
	"github.com/go-faster/errors"
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/openapi"
)

func (p *parser) parseContent(content map[string]ogen.Media) (map[string]*openapi.MediaType, error) {
	if content == nil {
		return nil, nil
	}

	result := make(map[string]*openapi.MediaType, len(content))
	for name, m := range content {
		mm, err := p.parseMediaType(m)
		if err != nil {
			return nil, errors.Wrap(err, name)
		}

		result[name] = mm
	}

	return result, nil
}

func (p *parser) parseMediaType(m ogen.Media) (*openapi.MediaType, error) {
	s, err := p.schemaParser.Parse(m.Schema.ToJSONSchema())
	if err != nil {
		return nil, errors.Wrap(err, "schema")
	}

	examples := make(map[string]*openapi.Example, len(m.Examples))
	for name, ex := range m.Examples {
		e, err := p.parseExample(ex, resolveCtx{})
		if err != nil {
			return nil, errors.Wrapf(err, "examples: %q", name)
		}

		examples[name] = e
	}

	// OpenAPI 3.0.3 doc says:
	//
	//   Furthermore, referencing a schema which contains an example,
	//   the example value SHALL override the example provided by the schema.
	//
	// Probably this will be rewritten later.
	// Kept for backward compatibility.
	s.AddExample(m.Example)
	for _, ex := range examples {
		s.AddExample(ex.Value)

	}

	return &openapi.MediaType{
		Schema:   s,
		Example:  m.Example,
		Examples: examples,
	}, nil
}