package dsl

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type EndpointType string

const (
	GET  EndpointType = "GET"
	POST EndpointType = "POST"
)

type Endpoint struct {
	RelativePath string
	Content      string
	Type         EndpointType
}

type Domain struct {
	Name      string
	Endpoints []Endpoint
}

func parseEndpointType(folderName string) (EndpointType, error) {
	switch folderName {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	default:
		return "", errors.New("invalid endpoint type")
	}
}

func parseEndpointsWithMethod(
	path string, endpoints *[]Endpoint, relativePath string, endpointType EndpointType,
) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			parseEndpointsWithMethod(
				filepath.Join(path, entry.Name()), endpoints,
				relativePath+"/"+entry.Name(), endpointType,
			)
		} else {
			if !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}

			data, err := os.ReadFile(filepath.Join(path, entry.Name()))

			if err != nil {
				continue
			}

			parts := strings.Split(entry.Name(), ".")

			endpoint := Endpoint{
				RelativePath: relativePath + "/" + strings.Join(parts[:len(parts)-1], "."),
				Content:      string(data),
				Type:         endpointType,
			}

			*endpoints = append(*endpoints, endpoint)
		}
	}
}

func parseEndpoints(path string) (endpoints []Endpoint, err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		endpointType, err := parseEndpointType(entry.Name())
		if err != nil {
			continue
		}

		parseEndpointsWithMethod(
			filepath.Join(path, entry.Name()), &endpoints, "", endpointType,
		)
	}

	return
}

func ParseDSL(dslPath string) (domains []Domain, err error) {
	entries, err := os.ReadDir(dslPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var endpoints []Endpoint
		endpoints, err = parseEndpoints(filepath.Join(dslPath, entry.Name()))
		if err == nil && len(endpoints) > 0 {
			domains = append(domains, Domain{Name: entry.Name(), Endpoints: endpoints})
		}
	}

	return
}
