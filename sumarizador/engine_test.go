package main

import (
	"fmt"
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
func Test_copySubsetJData_Exists(t *testing.T) {
	dataMocks := JData{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"}
	
	result := JData{"username": "admin_root", "tentativas": 1.0, "setor": "Diretoria"}
	
	fmt.Println("[INFO] Executando 'Test_copySubsetJData_Exists'")
	ret := copySubsetJData(dataMocks, []string{"username", "tentativas", "setor"})
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_copySubsetJData_Exists2(t *testing.T) {
	dataMocks := JData{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"}
	
	result := JData{"username": "admin_root", "tentativas": 1.0, "setor": "Diretoria"}
	
	fmt.Println("[INFO] Executando 'Test_copySubsetJData_Exists2'")
	ret := copySubsetJData(dataMocks, []string{"tentativas", "setor", "username"})
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_copySubsetJData_NotExists(t *testing.T) {
	dataMocks := JData{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"}
	
	result := JData{"inex":nil}
	
	fmt.Println("[INFO] Executando 'Test_copySubsetJData_NotExists'")
	ret := copySubsetJData(dataMocks, []string{"inex"})
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_copySubsetJData_ExistsNotExists(t *testing.T) {
	dataMocks := JData{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"}
	
	result := JData{"setor": "Diretoria", "inex":nil, "username": "admin_root"}
	
	fmt.Println("[INFO] Executando 'Test_copySubsetJData_ExistsNotExists'")
	ret := copySubsetJData(dataMocks, []string{"setor", "inex", "username"})
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

/*************************************************************************************************************
* função: genSliceHash                                                                                       *
* descrição: Cria um hash de 64 bits usando uma lista de arrays, o hash deve ser diferente mesmo se a mesma  *
* string for quebrada em pontos diferentes. Não é esperado caracteres não imprimíveis nessa função           *
*************************************************************************************************************/
func Test_genSliceHash_equal(t *testing.T) {
	strs := []string{"Lorem ipsum dolor sit amet"}
	
	fmt.Println("[INFO] Executando 'Test_genSliceHash_equal'")
	ret1 := genSliceHash(strs)
	ret2 := genSliceHash(strs)
	if ret1 != ret2 {
		t.Errorf("[ERROR] Erro: Esperado valores iguais, mas %d != %d", ret1, ret2)
	}
}

func Test_genSliceHash_equal2(t *testing.T) {
	strs := []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
	
	fmt.Println("[INFO] Executando 'Test_genSliceHash_equal2'")
	ret1 := genSliceHash(strs)
	ret2 := genSliceHash(strs)
	if ret1 != ret2 {
		t.Errorf("[ERROR] Erro: Esperado valores iguais, mas %d != %d", ret1, ret2)
	}
}

func Test_genSliceHash_notEqual(t *testing.T) {
	strs1 := []string{"Lorem", "ipsum"}
	strs2 := []string{"dolor", "sit", "amet"}
	
	fmt.Println("[INFO] Executando 'Test_genSliceHash_notEqual'")
	ret1 := genSliceHash(strs1)
	ret2 := genSliceHash(strs2)
	if ret1 == ret2 {
		t.Errorf("[ERROR] Erro: Esperado valores diferentes, mas %d == %d", ret1, ret2)
	}
}

func Test_genSliceHash_notEqual2(t *testing.T) {
	strs1 := []string{"Lorem", "ipsum"}
	strs2 := []string{"Loremipsum"}
	
	fmt.Println("[INFO] Executando 'Test_genSliceHash_notEqual2'")
	ret1 := genSliceHash(strs1)
	ret2 := genSliceHash(strs2)
	if ret1 == ret2 {
		t.Errorf("[ERROR] Erro: Esperado valores diferentes, mas %d == %d", ret1, ret2)
	}
}

func Test_genSliceHash_notEqual3(t *testing.T) {
	strs1 := []string{"Lorem", "ipsum"}
	strs2 := []string{"Lor", "emipsum"}
	
	fmt.Println("[INFO] Executando 'Test_genSliceHash_notEqual3'")
	ret1 := genSliceHash(strs1)
	ret2 := genSliceHash(strs2)
	if ret1 == ret2 {
		t.Errorf("[ERROR] Erro: Esperado valores diferentes, mas %d == %d", ret1, ret2)
	}
}

/*************************************************************************************************************
* função: groupBySanitization                                                                                *
* descrição: Retira itens iguais do array de strings                                                         *
*************************************************************************************************************/
func Test_groupBySanitization_happy(t *testing.T) {
	input := []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
	output:= []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
	
	fmt.Println("[INFO] Executando 'Test_groupBySanitization_happy'")
	ret :=groupBySanitization(input)
	if !slices.Equal(ret, output){
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, output)
	}
}

func Test_groupBySanitization_dupHappy(t *testing.T) {
	input := []string{"Lorem", "ipsum", "dolor", "ipsum", "sit", "amet", "Lorem", "amet"}
	output:= []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
	
	fmt.Println("[INFO] Executando 'Test_groupBySanitization_dupHappy'")
	ret :=groupBySanitization(input)
	if !slices.Equal(ret, output) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, output)
	}
}

/*************************************************************************************************************
* função: genKeyString                                                                                       *
* descrição: Retorna array de strings com os valores do par de chaves e valores 'data das chaves             *
* especificadas 'groupBy'. Resultado não pode ter caracteres especiais, deve ser sanitizado                  *
*************************************************************************************************************/

//agregação com zero colunas
func Test_genKeyString_zeroAgreg(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta_destino": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta_destino": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta_destino": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta_destino": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta_destino": nil,   "username": nil},
	}
	
	result := [][]string{
		{},
		{},
		{},
		{},
		{},
		{},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_zeroAgreg'")
	groupBy := groupBySanitization([]string{})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

//agregação com uma coluna
func Test_genKeyString_oneAgreg_Setor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta_destino": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta_destino": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta_destino": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta_destino": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta_destino": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"Diretoria\""},
		{"\"Diretoria\""},
		{"\"Devs\""},
		{"\"Devs\""},
		{"\"Devs\""},
		{"\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_oneAgreg_Setor'")
	groupBy := groupBySanitization([]string{"setor"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_genKeyString_oneAgreg_Porta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"443\""},
		{"\"443\""},
		{"\"80\""},
		{"\"22\""},
		{"\"8080\""},
		{"\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_oneAgreg_Porta'")
	groupBy := groupBySanitization([]string{"porta"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_genKeyString_oneAgreg_Inex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"\""},
		{"\"\""},
		{"\"\""},
		{"\"\""},
		{"\"\""},
		{"\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_oneAgreg_Inex'")
	groupBy := groupBySanitization([]string{"Inex"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_genKeyString_oneAgreg_Special(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Dire\ntoria", "payload_mb": 120.0, "tentativas": 1.0, "porta_destino": 443, "username": "admin_root"},
		{"setor": "Dire\\ntoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta_destino": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta_destino": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta_destino": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta_destino": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta_destino": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"Dire\\ntoria\""},
		{"\"Dire\\\\ntoria\""},
		{"\"Devs\""},
		{"\"Devs\""},
		{"\"Devs\""},
		{"\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_oneAgreg_Special'")
	groupBy := groupBySanitization([]string{"setor"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

//agregação com duas colunas
func Test_genKeyString_twoAgreg_SetorPorta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"Diretoria\"","\"443\""},
		{"\"Diretoria\"","\"443\""},
		{"\"Devs\"","\"80\""},
		{"\"Devs\"","\"22\""},
		{"\"Devs\"","\"8080\""},
		{"\"\"","\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_twoAgreg_SetorPorta'")
	groupBy := groupBySanitization([]string{"setor","porta"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_genKeyString_twoAgreg_PortaSetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"443\"","\"Diretoria\""},
		{"\"443\"","\"Diretoria\""},
		{"\"80\"","\"Devs\""},
		{"\"22\"","\"Devs\""},
		{"\"8080\"","\"Devs\""},
		{"\"\"","\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_twoAgreg_PortaSetor'")
	groupBy := groupBySanitization([]string{"porta","setor"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

func Test_genKeyString_twoAgreg_PortaInex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"443\"","\"\""},
		{"\"443\"","\"\""},
		{"\"80\"","\"\""},
		{"\"22\"","\"\""},
		{"\"8080\"","\"\""},
		{"\"\"","\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_twoAgreg_PortaSetor'")
	groupBy := groupBySanitization([]string{"porta","Inex"})
	
	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}

//agregação com duas colunas duplicada
func Test_genKeyString_threeAgreg_PortaSetorPorta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := [][]string{
		{"\"443\"","\"Diretoria\""},
		{"\"443\"","\"Diretoria\""},
		{"\"80\"","\"Devs\""},
		{"\"22\"","\"Devs\""},
		{"\"8080\"","\"Devs\""},
		{"\"\"","\"\""},
	}
	
	fmt.Println("[INFO] Executando 'Test_genKeyString_threeAgreg_PortaSetorPorta'")
	groupBy := groupBySanitization([]string{"porta","setor", "porta"})

	var ret [][]string
	for _, value := range dataMocks {
		ret = append(ret, genKeyString(value, groupBy))
	}
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %q, mas esperado %q", ret, result)
	}
}
