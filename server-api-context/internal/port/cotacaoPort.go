package port

import (
	"github.com/jimmmisss/server-api-context/internal/model"
	"net/http"
)

type CotacaoRepositoryInterface interface {
	Create(cotacao *model.Cotacao) error
}

type CotacaoAPIInterface interface {
	GetCotacaoAPI() (*model.CotacaoDTO, error)
}

type CotacaoServiceInterface interface {
	ObtemCotacaoViaApiSalva() (*model.BidResponse, error)
}

type CotacaoHandlerInterface interface {
	CotacaoHandler(w http.ResponseWriter, r *http.Request)
}
