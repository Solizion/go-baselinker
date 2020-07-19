package baselinker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetJournalListResponse struct {
	*BaseResponse
	Logs []Log `json:"logs"`
}

type GetJournalListParameters struct {
	OrderId int   `json:"order_id,omitempty"`
	LastId  int   `json:"last_log_id,omitempty"`
	Types   []int `json:"logs_types,omitempty"`
}

// Documentation: https://api.baselinker.com/index.php?method=getJournalList
func (baseLiner *BaseLinker) GetJournal(parameters GetJournalListParameters) ([]Log, Error) {
	var (
		response GetJournalListResponse
	)

	if !isProvidedAtLeastOneParam(parameters) {
		return response.Logs, NewSimpleError(
			fmt.Errorf(
				"Method %s is requiring at least one variable to be not empty",
				"'getJournalList'",
			),
		)
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

func isProvidedAtLeastOneParam(parameters GetJournalListParameters) bool {
	return 0 != parameters.OrderId || 0 != parameters.LastId || 0 != len(parameters.Types)
}
