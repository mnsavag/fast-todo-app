package log

import "github.com/go-playground/validator/v10"

type Config struct {
	LogEnv string `env:"LOG_ENV" validate:"required,oneof=local prod"`
}

func (conf *Config) Validate() error {
	return validator.New().Struct(conf)
}
