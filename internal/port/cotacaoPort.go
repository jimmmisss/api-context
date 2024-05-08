package port

import "api-context/internal/model"

type CotacaoRepository interface {
	Create(cotacao *model.Cotacao) error
}

type CotacaoAPIService interface {
	GetCotacaoAPI() (*model.CotacaoDTO, error)
}
