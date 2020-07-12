package baselinker

import (
	"encoding/json"
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
		logs     []Log
	)

	requestForm := baseLiner.createRequestForm("getJournalList", parameters)
	resp, err := http.PostForm(baseLiner.Url, requestForm)

	if nil != err {
		return logs, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return logs, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return logs, response
	}

	return logs, nil
}
