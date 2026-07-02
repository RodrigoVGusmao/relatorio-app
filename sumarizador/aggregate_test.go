package main

import (
	"fmt"
	"reflect"
	"testing"
)

//agregações

/*************************************************************************************************************
* função: groupBy                                                                                            *
* descrição: agrupa colunas com valores iguais                                                               *
*************************************************************************************************************/

func Test_groupBy_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_groupBy_zeroAgreg'")
	ret:= groupBy(dataMocks, []string{}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_groupBy_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_groupBy_zeroAgreg_tentativas'")
	ret:= groupBy(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_groupBy_oneAgreg_nil_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"setor":"Diretoria"},
		{"setor":"Devs"},
		{"setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_groupBy_oneAgreg_nil_bySetor'")
	ret:= groupBy(dataMocks, []string{"setor"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_groupBy_oneAgreg_nil_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"setor":"Diretoria", "username": "admin_root"},
		{"setor":"Diretoria", "username": "alice_silva"},
		{"setor":"Devs", "username": "bob_dev"},
		{"setor":"Devs", "username": "stranger"},
		{"setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_groupBy_oneAgreg_nil_bySetorUsername'")
	ret:= groupBy(dataMocks, []string{"setor", "username"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: count                                                                                              *
* descrição: conta ocorrencias dos valores definidos em groupBy                                              *
*************************************************************************************************************/
func Test_count_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_count_zeroAgreg'")
	ret:= count(dataMocks, []string{}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_count_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_count_zeroAgreg_tentativas'")
	ret:= count(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_count_oneAgreg_nil_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_": 2.0, "setor":"Diretoria"},
		{"count_": 3.0, "setor":"Devs"},
		{"count_": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_count_oneAgreg_nil_bySetor'")
	ret:= count(dataMocks, []string{"setor"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_count_oneAgreg_nil_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_": 1.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_count_oneAgreg_nil_bySetorUsername'")
	ret:= count(dataMocks, []string{"setor", "username"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_count_oneAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_": 1.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_count_oneAgreg_porta_bySetorUsername'")
	ret:= count(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_count2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_": 2.0, "setor":"Diretoria"},
		{"count_": 2.0, "setor":"Devs"},
		{"count_": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_count2_zeroAgreg_porta_bySetorUsername'")
	parc := count(dataMocks, []string{"setor", "username"}, "")
	ret := count(parc, []string{"setor"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_": 2.0, "setor":"Diretoria"},
		{"count_": 2.0, "setor":"Devs"},
		{"count_": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countsum_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := count(parc, []string{"setor"}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: countNotNil                                                                                        *
* descrição: conta ocorrencias dos valores definidos em groupBy desde que a coluna alvo não seja nula ou     *
* inexistente                                                                                                *
*************************************************************************************************************/
func Test_countNotNil_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_countNotNil_zeroAgreg'")
	ret:= countNotNil(dataMocks, []string{}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNil_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_countNotNil_zeroAgreg_tentativas'")
	ret:= countNotNil(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNil_oneAgreg_nil_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_not_null_porta": 2.0, "setor":"Diretoria"},
		{"count_not_null_porta": 3.0, "setor":"Devs"},
		{"count_not_null_porta": 0.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countNotNil_oneAgreg_nil_bySetor'")
	ret:= countNotNil(dataMocks, []string{"setor"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNil_oneAgreg_nil_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_not_null_tentativas": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_not_null_tentativas": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_not_null_tentativas": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_not_null_tentativas": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_not_null_tentativas": 1.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countNotNil_oneAgreg_nil_bySetorUsername'")
	ret:= countNotNil(dataMocks, []string{"setor", "username"}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNil_oneAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_not_null_porta": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_not_null_porta": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_not_null_porta": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_not_null_porta": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_not_null_porta": 0.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countNotNil_oneAgreg_porta_bySetorUsername'")
	ret:= countNotNil(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNil2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_not_null_count_not_null_porta": 2.0, "setor":"Diretoria"},
		{"count_not_null_count_not_null_porta": 2.0, "setor":"Devs"},
		{"count_not_null_count_not_null_porta": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countNotNil2_zeroAgreg_porta_bySetorUsername'")
	parc := countNotNil(dataMocks, []string{"setor", "username"}, "porta")
	ret := countNotNil(parc, []string{"setor"}, "count_not_null_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countNotNilsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_not_null_sum_tentativas": 2.0, "setor":"Diretoria"},
		{"count_not_null_sum_tentativas": 2.0, "setor":"Devs"},
		{"count_not_null_sum_tentativas": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countNotNilsum_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "tentativas")
	ret := countNotNil(parc, []string{"setor"}, "sum_tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: countDistinct                                                                                      *
* descrição: conta ocorrencias dos valores definidos em groupBy desde que o valor alvo seja único            *
*************************************************************************************************************/
func Test_countDistinct_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_countDistinct_zeroAgreg'")
	ret:= countDistinct(dataMocks, []string{}, "")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinct_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_countDistinct_zeroAgreg_tentativas'")
	ret:= countDistinct(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinct_oneAgreg_nil_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_distinct_porta": 1.0, "setor":"Diretoria"},
		{"count_distinct_porta": 3.0, "setor":"Devs"},
		{"count_distinct_porta": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countDistinct_oneAgreg_nil_bySetor'")
	ret:= countDistinct(dataMocks, []string{"setor"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinct_oneAgreg_nil_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_distinct_payload_mb": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_distinct_payload_mb": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_distinct_payload_mb": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_distinct_payload_mb": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_distinct_payload_mb": 1.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countDistinct_oneAgreg_nil_bySetorUsername'")
	ret:= countDistinct(dataMocks, []string{"setor", "username"}, "payload_mb")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinct_oneAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_distinct_porta": 1.0, "setor":"Diretoria", "username": "admin_root"},
		{"count_distinct_porta": 1.0, "setor":"Diretoria", "username": "alice_silva"},
		{"count_distinct_porta": 2.0, "setor":"Devs", "username": "bob_dev"},
		{"count_distinct_porta": 1.0, "setor":"Devs", "username": "stranger"},
		{"count_distinct_porta": 1.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countDistinct_oneAgreg_porta_bySetorUsername'")
	ret:= countDistinct(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinct2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_distinct_count_distinct_porta": 1.0, "setor":"Diretoria"},
		{"count_distinct_count_distinct_porta": 2.0, "setor":"Devs"},
		{"count_distinct_count_distinct_porta": 1.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countDistinct2_zeroAgreg_porta_bySetorUsername'")
	parc := countDistinct(dataMocks, []string{"setor", "username"}, "porta")
	ret := countDistinct(parc, []string{"setor"}, "count_distinct_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_countDistinctsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"count_distinct_setor": 1.0, "username":"admin_root"},
		{"count_distinct_setor": 1.0, "username":"alice_silva"},
		{"count_distinct_setor": 1.0, "username":"bob_dev"},
		{"count_distinct_setor": 1.0, "username":"stranger"},
		{"count_distinct_setor": 1.0, "username":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_countDistinctsum_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := countDistinct(parc, []string{"username"}, "setor")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: sum                                                                                                *
* descrição: Soma uma coluna númerica qualquer agrupada pelos vaores em groupBy                              *
*************************************************************************************************************/
func Test_sum_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_tentativas'")
	ret:= sum(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_tentativas": 24.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_tentativas'")
	ret:= sum(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_zeroAgreg_username(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_username": 0.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_username'")
	ret:= sum(dataMocks, []string{}, "username")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_zeroAgreg_porta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_porta": 9068.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_porta'")
	ret:= sum(dataMocks, []string{}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_oneAgreg_tentativas_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_tentativas": 6.0, "setor":"Diretoria"},
		{"sum_tentativas": 15.0, "setor":"Devs"},
		{"sum_tentativas": 3.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_tentativas_bySetor'")
	ret:= sum(dataMocks, []string{"setor"}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
		{"sum_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
		{"sum_porta": 102.0, "setor":"Devs", "username": "bob_dev"},
		{"sum_porta": 8080.0, "setor":"Devs", "username": "stranger"},
		{"sum_porta": 0.0, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_porta_bySetorUsername'")
	ret:= sum(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum_oneAgreg_porta_byInex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"inex": nil, "sum_payload_mb":570.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum_zeroAgreg_porta_byInex'")
	ret:= sum(dataMocks, []string{"inex"}, "payload_mb")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_sum2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"sum_sum_porta": 886.0, "setor":"Diretoria"},
		{"sum_sum_porta": 8182.0, "setor":"Devs"},
		{"sum_sum_porta": 0.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_sum2_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := sum(parc, []string{"setor"}, "sum_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: avg                                                                                                *
* descrição: Tira a média de uma coluna númerica qualquer agrupada pelos vaores em groupBy                   *
*************************************************************************************************************/
func Test_avg_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_avg_zeroAgreg_tentativas'")
	ret:= avg(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_tentativas": 4.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_zeroAgreg_tentativas'")
	ret:= avg(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_zeroAgreg_username(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_zeroAgreg_username'")
	ret:= avg(dataMocks, []string{}, "username")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_zeroAgreg_porta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_porta": 1813.6},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_zeroAgreg_porta'")
	ret:= avg(dataMocks, []string{}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_oneAgreg_tentativas_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_tentativas": 3.0, "setor":"Diretoria"},
		{"avg_tentativas": 5.0, "setor":"Devs"},
		{"avg_tentativas": 3.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_zeroAgreg_tentativas_bySetor'")
	ret:= avg(dataMocks, []string{"setor"}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
		{"avg_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
		{"avg_porta": 51.0, "setor":"Devs", "username": "bob_dev"},
		{"avg_porta": 8080.0, "setor":"Devs", "username": "stranger"},
		{"avg_porta": nil, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_twoAgreg_porta_bySetorUsername'")
	ret:= avg(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg_oneAgreg_porta_byInex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"inex": nil, "avg_payload_mb":95.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg_oneAgreg_porta_byInex'")
	ret:= avg(dataMocks, []string{"inex"}, "payload_mb")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avg2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_avg_porta": 443.0, "setor":"Diretoria"},
		{"avg_avg_porta": 4065.5, "setor":"Devs"},
		{"avg_avg_porta": nil, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg2_zeroAgreg_porta_bySetorUsername'")
	parc := avg(dataMocks, []string{"setor", "username"}, "porta")
	ret := avg(parc, []string{"setor"}, "avg_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_avgsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"avg_sum_porta": 443.0, "setor":"Diretoria"},
		{"avg_sum_porta": 4091.0, "setor":"Devs"},
		{"avg_sum_porta": 0.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avgsum_twoAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := avg(parc, []string{"setor"}, "sum_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: min                                                                                                *
* descrição: Retorna o valor mínimo de uma coluna númerica qualquer agrupada pelos vaores em groupBy         *
*************************************************************************************************************/
func Test_min_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg'")
	ret:= min(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_tentativas": 1.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg_tentativas'")
	ret:= min(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_zeroAgreg_username(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg_username'")
	ret:= min(dataMocks, []string{}, "username")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_zeroAgreg_porta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_porta": 22.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg_porta'")
	ret:= min(dataMocks, []string{}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_oneAgreg_tentativas_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_tentativas": 1.0, "setor":"Diretoria"},
		{"min_tentativas": 1.0, "setor":"Devs"},
		{"min_tentativas": 3.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg_tentativas_bySetor'")
	ret:= min(dataMocks, []string{"setor"}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
		{"min_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
		{"min_porta": 22.0, "setor":"Devs", "username": "bob_dev"},
		{"min_porta": 8080.0, "setor":"Devs", "username": "stranger"},
		{"min_porta": nil, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_twoAgreg_porta_bySetorUsername'")
	ret:= min(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min_oneAgreg_porta_byInex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"inex": nil, "min_payload_mb":5.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_min_zeroAgreg_porta_byInex'")
	ret:= min(dataMocks, []string{"inex"}, "payload_mb")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_min2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_min_porta": 443.0, "setor":"Diretoria"},
		{"min_min_porta": 22.0, "setor":"Devs"},
		{"min_min_porta": nil, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg2_zeroAgreg_porta_bySetorUsername'")
	parc := min(dataMocks, []string{"setor", "username"}, "porta")
	ret := min(parc, []string{"setor"}, "min_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_minsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"min_sum_porta": 443.0, "setor":"Diretoria"},
		{"min_sum_porta": 102.0, "setor":"Devs"},
		{"min_sum_porta": 0.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg2_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := min(parc, []string{"setor"}, "sum_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: max                                                                                                *
* descrição: Retorna o valor máximo de uma coluna númerica qualquer agrupada pelos vaores em groupBy         *
*************************************************************************************************************/
func Test_max_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_max_zeroAgreg_tentativas'")
	ret:= max(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_zeroAgreg_tentativas(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_tentativas": 12.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_zeroAgreg_tentativas'")
	ret:= max(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_zeroAgreg_username(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_zeroAgreg_username'")
	ret:= max(dataMocks, []string{}, "username")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_zeroAgreg_porta(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_porta": 8080.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_zeroAgreg_porta'")
	ret:= max(dataMocks, []string{}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_oneAgreg_tentativas_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_tentativas": 5.0, "setor":"Diretoria"},
		{"max_tentativas": 12.0, "setor":"Devs"},
		{"max_tentativas": 3.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_zeroAgreg_tentativas_bySetor'")
	ret:= max(dataMocks, []string{"setor"}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_porta": 443.0, "setor":"Diretoria", "username": "admin_root"},
		{"max_porta": 443.0, "setor":"Diretoria", "username": "alice_silva"},
		{"max_porta": 80.0, "setor":"Devs", "username": "bob_dev"},
		{"max_porta": 8080.0, "setor":"Devs", "username": "stranger"},
		{"max_porta": nil, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_twoAgreg_porta_bySetorUsername'")
	ret:= max(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max_oneAgreg_porta_byInex(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"inex": nil, "max_payload_mb":300.0},
	}
	
	fmt.Println("[INFO] Executando 'Test_max_oneAgreg_porta_byInex'")
	ret:= max(dataMocks, []string{"inex"}, "payload_mb")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_max2_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_max_porta": 443.0, "setor":"Diretoria"},
		{"max_max_porta": 8080.0, "setor":"Devs"},
		{"max_max_porta": nil, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_max2_twoAgreg_porta_bySetorUsername'")
	parc := max(dataMocks, []string{"setor", "username"}, "porta")
	ret := max(parc, []string{"setor"}, "max_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_maxsum_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"max_sum_porta": 443.0, "setor":"Diretoria"},
		{"max_sum_porta": 8080.0, "setor":"Devs"},
		{"max_sum_porta": 0.0, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_avg2_zeroAgreg_porta_bySetorUsername'")
	parc := sum(dataMocks, []string{"setor", "username"}, "porta")
	ret := max(parc, []string{"setor"}, "sum_porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: arrPack                                                                                            *
* descrição: Inclui todos os valores em uma coluna array                                                     *
*************************************************************************************************************/
func Test_arrPack_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_arrPack_zeroAgreg'")
	ret:= arrPack(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_arrPack_oneAgreg_porta_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"arr_porta": []any{443, 443}, "setor":"Diretoria"},
		{"arr_porta": []any{80, 22, 8080}, "setor":"Devs"},
		{"arr_porta": []any{nil}, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_arrPack_oneAgreg_porta_bySetor'")
	ret := arrPack(dataMocks, []string{"setor"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_arrPack_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"arr_porta": []any{443}, "setor":"Diretoria", "username": "admin_root"},
		{"arr_porta": []any{443}, "setor":"Diretoria", "username": "alice_silva"},
		{"arr_porta": []any{80, 22}, "setor":"Devs", "username": "bob_dev"},
		{"arr_porta": []any{8080}, "setor":"Devs", "username": "stranger"},
		{"arr_porta": []any{nil}, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_arrPack_twoAgreg_porta_bySetorUsername'")
	ret := arrPack(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

/*************************************************************************************************************
* função: arrDistinctPack                                                                                            *
* descrição: Inclui todos os valores em uma coluna array                                                     *
*************************************************************************************************************/
func Test_arrDistinctPack_zeroAgreg(t *testing.T) {
	dataMocks := []JData{}
	
	var result []JData = nil
	
	fmt.Println("[INFO] Executando 'Test_arrDistinctPack_zeroAgreg'")
	ret:= arrDistinctPack(dataMocks, []string{}, "tentativas")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_arrDistinctPack_oneAgreg_porta_bySetor(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"set_porta": []any{443}, "setor":"Diretoria"},
		{"set_porta": []any{80, 22, 8080}, "setor":"Devs"},
		{"set_porta": []any{nil}, "setor":nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_arrDistinctPack_oneAgreg_porta_bySetor'")
	ret := arrDistinctPack(dataMocks, []string{"setor"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}

func Test_arrDistinctPack_twoAgreg_porta_bySetorUsername(t *testing.T) {
	dataMocks := []JData{
		{"setor": "Diretoria", "payload_mb": 120.0, "tentativas": 1.0, "porta": 443, "username": "admin_root"},
		{"setor": "Diretoria", "payload_mb": 80.0,  "tentativas": 5.0, "porta": 443, "username": "alice_silva"},
		{"setor": "Devs",       "payload_mb": 300.0, "tentativas": 2.0, "porta": 80,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 15.0,  "tentativas": 12.0,"porta": 22,  "username": "bob_dev"},
		{"setor": "Devs",       "payload_mb": 5.0,   "tentativas": 1.0, "porta": 8080,"username": "stranger"},
		{"setor": nil,          "payload_mb": 50.0,  "tentativas": 3.0, "porta": nil,   "username": nil},
	}
	
	result := []JData{
		{"set_porta": []any{443}, "setor":"Diretoria", "username": "admin_root"},
		{"set_porta": []any{443}, "setor":"Diretoria", "username": "alice_silva"},
		{"set_porta": []any{80, 22}, "setor":"Devs", "username": "bob_dev"},
		{"set_porta": []any{8080}, "setor":"Devs", "username": "stranger"},
		{"set_porta": []any{nil}, "setor":nil, "username": nil},
	}
	
	fmt.Println("[INFO] Executando 'Test_arrDistinctPack_twoAgreg_porta_bySetorUsername'")
	ret := arrDistinctPack(dataMocks, []string{"setor", "username"}, "porta")
	
	if !reflect.DeepEqual(result, ret) {
		t.Errorf("[ERROR] Erro: Obtido %v, mas esperado %v", ret, result)
	}
}
