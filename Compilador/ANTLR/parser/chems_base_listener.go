// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseChemsListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseChemsListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterEnd_instr is called when production end_instr is entered.
func (s *BaseChemsListener) EnterEnd_instr(ctx *End_instrContext) {}

// ExitEnd_instr is called when production end_instr is exited.
func (s *BaseChemsListener) ExitEnd_instr(ctx *End_instrContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseChemsListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseChemsListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInstr_println is called when production instr_println is entered.
func (s *BaseChemsListener) EnterInstr_println(ctx *Instr_printlnContext) {}

// ExitInstr_println is called when production instr_println is exited.
func (s *BaseChemsListener) ExitInstr_println(ctx *Instr_printlnContext) {}

// EnterInstr_main is called when production instr_main is entered.
func (s *BaseChemsListener) EnterInstr_main(ctx *Instr_mainContext) {}

// ExitInstr_main is called when production instr_main is exited.
func (s *BaseChemsListener) ExitInstr_main(ctx *Instr_mainContext) {}

// EnterInstr_declaracion is called when production instr_declaracion is entered.
func (s *BaseChemsListener) EnterInstr_declaracion(ctx *Instr_declaracionContext) {}

// ExitInstr_declaracion is called when production instr_declaracion is exited.
func (s *BaseChemsListener) ExitInstr_declaracion(ctx *Instr_declaracionContext) {}

// EnterInstr_asignacion is called when production instr_asignacion is entered.
func (s *BaseChemsListener) EnterInstr_asignacion(ctx *Instr_asignacionContext) {}

// ExitInstr_asignacion is called when production instr_asignacion is exited.
func (s *BaseChemsListener) ExitInstr_asignacion(ctx *Instr_asignacionContext) {}

// EnterInstr_if is called when production instr_if is entered.
func (s *BaseChemsListener) EnterInstr_if(ctx *Instr_ifContext) {}

// ExitInstr_if is called when production instr_if is exited.
func (s *BaseChemsListener) ExitInstr_if(ctx *Instr_ifContext) {}

// EnterInstr_else_if is called when production instr_else_if is entered.
func (s *BaseChemsListener) EnterInstr_else_if(ctx *Instr_else_ifContext) {}

// ExitInstr_else_if is called when production instr_else_if is exited.
func (s *BaseChemsListener) ExitInstr_else_if(ctx *Instr_else_ifContext) {}

// EnterInstr_ternario is called when production instr_ternario is entered.
func (s *BaseChemsListener) EnterInstr_ternario(ctx *Instr_ternarioContext) {}

// ExitInstr_ternario is called when production instr_ternario is exited.
func (s *BaseChemsListener) ExitInstr_ternario(ctx *Instr_ternarioContext) {}

// EnterInstr_else_if_ternario is called when production instr_else_if_ternario is entered.
func (s *BaseChemsListener) EnterInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// ExitInstr_else_if_ternario is called when production instr_else_if_ternario is exited.
func (s *BaseChemsListener) ExitInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// EnterInstr_match is called when production instr_match is entered.
func (s *BaseChemsListener) EnterInstr_match(ctx *Instr_matchContext) {}

// ExitInstr_match is called when production instr_match is exited.
func (s *BaseChemsListener) ExitInstr_match(ctx *Instr_matchContext) {}

// EnterList_case is called when production list_case is entered.
func (s *BaseChemsListener) EnterList_case(ctx *List_caseContext) {}

// ExitList_case is called when production list_case is exited.
func (s *BaseChemsListener) ExitList_case(ctx *List_caseContext) {}

// EnterInstr_case is called when production instr_case is entered.
func (s *BaseChemsListener) EnterInstr_case(ctx *Instr_caseContext) {}

// ExitInstr_case is called when production instr_case is exited.
func (s *BaseChemsListener) ExitInstr_case(ctx *Instr_caseContext) {}

// EnterList_expre_case is called when production list_expre_case is entered.
func (s *BaseChemsListener) EnterList_expre_case(ctx *List_expre_caseContext) {}

// ExitList_expre_case is called when production list_expre_case is exited.
func (s *BaseChemsListener) ExitList_expre_case(ctx *List_expre_caseContext) {}

// EnterBlock_case is called when production block_case is entered.
func (s *BaseChemsListener) EnterBlock_case(ctx *Block_caseContext) {}

// ExitBlock_case is called when production block_case is exited.
func (s *BaseChemsListener) ExitBlock_case(ctx *Block_caseContext) {}

// EnterBlock_instr_match is called when production block_instr_match is entered.
func (s *BaseChemsListener) EnterBlock_instr_match(ctx *Block_instr_matchContext) {}

// ExitBlock_instr_match is called when production block_instr_match is exited.
func (s *BaseChemsListener) ExitBlock_instr_match(ctx *Block_instr_matchContext) {}

// EnterInstr_default is called when production instr_default is entered.
func (s *BaseChemsListener) EnterInstr_default(ctx *Instr_defaultContext) {}

// ExitInstr_default is called when production instr_default is exited.
func (s *BaseChemsListener) ExitInstr_default(ctx *Instr_defaultContext) {}

// EnterBlock_default is called when production block_default is entered.
func (s *BaseChemsListener) EnterBlock_default(ctx *Block_defaultContext) {}

// ExitBlock_default is called when production block_default is exited.
func (s *BaseChemsListener) ExitBlock_default(ctx *Block_defaultContext) {}

// EnterInstr_match_ter is called when production instr_match_ter is entered.
func (s *BaseChemsListener) EnterInstr_match_ter(ctx *Instr_match_terContext) {}

// ExitInstr_match_ter is called when production instr_match_ter is exited.
func (s *BaseChemsListener) ExitInstr_match_ter(ctx *Instr_match_terContext) {}

// EnterList_case_ternario is called when production list_case_ternario is entered.
func (s *BaseChemsListener) EnterList_case_ternario(ctx *List_case_ternarioContext) {}

// ExitList_case_ternario is called when production list_case_ternario is exited.
func (s *BaseChemsListener) ExitList_case_ternario(ctx *List_case_ternarioContext) {}

// EnterInstr_case_ter is called when production instr_case_ter is entered.
func (s *BaseChemsListener) EnterInstr_case_ter(ctx *Instr_case_terContext) {}

// ExitInstr_case_ter is called when production instr_case_ter is exited.
func (s *BaseChemsListener) ExitInstr_case_ter(ctx *Instr_case_terContext) {}

// EnterList_expre_case_ter is called when production list_expre_case_ter is entered.
func (s *BaseChemsListener) EnterList_expre_case_ter(ctx *List_expre_case_terContext) {}

// ExitList_expre_case_ter is called when production list_expre_case_ter is exited.
func (s *BaseChemsListener) ExitList_expre_case_ter(ctx *List_expre_case_terContext) {}

// EnterBlock_case_ter is called when production block_case_ter is entered.
func (s *BaseChemsListener) EnterBlock_case_ter(ctx *Block_case_terContext) {}

// ExitBlock_case_ter is called when production block_case_ter is exited.
func (s *BaseChemsListener) ExitBlock_case_ter(ctx *Block_case_terContext) {}

// EnterInstr_default_ter is called when production instr_default_ter is entered.
func (s *BaseChemsListener) EnterInstr_default_ter(ctx *Instr_default_terContext) {}

// ExitInstr_default_ter is called when production instr_default_ter is exited.
func (s *BaseChemsListener) ExitInstr_default_ter(ctx *Instr_default_terContext) {}

// EnterInstr_while is called when production instr_while is entered.
func (s *BaseChemsListener) EnterInstr_while(ctx *Instr_whileContext) {}

// ExitInstr_while is called when production instr_while is exited.
func (s *BaseChemsListener) ExitInstr_while(ctx *Instr_whileContext) {}

// EnterInstr_for_in is called when production instr_for_in is entered.
func (s *BaseChemsListener) EnterInstr_for_in(ctx *Instr_for_inContext) {}

// ExitInstr_for_in is called when production instr_for_in is exited.
func (s *BaseChemsListener) ExitInstr_for_in(ctx *Instr_for_inContext) {}

// EnterInstr_loop is called when production instr_loop is entered.
func (s *BaseChemsListener) EnterInstr_loop(ctx *Instr_loopContext) {}

// ExitInstr_loop is called when production instr_loop is exited.
func (s *BaseChemsListener) ExitInstr_loop(ctx *Instr_loopContext) {}

// EnterInstr_loop_ternario is called when production instr_loop_ternario is entered.
func (s *BaseChemsListener) EnterInstr_loop_ternario(ctx *Instr_loop_ternarioContext) {}

// ExitInstr_loop_ternario is called when production instr_loop_ternario is exited.
func (s *BaseChemsListener) ExitInstr_loop_ternario(ctx *Instr_loop_ternarioContext) {}

// EnterInstr_break is called when production instr_break is entered.
func (s *BaseChemsListener) EnterInstr_break(ctx *Instr_breakContext) {}

// ExitInstr_break is called when production instr_break is exited.
func (s *BaseChemsListener) ExitInstr_break(ctx *Instr_breakContext) {}

// EnterInstr_continue is called when production instr_continue is entered.
func (s *BaseChemsListener) EnterInstr_continue(ctx *Instr_continueContext) {}

// ExitInstr_continue is called when production instr_continue is exited.
func (s *BaseChemsListener) ExitInstr_continue(ctx *Instr_continueContext) {}

// EnterInstr_return is called when production instr_return is entered.
func (s *BaseChemsListener) EnterInstr_return(ctx *Instr_returnContext) {}

// ExitInstr_return is called when production instr_return is exited.
func (s *BaseChemsListener) ExitInstr_return(ctx *Instr_returnContext) {}

// EnterInstr_func is called when production instr_func is entered.
func (s *BaseChemsListener) EnterInstr_func(ctx *Instr_funcContext) {}

// ExitInstr_func is called when production instr_func is exited.
func (s *BaseChemsListener) ExitInstr_func(ctx *Instr_funcContext) {}

// EnterList_function_parameters is called when production list_function_parameters is entered.
func (s *BaseChemsListener) EnterList_function_parameters(ctx *List_function_parametersContext) {}

// ExitList_function_parameters is called when production list_function_parameters is exited.
func (s *BaseChemsListener) ExitList_function_parameters(ctx *List_function_parametersContext) {}

// EnterBlock_parameters_fn is called when production block_parameters_fn is entered.
func (s *BaseChemsListener) EnterBlock_parameters_fn(ctx *Block_parameters_fnContext) {}

// ExitBlock_parameters_fn is called when production block_parameters_fn is exited.
func (s *BaseChemsListener) ExitBlock_parameters_fn(ctx *Block_parameters_fnContext) {}

// EnterInstr_llamada is called when production instr_llamada is entered.
func (s *BaseChemsListener) EnterInstr_llamada(ctx *Instr_llamadaContext) {}

// ExitInstr_llamada is called when production instr_llamada is exited.
func (s *BaseChemsListener) ExitInstr_llamada(ctx *Instr_llamadaContext) {}

// EnterInstr_llamada_expre is called when production instr_llamada_expre is entered.
func (s *BaseChemsListener) EnterInstr_llamada_expre(ctx *Instr_llamada_expreContext) {}

// ExitInstr_llamada_expre is called when production instr_llamada_expre is exited.
func (s *BaseChemsListener) ExitInstr_llamada_expre(ctx *Instr_llamada_expreContext) {}

// EnterInstr_tipo is called when production instr_tipo is entered.
func (s *BaseChemsListener) EnterInstr_tipo(ctx *Instr_tipoContext) {}

// ExitInstr_tipo is called when production instr_tipo is exited.
func (s *BaseChemsListener) ExitInstr_tipo(ctx *Instr_tipoContext) {}

// EnterList_expression is called when production list_expression is entered.
func (s *BaseChemsListener) EnterList_expression(ctx *List_expressionContext) {}

// ExitList_expression is called when production list_expression is exited.
func (s *BaseChemsListener) ExitList_expression(ctx *List_expressionContext) {}

// EnterBlock_list_expression is called when production block_list_expression is entered.
func (s *BaseChemsListener) EnterBlock_list_expression(ctx *Block_list_expressionContext) {}

// ExitBlock_list_expression is called when production block_list_expression is exited.
func (s *BaseChemsListener) ExitBlock_list_expression(ctx *Block_list_expressionContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseChemsListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseChemsListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpre_log is called when production expre_log is entered.
func (s *BaseChemsListener) EnterExpre_log(ctx *Expre_logContext) {}

// ExitExpre_log is called when production expre_log is exited.
func (s *BaseChemsListener) ExitExpre_log(ctx *Expre_logContext) {}

// EnterExpre_rel is called when production expre_rel is entered.
func (s *BaseChemsListener) EnterExpre_rel(ctx *Expre_relContext) {}

// ExitExpre_rel is called when production expre_rel is exited.
func (s *BaseChemsListener) ExitExpre_rel(ctx *Expre_relContext) {}

// EnterExpre_arit is called when production expre_arit is entered.
func (s *BaseChemsListener) EnterExpre_arit(ctx *Expre_aritContext) {}

// ExitExpre_arit is called when production expre_arit is exited.
func (s *BaseChemsListener) ExitExpre_arit(ctx *Expre_aritContext) {}

// EnterExpre_valor is called when production expre_valor is entered.
func (s *BaseChemsListener) EnterExpre_valor(ctx *Expre_valorContext) {}

// ExitExpre_valor is called when production expre_valor is exited.
func (s *BaseChemsListener) ExitExpre_valor(ctx *Expre_valorContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseChemsListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseChemsListener) ExitPrimitivo(ctx *PrimitivoContext) {}

// EnterPrimitivo_casteo is called when production primitivo_casteo is entered.
func (s *BaseChemsListener) EnterPrimitivo_casteo(ctx *Primitivo_casteoContext) {}

// ExitPrimitivo_casteo is called when production primitivo_casteo is exited.
func (s *BaseChemsListener) ExitPrimitivo_casteo(ctx *Primitivo_casteoContext) {}

// EnterExpre_casteo is called when production expre_casteo is entered.
func (s *BaseChemsListener) EnterExpre_casteo(ctx *Expre_casteoContext) {}

// ExitExpre_casteo is called when production expre_casteo is exited.
func (s *BaseChemsListener) ExitExpre_casteo(ctx *Expre_casteoContext) {}

// EnterType_casteo is called when production type_casteo is entered.
func (s *BaseChemsListener) EnterType_casteo(ctx *Type_casteoContext) {}

// ExitType_casteo is called when production type_casteo is exited.
func (s *BaseChemsListener) ExitType_casteo(ctx *Type_casteoContext) {}

// EnterNativa_expre is called when production nativa_expre is entered.
func (s *BaseChemsListener) EnterNativa_expre(ctx *Nativa_expreContext) {}

// ExitNativa_expre is called when production nativa_expre is exited.
func (s *BaseChemsListener) ExitNativa_expre(ctx *Nativa_expreContext) {}

// EnterNative_block_abs is called when production native_block_abs is entered.
func (s *BaseChemsListener) EnterNative_block_abs(ctx *Native_block_absContext) {}

// ExitNative_block_abs is called when production native_block_abs is exited.
func (s *BaseChemsListener) ExitNative_block_abs(ctx *Native_block_absContext) {}
