package baselinker

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	JournalLogTypeOrderCreate                  = 1
	JournalLogTypeOrderDownloadFod             = 2
	JournalLogTypeOrderPaid                    = 3
	JournalLogTypeOrderOrPaymentDocumentDelete = 4
	JournalLogTypeOrdersMerge                  = 5
	JournalLogTypeOrdersSeparate               = 6
	JournalLogTypeOrderToInvoice               = 7
	JournalLogTypeOrderToRecipt                = 8
	JournalLogTypeShippmentCreate              = 9
	JournalLogTypeShippmentDelete              = 10
	JournalLogTypeShippmentEdit                = 11
	JournalLogTypeOrderAddProduct              = 12
	JournalLogTypeOrderEditProduct             = 13
	JournalLogTypeOrderDeleteProduct           = 14
	JournalLogTypeBuyerToBlackList             = 15
	JournalLogTypeOrderEdit                    = 16
	JournalLogTypeOrderCopy                    = 17
	JournalLogTypeOrderStatusChange            = 18
	JournalLogTypeInvoiceDelete                = 19
	JournalLogTypeReciptDelete                 = 20
)

var journalLogTypes = intArray{
	JournalLogTypeOrderCreate,
	JournalLogTypeOrderDownloadFod,
	JournalLogTypeOrderPaid,
	JournalLogTypeOrderOrPaymentDocumentDelete,
	JournalLogTypeOrdersMerge,
	JournalLogTypeOrdersSeparate,
	JournalLogTypeOrderToInvoice,
	JournalLogTypeOrderToRecipt,
	JournalLogTypeShippmentCreate,
	JournalLogTypeShippmentDelete,
	JournalLogTypeShippmentEdit,
	JournalLogTypeOrderAddProduct,
	JournalLogTypeOrderEditProduct,
	JournalLogTypeOrderDeleteProduct,
	JournalLogTypeBuyerToBlackList,
	JournalLogTypeOrderEdit,
	JournalLogTypeOrderCopy,
	JournalLogTypeOrderStatusChange,
	JournalLogTypeInvoiceDelete,
	JournalLogTypeReciptDelete,
}

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
