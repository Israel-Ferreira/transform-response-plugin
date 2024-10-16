package models

type CountriesApiResponse struct {
	FlagImgUrl  string   `json:"flag_img_url"`
	CountryName string   `json:"name"`
	Capital     []string `json:"capital"`
	Languages   []string `json:"languages"`
}

func NewResponse(flagImageUrl string, countryName string, capital []string, languages []string) CountriesApiResponse {
	return CountriesApiResponse{
		FlagImgUrl:  flagImageUrl,
		CountryName: countryName,
		Capital:     capital,
		Languages:   languages,
	}
}

func ConvertServiceResponse(respItem map[string]any) CountriesApiResponse {
	flagImgUrl := respItem["flags"].(map[string]interface{})["png"].(string)
	countryName := respItem["name"].(map[string]interface{})["common"].(string)
	capital := respItem["capital"].([]string)

	languages_map := respItem["languages"].(map[string]any)

	var languages []string

	for _, v := range languages_map {
		languages = append(languages, v.(string))
	}

	return NewResponse(flagImgUrl, countryName, capital, languages)
}
