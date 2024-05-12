package service

import (
	"github.com/jimmmisss/server-api-context/internal/model"
	"github.com/jimmmisss/server-api-context/internal/port"
)

type CotacaoService struct {
	repository port.CotacaoRepositoryInterface
	apiService port.CotacaoAPIInterface
}

func NewCotacaoService(
	repository port.CotacaoRepositoryInterface,
	apiRepository port.CotacaoAPIInterface) *CotacaoService {
	return &CotacaoService{
		repository: repository,
		apiService: apiRepository,
	}
}

func (c *CotacaoService) ObtemCotacaoESalva() (*model.BidResponse, error) {
	cotacaoDTO, err := c.apiService.GetCotacaoAPI()
	if err != nil {
		return nil, err
	}

	cotacao := &model.Cotacao{
		Code:       cotacaoDTO.USDBRL.Code,
		Codein:     cotacaoDTO.USDBRL.Codein,
		Name:       cotacaoDTO.USDBRL.Name,
		High:       cotacaoDTO.USDBRL.High,
		Low:        cotacaoDTO.USDBRL.Low,
		VarBid:     cotacaoDTO.USDBRL.VarBid,
		PctChange:  cotacaoDTO.USDBRL.PctChange,
		Bid:        cotacaoDTO.USDBRL.Bid,
		Ask:        cotacaoDTO.USDBRL.Ask,
		Timestamp:  cotacaoDTO.USDBRL.Timestamp,
		CreateDate: cotacaoDTO.USDBRL.CreateDate,
	}

	err = c.repository.Create(cotacao)
	if err != nil {
		return nil, err
	}

	bid := &model.BidResponse{
		Bid: cotacaoDTO.USDBRL.Bid,
	}

	return bid, nil
}
