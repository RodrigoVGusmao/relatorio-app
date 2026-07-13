package correlacionador

import (
	"relatorio/conector"
)

func joinRecord(data1 conector.NormData, data2 conector.NormData, joinColumn string) {
	for k, v := range data2 {
			data1[k] = v
	}
	delete(data1, joinColumn)
}

func joinRecords(dstRecord conector.NormRecord, column string, cmpRecord conector.NormRecord, foreignColumn string) {
	for i, value := range dstRecord {
		for j, value2 := range cmpRecord {
			if value[column] == value2[foreignColumn] {
				joinRecord(dstRecord[i], cmpRecord[j], column)
				break
			}
		}
	}
}

func MakeUnifiedData(normRecords conector.NormRecords, config conector.ConfigEnvironment, parentTable string) conector.NormRecord {
	for endpoint, configItem := range config {
		for _, info := range configItem.(conector.ConfigEndpoint).Schema {
			if info.ForeignKey.Column == "" {
				continue
			}
			joinRecords(normRecords[endpoint], endpoint+"."+info.Rename, normRecords[info.ForeignKey.Table], info.ForeignKey.Table+"."+info.ForeignKey.Column)
			normRecords[info.ForeignKey.Table] = normRecords[endpoint]
		}
	}
	return normRecords[parentTable]
}
