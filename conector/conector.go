package conector

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"strings"
)

func LoadConfigFiles(fds fs.FS, dirPath string) (ConfigEnvironment, []error) {
	result := make(ConfigEnvironment)
	var errProp []error
	
	files, err:=  matchDir(fds, dirPath, matchYaml)
	if err != nil {
		errProp = append(errProp, err)
	}

	for _, file := range files {
		bytes, err := fs.ReadFile(fds, file)
		if err != nil {
			errProp = append(errProp, err)
			continue
		}
		
		err = yaml.Unmarshal(bytes, &result)
		if errArr, exists := result["\x12err"].([]error); exists || err != nil{
			if exists {
				errProp = append(errProp, errArr...)
				delete(result, "\x12err")
			}
			if err != nil {
				errProp = append(errProp, err)
			}
			continue
		}
	}
	
	return result, errProp
}

func NormalizeData(rawData []byte, endpoint ConfigEndpoint, tableName string) (NormRecord, error) {	
	size:=len(rawData)
	if size == 0 {
		rawData = []byte("[]")
	} else if rawData[0] == '{' {
		result := make([]byte, size+2)
		result[0] = '['
		copy(result[1:], rawData)
		result[size+1] = ']'
		rawData = result
	}
	
	var jsonDataArr []map[string]any
	if err := json.Unmarshal(rawData, &jsonDataArr); err != nil {
		return nil, err
	}
	
	var normalizedArr NormRecord
	for _, jsonData := range jsonDataArr {
		normalized := make(NormData)
		for column, meta := range endpoint.Schema {
			normalized[tableName+"."+meta.Rename] = jsonData[column]
		}
		normalizedArr = append(normalizedArr, normalized)
	}
	
	return normalizedArr, nil
}

func CompileConfigEnvironment(env ConfigEnvironment) (string, []error) {
	var errArr []error
	var ret string
	var rootNodeLen int
	for key, value := range env {
		if strings.ContainsRune(key, '.') {
			return "", append(errArr, fmt.Errorf("found invalid character '.' on endpoint %q", key))
		}
		for _, column := range value.(ConfigEndpoint).Schema {
			if strings.ContainsRune(column.Rename, '.') {
				return "", append(errArr, fmt.Errorf("found invalid character '.' on column %q", key+"."+column.Rename))
			}
		}
		
		seen := make(map[string]bool)
		nodeProp, err := depthFirstSearch(env, key, seen)
		if err != nil {
			errArr = append(errArr, err...)
		}
		if nodeProp > rootNodeLen {
			rootNodeLen = nodeProp
			ret = key
		}
	}
	
	return ret, errArr
}
