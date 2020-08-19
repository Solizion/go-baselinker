package baselinker

import (
	"encoding/json"
	"net/http"
)

type GetCategoriesListResponse struct {
	*BaseResponse
	Categories []Category `json:"categories"`
}

type GetCategoriesListParameters struct {
	StorageId string `json:"sotrage_id" validate:"required"`
}

// Documentation: https://api.baselinker.com/index.php?method=getCGetCategories
func (baseLinker *BaseLinker) GetCategories(parameters GetCategoriesListParameters) ([]Category, Error) {
	var (
		response GetCategoriesListResponse
	)

	err := baseLinker.validator.Struct(parameters)
	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	requestForm := baseLinker.createRequestForm("getJournalList", parameters)
	resp, err := http.PostForm(baseLinker.Url, requestForm)

	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Categories, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Categories, response
	}

	return response.Categories, nil
}
