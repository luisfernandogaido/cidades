package modelo

import (
	"fmt"
	"strings"
	"sync"
)

type Municipio struct {
	Codigo       string
	Uf           string
	CodUf        string
	CodMunicipio string
	Nome         string
	Populacao    int
	Indice       string
}

var (
	muMapa sync.RWMutex
	mapa   = make(map[string][]Municipio)
)

func MunicipioInsere(uf, coduf, codMunicipio, nome string, populacao int) (Municipio, error) {
	m := Municipio{
		Codigo:       coduf + codMunicipio,
		Uf:           uf,
		CodUf:        coduf,
		CodMunicipio: codMunicipio,
		Nome:         nome,
		Populacao:    populacao,
		Indice:       coduf + codMunicipio + " " + uf + " " + nome,
	}
	fmt.Println(uf, len(uf))
	_, err := db.Exec(
		"CALL municipio_insere(?,?,?,?,?,?,?)",
		m.Codigo,
		m.Uf,
		m.CodUf,
		m.CodMunicipio,
		m.Nome,
		m.Populacao,
		m.Indice,
	)
	return m, err
}

func MunicipiosSeleciona(search string) ([]Municipio, error) {
	search = strings.ToLower(search)
	muMapa.RLock()
	if municipios, ok := mapa[search]; ok {
		muMapa.RUnlock()
		return municipios, nil
	}
	muMapa.RUnlock()
	searchFt := ft(search)
	rows, err := db.Query("CALL municipios_seleciona(?)", searchFt)
	if err != nil {
		return nil, err
	}
	municipios := make([]Municipio, 0)
	for rows.Next() {
		m := Municipio{}
		err = rows.Scan(
			&m.Codigo,
			&m.Uf,
			&m.CodUf,
			&m.CodMunicipio,
			&m.Nome,
			&m.Populacao,
			&m.Indice,
		)
		if err != nil {
			return nil, err
		}
		municipios = append(municipios, m)
	}
	muMapa.Lock()
	mapa[search] = municipios
	muMapa.Unlock()
	return municipios, nil
}
