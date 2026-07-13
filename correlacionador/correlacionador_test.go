package correlacionador

import (
	"reflect"
	"relatorio/conector"
	"testing"
)

func deColumn(rename string, isPK bool) conector.ConfigColumn {
	return conector.ConfigColumn{
		DataType:     "dummy",
		IsPrimaryKey: isPK,
		Rename:       rename,
	}
}
func fkColumn(rename string, fkTable, fkCol string) conector.ConfigColumn {
	col := deColumn(rename, false)
	col.ForeignKey = conector.FKey{Table: fkTable, Column: fkCol}
	return col
}
func env(schema map[string]conector.ConfigColumn) conector.ConfigEndpoint {
	return conector.ConfigEndpoint{
		EndpointLocation: "dummy",
		Schema:           schema,
	}
}

func getDoubleKeyEnv() conector.ConfigEnvironment {
	return conector.ConfigEnvironment{
		"env1": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
			"key3": fkColumn("fk1", "env2", "pk"),
			"key4": fkColumn("fk2", "env3", "pk"),
		}),
		"env2": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy2", false),
		}),
		"env3": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
		}),
	}
}
func getSeriesKeyEnv() conector.ConfigEnvironment {
	return conector.ConfigEnvironment{
		"env1": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
			"key3": fkColumn("fk", "env2", "pk"),
		}),
		"env2": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy2", false),
			"key3": fkColumn("fk", "env3", "pk"),
		}),
		"env3": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
		}),
	}
}
func getReverseSeriesKeyEnv() conector.ConfigEnvironment {
	return conector.ConfigEnvironment{
		"env3": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
			"key3": fkColumn("fk", "env2", "pk"),
		}),
		"env2": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy2", false),
			"key3": fkColumn("fk", "env1", "pk"),
		}),
		"env1": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy3", false),
		}),
	}
}
func getCollisionEnv() conector.ConfigEnvironment {
	return conector.ConfigEnvironment{
		"en.vTest": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
			"key3": fkColumn("fk", "env.Test", "pk"),
		}),
		"env.Test": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
		}),
	}
}
func getUnrelatedEnv() conector.ConfigEnvironment {
	return conector.ConfigEnvironment{
		"env1": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
		}),
		"env2": env(map[string]conector.ConfigColumn{
			"key1": deColumn("pk", true),
			"key2": deColumn("dummy", false),
		}),
	}
}

func makeUnifiedDataResult(t *testing.T, name string, records normRecords, config conector.ConfigEnvironment, parent string, result normRecord) {
	t.Logf("[INFO] Executando %q", name)
	ret := makeUnifiedData(records, config, parent)
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}
func Test_makeUnifiedDataResult(t *testing.T) {
	doubleKeyEnvironment := getDoubleKeyEnv()
	seriesKeyEnvironment := getSeriesKeyEnv()
	reverseSeriesKeyEnvironment := getReverseSeriesKeyEnv()
	collisionEnvironment := getCollisionEnv()
	unrelatedEnvironment := getUnrelatedEnv()
	
	tests := []struct {
		name string
		config conector.ConfigEnvironment
		data normRecords
		parent string
		result normRecord
	} {
		{
			name: "makeUnifiedData_doubleKey",
			config: doubleKeyEnvironment,
			parent: "env1",
			data: normRecords {
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy": "dummy", "env1.fk1": "id1", "env1.fk2": "id1",},
					{"env1.pk": "id2", "env1.dummy": "dummy", "env1.fk1": "id1", "env1.fk2": "id2",},
					{"env1.pk": "id3", "env1.dummy": "dummy", "env1.fk1": "id3", "env1.fk2": "id1",},
				},
				"env2": normRecord{
					{"env2.pk": "id1", "env2.dummy2": "dummy1",},
					{"env2.pk": "id2", "env2.dummy2": "dummy2",},
					{"env2.pk": "id3", "env2.dummy2": "dummy3",},
				},
				"env3": normRecord{
					{"env3.pk": "id1", "env3.dummy3": "dummy4",},
					{"env3.pk": "id2", "env3.dummy3": "dummy5",},
					{"env3.pk": "id3", "env3.dummy3": "dummy6",},
				},
			},
			result: normRecord {
				{"env1.pk": "id1", "env1.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env3.pk": "id1", "env3.dummy3": "dummy4",},
				{"env1.pk": "id2", "env1.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env3.pk": "id2", "env3.dummy3": "dummy5",},
				{"env1.pk": "id3", "env1.dummy": "dummy", "env2.pk": "id3", "env2.dummy2": "dummy3", "env3.pk": "id1", "env3.dummy3": "dummy4",},
			},
		},
		{
			name: "makeUnifiedData_seriesKey",
			config: seriesKeyEnvironment,
			parent: "env1",
			data: normRecords {
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy": "dummy", "env1.fk": "id1",},
					{"env1.pk": "id2", "env1.dummy": "dummy", "env1.fk": "id1",},
					{"env1.pk": "id3", "env1.dummy": "dummy", "env1.fk": "id3",},
				},
				"env2": normRecord{
					{"env2.pk": "id1", "env2.dummy2": "dummy1", "env2.fk": "id1",},
					{"env2.pk": "id2", "env2.dummy2": "dummy2", "env2.fk": "id3",},
					{"env2.pk": "id3", "env2.dummy2": "dummy3", "env2.fk": "id2",},
				},
				"env3": normRecord{
					{"env3.pk": "id1", "env3.dummy3": "dummy4",},
					{"env3.pk": "id2", "env3.dummy3": "dummy5",},
					{"env3.pk": "id3", "env3.dummy3": "dummy6",},
				},
			},
			result: normRecord {
				{"env1.pk": "id1", "env1.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env3.pk": "id1", "env3.dummy3": "dummy4",},
				{"env1.pk": "id2", "env1.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env3.pk": "id1", "env3.dummy3": "dummy4",},
				{"env1.pk": "id3", "env1.dummy": "dummy", "env2.pk": "id3", "env2.dummy2": "dummy3", "env3.pk": "id2", "env3.dummy3": "dummy5",},
			},
		},
		{
			name: "makeUnifiedData_reverseSeriesKey",
			config: reverseSeriesKeyEnvironment,
			parent: "env3",
			data: normRecords {
				"env3": normRecord{
					{"env3.pk": "id1", "env3.dummy": "dummy", "env3.fk": "id1",},
					{"env3.pk": "id2", "env3.dummy": "dummy", "env3.fk": "id1",},
					{"env3.pk": "id3", "env3.dummy": "dummy", "env3.fk": "id3",},
				},
				"env2": normRecord{
					{"env2.pk": "id1", "env2.dummy2": "dummy1", "env2.fk": "id1",},
					{"env2.pk": "id2", "env2.dummy2": "dummy2", "env2.fk": "id3",},
					{"env2.pk": "id3", "env2.dummy2": "dummy3", "env2.fk": "id2",},
				},
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy3": "dummy4",},
					{"env1.pk": "id2", "env1.dummy3": "dummy5",},
					{"env1.pk": "id3", "env1.dummy3": "dummy6",},
				},
			},
			result: normRecord {
				{"env3.pk": "id1", "env3.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env1.pk": "id1", "env1.dummy3": "dummy4",},
				{"env3.pk": "id2", "env3.dummy": "dummy", "env2.pk": "id1", "env2.dummy2": "dummy1", "env1.pk": "id1", "env1.dummy3": "dummy4",},
				{"env3.pk": "id3", "env3.dummy": "dummy", "env2.pk": "id3", "env2.dummy2": "dummy3", "env1.pk": "id2", "env1.dummy3": "dummy5",},
			},
		},
		{
			name: "makeUnifiedData_noForeignKey",
			config: doubleKeyEnvironment,
			parent: "env1",
			data: normRecords {
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy": "dummy", "env1.fk1": "id4", "env1.fk2": "id4",},
				},
				"env2": normRecord{
					{"env2.pk": "id1", "env2.dummy2": "dummy1",},
					{"env2.pk": "id2", "env2.dummy2": "dummy2",},
					{"env2.pk": "id3", "env2.dummy2": "dummy3",},
				},
				"env3": normRecord{
					{"env3.pk": "id1", "env3.dummy3": "dummy4",},
					{"env3.pk": "id2", "env3.dummy3": "dummy5",},
					{"env3.pk": "id3", "env3.dummy3": "dummy6",},
				},
			},
			result: normRecord {
				{"env1.pk": "id1", "env1.dummy": "dummy", "env1.fk1": "id4", "env1.fk2": "id4",},
			},
		},
		{
			name: "makeUnifiedData_nameCollision",
			config: collisionEnvironment,
			parent: "env.Test",
			data: normRecords {
				"en.vTest": normRecord{
					{"en.vTest.pk": "id1", "en.vTest.dummy": "dummy", "en.vTest.fk": "id1",},
				},
				"env.Test": normRecord{
					{"env.Test.pk": "id1", "env.Test.dummy": "dummy",},
				},
			},
			result: normRecord {
				{"env.Test.pk": "id1", "env.Test.dummy": "dummy", "en.vTest.pk": "id1", "en.vTest.dummy": "dummy",},
			},
		},
		{
			name: "makeUnifiedData_unrelated",
			config: unrelatedEnvironment,
			parent: "env1",
			data: normRecords {
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy": "dummy",},
				},
				"env2": normRecord{
					{"env2.pk": "id2", "env2.dummy": "dummy",},
				},
			},
			result: normRecord {
				{"env1.pk": "id1", "env1.dummy": "dummy",},
			},
		},
		{
			name: "makeUnifiedData_unrelated2",
			config: unrelatedEnvironment,
			parent: "env2",
			data: normRecords {
				"env1": normRecord{
					{"env1.pk": "id1", "env1.dummy": "dummy",},
				},
				"env2": normRecord{
					{"env2.pk": "id2", "env2.dummy": "dummy",},
				},
			},
			result: normRecord {
				{"env2.pk": "id2", "env2.dummy": "dummy",},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			makeUnifiedDataResult(t, test.name, test.data, test.config, test.parent, test.result)
		})
	}
}
