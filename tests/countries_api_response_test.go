package tests

import (
	"testing"

	"github.com/Israel-Ferreira/transform-response-plugin/pkg/models"
)

func TestConvertServiceResponse(t *testing.T) {
	respItem := map[string]any{
		"flags": map[string]any{
			"png": "https://flagcdn.com/w320/sl.png",
			"svg": "https://flagcdn.com/sl.svg",
			"alt": "The flag of Sierra Leone is composed of three equal horizontal bands of green, white and blue.",
		},
		"name": map[string]any{
			"common":   "Sierra Leone",
			"official": "Republic of Sierra Leone",
			"nativeName": map[string]any{
				"eng": map[string]any{
					"official": "Republic of Sierra Leone",
					"common":   "Sierra Leone",
				},
			},
		},
		"capital": []string{"Freetown"},
		"languages": map[string]any{
			"eng": "English",
		},
	}

	result := models.ConvertServiceResponse(respItem)

	if len(result.Capital) == 0 && result.CountryName == "" {
		t.Errorf("Falha: O Valor não foi preenchido")
	}

}
