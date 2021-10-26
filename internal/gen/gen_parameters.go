package gen

import (
	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen/internal/ir"
	"github.com/ogen-go/ogen/internal/oas"
)

func (g *Generator) generateParameters(params []*oas.Parameter) ([]*ir.Parameter, error) {
	var result []*ir.Parameter
	for _, p := range params {
		param, err := g.generateParameter(p)
		if err != nil {
			return nil, xerrors.Errorf("'%s': %w", p.Name, err)
		}

		result = append(result, param)
	}

	return result, nil
}

func (g *Generator) generateParameter(p *oas.Parameter) (*ir.Parameter, error) {
	typ, err := g.generateSchema(pascal(p.Name), p.Schema)
	if err != nil {
		return nil, xerrors.Errorf("'%s': %w", p.Name, err)
	}

	if !isUnderlyingPrimitive(typ) {
		return nil, &ErrNotImplemented{"complex parameter types"}
	}

	return &ir.Parameter{
		Name: pascal(p.Name),
		Type: typ,
		Spec: p,
	}, nil
}

func isUnderlyingPrimitive(typ *ir.Type) bool {
	switch typ.Kind {
	case ir.KindPrimitive, ir.KindEnum:
		return true
	case ir.KindArray:
		return isUnderlyingPrimitive(typ.Item)
	case ir.KindAlias:
		return isUnderlyingPrimitive(typ.AliasTo)
	case ir.KindPointer:
		return isUnderlyingPrimitive(typ.PointerTo)
	default:
		return false
	}
}
