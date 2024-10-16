package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/transform-response-plugin/pkg/models"
	"github.com/Kong/go-pdk"
)

type TransformResponsePlugin struct{}

func (tr TransformResponsePlugin) Response(kong *pdk.PDK) {
	body, err := kong.ServiceResponse.GetRawBody()

	var bodyResponse []map[string]any

	if err != nil {
		errorMessage := map[string]string{
			"msg": "Erro ao capturar a resposta do serviço",
		}

		jsonMsg, _ := json.Marshal(errorMessage)

		responseHeader, _ := kong.Response.GetHeaders(10)

		kong.Log.Notice("Erro ao decodificar a resposta")

		kong.Response.Exit(400, jsonMsg, responseHeader)
	}

	if err := json.Unmarshal(body, &bodyResponse); err != nil {
		errorMessage := map[string]string{
			"msg": "Erro ao capturar a resposta do serviço",
		}

		jsonMsg, _ := json.Marshal(errorMessage)

		responseHeader, _ := kong.Response.GetHeaders(10)

		kong.Log.Notice("Erro ao decodificar a resposta: %v", err)

		kong.Response.Exit(400, jsonMsg, responseHeader)
		return
	}

	var countries []models.CountriesApiResponse

	for _, respItem := range bodyResponse {

		response := models.ConvertServiceResponse(respItem)

		countries = append(countries, response)
	}

	defer func() {
		if r := recover(); r != nil {
			kong.Log.Err("Recovered. Error: ", r)
			errorMessage := map[string]string{
				"msg": fmt.Sprintf("%s", r),
			}
			jsonMsg, _ := json.Marshal(errorMessage)

			kong.Response.Exit(500, jsonMsg, nil)
		}
	}()

	marshallResponse, err := json.Marshal(&countries)

	if err != nil {
		kong.Response.Exit(500, nil, nil)
		return
	}

	headers, _ := kong.Response.GetHeaders(10)

	kong.Log.Notice("Response: %s", string(marshallResponse))

	kong.Response.Exit(http.StatusOK, marshallResponse, headers)

}

func New() any {
	return &TransformResponsePlugin{}
}
