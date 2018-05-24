package main

import (
	"log"
	"github.com/luisfernandogaido/cidades/modelo"
	"github.com/luisfernandogaido/cidades/conf"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	if err := conf.Load("./conf.json"); err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", conf.Conf.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	modelo.Db(db)
	http.HandleFunc("/municipios", municipios)
	http.ListenAndServe(conf.Conf.Porta, nil)
}

func municipios(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if search == "" {
		http.Error(w, "cidade não encontrada", http.StatusNotFound)
		return
	}
	municipios, err := modelo.MunicipiosSeleciona(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	printJson(w, municipios)
}

func printJson(w http.ResponseWriter, i interface{}) error {
	bytes, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json; charset=utf8")
	fmt.Fprintln(w, string(bytes))
	return nil
}
