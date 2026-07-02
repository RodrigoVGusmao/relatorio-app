package main

import (
	"math"
)

func zeroPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = 0.0
}

//agrupamento
func groupBy(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, "", "", nil, nil, nil)
}
func groupByPosProc(accumulator JData, currentData JData, column string, target string) {
	delete(accumulator, column)
}

//contagem
func count(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, "", "count", zeroPreProc, countProc, nil)
}
func countProc(accumulator JData, currentData JData, column string, target string) {
	currentCount, _ := toFloat64(accumulator[column])
	accumulator[column] = currentCount + 1.0
}

//contagem não nula
func countNotNil(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "count_not_null", zeroPreProc, countNotNilProc, nil)
}
func countNotNilProc(accumulator JData, currentData JData, column string, target string) {
	if val, exists := currentData[target]; exists && val != nil {
		countProc(accumulator, currentData, column, target)
	}
}

//contagem de valores únicos
func countDistinct(data []JData, groupByVar []string, target string) []JData {
	ret := groupBy(data, append(groupByVar, target), "")
	return aggregate(ret, groupByVar, target, "count_distinct", zeroPreProc, countProc, nil)
}

//soma
func sum(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "sum", zeroPreProc, sumProc, nil)
}
func sumProc(accumulator JData, currentData JData, column string, target string) {
	var targetVal float64
	if val, ok := toFloat64(currentData[target]); ok {
		targetVal = val
	} else {
		return
	}
	currentSum, _ := toFloat64(accumulator[column])
	accumulator[column] = currentSum + targetVal
}

//média
func avg(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "avg", zeroPreProc, avgProc, avgPosProc)
}
func avgProc(accumulator JData, currentData JData, column string, target string) {
	if _, ok := toFloat64(currentData[target]); !ok {
		return
	}
	sumProc(accumulator, currentData, "_sum\x1f", target)
	countProc(accumulator, currentData, "_count\x1f", "")
}
func avgPosProc(accumulator JData, currentData JData, column string, target string) {
	if _, exists := accumulator["_count\x1f"]; !exists {
		accumulator[column] = nil
		return
	}
	currentSum, _ := toFloat64(accumulator["_sum\x1f"])
	currentCount, _ := toFloat64(accumulator["_count\x1f"])
	accumulator[column] = currentSum/currentCount
	delete(accumulator, "_sum\x1f")
	delete(accumulator, "_count\x1f")
}

//mínimo
func min(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "min", minPreProc, minProc, minMaxPosProc)
}
func minPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = math.MaxFloat64
	accumulator["_current\x1f"] = column
}
func minProc(accumulator JData, currentData JData, column string, target string) {
	var targetVal float64
	if val, ok := toFloat64(currentData[target]); ok {
		targetVal = val
		accumulator["_valid\x1f"] = true
	} else {
		return
	}
	currentMin, _ := toFloat64(accumulator[column])
	accumulator[column] = math.Min(currentMin, targetVal)
}
func minMaxPosProc(accumulator JData, currentData JData, column string, target string) {
	if _, exists := accumulator["_valid\x1f"]; !exists {
		accumulator[accumulator["_current\x1f"].(string)] = nil
	}
	delete(accumulator, "_valid\x1f")
	delete(accumulator, "_current\x1f")
}

//máximo
func max(data []JData, groupBy []string, target string) []JData {
	return aggregate(data, groupBy, target, "max", maxPreProc, maxProc, minMaxPosProc)
}
func maxPreProc(accumulator JData, currentData JData, column string, target string) {
	accumulator[column] = -math.MaxFloat64
	accumulator["_current\x1f"] = column
}
func maxProc(accumulator JData, currentData JData, column string, target string) {
	var targetVal float64
	if val, ok := toFloat64(currentData[target]); ok {
		targetVal = val
		accumulator["_valid\x1f"] = true
	} else {
		return
	}
	currentMax, _ := toFloat64(accumulator[column])
	accumulator[column] = math.Max(currentMax, targetVal)
}

//agregação de valores em arrays
func arrPack(data []JData, groupBy []string, target string) []JData {
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
func arrDistinctPack(data []JData, groupByVar []string, target string) []JData {
	ret := groupBy(data, append(groupByVar, target), "")
	return aggregate(ret, groupByVar, target, "set", arrPreProc, arrProc, arrPosProc)
}
