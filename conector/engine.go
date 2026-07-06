package main

import (
	"io/fs"
	"path"
	"regexp"
)

type compFunc func(filename string) bool

var yamlRegex = regexp.MustCompile("\\.ya?ml$")
func matchYaml(filename string) bool {
	return yamlRegex.MatchString(filename)
}

func matchDir(fds fs.FS, dir string, compFunc compFunc) ([]string, error) {
	var ret []string
	files, err := fs.ReadDir(fds, dir)
	for _, val := range files {
		if !val.IsDir() && compFunc(val.Name()) {
			ret = append(ret, path.Join(dir, val.Name()))
		}
	}
	return ret, err
}
