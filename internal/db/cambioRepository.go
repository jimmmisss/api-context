package db

import (
	"api-context/internal/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type CotacaoRepository struct {
	db *sqlx.DB
}

func NewCotacaoRepository(db *sqlx.DB) *CotacaoRepository {
	return &CotacaoRepository{db: db}
}

func (repository *CotacaoRepository) Create(c *model.Cotacao) error {
	_, err := repository.db.Exec(`
		INSERT INTO cotacoes (code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, c.Code, c.Codein, c.Name, c.High, c.Low, c.VarBid, c.PctChange, c.Bid, c.Ask, c.Timestamp, c.CreateDate)
	if err != nil {
		log.Println("Erro ao inserir dados no banco de dados:", err)
	}

	return nil
}
