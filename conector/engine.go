package conector

import (
	"fmt"
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

func depthFirstSearch (env ConfigEnvironment, currentPoint string, seen map[string]bool) (int, []error) {
	if seen[currentPoint] {
		return 0, []error{fmt.Errorf("%s", currentPoint)}
	}
	
	seen[currentPoint] = true
	
	nodes := 0
	var errArr []error
	endpoint := env[currentPoint].(ConfigEndpoint)
	for column, value := range endpoint.Schema{
		if value.ForeignKey.Table != "" {
			nodesProp, err := depthFirstSearch(env, value.ForeignKey.Table, seen)
			if err != nil {
				for _, val := range err {
					errArr = append(errArr, fmt.Errorf("%s -> %v", currentPoint, val))
				}
				value.ForeignKey.Table = ""
				value.ForeignKey.Column = ""
				endpoint.Schema[column] = value
				env[currentPoint] = endpoint
				continue
			}
			nodes = max(nodes, nodesProp)
		}
	}
	
	seen[currentPoint] = false
	return nodes+1, errArr
}
