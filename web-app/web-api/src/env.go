package main

import (
	"github.com/EikenDram/kube-utils/web-api/config"
)

type Env struct {
	config *config.ConfigStruct
}

func Init() (*Env, error) {
	//load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initalise Env with a models.BookModel instance (which in turn wraps
	// the connection pool).
	env := &Env{
		config: cfg,
	}

	return env, nil
}
