package db

import (
	"api-context/internal/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type cotacaoRepository struct {
	db *sqlx.DB
}

func NewCotacaoRepository(db *sqlx.DB) model.CotacaoRepository {
	return &cotacaoRepository{db: db}
}

func (c *cotacaoRepository) Create(m *model.Cotacao) error {
	_, err := c.db.Exec(`
		INSERT INTO cotacoes (code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, m.Code, m.Codein, m.Name, m.High, m.Low, m.VarBid, m.PctChange, m.Bid, m.Ask, m.Timestamp, m.CreateDate)
	if err != nil {
		log.Println("Erro ao inserir dados no banco de dados:", err)
	}

	return nil
}
