package main

import (
	_ "github.com/caddyserver/caddy/v2"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	_ "github.com/hashicorp/go-retryablehttp"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/julienschmidt/httprouter"
	_ "github.com/logzio/logzio-go"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/minio/minio-go/v7"
	_ "github.com/pkg/errors"
	_ "github.com/sendgrid/sendgrid-go"
	_ "github.com/sethvargo/go-envconfig"
	_ "github.com/spf13/cobra"
	_ "github.com/stretchr/testify"
	_ "github.com/suborbital/atmo"
	_ "github.com/suborbital/grav"
	_ "github.com/suborbital/reactr"
	_ "github.com/suborbital/sat"
	_ "github.com/suborbital/subo"
	_ "github.com/suborbital/vektor"
	_ "github.com/testcontainers/testcontainers-go"
	_ "gocloud.dev"
	_ "golang.org/x/mod"
	_ "gopkg.in/segmentio/analytics-go.v3"
	_ "gopkg.in/yaml.v2"
)

func main() {

}
