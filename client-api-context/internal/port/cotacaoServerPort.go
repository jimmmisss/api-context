package port

import "github.com/jimmmisss/client-api-context/internal/model"

type CotacaoAPIInterface interface {
	GetCotacaoAPI() (*model.Bid, error)
}

type CotacaoCriaArquivoInterface interface {
	GeraArquivoTxt(cotacao *model.Bid) error
}
