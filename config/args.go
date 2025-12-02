package config

import (
	"errors"
	"strconv"

	"github.com/namsral/flag"
)
import "os"

func ParseArgs() (bind string, port int, dslPath string, dbUri string, err error) {
	bind = os.Getenv("BIND")
	if len(bind) == 0 {
		bind = "127.0.0.1"
	}

	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	err = nil

	dslPath = os.Getenv("DSL_PATH")
	if len(dslPath) == 0 {
		dslPath = "/DSL"
	}

	dbUri = os.Getenv("DB_URI")

	flag.IntVar(&port, "port", port, "Port to listen on")
	flag.StringVar(&bind, "bind", bind, "Address to listen on")
	flag.StringVar(&dslPath, "dslpath", dslPath, "Path to DSL folder")
	flag.StringVar(&dbUri, "dburi", dbUri, "DB URI")
	flag.Parse()

	if info, err := os.Stat(dslPath); err != nil || !info.IsDir() {
		return "", 0, "", "", err
	}

	if len(dbUri) == 0 {
		return "", 0, "", "", errors.New("DB URI is required")
	}

	return
}
