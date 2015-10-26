package ir

import (
	"github.com/rthornton128/calc/ast"
	"github.com/rthornton128/calc/token"
)

type Object interface {
	Kind() ast.Kind
	Name() string
	Pos() token.Pos
	Scope() *Scope
	String() string
	Type() Type
}

type IDer interface {
	ID() int
	SetID(int)
}

type object struct {
	kind  ast.Kind
	name  string
	pos   token.Pos
	scope *Scope
	typ   Type
}

func newObject(name, t string, p token.Pos, k ast.Kind, s *Scope) object {
	return object{
		kind:  k,
		name:  name,
		pos:   p,
		scope: s,
		typ:   typeFromString(t),
	}
}

func (o object) Kind() ast.Kind { return o.kind }
func (o object) Name() string   { return o.name }
func (o object) Pos() token.Pos { return o.pos }
func (o object) Scope() *Scope  { return o.scope }
func (o object) Type() Type     { return o.typ }