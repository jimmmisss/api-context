package config

import (
	"github.com/google/wire"
	"github.com/jimmmisss/server-api-context/internal/api"
	"github.com/jimmmisss/server-api-context/internal/handler"
	"github.com/jimmmisss/server-api-context/internal/port"
	"github.com/jimmmisss/server-api-context/internal/repository"
	"github.com/jimmmisss/server-api-context/internal/service"
	"sync"
)

var (
	handlerCotacao     *handler.CotacaoHandler
	handlerCotacaoOnce sync.Once

	serviceCotacao     *service.CotacaoService
	serviceCotacaoOnce sync.Once

	apiCotacao     *api.CotacaoAPI
	apiCotacaoOnce sync.Once

	repositoryCotacao     *repository.CotacaoRepository
	repositoryCotacaoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProviderHandlerCotacao,
		ProviderServiceCotacao,
		ProviderApiServiceCotacao,
		ProviderRepositoryCotacao,

		wire.Bind(new(port.CotacaoHandlerInterface), new(*handler.CotacaoHandler)),
		wire.Bind(new(port.CotacaoServiceInterface), new(*service.CotacaoService)),
		wire.Bind(new(port.CotacaoAPIInterface), new(*api.CotacaoAPI)),
		wire.Bind(new(port.CotacaoRepositoryInterface), new(*repository.CotacaoRepository)),
	)
)

func ProviderHandlerCotacao(service port.CotacaoServiceInterface) *handler.CotacaoHandler {
	handlerCotacaoOnce.Do(func() {
		handlerCotacao = &handler.CotacaoHandler{
			Service: service,
		}
	})
	return handlerCotacao
}

func ProviderServiceCotacao(repo port.CotacaoRepositoryInterface, api port.CotacaoAPIInterface) *service.CotacaoService {
	serviceCotacaoOnce.Do(func() {
		serviceCotacao = &service.CotacaoService{
			Repo: repo,
			Api:  api,
		}
	})
	return serviceCotacao
}

func ProviderApiServiceCotacao() *api.CotacaoAPI {
	apiCotacaoOnce.Do(func() {
		apiCotacao = &api.CotacaoAPI{}
	})
	return apiCotacao
}

func ProviderRepositoryCotacao() *repository.CotacaoRepository {
	repositoryCotacaoOnce.Do(func() {
		repositoryCotacao = &repository.CotacaoRepository{}
	})
	return repositoryCotacao
}
