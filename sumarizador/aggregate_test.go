package sumarizador

import (
	"reflect"
	"testing"
)

//agregações

/*************************************************************************************************************
* função: groupBy                                                                                            *
* descrição: agrupa colunas com valores iguais                                                               *
*************************************************************************************************************/
func groupByResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= groupBy(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_groupBy(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "groupBy_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "groupBy_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "groupBy_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "groupBy_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "groupBy_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{{"setor":"Diretoria"},{"setor":"Devs"},{"setor":nil},},
			
		}, {
			name: "groupBy_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{{"setor":"Diretoria"},{"setor":"Devs"},{"setor":nil},},
			
		}, {
			name: "groupBy_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{{"setor":"Diretoria"},{"setor":"Devs"},{"setor":nil},},
			
		}, {
			name: "groupBy_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"setor":"Diretoria", "username": "admin_root"},
				{"setor":"Diretoria", "username": "alice_silva"},
				{"setor":"Devs", "username": "bob_dev"},
				{"setor":"Devs", "username": "stranger"},
				{"setor":nil, "username": nil},
			},
		}, {
			name: "groupBy_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"setor":"Diretoria", "username": "admin_root"},
				{"setor":"Diretoria", "username": "alice_silva"},
				{"setor":"Devs", "username": "bob_dev"},
				{"setor":"Devs", "username": "stranger"},
				{"setor":nil, "username": nil},
			},
		}, {
			name: "groupBy_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"setor":"Diretoria", "username": "admin_root"},
				{"setor":"Diretoria", "username": "alice_silva"},
				{"setor":"Devs", "username": "bob_dev"},
				{"setor":"Devs", "username": "stranger"},
				{"setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupByResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: count                                                                                              *
* descrição: conta ocorrencias dos valores definidos em groupBy                                              *
*************************************************************************************************************/
func countResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= count(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_count(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "count_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "count_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "count_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "count_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "count_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"count_": 2.0, "setor":"Diretoria"},
				{"count_": 3.0, "setor":"Devs"},
				{"count_": 1.0, "setor":nil},
			},
			
		}, {
			name: "count_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"count_": 2.0, "setor":"Diretoria"},
				{"count_": 3.0, "setor":"Devs"},
				{"count_": 1.0, "setor":nil},
			},
			
		}, {
			name: "count_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"count_": 2.0, "setor":"Diretoria"},
				{"count_": 3.0, "setor":"Devs"},
				{"count_": 1.0, "setor":nil},
			},
			
		}, {
			name: "count_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result:[]JData{
				{"count_": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_": 1.0, "setor":nil, "username": nil},
			},
		}, {
			name: "count_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result:[]JData{
				{"count_": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_": 1.0, "setor":nil, "username": nil},
			},
		}, {
			name: "count_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"count_": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_": 1.0, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			countResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: countNotNil                                                                                        *
* descrição: conta ocorrencias dos valores definidos em groupBy desde que a coluna alvo não seja nula ou     *
* inexistente                                                                                                *
*************************************************************************************************************/
func countNotNilResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= countNotNil(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_countNotNil(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "countNotNil_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "countNotNil_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "countNotNil_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "countNotNil_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "countNotNil_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"count_not_null_": 0.0, "setor":"Diretoria"},
				{"count_not_null_": 0.0, "setor":"Devs"},
				{"count_not_null_": 0.0, "setor":nil},
			},
			
		}, {
			name: "countNotNil_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"count_not_null_username": 2.0, "setor":"Diretoria"},
				{"count_not_null_username": 3.0, "setor":"Devs"},
				{"count_not_null_username": 0.0, "setor":nil},
			},
			
		}, {
			name: "countNotNil_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"count_not_null_setor": 2.0, "setor":"Diretoria"},
				{"count_not_null_setor": 3.0, "setor":"Devs"},
				{"count_not_null_setor": 0.0, "setor":nil},
			},
			
		}, {
			name: "countNotNil_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"count_not_null_": 0.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_not_null_": 0.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_not_null_": 0.0, "setor":"Devs", "username": "bob_dev"},
				{"count_not_null_": 0.0, "setor":"Devs", "username": "stranger"},
				{"count_not_null_": 0.0, "setor":nil, "username": nil},
			},
		}, {
			name: "countNotNil_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"count_not_null_username": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_not_null_username": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_not_null_username": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_not_null_username": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_not_null_username": 0.0, "setor":nil, "username": nil},
			},
		}, {
			name: "countNotNil_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"count_not_null_porta": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_not_null_porta": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_not_null_porta": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_not_null_porta": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_not_null_porta": 0.0, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			countNotNilResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: countDistinct                                                                                      *
* descrição: conta ocorrencias dos valores definidos em groupBy desde que o valor alvo seja único            *
*************************************************************************************************************/
func countDistinctResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= countDistinct(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_countDistinct(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "countDistinct_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "countDistinct_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "countDistinct_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "countDistinct_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "countDistinct_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"count_distinct_": 1.0, "setor":"Diretoria"},
				{"count_distinct_": 1.0, "setor":"Devs"},
				{"count_distinct_": 1.0, "setor":nil},
			},
			
		}, {
			name: "countDistinct_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"count_distinct_username": 2.0, "setor":"Diretoria"},
				{"count_distinct_username": 2.0, "setor":"Devs"},
				{"count_distinct_username": 1.0, "setor":nil},
			},
			
		}, {
			name: "countDistinct_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"count_distinct_setor": 1.0, "setor":"Diretoria"},
				{"count_distinct_setor": 1.0, "setor":"Devs"},
				{"count_distinct_setor": 1.0, "setor":nil},
			},
			
		}, {
			name: "countDistinct_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"count_distinct_": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_distinct_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_distinct_": 1.0, "setor":"Devs", "username": "bob_dev"},
				{"count_distinct_": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_distinct_": 1.0, "setor":nil, "username": nil},
			},
		}, {
			name: "countDistinct_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"count_distinct_username": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_distinct_username": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_distinct_username": 1.0, "setor":"Devs", "username": "bob_dev"},
				{"count_distinct_username": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_distinct_username": 1.0, "setor":nil, "username": nil},
			},
		}, {
			name: "countDistinct_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"count_distinct_porta": 1.0, "setor":"Diretoria", "username": "admin_root"},
				{"count_distinct_porta": 1.0, "setor":"Diretoria", "username": "alice_silva"},
				{"count_distinct_porta": 2.0, "setor":"Devs", "username": "bob_dev"},
				{"count_distinct_porta": 1.0, "setor":"Devs", "username": "stranger"},
				{"count_distinct_porta": 1.0, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			countDistinctResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: sum                                                                                                *
* descrição: Soma uma coluna númerica qualquer agrupada pelos vaores em groupBy                              *
*************************************************************************************************************/
func sumResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= sum(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_sum(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "sum_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "sum_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "sum_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "sum_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "sum_nilAgreg_porta",
			input: dataMock,
			keys: nil,
			target: "porta",
			result: []JData{
				{"sum_porta": 9068.0},
			},
			
		}, {
			name: "sum_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"sum_": nil, "setor":"Diretoria"},
				{"sum_": nil, "setor":"Devs"},
				{"sum_": nil, "setor":nil},
			},
			
		}, {
			name: "sum_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"sum_username": nil, "setor":"Diretoria"},
				{"sum_username": nil, "setor":"Devs"},
				{"sum_username": nil, "setor":nil},
			},
			
		}, {
			name: "sum_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"sum_setor": nil, "setor":"Diretoria"},
				{"sum_setor": nil, "setor":"Devs"},
				{"sum_setor": nil, "setor":nil},
			},
			
		}, {
			name: "sum_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"sum_porta": 886.0, "setor":"Diretoria"},
				{"sum_porta": 8182.0, "setor":"Devs"},
				{"sum_porta": nil, "setor":nil},
			},
			
		}, {
			name: "sum_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"sum_": nil, "setor":"Diretoria", "username": "admin_root"},
				{"sum_": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"sum_": nil, "setor":"Devs", "username": "bob_dev"},
				{"sum_": nil, "setor":"Devs", "username": "stranger"},
				{"sum_": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "sum_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"sum_username": nil, "setor":"Diretoria", "username": "admin_root"},
				{"sum_username": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"sum_username": nil, "setor":"Devs", "username": "bob_dev"},
				{"sum_username": nil, "setor":"Devs", "username": "stranger"},
				{"sum_username": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "sum_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"sum_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
				{"sum_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
				{"sum_porta": 102.0, "setor":"Devs", "username": "bob_dev"},
				{"sum_porta": 8080.0, "setor":"Devs", "username": "stranger"},
				{"sum_porta": nil, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sumResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: avg                                                                                                *
* descrição: Tira a média de uma coluna númerica qualquer agrupada pelos vaores em groupBy                   *
*************************************************************************************************************/
func avgResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= avg(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_avg(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "avg_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "avg_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "avg_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "avg_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		},  {
			name: "avg_nilAgreg_porta",
			input: dataMock,
			keys: nil,
			target: "porta",
			result: []JData{
				{"avg_porta": 1813.6},
			},
			
		}, {
			name: "avg_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"avg_": nil, "setor":"Diretoria"},
				{"avg_": nil, "setor":"Devs"},
				{"avg_": nil, "setor":nil},
			},
			
		}, {
			name: "avg_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"avg_username": nil, "setor":"Diretoria"},
				{"avg_username": nil, "setor":"Devs"},
				{"avg_username": nil, "setor":nil},
			},
			
		}, {
			name: "avg_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"avg_setor": nil, "setor":"Diretoria"},
				{"avg_setor": nil, "setor":"Devs"},
				{"avg_setor": nil, "setor":nil},
			},
			
		},  {
			name: "avg_setorAgreg_porta",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"avg_porta": 443.0, "setor":"Diretoria"},
				{"avg_porta": 2727.333333333333333333333333333333333, "setor":"Devs"},
				{"avg_porta": nil, "setor":nil},
			},
			
		}, {
			name: "avg_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"avg_": nil, "setor":"Diretoria", "username": "admin_root"},
				{"avg_": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"avg_": nil, "setor":"Devs", "username": "bob_dev"},
				{"avg_": nil, "setor":"Devs", "username": "stranger"},
				{"avg_": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "avg_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"avg_username": nil, "setor":"Diretoria", "username": "admin_root"},
				{"avg_username": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"avg_username": nil, "setor":"Devs", "username": "bob_dev"},
				{"avg_username": nil, "setor":"Devs", "username": "stranger"},
				{"avg_username": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "avg_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"avg_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
				{"avg_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
				{"avg_porta": 51.0, "setor":"Devs", "username": "bob_dev"},
				{"avg_porta": 8080.0, "setor":"Devs", "username": "stranger"},
				{"avg_porta": nil, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avgResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: min                                                                                                *
* descrição: Retorna o valor mínimo de uma coluna númerica qualquer agrupada pelos vaores em groupBy         *
*************************************************************************************************************/
func minResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= min(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_min(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "min_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "min_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "min_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "min_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		},  {
			name: "min_nilAgreg_porta",
			input: dataMock,
			keys: nil,
			target: "porta",
			result: []JData{
				{"min_porta": 22.0},
			},
			
		}, {
			name: "min_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"min_": nil, "setor":"Diretoria"},
				{"min_": nil, "setor":"Devs"},
				{"min_": nil, "setor":nil},
			},
			
		}, {
			name: "min_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"min_username": nil, "setor":"Diretoria"},
				{"min_username": nil, "setor":"Devs"},
				{"min_username": nil, "setor":nil},
			},
			
		}, {
			name: "min_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"min_setor": nil, "setor":"Diretoria"},
				{"min_setor": nil, "setor":"Devs"},
				{"min_setor": nil, "setor":nil},
			},
			
		},  {
			name: "min_setorAgreg_porta",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"min_porta": 443.0, "setor":"Diretoria"},
				{"min_porta": 22.0, "setor":"Devs"},
				{"min_porta": nil, "setor":nil},
			},
			
		}, {
			name: "min_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"min_": nil, "setor":"Diretoria", "username": "admin_root"},
				{"min_": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"min_": nil, "setor":"Devs", "username": "bob_dev"},
				{"min_": nil, "setor":"Devs", "username": "stranger"},
				{"min_": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "min_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"min_username": nil, "setor":"Diretoria", "username": "admin_root"},
				{"min_username": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"min_username": nil, "setor":"Devs", "username": "bob_dev"},
				{"min_username": nil, "setor":"Devs", "username": "stranger"},
				{"min_username": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "min_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"min_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
				{"min_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
				{"min_porta": 22.0, "setor":"Devs", "username": "bob_dev"},
				{"min_porta": 8080.0, "setor":"Devs", "username": "stranger"},
				{"min_porta": nil, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			minResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: max                                                                                                *
* descrição: Retorna o valor máximo de uma coluna númerica qualquer agrupada pelos vaores em groupBy         *
*************************************************************************************************************/
func maxResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= max(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_max(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "max_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "max_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "max_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "max_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		},  {
			name: "max_nilAgreg_porta",
			input: dataMock,
			keys: nil,
			target: "porta",
			result: []JData{
				{"max_porta": 8080.0},
			},
			
		}, {
			name: "max_setorAgreg_zero",
			input: dataMock,
			keys: []string{"setor"},
			target: "",
			result: []JData{
				{"max_": nil, "setor":"Diretoria"},
				{"max_": nil, "setor":"Devs"},
				{"max_": nil, "setor":nil},
			},
			
		}, {
			name: "max_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"max_username": nil, "setor":"Diretoria"},
				{"max_username": nil, "setor":"Devs"},
				{"max_username": nil, "setor":nil},
			},
			
		}, {
			name: "max_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"max_setor": nil, "setor":"Diretoria"},
				{"max_setor": nil, "setor":"Devs"},
				{"max_setor": nil, "setor":nil},
			},
			
		},  {
			name: "max_setorAgreg_porta",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"max_porta": 443.0, "setor":"Diretoria"},
				{"max_porta": 8080.0, "setor":"Devs"},
				{"max_porta": nil, "setor":nil},
			},
			
		}, {
			name: "max_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"max_": nil, "setor":"Diretoria", "username": "admin_root"},
				{"max_": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"max_": nil, "setor":"Devs", "username": "bob_dev"},
				{"max_": nil, "setor":"Devs", "username": "stranger"},
				{"max_": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "max_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"max_username": nil, "setor":"Diretoria", "username": "admin_root"},
				{"max_username": nil, "setor":"Diretoria", "username": "alice_silva"},
				{"max_username": nil, "setor":"Devs", "username": "bob_dev"},
				{"max_username": nil, "setor":"Devs", "username": "stranger"},
				{"max_username": nil, "setor":nil, "username": nil},
			},
		}, {
			name: "max_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"max_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
				{"max_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
				{"max_porta": 80.0, "setor":"Devs", "username": "bob_dev"},
				{"max_porta": 8080.0, "setor":"Devs", "username": "stranger"},
				{"max_porta": nil, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			maxResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: arrPack                                                                                            *
* descrição: Inclui todos os valores em uma coluna array                                                     *
*************************************************************************************************************/
func arrPackResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= arrPack(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_arrPack(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "arrPack_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "arrPack_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "arrPack_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "arrPack_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "arrPack_setorAgreg_porta",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"arr_porta": []any{443.0, 443.0}, "setor":"Diretoria"},
				{"arr_porta": []any{80.0, 22.0, 8080.0}, "setor":"Devs"},
				{"arr_porta": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrPack_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"arr_username": []any{"admin_root","alice_silva"}, "setor":"Diretoria"},
				{"arr_username": []any{"bob_dev","bob_dev","stranger"}, "setor":"Devs"},
				{"arr_username": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrPack_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"arr_setor": []any{"Diretoria","Diretoria"}, "setor":"Diretoria"},
				{"arr_setor": []any{"Devs","Devs","Devs"}, "setor":"Devs"},
				{"arr_setor": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrPack_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"arr_": []any{nil}, "setor":"Diretoria", "username": "admin_root"},
				{"arr_": []any{nil}, "setor":"Diretoria", "username": "alice_silva"},
				{"arr_": []any{nil, nil}, "setor":"Devs", "username": "bob_dev"},
				{"arr_": []any{nil}, "setor":"Devs", "username": "stranger"},
				{"arr_": []any{nil}, "setor":nil, "username": nil},
			},
		}, {
			name: "arrPack_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"arr_username": []any{"admin_root"}, "setor":"Diretoria", "username": "admin_root"},
				{"arr_username": []any{"alice_silva"}, "setor":"Diretoria", "username": "alice_silva"},
				{"arr_username": []any{"bob_dev", "bob_dev"}, "setor":"Devs", "username": "bob_dev"},
				{"arr_username": []any{"stranger"}, "setor":"Devs", "username": "stranger"},
				{"arr_username": []any{nil}, "setor":nil, "username": nil},
			},
		}, {
			name: "arrPack_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"arr_porta": []any{443.0}, "setor":"Diretoria", "username": "admin_root"},
				{"arr_porta": []any{443.0}, "setor":"Diretoria", "username": "alice_silva"},
				{"arr_porta": []any{80.0, 22.0}, "setor":"Devs", "username": "bob_dev"},
				{"arr_porta": []any{8080.0}, "setor":"Devs", "username": "stranger"},
				{"arr_porta": []any{nil}, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arrPackResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}

/*************************************************************************************************************
* função: arrDistinctPack                                                                                    *
* descrição: Inclui todos os valores em uma coluna array                                                     *
*************************************************************************************************************/
func arrDistinctPackResult(t *testing.T, name string, dataMock []JData, keys []string, target string, result []JData) {
	t.Logf("[INFO] Executando %q", name)
	ret:= arrDistinctPack(dataMock, keys, target)
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
func Test_arrDistinctPack(t *testing.T) {
	dataMock := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443.0, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443.0, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22.0,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080.0,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		input []JData
		keys []string
		target string
		result []JData
	} {
		{
			name: "arrDistinctPack_zeroAgreg_zero",
			input: []JData{},
			keys: []string{},
			target: "",
			result: nil,
			
		}, {
			name: "arrDistinctPack_nilAgreg_zero",
			input: []JData{},
			keys: nil,
			target: "",
			result: nil,
			
		}, {
			name: "arrDistinctPack_zeroAgreg_tentativas",
			input: []JData{},
			keys: []string{},
			target: "tentativas",
			result: nil,
			
		}, {
			name: "arrDistinctPack_nilAgreg_tentativas",
			input: []JData{},
			keys: nil,
			target: "tentativas",
			result: nil,
			
		}, {
			name: "arrDistinctPack_setorAgreg_porta",
			input: dataMock,
			keys: []string{"setor"},
			target: "porta",
			result: []JData{
				{"set_porta": []any{443.0}, "setor":"Diretoria"},
				{"set_porta": []any{80.0, 22.0, 8080.0}, "setor":"Devs"},
				{"set_porta": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrDistinctPack_setorAgreg_username",
			input: dataMock,
			keys: []string{"setor"},
			target: "username",
			result: []JData{
				{"set_username": []any{"admin_root","alice_silva"}, "setor":"Diretoria"},
				{"set_username": []any{"bob_dev","stranger"}, "setor":"Devs"},
				{"set_username": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrDistinctPack_setorAgreg_setor",
			input: dataMock,
			keys: []string{"setor"},
			target: "setor",
			result: []JData{
				{"set_setor": []any{"Diretoria"}, "setor":"Diretoria"},
				{"set_setor": []any{"Devs"}, "setor":"Devs"},
				{"set_setor": []any{nil}, "setor":nil},
			},
			
		}, {
			name: "arrDistinctPack_setorUsernameAgreg_zero",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "",
			result: []JData{
				{"set_": []any{nil}, "setor":"Diretoria", "username": "admin_root"},
				{"set_": []any{nil}, "setor":"Diretoria", "username": "alice_silva"},
				{"set_": []any{nil}, "setor":"Devs", "username": "bob_dev"},
				{"set_": []any{nil}, "setor":"Devs", "username": "stranger"},
				{"set_": []any{nil}, "setor":nil, "username": nil},
			},
		}, {
			name: "arrDistinctPack_setorUsernameAgreg_username",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "username",
			result: []JData{
				{"set_username": []any{"admin_root"}, "setor":"Diretoria", "username": "admin_root"},
				{"set_username": []any{"alice_silva"}, "setor":"Diretoria", "username": "alice_silva"},
				{"set_username": []any{"bob_dev"}, "setor":"Devs", "username": "bob_dev"},
				{"set_username": []any{"stranger"}, "setor":"Devs", "username": "stranger"},
				{"set_username": []any{nil}, "setor":nil, "username": nil},
			},
		}, {
			name: "arrDistinctPack_setorUsernameAgreg_porta",
			input: dataMock,
			keys: []string{"setor", "username"},
			target: "porta",
			result: []JData{
				{"set_porta": []any{443.0}, "setor":"Diretoria", "username": "admin_root"},
				{"set_porta": []any{443.0}, "setor":"Diretoria", "username": "alice_silva"},
				{"set_porta": []any{80.0, 22.0}, "setor":"Devs", "username": "bob_dev"},
				{"set_porta": []any{8080.0}, "setor":"Devs", "username": "stranger"},
				{"set_porta": []any{nil}, "setor":nil, "username": nil},
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arrDistinctPackResult(t, test.name, test.input, test.keys, test.target, test.result)
		})
	}
}
