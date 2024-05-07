package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ConnDB struct {
	DB *sqlx.DB
}

func (c *ConnDB) Close() {
	c.DB.Close()
}

func NewConnectDB() (*ConnDB, error) {
	db, err := sqlx.Connect("sqlite3", "cotacoes.db")
	if err != nil {
		log.Fatalln("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	criarTabela := `CREATE TABLE IF NOT EXISTS cotacao (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code TEXT,
			codein TEXT,
			name TEXT,
			high TEXT,
			low TEXT,
			var_bid TEXT,
			pct_change TEXT,
			bid TEXT,
			ask TEXT,
			timestamp TEXT,
			create_date TEXT
		);`
	_, err = db.Exec(criarTabela)
	if err != nil {
		log.Fatalln("Erro ao criar a tabela:", err)
	}

	return &ConnDB{DB: db}, nil
}
