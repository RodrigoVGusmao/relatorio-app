package main

import (
	"gopkg.in/yaml.v3"
	"io/fs"
)

func loadConfigFiles(fds fs.FS, dirPath string) (configEnvironment, []error) {
	result := make(configEnvironment)
	var errProp []error
	
	files, err:= matchDir(fds, dirPath, matchYaml)
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
