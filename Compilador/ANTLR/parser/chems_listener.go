// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterEnd_instr is called when entering the end_instr production.
	EnterEnd_instr(c *End_instrContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstr_println is called when entering the instr_println production.
	EnterInstr_println(c *Instr_printlnContext)

	// EnterInstr_main is called when entering the instr_main production.
	EnterInstr_main(c *Instr_mainContext)

	// EnterInstr_declaracion is called when entering the instr_declaracion production.
	EnterInstr_declaracion(c *Instr_declaracionContext)

	// EnterInstr_asignacion is called when entering the instr_asignacion production.
	EnterInstr_asignacion(c *Instr_asignacionContext)

	// EnterInstr_if is called when entering the instr_if production.
	EnterInstr_if(c *Instr_ifContext)

	// EnterInstr_else_if is called when entering the instr_else_if production.
	EnterInstr_else_if(c *Instr_else_ifContext)

	// EnterInstr_ternario is called when entering the instr_ternario production.
	EnterInstr_ternario(c *Instr_ternarioContext)

	// EnterInstr_else_if_ternario is called when entering the instr_else_if_ternario production.
	EnterInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// EnterInstr_match is called when entering the instr_match production.
	EnterInstr_match(c *Instr_matchContext)

	// EnterList_case is called when entering the list_case production.
	EnterList_case(c *List_caseContext)

	// EnterInstr_case is called when entering the instr_case production.
	EnterInstr_case(c *Instr_caseContext)

	// EnterList_expre_case is called when entering the list_expre_case production.
	EnterList_expre_case(c *List_expre_caseContext)

	// EnterBlock_case is called when entering the block_case production.
	EnterBlock_case(c *Block_caseContext)

	// EnterBlock_instr_match is called when entering the block_instr_match production.
	EnterBlock_instr_match(c *Block_instr_matchContext)

	// EnterInstr_default is called when entering the instr_default production.
	EnterInstr_default(c *Instr_defaultContext)

	// EnterBlock_default is called when entering the block_default production.
	EnterBlock_default(c *Block_defaultContext)

	// EnterInstr_match_ter is called when entering the instr_match_ter production.
	EnterInstr_match_ter(c *Instr_match_terContext)

	// EnterList_case_ternario is called when entering the list_case_ternario production.
	EnterList_case_ternario(c *List_case_ternarioContext)

	// EnterInstr_case_ter is called when entering the instr_case_ter production.
	EnterInstr_case_ter(c *Instr_case_terContext)

	// EnterList_expre_case_ter is called when entering the list_expre_case_ter production.
	EnterList_expre_case_ter(c *List_expre_case_terContext)

	// EnterBlock_case_ter is called when entering the block_case_ter production.
	EnterBlock_case_ter(c *Block_case_terContext)

	// EnterInstr_default_ter is called when entering the instr_default_ter production.
	EnterInstr_default_ter(c *Instr_default_terContext)

	// EnterInstr_while is called when entering the instr_while production.
	EnterInstr_while(c *Instr_whileContext)

	// EnterInstr_for_in is called when entering the instr_for_in production.
	EnterInstr_for_in(c *Instr_for_inContext)

	// EnterInstr_loop is called when entering the instr_loop production.
	EnterInstr_loop(c *Instr_loopContext)

	// EnterInstr_loop_ternario is called when entering the instr_loop_ternario production.
	EnterInstr_loop_ternario(c *Instr_loop_ternarioContext)

	// EnterInstr_break is called when entering the instr_break production.
	EnterInstr_break(c *Instr_breakContext)

	// EnterInstr_continue is called when entering the instr_continue production.
	EnterInstr_continue(c *Instr_continueContext)

	// EnterInstr_return is called when entering the instr_return production.
	EnterInstr_return(c *Instr_returnContext)

	// EnterInstr_func is called when entering the instr_func production.
	EnterInstr_func(c *Instr_funcContext)

	// EnterList_function_parameters is called when entering the list_function_parameters production.
	EnterList_function_parameters(c *List_function_parametersContext)

	// EnterBlock_parameters_fn is called when entering the block_parameters_fn production.
	EnterBlock_parameters_fn(c *Block_parameters_fnContext)

	// EnterInstr_llamada is called when entering the instr_llamada production.
	EnterInstr_llamada(c *Instr_llamadaContext)

	// EnterInstr_llamada_expre is called when entering the instr_llamada_expre production.
	EnterInstr_llamada_expre(c *Instr_llamada_expreContext)

	// EnterInstr_tipo is called when entering the instr_tipo production.
	EnterInstr_tipo(c *Instr_tipoContext)

	// EnterList_expression is called when entering the list_expression production.
	EnterList_expression(c *List_expressionContext)

	// EnterBlock_list_expression is called when entering the block_list_expression production.
	EnterBlock_list_expression(c *Block_list_expressionContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpre_log is called when entering the expre_log production.
	EnterExpre_log(c *Expre_logContext)

	// EnterExpre_rel is called when entering the expre_rel production.
	EnterExpre_rel(c *Expre_relContext)

	// EnterExpre_arit is called when entering the expre_arit production.
	EnterExpre_arit(c *Expre_aritContext)

	// EnterExpre_valor is called when entering the expre_valor production.
	EnterExpre_valor(c *Expre_valorContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// EnterPrimitivo_casteo is called when entering the primitivo_casteo production.
	EnterPrimitivo_casteo(c *Primitivo_casteoContext)

	// EnterExpre_casteo is called when entering the expre_casteo production.
	EnterExpre_casteo(c *Expre_casteoContext)

	// EnterType_casteo is called when entering the type_casteo production.
	EnterType_casteo(c *Type_casteoContext)

	// EnterNativa_expre is called when entering the nativa_expre production.
	EnterNativa_expre(c *Nativa_expreContext)

	// EnterNative_block_abs is called when entering the native_block_abs production.
	EnterNative_block_abs(c *Native_block_absContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitEnd_instr is called when exiting the end_instr production.
	ExitEnd_instr(c *End_instrContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstr_println is called when exiting the instr_println production.
	ExitInstr_println(c *Instr_printlnContext)

	// ExitInstr_main is called when exiting the instr_main production.
	ExitInstr_main(c *Instr_mainContext)

	// ExitInstr_declaracion is called when exiting the instr_declaracion production.
	ExitInstr_declaracion(c *Instr_declaracionContext)

	// ExitInstr_asignacion is called when exiting the instr_asignacion production.
	ExitInstr_asignacion(c *Instr_asignacionContext)

	// ExitInstr_if is called when exiting the instr_if production.
	ExitInstr_if(c *Instr_ifContext)

	// ExitInstr_else_if is called when exiting the instr_else_if production.
	ExitInstr_else_if(c *Instr_else_ifContext)

	// ExitInstr_ternario is called when exiting the instr_ternario production.
	ExitInstr_ternario(c *Instr_ternarioContext)

	// ExitInstr_else_if_ternario is called when exiting the instr_else_if_ternario production.
	ExitInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// ExitInstr_match is called when exiting the instr_match production.
	ExitInstr_match(c *Instr_matchContext)

	// ExitList_case is called when exiting the list_case production.
	ExitList_case(c *List_caseContext)

	// ExitInstr_case is called when exiting the instr_case production.
	ExitInstr_case(c *Instr_caseContext)

	// ExitList_expre_case is called when exiting the list_expre_case production.
	ExitList_expre_case(c *List_expre_caseContext)

	// ExitBlock_case is called when exiting the block_case production.
	ExitBlock_case(c *Block_caseContext)

	// ExitBlock_instr_match is called when exiting the block_instr_match production.
	ExitBlock_instr_match(c *Block_instr_matchContext)

	// ExitInstr_default is called when exiting the instr_default production.
	ExitInstr_default(c *Instr_defaultContext)

	// ExitBlock_default is called when exiting the block_default production.
	ExitBlock_default(c *Block_defaultContext)

	// ExitInstr_match_ter is called when exiting the instr_match_ter production.
	ExitInstr_match_ter(c *Instr_match_terContext)

	// ExitList_case_ternario is called when exiting the list_case_ternario production.
	ExitList_case_ternario(c *List_case_ternarioContext)

	// ExitInstr_case_ter is called when exiting the instr_case_ter production.
	ExitInstr_case_ter(c *Instr_case_terContext)

	// ExitList_expre_case_ter is called when exiting the list_expre_case_ter production.
	ExitList_expre_case_ter(c *List_expre_case_terContext)

	// ExitBlock_case_ter is called when exiting the block_case_ter production.
	ExitBlock_case_ter(c *Block_case_terContext)

	// ExitInstr_default_ter is called when exiting the instr_default_ter production.
	ExitInstr_default_ter(c *Instr_default_terContext)

	// ExitInstr_while is called when exiting the instr_while production.
	ExitInstr_while(c *Instr_whileContext)

	// ExitInstr_for_in is called when exiting the instr_for_in production.
	ExitInstr_for_in(c *Instr_for_inContext)

	// ExitInstr_loop is called when exiting the instr_loop production.
	ExitInstr_loop(c *Instr_loopContext)

	// ExitInstr_loop_ternario is called when exiting the instr_loop_ternario production.
	ExitInstr_loop_ternario(c *Instr_loop_ternarioContext)

	// ExitInstr_break is called when exiting the instr_break production.
	ExitInstr_break(c *Instr_breakContext)

	// ExitInstr_continue is called when exiting the instr_continue production.
	ExitInstr_continue(c *Instr_continueContext)

	// ExitInstr_return is called when exiting the instr_return production.
	ExitInstr_return(c *Instr_returnContext)

	// ExitInstr_func is called when exiting the instr_func production.
	ExitInstr_func(c *Instr_funcContext)

	// ExitList_function_parameters is called when exiting the list_function_parameters production.
	ExitList_function_parameters(c *List_function_parametersContext)

	// ExitBlock_parameters_fn is called when exiting the block_parameters_fn production.
	ExitBlock_parameters_fn(c *Block_parameters_fnContext)

	// ExitInstr_llamada is called when exiting the instr_llamada production.
	ExitInstr_llamada(c *Instr_llamadaContext)

	// ExitInstr_llamada_expre is called when exiting the instr_llamada_expre production.
	ExitInstr_llamada_expre(c *Instr_llamada_expreContext)

	// ExitInstr_tipo is called when exiting the instr_tipo production.
	ExitInstr_tipo(c *Instr_tipoContext)

	// ExitList_expression is called when exiting the list_expression production.
	ExitList_expression(c *List_expressionContext)

	// ExitBlock_list_expression is called when exiting the block_list_expression production.
	ExitBlock_list_expression(c *Block_list_expressionContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpre_log is called when exiting the expre_log production.
	ExitExpre_log(c *Expre_logContext)

	// ExitExpre_rel is called when exiting the expre_rel production.
	ExitExpre_rel(c *Expre_relContext)

	// ExitExpre_arit is called when exiting the expre_arit production.
	ExitExpre_arit(c *Expre_aritContext)

	// ExitExpre_valor is called when exiting the expre_valor production.
	ExitExpre_valor(c *Expre_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)

	// ExitPrimitivo_casteo is called when exiting the primitivo_casteo production.
	ExitPrimitivo_casteo(c *Primitivo_casteoContext)

	// ExitExpre_casteo is called when exiting the expre_casteo production.
	ExitExpre_casteo(c *Expre_casteoContext)

	// ExitType_casteo is called when exiting the type_casteo production.
	ExitType_casteo(c *Type_casteoContext)

	// ExitNativa_expre is called when exiting the nativa_expre production.
	ExitNativa_expre(c *Nativa_expreContext)

	// ExitNative_block_abs is called when exiting the native_block_abs production.
	ExitNative_block_abs(c *Native_block_absContext)
}
