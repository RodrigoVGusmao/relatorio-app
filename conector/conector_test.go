package conector

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func loadConfigFilesResult(t *testing.T, name string, fds fs.FS, dir string, result ConfigEnvironment) {
	t.Logf("[INFO] Executando %q", name)
	ret1, ret2 := loadConfigFiles(fds, dir)
	if !reflect.DeepEqual(result, ret1) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q, log: %q", ret1, result, ret2)
	}
}
func Test_loadConfigFiles(t *testing.T) {
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
		"test/folder.yaml/next.yaml": &fstest.MapFile{Data: file[2]},
		"test/folder.yaml/teste.txt": &fstest.MapFile{Data: file[0]},
	}
	
	tests := []struct {
		name string
		dir string
		result ConfigEnvironment
	} {
		{
			name: "loadConfigFiles_happy",
			dir: "test",
			result: ConfigEnvironment{
				"usuarios": env(map[string]ConfigColumn{
					"id": uuidColumn("user_id"),
				}),
				"endpoints": env(map[string]ConfigColumn{
					"id": uuidColumn("user_id"),
				}),
				"teste": env(map[string]ConfigColumn{
					"id": uuidColumn("user_id"),
				}),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loadConfigFilesResult(t, test.name, mockFS, test.dir, test.result)
		})
	}
}

func getInvalidEnv() ConfigEnvironment {
	return ConfigEnvironment{
		"env.3": env(map[string]ConfigColumn{
			"key.1": deColumn("pk", true),
			"key.2": deColumn("dummy", false),
			"key.3": fkColumn("fk", "env2", "pk"),
		}),
		"env2": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy2", false),
			"key3": fkColumn("fk", "env.3", "pk"),
		}),
		"env1": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
		}),
	}
}

func getLongEnv() ConfigEnvironment {
	return ConfigEnvironment{
		"env3": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
			"key3": fkColumn("fk", "env2", "pk"),
		}),
		"env2": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy2", false),
			"key3": fkColumn("fk", "env1", "pk"),
		}),
		"env1": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
			"key3": fkColumn("fk", "env4", "pk"),
		}),
		"env4": env(map[string]ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
		}),
	}
}

func normalizeDataResult(t *testing.T, name string, input []byte, endpoint ConfigEndpoint, tableName string, result normRecord) {
	t.Logf("[INFO] Executando %q", name)
	ret, _ := normalizeData(input, endpoint, tableName)
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}
func Test_normalizeData(t *testing.T) {
	endpoint := getLongEnv()["env4"].(ConfigEndpoint)
	
	tests := []struct {
		name string
		rawData []byte
		config ConfigEndpoint
		table string
		result normRecord
	} {
		{
			name: "normalizeData_unique",
			rawData: jsonData(jsonRecord(
				jsonString("key1", "id1"),
				jsonString("key2", "dummy"),
			)),
			config: endpoint,
			table: "env4",
			result: normRecord{
				normData{"env4.pk": "id1", "env4.dummy3": "dummy",},
			},
		},
		{
			name: "normalizeData_multiple",
			rawData: jsonData(
				jsonRecord(
					jsonString("key1", "id1"),
					jsonString("key2", "dummy"),),
				jsonRecord(
					jsonString("key1", "id2"),
					jsonString("key2", "dummy"),),),
			config: endpoint,
			table: "env4",
			result: normRecord{
				normData{"env4.pk": "id1", "env4.dummy3": "dummy",},
				normData{"env4.pk": "id2", "env4.dummy3": "dummy",},
			},
		},
		{
			name: "normalizeData_invalidJson",
			rawData: []byte(""),
			config: endpoint,
			table: "env4",
			result: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			normalizeDataResult(t, test.name, test.rawData, test.config, test.table, test.result)
		})
	}
}

func compileConfigEnvironmentResult(t *testing.T, name string, config ConfigEnvironment, root string) {
	t.Logf("[INFO] Executando %q", name)
	ret, _ := compileConfigEnvironment(config)
	if !reflect.DeepEqual(root, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, root)
	}
}
func Test_compileConfigEnvironment(t *testing.T) {
	tests := []struct {
		name string
		config ConfigEnvironment
		root string
	} {
		{
			name: "compileConfigEnvironment_invalid",
			config: getInvalidEnv(),
			root: "",
		},
		{
			name: "compileConfigEnvironment_long",
			config: getLongEnv(),
			root: "env3",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compileConfigEnvironmentResult(t, test.name, test.config, test.root)
		})
	}
}
