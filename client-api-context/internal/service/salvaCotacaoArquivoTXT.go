package service

import (
	"fmt"
	"github.com/jimmmisss/client-api-context/internal/model"
	"log"
	"os"
)

type GeraArquivoTxtService struct {
}

func NewGeraArquivoTxtService() *GeraArquivoTxtService {
	return &GeraArquivoTxtService{}
}

func (g *GeraArquivoTxtService) GeraArquivoTxt(cotacao *model.Bid) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Erro ao abrir arquivo:", err)
	}
	defer file.Close()

	text := fmt.Sprintf("Bid do dolar: %s\n", cotacao.Bid)
	_, err = file.WriteString(text)
	if err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
		return err
	}
	return nil
}
