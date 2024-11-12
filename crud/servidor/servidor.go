package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome) values (?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome)
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement/ insert"))
		return
	}

	iduser, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter ID inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("ID: %d", iduser)))
}

// BuscarUsuarios retorna todos os usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT id, nome FROM usuarios")
	if erro != nil {
		http.Error(w, "Erro ao buscar os usuários", http.StatusInternalServerError)
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome); erro != nil {
			http.Error(w, "Erro ao ler os dados dos usuários", http.StatusInternalServerError)
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if erro := linhas.Err(); erro != nil {
		http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		http.Error(w, "Erro ao enviar os dados", http.StatusInternalServerError)
	}
}

// BuscarUsuario retorna um usuario do banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		http.Error(w, "Erro ao buscar usuario", http.StatusInternalServerError)
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome); erro != nil {
			http.Error(w, "Erro ao escanear usuario", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		http.Error(w, "Erro ao converter os dados", http.StatusInternalServerError)
	}
}

// DeletarUsuario deleta o usuario passado por parametro pela URL
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linha, erro := db.Query("delete from usuarios where id = ?", ID)
	if erro != nil {
		http.Error(w, "Erro ao deletar usuario", http.StatusInternalServerError)
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome); erro != nil {
			http.Error(w, "Erro ao escanear usuario", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		http.Error(w, "Erro ao converter os dados", http.StatusInternalServerError)
	}

}

// AtualizarUsuario vai atualizar o usuario que foi passado pelo parametro
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obter os parâmetros da URL
	parametros := mux.Vars(r)

	// Converter o parâmetro "id" para inteiro
	ID, erro := strconv.ParseInt(parametros["id"], 10, 64)
	if erro != nil {
		http.Error(w, "Erro ao converter o parametro para inteiro", http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Ler o corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Deserializar o corpo para a estrutura "usuario"
	var usuarioBody usuario
	if erro = json.Unmarshal(corpoReq, &usuarioBody); erro != nil {
		http.Error(w, "Erro ao converter o usuário para struct", http.StatusBadRequest)
		return
	}

	// Preparar a query de atualização
	statement, erro := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	if erro != nil {
		http.Error(w, "Erro ao criar o statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	// Executar a query com os parâmetros fornecidos
	_, erro = statement.Exec(usuarioBody.Nome, ID)
	if erro != nil {
		http.Error(w, "Erro ao executar a query de atualização", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário atualizado com sucesso"))
}
