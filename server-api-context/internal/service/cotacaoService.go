package service

import (
	"github.com/jimmmisss/server-api-context/internal/model"
	"github.com/jimmmisss/server-api-context/internal/port"
	"log"
)

type CotacaoService struct {
	Repo port.CotacaoRepositoryInterface
	Api  port.CotacaoAPIInterface
}

func (c *CotacaoService) ObtemCotacaoViaApiSalva() (*model.BidResponse, error) {
	cotacaoDTO, err := c.Api.GetCotacaoAPI()
	if err != nil {
		log.Printf("================== ERRO NA API - ObtemCotacaoViaApiSalva: %v", err)
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

	err = c.Repo.Create(cotacao)
	if err != nil {
		log.Printf("================== ERRO NO REPO - ObtemCotacaoViaApiSalva - Create: %v", err)
		return nil, err
	}

	bid := &model.BidResponse{
		Bid: cotacaoDTO.USDBRL.Bid,
	}

	return bid, nil
}
