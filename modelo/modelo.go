package modelo

import (
	"database/sql"
	"strings"
)

var db *sql.DB

func Db(d *sql.DB) {
	db = d
}

func ft(texto string) string {
	palavras := strings.Split(texto, " ")
	p := make([]string, 0)
	for _, palavra := range palavras {
		if palavra == "" {
			continue
		}
		p = append(p, "+"+palavra+"*")
	}
	return strings.Join(p, " ")
}
