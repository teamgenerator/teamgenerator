package conf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

var (
	pgUser     = os.Getenv("TG_PG_USER")
	pgPassword = os.Getenv("TG_PG_PASSWORD")
	pgDatabase = os.Getenv("TG_PG_DATABASE")
	pgHost     = os.Getenv("TG_PG_HOST")
	pgPort     = os.Getenv("TG_PG_PORT")
	port       = fmt.Sprintf(":%s", os.Getenv("TG_API_PORT"))
)

type Conf struct {
	PgUser     string `envconfig:"TG_PG_USER" default:"postgres"`
	PgPassword string `envconfig:"TG_PG_PASSWORD" default:"password"`
	PgDatabase string `envconfig:"TG_PG_DATABASE" default:"postgres"`
	PgHost     string `envconfig:"TG_PG_HOST" default:"localhost"`
	PgPort     string `envconfig:"TG_PG_PORT" default:"5432"`
	ApiPort    string `envconfig:"TG_API_PORT" default:"3030"`
}

func NewConf() Conf {
	c := Conf{}
	envconfig.MustProcess("", &c)

	return c
}
