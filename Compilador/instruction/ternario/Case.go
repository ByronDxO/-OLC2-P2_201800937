package ternario

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Case struct {
	Condition      interfaces.Expression
	ListaExpresion *arrayList.List
	Instrucciones  interfaces.Expression
}

func NewCase(cond interfaces.Expression, listaExpresion *arrayList.List, instrucciones interfaces.Expression) Case {
	instr := Case{cond, listaExpresion, instrucciones}
	return instr
}

func (p Case) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {

	var result interfaces.Value
	if p.Condition != nil {
		result = p.Condition.Compilar(env, tree, gen)
		return result

	}
	return result
}
