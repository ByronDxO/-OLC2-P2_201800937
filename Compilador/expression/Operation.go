package expression

import (
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	"fmt"	
	// "math"
	// "strconv"
	// "reflect"
)

type Aritmetica struct {
	left      	interfaces.Expression
	Operator 	string
	right      	interfaces.Expression
	Unario   	bool
	Row 		int
	Column 		int
}

func NewOperacion(left interfaces.Expression, Operator string, right interfaces.Expression, unario bool, row int, column int) Aritmetica {
	exp := Aritmetica{left, Operator, right, unario, row, column}
	return exp
}

func (p Aritmetica) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {
	

	var exp_left interfaces.Value
	var exp_right interfaces.Value


	switch p.Operator {
	case "+":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Aritmetica +")
			temp 	:= gen.NewTemp()
			isType  := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "+")
				isType = interfaces.INTEGER
				
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "+")
				isType = interfaces.FLOAT
				
			} else if exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING {
				if !tree.IsCocant {
					gen.AddConcatString()
					tree.IsCocant = true
				}
				gen.AddExpression(temp,"P",fmt.Sprintf("%v", tree.GetPos()),"+")
				gen.AddExpression(temp,temp,"1","+")
				gen.AddStack(temp,exp_left.Value)
				gen.AddExpression(temp,temp,"1","+")
				gen.AddStack(temp,exp_right.Value)
				gen.AddExpression("P","P",fmt.Sprintf("%v", tree.GetPos()),"+")
				gen.ConcatString()
				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				gen.AddExpression("P","P",fmt.Sprintf("%v", tree.GetPos()),"-")


				isType = interfaces.STRING



			}else {
				excep := ast.NewException("Semantico","No es posible Sumar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "-":
		{
			temp 	:= gen.NewTemp()
			isType  := interfaces.NULL

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
				gen.AddComment("Aritmetica -")
				/* ************************************************************** INTEGER ************************************************************** */
				if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
					gen.AddExpression(temp, "0", exp_left.Value, "-")
					isType = interfaces.INTEGER
				/* ************************************************************** FLOAT ************************************************************** */
				}  else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
					gen.AddExpression(temp, "0", exp_left.Value, "-")
					isType = interfaces.FLOAT
					
				} else {
					excep := ast.NewException("Semantico","No es posible Restar, Tipos de datos Incorrectos.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}


			} else {

				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
				gen.AddComment("Aritmetica -")

				/* ************************************************************** INTEGER ************************************************************** */
				if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
					gen.AddExpression(temp, exp_left.Value, exp_right.Value, "-")
					isType = interfaces.INTEGER
				/* ************************************************************** FLOAT ************************************************************** */
				}  else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
					gen.AddExpression(temp, exp_left.Value, exp_right.Value, "-")
					isType = interfaces.FLOAT
					
				} else {
					excep := ast.NewException("Semantico","No es posible Restar, Tipos de datos Incorrectos.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}

			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "*":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			
			gen.AddComment("Aritmetica *")
			
			temp 	:= gen.NewTemp()
			isType  := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "*")
				isType = interfaces.INTEGER
			/* ************************************************************** FLOAT ************************************************************** */	
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "*")
				isType = interfaces.FLOAT
				
			} else {
				excep := ast.NewException("Semantico","No es posible Multiplicar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "%":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			
			gen.AddComment("Aritmetica %")
			
			temp 	:= gen.NewTemp()
			isType  := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				val := "fmod(" + fmt.Sprintf("%v", exp_left.Value) + "," + fmt.Sprintf("%v", exp_right.Value) + ")"
				gen.AddExpression(temp, val, "0", "+")
				isType = interfaces.INTEGER
			/* ************************************************************** FLOAT ************************************************************** */	
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				val := "fmod(" + fmt.Sprintf("%v", exp_left.Value) + "," + fmt.Sprintf("%v", exp_right.Value) + ")"
				gen.AddExpression(temp, val, "0", "+")
				isType = interfaces.FLOAT
				
			} else {
				excep := ast.NewException("Semantico","No es posible Multiplicar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "/":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Aritmetica /")

			if !exp_right.IsTemp {
				if exp_right.Value == "0" {
					excep := ast.NewException("Semantico","No es posible Dividir entre 0.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}
			}
			temp 	:= gen.NewTemp()
			isType  := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "/")
				isType = interfaces.INTEGER
			/* ************************************************************** FLOAT ************************************************************** */	
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "/")
				isType = interfaces.FLOAT
				
			} else {
				excep := ast.NewException("Semantico","No es posible Dividir.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "<":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional <")
			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "<", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "<", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar <.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}

		}
	case ">":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional >")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, ">", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, ">", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar >.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}

		}
	case "<=":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional <=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "<=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "<=", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar <=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel:  EV, FalseLabel: EF}

		}
	case ">=":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional >=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, ">=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, ">=", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar >=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel:  EV, FalseLabel: EF}

		}
	case "==":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional ==")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "==", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar ==.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}

		}
	case "!=":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional !=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "!=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "!=", EV)
				gen.AddGoto(EF)

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar !=.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}

		}
	case "&&":
		{

			exp_left = p.left.Compilar(env, tree, gen)

			if exp_left.Type == interfaces.BOOLEAN {
				

				newLabel := gen.NewLabel()
				leftTemp := gen.NewTemp()	
				gen.Boolean(exp_left.TrueLabel, exp_left.FalseLabel, newLabel, leftTemp)

				exp_right 	:= p.right.Compilar(env, tree, gen)

				if exp_right.Type == interfaces.BOOLEAN {

					newLabel   = gen.NewLabel()
					rightTemp := gen.NewTemp()	
					gen.Boolean(exp_right.TrueLabel, exp_right.FalseLabel, newLabel, rightTemp)


					EV   := gen.NewLabel()
					EF 	 := gen.NewLabel()
					gen.AddIf(leftTemp, rightTemp, "&&", EV)
					gen.AddGoto(EF)

					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}
				} else {
					excep := ast.NewException("Semantico","No es posible Comparar &&; Tipo de Dato no Booleano en expresion Der.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

				}
				

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar &&; Tipo de Dato no Booleano en expresion Izq.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

		}
	case "||":
		{

			exp_left = p.left.Compilar(env, tree, gen)

			if exp_left.Type == interfaces.BOOLEAN {
				

				newLabel := gen.NewLabel()
				leftTemp := gen.NewTemp()	
				gen.Boolean(exp_left.TrueLabel, exp_left.FalseLabel, newLabel, leftTemp)

				exp_right 	:= p.right.Compilar(env, tree, gen)

				if exp_right.Type == interfaces.BOOLEAN {

					newLabel   = gen.NewLabel()
					rightTemp := gen.NewTemp()	
					gen.Boolean(exp_right.TrueLabel, exp_right.FalseLabel, newLabel, rightTemp)


					EV   := gen.NewLabel()
					EF 	 := gen.NewLabel()
					gen.AddIf(leftTemp, rightTemp, "||", EV)
					gen.AddGoto(EF)

					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: EV, FalseLabel: EF}
				} else {
					excep := ast.NewException("Semantico","No es posible Comparar ||; Tipo de Dato no Booleano en expresion Der.", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

				}
				

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar ||; Tipo de Dato no Booleano en expresion Izq.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

		}
	case "!":
		{
			if p.Unario {

				exp_left = p.left.Compilar(env, tree, gen)
			}

			if exp_left.Type == interfaces.BOOLEAN {

				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: exp_left.FalseLabel, FalseLabel: exp_left.TrueLabel}
			
				

			} else {
				excep := ast.NewException("Semantico","No es posible Comparar ||; Tipo de Dato no Booleano.", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo:excep.Tipo, Descripcion:excep.Descripcion, Row:excep.Row, Column:excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

		}

	}

	

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

}

