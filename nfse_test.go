package plugnotas_test

import (
	"os"
	"testing"

	"github.com/Stratumhk/plugnotas-sdk-go"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCreateNfse(t *testing.T) {
	godotenv.Load("./.env")
	client := plugnotas.NewClient(os.Getenv("API_KEY"))

	endereco := &plugnotas.Endereco{
		TipoLogradouro: "RUA",
		Logradouro:     "MAZ",
		Numero:         "2111",
		TipoBairro:     "Bairro",
		Bairro:         "Azul",
		CodigoCidade:   "3550308",
		Estado:         "SP",
		Cep:            "88036-280",
	}

	tomador := &plugnotas.Tomador{
		CpfCnpj:     "934.439.970-07",
		RazaoSocial: "Joao Pedro",
	}

	prestador := &plugnotas.Prestador{
		CpfCnpj:            "63.866.907/0001-58",
		InscricaoMunicipal: "11111111",
		RazaoSocial:        "TESTE LTDA",
		NomeFantasia:       "TESTE",
		Endereco:           endereco,
		SimplesNacional:    true,
	}
	servico := &plugnotas.Servico{
		Codigo:        "1.02",
		IDIntegracao:  "A001XT",
		Discriminacao: "Programação de software",
		Cnae:          "4751201",
		Iss: &plugnotas.Iss{
			Aliquota: 3,
		},
		Valor: &plugnotas.Valor{
			Servico: 10,
		},
	}

	nfse := &plugnotas.Nfse{
		Tomador:   tomador,
		Prestador: prestador,
		Servico:   servico,
	}
	var list = []*plugnotas.Nfse{nfse}

	result, err := client.CreateNfse(list)
	if err != nil {
		t.Errorf("TestCreateNfse:%#v", err.Error)
	}
	assert.Equal(t, len(result.Documents), 1, "return array should is 1")
	assert.NotEmpty(t, result.Protocol, "Protocol can't be empty")

}

func TestGetNfseByID(t *testing.T) {
	godotenv.Load("./.env")
	id := "5d5ad649c7de482dee17451c"
	client := plugnotas.NewClient(os.Getenv("API_KEY"))
	result, err := client.GetNfseByID(id)
	if err != nil {
		t.Errorf("TestCreateNfse:%#v", err.Error)
	}
	assert.Equal(t, result.ID, id, "return id should is id")
}

func TestConsultarNfse(t *testing.T) {
	godotenv.Load("./.env")
	id := "5d5ad649c7de482dee17451c"
	client := plugnotas.NewClient(os.Getenv("API_KEY"))
	result, err := client.ConsultarNfse(id)
	if err != nil {
		t.Errorf("TestConsultarNfse:%#v", err.Error)
	}
	assert.Equal(t, len(result), 1, "return array should is 1")
}

func TestCancelarNfse(t *testing.T) {
	godotenv.Load("./.env")
	id := "5d5ad649c7de482dee17451c"
	client := plugnotas.NewClient(os.Getenv("API_KEY"))
	_, err := client.CancelarNfse(id)
	if err != nil {
		t.Errorf("TestCancelarNfse:%#v", err.Error)
	}
}
