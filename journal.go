package baselinker

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type GetJournalListResponse struct {
	*BaseResponse
	Logs []Log `json:"logs"`
}

type GetJournalListParameters struct {
	OrderId int   `json:"order_id,omitempty"`
	LastId  int   `json:"last_log_id,omitempty"`
	Types   []int `json:"logs_types,omitempty" validate:"is-journal-log-types"`
}

// Documentation: https://api.baselinker.com/index.php?method=getJournalList
func (baseLiner *BaseLinker) GetJournal(parameters GetJournalListParameters) ([]Log, Error) {
	var (
		response GetJournalListResponse
	)

	err := baseLiner.validator.Struct(parameters)
	if nil != err {
		return response.Logs, NewSimpleError(err)
	}

	requestForm := baseLiner.createRequestForm("getJournalList", parameters)
	resp, err := http.PostForm(baseLiner.Url, requestForm)

	if nil != err {
		return response.Logs, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Logs, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Logs, response
	}

	return response.Logs, nil
}

// TODO check if this validation works with nil value (it must return true)
func validateJournalTypes(field validator.FieldLevel) bool {
	slice, ok := field.Field().Interface().([]int)
	if !ok {
		return false
	}

	for _, logType := range slice {
		if !journalLogTypes.has(logType) {
			return false
		}
	}

	return true
}
