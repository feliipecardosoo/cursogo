package servidor

import (
	"crud/banco"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID   uint32 `json:"id"`
	Nome string `json:"nome"`
}

// CriarUsuario insere um user no BD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição, parece estar vazio"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
}
