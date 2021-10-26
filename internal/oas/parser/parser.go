package parser

import (
	"strings"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen"
	ast "github.com/ogen-go/ogen/internal/oas"
)

type parser struct {
	ops Options
	// schema specification, immutable.
	spec *ogen.Spec
	// parsed operations.
	operations []*ast.Operation
	// refs contains lazy-initialized referenced components.
	refs struct {
		schemas       map[string]*ast.Schema
		requestBodies map[string]*ast.RequestBody
		responses     map[string]*ast.Response
		parameters    map[string]*ast.Parameter
	}
}

type Options struct {
	IgnoreUnspecifiedParams bool
}

func Parse(spec *ogen.Spec, opts Options) ([]*ast.Operation, error) {
	spec.Init()
	p := &parser{
		ops:  opts,
		spec: spec,
		refs: struct {
			schemas       map[string]*ast.Schema
			requestBodies map[string]*ast.RequestBody
			responses     map[string]*ast.Response
			parameters    map[string]*ast.Parameter
		}{
			schemas:       map[string]*ast.Schema{},
			requestBodies: map[string]*ast.RequestBody{},
			responses:     map[string]*ast.Response{},
			parameters:    map[string]*ast.Parameter{},
		},
	}

	err := p.parse()
	return p.operations, err
}

func (p *parser) parse() error {
	for path, item := range p.spec.Paths {
		if item.Ref != "" {
			return xerrors.New("referenced paths are not supported")
		}

		if err := forEachOps(item, func(method string, op ogen.Operation) error {
			parsedOp, err := p.parseOp(path, strings.ToUpper(method), op)
			if err != nil {
				var paramNotSpecified *ErrPathParameterNotSpecified
				if xerrors.As(err, &paramNotSpecified) {
					if p.ops.IgnoreUnspecifiedParams {
						return nil
					}
				}
				return err
			}

			p.operations = append(p.operations, parsedOp)
			return nil
		}); err != nil {
			return xerrors.Errorf("paths: %s: %w", path, err)
		}
	}

	return nil
}

func (p *parser) parseOp(path, httpMethod string, spec ogen.Operation) (_ *ast.Operation, err error) {
	op := &ast.Operation{
		OperationID: spec.OperationID,
		HTTPMethod:  httpMethod,
	}

	op.Parameters, err = p.parseParams(spec.Parameters)
	if err != nil {
		return nil, xerrors.Errorf("parameters: %w", err)
	}

	op.PathParts, err = parsePath(path, op.Parameters)
	if err != nil {
		return nil, xerrors.Errorf("parse path: %w", err)
	}

	if spec.RequestBody != nil {
		op.RequestBody, err = p.parseRequestBody(spec.RequestBody)
		if err != nil {
			return nil, xerrors.Errorf("requestBody: %w", err)
		}
	}

	if len(spec.Responses) > 0 {
		op.Responses, err = p.parseResponses(spec.Responses)
		if err != nil {
			return nil, xerrors.Errorf("responses: %w", err)
		}
	}

	return op, nil
}
