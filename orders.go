package baselinker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetOrdersListResponse struct {
	*BaseResponse
	Orders []Order `json:"orders"`
}

type GetOrdersListParameters struct {
	Unconfirmed bool `json:"get_unconfirmed_orders"`
	ConfirmedOn int  `json:"date_confirmed_from,omitempty"`
	FromOn      int  `json:"date_from,omitempty"`
	FromId      int  `json:"id_from,omitempty"`
	OrderId     int  `json:"order_id,omitempty"`
	StatusId    int  `json:"status_id,omitempty"`
}

// Documentation: https://api.baselinker.com/index.php?method=getOrders
func (baseLiner *BaseLinker) GetOrders(parameters GetOrdersListParameters) ([]Order, Error) {
	var (
		response GetOrdersListResponse
		orders   []Order
	)

	formData := baseLiner.createRequestForm("getOrders", parameters)
	resp, err := http.PostForm(baseLiner.Url, formData)

	if nil != err {
		return orders, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return orders, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return orders, response
	}

	return orders, nil
}

func (baseLiner *BaseLinker) GetOrder(orderId int) (Order, error) {
	orders, err := baseLiner.GetOrders(GetOrdersListParameters{OrderId: orderId})
	if nil != err {
		return Order{}, err
	}

	if len(orders) == 0 {
		return Order{}, NewSimpleError(fmt.Errorf("Order %d not found", orderId))
	}

	return orders[0], nil
}
