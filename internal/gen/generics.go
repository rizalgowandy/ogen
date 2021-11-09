package gen

import (
	"fmt"
	"strings"

	"github.com/ogen-go/ogen/internal/ir"
)

func (g *Generator) wrapGenerics() {
	for _, typ := range g.types {
		if typ.Is(ir.KindStruct) {
			g.boxStructFields(typ)
		}
	}

	for _, op := range g.operations {
		for _, param := range op.Params {
			v := ir.GenericVariant{
				Nullable: param.Spec.Schema.Nullable,
				Optional: !param.Spec.Required,
			}

			if v.Any() {
				param.Type = g.boxType(v, param.Type)
			}
		}
	}
}

func (g *Generator) boxStructFields(s *ir.Type) {
	for _, field := range s.Fields {
		if field.Spec == nil {
			continue
		}

		v := ir.GenericVariant{
			Nullable: field.Spec.Schema.Nullable,
			Optional: !field.Spec.Required,
		}

		field.Type = func(typ *ir.Type) *ir.Type {
			if s.RecursiveTo(typ) {
				switch {
				case v.OnlyOptional():
					return ir.Pointer(typ, ir.NilOptional)
				case v.OnlyNullable():
					return ir.Pointer(typ, ir.NilNull)
				case v.NullableOptional():
					return ir.Pointer(g.boxType(ir.GenericVariant{
						Optional: true,
					}, typ), ir.NilNull)
				default:
					// Required.
					panic(fmt.Sprintf("recursion: %s.%s", s.Name, field.Name))
				}
			}
			if v.Any() {
				return g.boxType(v, typ)
			}
			return typ
		}(field.Type)
	}
}

func (g *Generator) boxType(v ir.GenericVariant, t *ir.Type) *ir.Type {
	if t.IsArray() || t.Primitive == ir.ByteSlice {
		// Using special case for array nil value if possible.
		switch {
		case v.OnlyOptional():
			t.NilSemantic = ir.NilOptional
		case v.OnlyNullable():
			t.NilSemantic = ir.NilNull
		default:
			t = ir.Generic(genericPostfix(t.Go()),
				t, v,
			)
			g.saveType(t)
		}

		return t
	}

	if t.CanGeneric() {
		t = ir.Generic(genericPostfix(t.Go()), t, v)
		g.saveType(t)
		return t
	}

	switch {
	case v.OnlyOptional():
		return t.Pointer(ir.NilOptional)
	case v.OnlyNullable():
		return t.Pointer(ir.NilNull)
	default:
		t = ir.Generic(genericPostfix(t.Go()),
			t.Pointer(ir.NilNull), ir.GenericVariant{Optional: true},
		)
		g.saveType(t)
		return t
	}
}

func genericPostfix(name string) string {
	if idx := strings.Index(name, "."); idx > 0 {
		name = name[idx+1:]
	}
	return pascal(name)
}
