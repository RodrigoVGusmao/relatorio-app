package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
)

type configColumn struct {
	DataType string		`yaml:"type"`
	IsPrimaryKey bool	`yaml:"primary_key"`
	Rename string		`yaml:"name"`
	ForeignKey string	`yaml:"foreign_key"`
}

type configEndpoint struct {
	EndpointLocation string		`yaml:"endpoint"`
	Schema map[string]configColumn	`yaml:"schema"`
}

type configEnvironment map[string]any

func (m *configEnvironment) UnmarshalYAML(value *yaml.Node) error {
	if *m == nil {
		*m = make(configEnvironment)
	}
	if value.Kind != yaml.MappingNode {
		return errors.New("yaml file is invalid or empty")
	}
	
	elements := value.Content
	for i:=0;i<len(elements);i+=2 {
		if _, exists := (*m)[elements[i].Value]; exists {
			errs, _ := (*m)["\x12err"].([]error)
			(*m)["\x12err"] = append(errs, fmt.Errorf("endpoint %q already declared", elements[i].Value))
			continue
		}
		
		var endpoint configEndpoint
		if err := elements[i+1].Decode(&endpoint); err != nil {
			errs, _ := (*m)["\x12err"].([]error)
			(*m)["\x12err"] = append(errs, err)
			continue
		}
		(*m)[elements[i].Value] = endpoint
	}
	
	return nil
}
