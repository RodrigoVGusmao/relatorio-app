package main

import (
	"reflect"
	"slices"
	"testing"
)

//utilitários

/*************************************************************************************************************
* função: copySubsetJData                                                                                    *
* descrição: Cópia os pares de chaves e valores do argumento 'src'. Se chave não existir, criar com valor    *
* padrão 'nil'                                                                                               *
*************************************************************************************************************/
func copySubsetJDataResult(t *testing.T, name string, input JData, keys []string, expected JData) {
	t.Logf("[INFO] Executando %q", name)
	ret := copySubsetJData(input, keys)
	if !reflect.DeepEqual(expected, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, expected)
	}
}
func Test_copySubsetJData(t *testing.T) {
	dataMocks := JData{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"}
	
	tests := []struct {
		name string
		result JData
		keys []string
	} {
		{
			name: "copySubsetJData_empty",
			result: JData{},
			keys: []string{},
		}, {
			name: "copySubsetJData_Happy",
			result: JData{"username": "admin_root", "tentativas": 1.0, "setor": "Diretoria"},
			keys: []string{"username", "tentativas", "setor"},
		}, {
			name: "copySubsetJData_Happy2",
			result: JData{"username": "admin_root", "tentativas": 1.0, "setor": "Diretoria"},
			keys: []string{"tentativas", "setor", "username"},
		}, {
			name: "copySubsetJData_Inex",
			result: JData{"inex":nil},
			keys: []string{"inex"},
		}, {
			name: "copySubsetJData_HappyInex",
			result: JData{"setor": "Diretoria", "inex":nil, "username": "admin_root"},
			keys: []string{"setor", "inex", "username"},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			copySubsetJDataResult(t, test.name, dataMocks, test.keys, test.result)
		})
	}
}

/*************************************************************************************************************
* função: genSliceHash                                                                                       *
* descrição: Cria um hash de 64 bits usando uma lista de arrays, o hash deve ser diferente mesmo se a mesma  *
* string for quebrada em pontos diferentes. Não é esperado caracteres não imprimíveis nessa função           *
*************************************************************************************************************/
func genSliceHashResult(t *testing.T, name string, input1 []string, input2 []string, equal bool) {
	t.Logf("[INFO] Executando %q", name)
	ret1 := genSliceHash(input1)
	ret2 := genSliceHash(input2)
	if (ret1 == ret2) != equal {
		t.Errorf("[ERROR] Erro: Esperado valores iguais, mas %d != %d", ret1, ret2)
	}
}
func Test_genSliceHash(t *testing.T) {
	tests := []struct {
		name string
		input1 []string
		input2 []string
		equal bool
	} {
		{
			name: "genSliceHash_empty",
			input1: []string{},
			input2: []string{},
			equal: true,
		}, {
			name: "genSliceHash_empty2",
			input1: []string{},
			input2: []string{"Lorem ipsum dolor sit amet"},
			equal: false,
		}, {
			name: "genSliceHash_equal",
			input1: []string{"Lorem ipsum dolor sit amet"},
			input2: []string{"Lorem ipsum dolor sit amet"},
			equal: true,
		}, {
			name: "genSliceHash_equal2",
			input1: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
			input2: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
			equal: true,
		}, {
			name: "genSliceHash_inequal",
			input1: []string{"Lorem", "ipsum"},
			input2: []string{"dolor", "sit", "amet"},
			equal: false,
		}, {
			name: "genSliceHash_inequal2",
			input1: []string{"Lorem", "ipsum"},
			input2: []string{"Loremipsum"},
			equal: false,
		}, {
			name: "genSliceHash_inequal3",
			input1: []string{"Lorem", "ipsum"},
			input2: []string{"Lor", "emipsum"},
			equal: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			genSliceHashResult(t, test.name, test.input1, test.input2, test.equal)
		})
	}
}

/*************************************************************************************************************
* função: groupBySanitization                                                                                *
* descrição: Retira itens iguais do array de strings                                                         *
*************************************************************************************************************/
func groupBySanitizationCopy(t *testing.T, name string, input []string, output []string) {
	t.Logf("[INFO] Executando %q", name)
	ret :=groupBySanitization(input)
	if !slices.Equal(ret, output){
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, output)
	}
}
func Test_groupBySanitization(t *testing.T) {
	tests := []struct {
		name string
		input []string
		output []string
	} {
		{
			name: "groupBySanitization_happy",
			input: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
			output: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
		}, {
			name: "groupBySanitization_dupHappy",
			input: []string{"Lorem", "ipsum", "dolor", "ipsum", "sit", "amet", "Lorem", "amet"},
			output: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
		}, {
			name: "groupBySanitization_junk",
			input: []string{"Lorem", "ipsum", "dolor\x11", "sit", "amet"},
			output: []string{"Lorem", "ipsum", "dolor\x11", "sit", "amet"},
		}, {
			name: "groupBySanitization_junk2",
			input: []string{"Lorem\x11", "ipsum", "dolor", "ipsum", "sit", "amet", "Lorem\x12", "amet"},
			output: []string{"Lorem\x11", "ipsum", "dolor", "sit", "amet", "Lorem\x12"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			groupBySanitizationCopy(t, test.name, test.input, test.output)
		})
	}
}

/*************************************************************************************************************
* função: genKeyString                                                                                       *
* descrição: Retorna array de strings com os valores do par de chaves e valores 'data das chaves             *
* especificadas 'groupBy'. Resultado não pode ter caracteres especiais, deve ser sanitizado                  *
*************************************************************************************************************/
func genKeyStringCopy(t *testing.T, name string, dataMocks []JData, keys []string, result [][]string) {
	t.Logf("[INFO] Executando %q", name)
	groupBy := groupBySanitization(keys)
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}
func Test_genKeyString(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob\n_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob\\n_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	tests := []struct {
		name string
		keys []string
		output [][]string
	} {
		{
			name: "genKeyString_zeroAgreg",
			keys: []string{},
			output: [][]string{{},{},{},{},{},{},},
		}, {
			name: "genKeyString_oneAgreg_Setor",
			keys: []string{"setor"},
			output: [][]string{{"\"Diretoria\""},{"\"Diretoria\""},{"\"Devs\""},{"\"Devs\""},{"\"Devs\""},{"\"\""},},
		}, {
			name: "genKeyString_oneAgreg_Porta",
			keys: []string{"porta"},
			output: [][]string{{"\"443\""},{"\"443\""},{"\"80\""},{"\"22\""},{"\"8080\""},{"\"\""},},
		}, {
			name: "genKeyString_oneAgreg_Inex",
			keys: []string{"Inex"},
			output: [][]string{{"\"\""},{"\"\""},{"\"\""},{"\"\""},{"\"\""},{"\"\""},},
		}, {
			name: "genKeyString_oneAgreg_Special",
			keys: []string{"username"},
			output: [][]string{{"\"admin_root\""},{"\"alice_silva\""},{"\"bob\\n_dev\""},{"\"bob\\\\n_dev\""},{"\"stranger\""},{"\"\""},},
		}, {
			name: "genKeyString_twoAgreg_SetorPorta",
			keys: []string{"setor","porta"},
			output: [][]string{{"\"Diretoria\"","\"443\""},{"\"Diretoria\"","\"443\""},{"\"Devs\"","\"80\""},{"\"Devs\"","\"22\""},{"\"Devs\"","\"8080\""},{"\"\"","\"\""},},
		}, {
			name: "genKeyString_twoAgreg_PortaSetor",
			keys: []string{"porta","setor"},
			output: [][]string{{"\"443\"","\"Diretoria\""},{"\"443\"","\"Diretoria\""},{"\"80\"","\"Devs\""},{"\"22\"","\"Devs\""},{"\"8080\"","\"Devs\""},{"\"\"","\"\""},},
		}, {
			name: "genKeyString_twoAgreg_PortaInex",
			keys: []string{"porta","Inex"},
			output: [][]string{{"\"443\"","\"\""},{"\"443\"","\"\""},{"\"80\"","\"\""},{"\"22\"","\"\""},{"\"8080\"","\"\""},{"\"\"","\"\""},},
		}, {
			name: "genKeyString_threeAgreg_PortaSetorPorta",
			keys: []string{"porta","setor", "porta"},
			output: [][]string{{"\"443\"","\"Diretoria\""},{"\"443\"","\"Diretoria\""},{"\"80\"","\"Devs\""},{"\"22\"","\"Devs\""},{"\"8080\"","\"Devs\""},{"\"\"","\"\""},},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			genKeyStringCopy(t, test.name, dataMocks, test.keys, test.output)
		})
	}
}
