package db

import (
	"api-context/internal/model"
	"github.com/jmoiron/sqlx"
)

type cotacaoRepository struct {
	db *sqlx.DB
}

func NewCotacaoRepository(aDB *sqlx.DB) model.CotacaoRepository {
	return &cotacaoRepository{db: aDB}
}

func (c *cotacaoRepository) Create(dolar *model.Cotacao) error {
	_, err := c.db.Exec("INSERT INTO cotacao ("+
		"code,"+
		" codein,"+
		" name,"+
		" high,"+
		" low,"+
		" varBid,"+
		" pctChange,"+
		" bid,"+
		" ask,"+
		" timestamp,"+
		" createdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, )",
		dolar.Code,
		dolar.Codein,
		dolar.Name,
		dolar.High,
		dolar.Low,
		dolar.VarBid,
		dolar.PctChange,
		dolar.Bid,
		dolar.Ask,
		dolar.Timestamp,
		dolar.CreateDate)
	return err
}
