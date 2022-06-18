package env

type Env struct{}

type Evaluator struct {
	env *Env
}

func New() *Evaluator {
	e := &Evaluator{env: &Env{}}
	return e
}
