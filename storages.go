package baselinker

import (
	"encoding/json"
	"net/http"
)

type GetStoragesListResponse struct {
	*BaseResponse
	Storages []Storage `json:"storages"`
}

// Documentation: https://api.baselinker.com/index.php?method=getStoragesList
func (baseLiner *BaseLinker) GetStorages() ([]Storage, Error) {
	var (
		response GetStoragesListResponse
	)

	requestForm := baseLiner.createRequestForm("getStoragesList", nil)
	resp, err := http.PostForm(baseLiner.Url, requestForm)

	if nil != err {
		return response.Storages, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Storages, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Storages, response
	}

	return response.Storages, nil
}
