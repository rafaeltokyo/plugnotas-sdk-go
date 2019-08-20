package plugnotas

import (
	"encoding/json"
	"fmt"
)

// NfseResponse - Struct para definir o objeto NfseResponse
type NfseResponse struct {
	Documents []*Nfse `json:"documents"`
	Protocol  string  `json:"protocol"`
}

// Servico - Struct para definir o objeto serviço
type Servico struct {
	Codigo                    string `json:"codigo"`
	IDIntegracao              string `json:"idIntegracao"`
	Discriminacao             string `json:"discriminacao"`
	CodigoTributacao          string `json:"codigoTributacao"`
	Cnae                      string `json:"cnae"`
	CodigoCidadeIncidencia    string `json:"codigoCidadeIncidencia"`
	DescricaoCidadeIncidencia string `json:"descricaoCidadeIncidencia"`
	Iss                       *Iss   `json:"iss"`
	Valor                     *Valor `json:"valor"`
}

// Nfse - Struct para definir o objeto Nfse
type Nfse struct {
	IDIntegracao string     `json:"IdIntegracao"`
	EnviarEmail  bool       `json:"enviarEmail"`
	Prestador    *Prestador `json:"prestador"`
	Tomador      *Tomador   `json:"tomador"`
	Servico      *Servico   `json:"servico"`
	ID           string     `json:"id"`
}
type ResumoNfse struct {
}

// CreateNfse - enviar uma lista de notas
func (plugnotas *Client) CreateNfse(req []*Nfse) (*NfseResponse, *ErrorResponse) {

	data, err := json.Marshal(req)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Error{
				Message: err.Error(),
			},
		}
	}
	var result = &NfseResponse{}
	err, errAPI := plugnotas.Request("POST", "/nfse", data, result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Error{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}

// GetNfseByID buscar nota por id
func (plugnotas *Client) GetNfseByID(id string) (*Nfse, *ErrorResponse) {
	var result = &Nfse{}
	err, errAPI := plugnotas.Request("GET", fmt.Sprintf("/nfse/%s", id), nil, result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Error{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}

// GetNfseByID buscar nota por id
func (plugnotas *Client) ConsultarNfse(id string) ([]*ResumoNfse, *ErrorResponse) {
	var result []*ResumoNfse
	err, errAPI := plugnotas.Request("GET", fmt.Sprintf("/nfse/consultar/%s", id), nil, result)
	if err != nil {
		return nil, &ErrorResponse{
			Error: &Error{
				Message: err.Error(),
			},
		}
	}
	if errAPI != nil {
		return nil, errAPI
	}
	return result, nil

}
