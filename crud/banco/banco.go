package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver conn mysql
)

// Conectar responsavel por toda a parte de conn do banco de dados!
func Conectar() (*sql.DB, error) {
	stringCon := "root:Fg310303@/devbook?charset=utf8&parseTime=True&Loc=Local"

	db, erro := sql.Open("mysql", stringCon)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
