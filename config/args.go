package config

import "github.com/namsral/flag"
import "os"

func ParseArgs() (bind string, port int, dslPath string, err error) {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&bind, "bind", "127.0.0.1", "Address to listen on")
	flag.StringVar(&dslPath, "dsl_path", "/DSL", "Path to DSL folder")
	flag.Parse()

	if info, err := os.Stat(dslPath); err != nil || !info.IsDir() {
		return "", 0, "", err
	}

	return
}
