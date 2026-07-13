package gerador

import (
	"relatorio/sumarizador"
)

type AggArgs struct {
	Exec sumarizador.AggProc
	Proc string		`yaml:"operation"`
	Columns []string	`yaml:"columns"`
	Target string		`yaml:"target"`
}
type Report struct {
	Pipeline []AggArgs	`yaml:"pipeline"`
}
type Reports map[string]Report

type HTMLTable struct {
	Headers []string
	Rows    []sumarizador.JData
}
type PageData struct {
	GeneratedAt string
	Reports     map[string]HTMLTable
}
