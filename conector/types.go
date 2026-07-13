package conector

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
)

type FKey struct {
	Table string	`yaml:"table"`
	Column string	`yaml:"column"`
}

type ConfigColumn struct {
	DataType string		`yaml:"type"`
	IsPrimaryKey bool	`yaml:"primary_key"`
	Rename string		`yaml:"name"`
	ForeignKey FKey		`yaml:"foreign_key"`
}

type ConfigEndpoint struct {
	EndpointLocation string		`yaml:"endpoint"`
	Schema map[string]ConfigColumn	`yaml:"schema"`
}

type ConfigEnvironment map[string]any

type NormData map[string]any

type NormRecord []NormData

type NormRecords map[string]NormRecord

func (m *ConfigEnvironment) UnmarshalYAML(value *yaml.Node) error {
	if *m == nil {
		*m = make(ConfigEnvironment)
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
		
		var endpoint ConfigEndpoint
		if err := elements[i+1].Decode(&endpoint); err != nil {
			errs, _ := (*m)["\x12err"].([]error)
			(*m)["\x12err"] = append(errs, err)
			continue
		}
		(*m)[elements[i].Value] = endpoint
	}
	
	return nil
}
