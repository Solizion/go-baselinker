package baselinker

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/url"
)

type BaseLinker struct {
	Url       string
	Token     string
	validator *validator.Validate
}

type BaseResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}

func (baseLinkerResponse *BaseResponse) IsSuccess() bool {
	return baseLinkerResponse.Status == "SUCCESS"
}

func (baseLinkerResponse *BaseResponse) Error() string {
	return baseLinkerResponse.ErrorMessage
}

func (baseLinkerResponse *BaseResponse) CodeError() string {
	return baseLinkerResponse.ErrorCode
}

func NewBaseLinker(url string, token string) *BaseLinker {
	baseLinker := &BaseLinker{Url: url, Token: token, validator: validator.New()}
	baseLinker.registerValidationMethods()

	return baseLinker
}

func (baseLiner *BaseLinker) createRequestForm(method string, parameters interface{}) url.Values {
	parametersJson, err := json.Marshal(parameters)

	if nil != err {
		return url.Values{}
	}

	formData := url.Values{
		"token":      {baseLiner.Token},
		"method":     {method},
		"parameters": {string(parametersJson)},
	}

	return formData
}
