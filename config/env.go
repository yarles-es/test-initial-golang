package env

import (
	"github.com/Netflix/go-env"
)

type Environment struct {
	Port   int `env:"PORT"`
	Extras env.EnvSet
}

var Env Environment

func GetEnv() Environment {
	return Env
}

func init() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		panic(err)
	}
}
