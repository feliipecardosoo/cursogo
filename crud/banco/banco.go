package banco

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver para o MySQL
)

// Conectar é responsável pela inicialização da conexão com o banco de dados MySQL.
func Conectar() (*sql.DB, error) {
	// Variáveis de configuração para maior flexibilidade
	user := "root"
	password := "Fg310303"
	host := "localhost"
	port := "3306"
	database := "devbook"

	// Montagem da string de conexão
	stringCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	// Abre a conexão com o banco de dados
	db, erro := sql.Open("mysql", stringCon)
	if erro != nil {
		// Se não conseguiu abrir a conexão, retorna erro
		log.Println("Erro ao abrir a conexão:", erro)
		return nil, erro
	}

	// Verifica se a conexão com o banco está realmente funcionando
	if erro = db.Ping(); erro != nil {
		// Se falhar ao conectar, retorna erro
		log.Println("Erro ao conectar com o banco de dados:", erro)
		return nil, erro
	}

	// Retorna a conexão com o banco se tudo correr bem
	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db, nil
}
