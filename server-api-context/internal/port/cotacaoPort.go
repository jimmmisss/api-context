package port

import "github.com/jimmmisss/server-api-context/internal/model"

type CotacaoRepositoryInterface interface {
	Create(cotacao *model.Cotacao) error
}

type CotacaoAPIInterface interface {
	GetCotacaoAPI() (*model.CotacaoDTO, error)
}

type CotacaoServiceInterface interface {
	ObtemCotacaoESalva() (*model.BidResponse, error)
}