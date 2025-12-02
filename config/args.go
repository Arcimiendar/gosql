package config

import (
	"errors"
	"strconv"

	"github.com/namsral/flag"
)
import "os"

func ParseArgs() (bind string, port int, dslPath string, dbUri string, err error) {
	defBind, defPort, defDslPath, defDbUri := "127.0.0.1", 8080, ".", ""

	envBind := os.Getenv("BIND")
	if envBind == "" {
		envBind = defBind
	}

	envPort, envErr := strconv.Atoi(os.Getenv("PORT"))
	if envErr != nil {
		envPort = defPort
	}

	envDslPath := os.Getenv("DSL_PATH")
	if envDslPath == "" {
		envDslPath = defDslPath
	}
	envDbUri := os.Getenv("DB_URI")
	if envDbUri == "" {
		envDbUri = defDbUri
	}

	var portString string

	flag.StringVar(&portString, "port", "", "Port to listen on")
	flag.StringVar(&bind, "bind", "", "Address to listen on")
	flag.StringVar(&dslPath, "dslpath", "", "Path to DSL folder")
	flag.StringVar(&dbUri, "dburi", "", "DB URI")
	flag.Parse()

	if portString != "" {
		port, err = strconv.Atoi(portString)
		if err != nil {
			err = nil
			port = 8080
		}
	} else {
		port = envPort
	}

	if dslPath == "" {
		dslPath = envDslPath
	}
	if dbUri == "" {
		dbUri = envDbUri
	}
	if bind == "" {
		bind = envBind
	}

	if info, err := os.Stat(dslPath); err != nil || !info.IsDir() {
		return "", 0, "", "", err
	}

	if len(dbUri) == 0 {
		return "", 0, "", "", errors.New("DB URI is required")
	}

	return
}
