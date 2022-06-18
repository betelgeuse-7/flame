package typecheck

import (
	"flame/ast"
	"flame/env"
)

type TypecheckError struct {
	PosX, PosY int
	Message    string
}

type Typechecker struct {
	program *ast.Program
	env     *env.Env
	errors  []TypecheckError
}

func New(program *ast.Program, env *env.Env) *Typechecker {
	t := &Typechecker{program: program, env: env, errors: []TypecheckError{}}
	return t
}

func (t *Typechecker) Check() {

}

func (t *Typechecker) Errors() []TypecheckError {
	return t.errors
}
