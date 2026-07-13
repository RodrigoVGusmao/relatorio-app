package conector

import(
	"strconv"
)

func makeColumn(rename string, isPK bool, dataType string) ConfigColumn {
	return ConfigColumn{
		DataType:     dataType,
		IsPrimaryKey: isPK,
		Rename:       rename,
	}
}
func deColumn(rename string, isPK bool) ConfigColumn {
	return makeColumn(rename, isPK, "dummy")
}
func uuidColumn(rename string) ConfigColumn {
	return makeColumn(rename, true, "uuid")
}
func fkColumn(rename string, fkTable, fkCol string) ConfigColumn {
	col := deColumn(rename, false)
	col.ForeignKey = FKey{Table: fkTable, Column: fkCol}
	return col
}
func env(schema map[string]ConfigColumn) ConfigEndpoint {
	return ConfigEndpoint{
		EndpointLocation: "dummy",
		Schema:           schema,
	}
}

func jsonData(data ...string) []byte {
	jsonStr := "["
	temp := ""
	for _, dat := range data {
		jsonStr += temp + dat
		temp = ","
	}
	return []byte(jsonStr+"]")
}
func jsonRecord(data ...string) string {
	jsonStr := "{"
	temp := ""
	for _, dat := range data {
		jsonStr += temp + dat
		temp = ","
	}
	return jsonStr + "}"
}
func jsonString(name string, data string) string {
	return "\"" + name + "\": \"" + data + "\""
}
func jsonFloat(name string, data float64) string {
	return "\"" + name + "\": " + strconv.FormatFloat(data,'f', -1, 64)
}
func jsonBool(name string, data bool) string {
	return "\"" + name + "\": " + strconv.FormatBool(data)
}

func fileRoot(envs ...string) []byte {
	yamlStr := "# version: \"1.0\"\n"
	for _, env := range envs {
		yamlStr += env
	}
	return []byte(yamlStr)
}
func fileEnv(name string, schemas ...string) string {
	yamlStr := name+":\n    endpoint: \"dummy\"\n    schema:\n"
	for _, schema := range schemas {
		yamlStr += schema+"\n"
	}
	return yamlStr
}
func fileMakeColumn(name string, rename string, isPK bool, dataType string) string {
	return "        " +name + ":\n            name: \"" + rename + "\"\n            primary_key: " + strconv.FormatBool(isPK) + "\n            type: \"" + dataType + "\""
}
func fileDeColumn(name string, rename string, isPK bool) string {
	return fileMakeColumn(name, rename, isPK, "dummy")
}
func fileUuidColumn(name string, rename string) string {
	return fileMakeColumn(name, rename, true, "uuid")
}
func fileFkColumn(name string, rename string, fkTable string, fkCol string) string {
	return fileDeColumn(name, rename, false) + "\n            foreign_key:\n                table: \"" + fkTable + "\"\n                column: \"" + fkCol + "\""
}
