package models

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	GlobalPrefix string `mapstructure:"GLOBAL_PREFIX"`
	Email        struct {
		Password string `mapstructure:"PASSWORD"`
		Username string `mapstructure:"USERNAME"`
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
	} `mapstructure:"EMAIL"`
	JWT struct {
		SecretKey   string `mapstructure:"SECRET_KEY"`
		ExpiresTime int    `mapstructure:"EXPIRES_TIME"`
	} `mapstructure:"JWT"`
	Port int `mapstructure:"PORT"`
}

type CommonStatusEnum string

const (
	INACTIVE CommonStatusEnum = "inactive"
	ACTIVE   CommonStatusEnum = "active"
)

type QueryStringParams struct {
	OrderBy        string `form:"order_by" validate:"omitempty"`
	OrderDirection string `form:"order_direction" validate:"omitempty,alpha"`
	Page           int    `form:"page" validate:"omitempty,min=1"`
	PageSize       int    `form:"page_size" validate:"omitempty,min=1,max=100"`
}

type Controller struct {
	Routes map[string]func(http.ResponseWriter, *http.Request)
	Prefix string
}

type PostgresConfig struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}
