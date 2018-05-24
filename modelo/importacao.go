package modelo

import (
	"os"
	"encoding/csv"
	"io"
	"strings"
	"strconv"
	"fmt"
)

type Cidade struct {
	Uf                string
	CodUf             string
	CodMunicipio      string
	Nome              string
	PopulacaoEstimada int
}

func LeArquivo(nome string) ([]Cidade, error) {
	f, err := os.Open(nome)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	reader.Comma = '\t'
	cidades := make([]Cidade, 0)
	for {
		l, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		strpop := strings.TrimSpace(strings.Replace(l[4], ".", "", -1))
		strnum := ""
		for i := 0; i < len(strpop); i++ {
			d := strpop[i : i+1]
			if strings.Index("0123456789", d) == -1 {
				break
			}
			strnum += d
		}
		populacaoEstimada, err := strconv.Atoi(strnum)
		if err != nil {
			return nil, err
		}
		cidade := Cidade{
			Uf:                l[0],
			CodUf:             l[1],
			CodMunicipio:      l[2],
			Nome:              l[3],
			PopulacaoEstimada: populacaoEstimada,
		}
		cidades = append(cidades, cidade)
	}
	return cidades, nil
}

func Importa(path string) error {
	cidades, err := LeArquivo(path)
	if err != nil {
		return err
	}
	for _, c := range cidades {
		fmt.Printf("%#v\n", c)
		m, err := MunicipioInsere(c.Uf, c.CodUf, c.CodMunicipio, c.Nome, c.PopulacaoEstimada)
		if err != nil {
			return err
		}
		fmt.Println(m)
	}

	return nil
}
