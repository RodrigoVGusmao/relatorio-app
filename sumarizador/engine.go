package sumarizador

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

func groupBySanitization(groupBy []string) []string {
	seen := make(map[string]bool)
	var unique []string
	
	for _, col := range groupBy {
		if !seen[col] {
			seen[col] = true
			unique = append(unique, col)
		}
	}
	return unique
}

func genKeyString(data JData, groupBy []string) []string {
	if len(groupBy) == 0 {
		return []string{}
	}
	
	result := make([]string, len(groupBy))
	for i, column := range groupBy {
		val, ok := data[column]
		if ok && val != nil {
			result[i] = strconv.Quote(fmt.Sprintf("%v", val))
		} else if ok {
			result[i] = "\"\""
		} else {
			result[i] = "\"\""
		}
	}
	return result
}

func genSliceHash(strs []string) uint64 {
	h := fnv.New64a()
	
	for _, val := range strs {
		h.Write([]byte(val))
		h.Write([]byte{0})
	}
	
	return h.Sum64()
}

func copySubsetJData(src JData, keys []string) JData {
	copyJData := make(JData, len(keys))
	
	for _, key := range keys {
		if val, ok := src[key]; ok {
			copyJData[key] = val
		} else
		{
			copyJData[key] = nil
		}
	}
	
	return copyJData
}

type aggFunc func(accumulator JData, currentData JData, column string, target string)

func aggregate(data []JData, groupBy []string, target string, prefix string, pre aggFunc, proc aggFunc, post aggFunc) []JData {
	indices := make(map[uint64]uint32)
	var result []JData
	column := prefix + "_" + target
	
	sanGroupBy := groupBySanitization(groupBy)
	
	for _, value := range data {
		valueCopy := value.Clone()
		group := genKeyString(valueCopy, sanGroupBy)
		hash := genSliceHash(group)
		
		idx, exists := indices[hash]
		if !exists {
			idx = uint32(len(result))
			indices[hash] = idx
			newData := copySubsetJData(valueCopy, sanGroupBy)
			if pre != nil {
				pre(newData, valueCopy, column, target)
			}
			result = append(result, newData)
		}
		
		if proc != nil {
			proc(result[idx], valueCopy, column, target)
		}
	}
	
	if post != nil {
		for _, value := range result {
			post(value, nil, column, target)
		}
	}
	
	return result
}
