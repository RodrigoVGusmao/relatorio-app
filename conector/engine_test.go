package conector

import (
	"io/fs"
	"reflect"
	"regexp"
	"testing"
	"testing/fstest"
)

func matchDirResult(t *testing.T, name string, fds fs.FS, dir string, compFunc compFunc, result []string, err bool) {
	t.Logf("[INFO] Executando %q", name)
	ret1, ret2 := matchDir(fds, dir, compFunc)
	if !reflect.DeepEqual(result, ret1) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret1, result)
	}
	if (ret2 != nil) != err {
		t.Errorf("[ERROR] Erro: Obtido %q", ret2)
	}
}
func Test_matchDir(t *testing.T) {
	file :=[][]byte{
		[]byte("Lorem Ipsum"), 
		fileRoot(
				fileEnv("usuarios", 
					fileUuidColumn("id", "user_id"),
				fileEnv("endpoints", 
					fileUuidColumn("id", "user_id"),),),),
		fileRoot(
				fileEnv("endpoints", 
					fileUuidColumn("id", "user_id"),),),
		fileRoot(
				fileEnv("teste", 
					fileUuidColumn("id", "user_id"),), 
				fileEnv("teste", 
					fileUuidColumn("id", "user_id"),),),}
	mockFS := fstest.MapFS{
		"test/app.yaml": &fstest.MapFile{Data: file[1]},
		"test/db.yml": &fstest.MapFile{Data: file[2]},
		"test/readme.txt": &fstest.MapFile{Data: file[0]},
		"test/db2.yml": &fstest.MapFile{Data: file[3]},
		"test/other.yaml": &fstest.MapFile{Data: file[0]},
		"test/double.yaml.url": &fstest.MapFile{Data: file[0]},
		"test/what.yaaml": &fstest.MapFile{Data: file[0]},
		"test/what.ybml": &fstest.MapFile{Data: file[0]},
		"test/folder.yaml/next.yaml": &fstest.MapFile{Data: file[1]},
		"test/folder.yaml/teste.txt": &fstest.MapFile{Data: file[0]},
	}
	
	var yamlRegex = regexp.MustCompile("\\.yaml$")
	
	tests := []struct {
		name string
		filesystem fs.FS
		dir string
		result []string
		err bool
		compFunc compFunc
	} {
		{
			name: "matchDir_happy",
			dir: "test",
			result: []string{"test/app.yaml","test/other.yaml"},
			err: false,
			compFunc: yamlRegex.MatchString,
		}, {
			name: "matchDir_happy2",
			dir: "test/folder.yaml",
			result: []string{"test/folder.yaml/next.yaml"},
			err: false,
			compFunc: yamlRegex.MatchString,
		}, {
			name: "matchDir_fail",
			dir: "config",
			result: nil,
			err: true,
			compFunc: yamlRegex.MatchString,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matchDirResult(t, test.name, mockFS, test.dir, test.compFunc, test.result, test.err)
		})
	}
}
