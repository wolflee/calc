// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package ast

type Visitor interface {
	Visit(node Node) bool
}

func Walk(node Node, v Visitor) {
	if !v.Visit(node) {
		return
	}

	switch n := node.(type) {
	case *AssignExpr:
		Walk(n.Value, v)
	case *BasicLit: /* do nothing */
	case *BinaryExpr:
		for _, x := range n.List {
			Walk(x, v)
		}
	case *CallExpr:
		for _, arg := range n.Args {
			Walk(arg, v)
		}
	case *DeclExpr:
		Walk(n.Body, v)
	case *ExprList:
		for _, expr := range n.List {
			Walk(expr, v)
		}
	case *File:
		for _, decl := range n.Decls {
			Walk(decl, v)
		}
	case *Ident: /* do nothing */
	case *IfExpr:
		Walk(n.Cond, v)
		Walk(n.Then, v)
		if n.Else != nil {
			Walk(n.Else, v)
		}
	case *Package:
		for _, file := range n.Files {
			Walk(file, v)
		}
	case *UnaryExpr:
		Walk(n.Value, v)
	case *VarExpr:
		if n.Value != nil {
			Walk(n.Value, v)
		}
	}
}
