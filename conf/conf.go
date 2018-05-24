package conf

import (
	"io/ioutil"
	"encoding/json"
)

type conf struct {
	Dsn string `json:"dsn"`
	Porta string `json:"porta"`
}

var Conf conf

func Load(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &Conf)
	return err
}
