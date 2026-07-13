package gerador

import (
	"gopkg.in/yaml.v3"
	"html/template"
	"io/fs"
	"os"
	"relatorio/sumarizador"
	"sort"
	"time"
)

func LoadReportConfigFiles(fds fs.FS, dirPath string) (Reports, error) {
	var result Reports
	
	data, err := fs.ReadFile(fds, dirPath)
	if err != nil {
	    return nil, err
	}
	
	err = yaml.Unmarshal(data, &result)
	
	for key, value := range result {
		for i, function := range value.Pipeline {
			result[key].Pipeline[i].Exec = sumarizador.StrToProc(function.Proc)
		}
	}
	
	return result, err
}

func RunReportConfig(config Report, data []sumarizador.JData) []sumarizador.JData {
	result := data
	for _, function := range config.Pipeline {
		result = function.Exec(result, function.Columns, function.Target)
	}
	return result
}

func ExportToHTML(outputs map[string][]sumarizador.JData, outputPath string) error {
	reportTables := make(map[string]HTMLTable)

	for reportName, rows := range outputs {
		var headers []string

		if len(rows) > 0 {
			for k := range rows[0] {
				headers = append(headers, k)
			}
			sort.Strings(headers)
		}

		reportTables[reportName] = HTMLTable{
			Headers: headers,
			Rows:    rows,
		}
	}

	data := PageData{
		GeneratedAt: time.Now().Format("02/01/2006 15:04:05"),
		Reports:     reportTables,
	}

	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
