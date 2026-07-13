package sumarizador

import (
	"math"
)

func zeroPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = 0.0
}
func nilPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = nil
}

//agrupamento
func GroupBy(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, "", "", nil, nil, nil)
}

//contagem
func Count(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, "", "count", zeroPreProc, countProc, nil)
}
func countProc(accumulator JData, currentData JData, column string, target string) {
	currentCount, _ := toFloat64(accumulator[column])
	accumulator[column] = currentCount + 1.0
}

//contagem não nula
func CountNotNil(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "count_not_null", zeroPreProc, countNotNilProc, nil)
}
func countNotNilProc(accumulator JData, currentData JData, column string, target string) {
	if val, exists := currentData[target]; exists && val != nil {
		countProc(accumulator, currentData, column, target)
	}
}

//contagem de valores únicos
func CountDistinct(data []JData, groupByVar []string, target string) []JData {
	ret := GroupBy(data, append(groupByVar, target), "")
	return aggregate(ret, groupByVar, target, "count_distinct", zeroPreProc, countProc, nil)
}

//soma
func Sum(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "sum", nilPreProc, sumProc, nil)
}
func sumProc(accumulator JData, currentData JData, column string, target string) {
	targetVal, ok1 := toFloat64(currentData[target])
	currentSum, ok2 := toFloat64(accumulator[column])
	if !ok2 && !ok1 {
		return
	}
	accumulator[column] = currentSum + targetVal
}

//média
func Avg(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "avg", nil, avgProc, avgPosProc)
}
func avgProc(accumulator JData, currentData JData, column string, target string) {
	if _, ok := toFloat64(currentData[target]); !ok {
		return
	}
	sumProc(accumulator, currentData, column, target)
	countProc(accumulator, currentData, "\x12count", "")
}
func avgPosProc(accumulator JData, currentData JData, column string, target string) {
	if _, exists := accumulator["\x12count"]; !exists {
		accumulator[column] = nil
		return
	}
	currentSum, _ := toFloat64(accumulator[column])
	currentCount, _ := toFloat64(accumulator["\x12count"])
	accumulator[column] = currentSum/currentCount
	delete(accumulator, "\x12count")
}

//mínimo
func Min(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "min", minPreProc, minProc, minMaxPosProc)
}
func minPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = math.MaxFloat64
	accumulator["\x12current"] = column
}
func minProc(accumulator JData, currentData JData, column string, target string) {
	targetVal, ok := toFloat64(currentData[target])
	if !ok {
		return
	}
	var currentMin float64
	currentMin, accumulator["\x12valid"] = toFloat64(accumulator[column])
	accumulator[column] = math.Min(currentMin, targetVal)
}
func minMaxPosProc(accumulator JData, currentData JData, column string, target string) {
	if _, exists := accumulator["\x12valid"]; !exists {
		accumulator[accumulator["\x12current"].(string)] = nil
	}
	delete(accumulator, "\x12valid")
	delete(accumulator, "\x12current")
}

//máximo
func Max(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "max", maxPreProc, maxProc, minMaxPosProc)
}
func maxPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = -math.MaxFloat64
	accumulator["\x12current"] = column
}
func maxProc(accumulator JData, currentData JData, column string, target string) {
	targetVal, ok := toFloat64(currentData[target])
	if !ok {
		return
	}
	var currentMax float64
	currentMax, accumulator["\x12valid"] = toFloat64(accumulator[column])
	accumulator[column] = math.Max(currentMax, targetVal)
}

//agregação de valores em arrays
func ArrPack(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "arr", arrPreProc, arrProc, arrPosProc)
}
func arrPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = &[]any{}
}
func arrProc(accumulator JData, currentData JData, column string, target string) {
	arrPtr := (accumulator[column].(*[]any))
	*arrPtr = append(*arrPtr, currentData[target])
}
func arrPosProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = *(accumulator[column].(*[]any))
}

//agregação de valores distintos em array
func ArrDistinctPack(data []JData, groupByVar []string, target string) []JData {
	ret := GroupBy(data, append(groupByVar, target), "")
	return aggregate(ret, groupByVar, target, "set", arrPreProc, arrProc, arrPosProc)
}
