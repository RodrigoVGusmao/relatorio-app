package main

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func loadConfigFilesResult(t *testing.T, name string, fds fs.FS, dir string, result configEnvironment) {
	t.Logf("[INFO] Executando %q", name)
	ret1, ret2 := loadConfigFiles(fds, dir)
	if !reflect.DeepEqual(result, ret1) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q, log: %q", ret1, result, ret2)
	}
}
func Test_loadConfigFiles(t *testing.T) {
	file :=[][]byte{[] byte("Lorem Ipsum"), []byte(`# version: "1.0"
usuarios:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"
endpoints:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"`), []byte(`endpoints:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"`), []byte(`teste:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"
teste:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"`), []byte(`endpoints:
    endpoint: "http://mock-api-1/usuarios/"
    schema:
        id:
            name: "user_id"
            primary_key: true
            type: "uuid"`)}
	mockFS := fstest.MapFS{
		"test/app.yaml": &fstest.MapFile{Data: file[1]},
		"test/db.yml": &fstest.MapFile{Data: file[2]},
		"test/readme.txt": &fstest.MapFile{Data: file[0]},
		"test/db2.yml": &fstest.MapFile{Data: file[3]},
		"test/other.yaml": &fstest.MapFile{Data: file[0]},
		"test/double.yaml.url": &fstest.MapFile{Data: file[0]},
		"test/what.yaaml": &fstest.MapFile{Data: file[0]},
		"test/what.ybml": &fstest.MapFile{Data: file[0]},
		"test/folder.yaml/next.yaml": &fstest.MapFile{Data: file[4]},
		"test/folder.yaml/teste.txt": &fstest.MapFile{Data: file[0]},
	}
	
	tests := []struct {
		name string
		dir string
		result configEnvironment
	} {
		{
			name: "loadConfigFiles_happy",
			dir: "test",
			result: configEnvironment{
				"usuarios": configEndpoint{
					EndpointLocation: "http://mock-api-1/usuarios/",
					Schema: map[string]configColumn{
						"id": {
							Rename: "user_id",
							IsPrimaryKey: true,
							DataType: "uuid",
						},
					},
				},
				"endpoints": configEndpoint{
					EndpointLocation: "http://mock-api-1/usuarios/",
					Schema: map[string]configColumn{
						"id": {
							Rename: "user_id",
							IsPrimaryKey: true,
							DataType: "uuid",
						},
					},
				},
				"teste": configEndpoint{
					EndpointLocation: "http://mock-api-1/usuarios/",
					Schema: map[string]configColumn{
						"id": {
							Rename: "user_id",
							IsPrimaryKey: true,
							DataType: "uuid",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loadConfigFilesResult(t, test.name, mockFS, test.dir, test.result)
		})
	}
}
