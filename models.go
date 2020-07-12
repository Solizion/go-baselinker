package baselinker

const (
	LogTypeCreateNewOrder    = 1
	LogTypeChangeOrderStatus = 18
)

type Product struct {
	Id        string `json:"product_id"`
	Sku       string `json:"sku"`
	Quantity  int    `json:"quantity"`
	Storage   string `json:"storage"`
	StorageId string `json:"storage_id"`
}

type Order struct {
	Id       int       `json:"order_id"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Products []Product `json:"products"`

	// Invoice
	InvoiceNip         string `json:"invoice_nip"`
	InvoiceClientName  string `json:"invoice_fullname"`
	InvoiceCompanyName string `json:"invoice_company"`

	// Client
	ClientName        string `json:"delivery_fullname"`
	ClientCompanyName string `json:"delivery_company"`
	ClientCity        string `json:"delivery_city"`
	ClientStreet      string `json:"delivery_address"`
	ClientPostalCode  string `json:"delivery_postcode"`
}

type Log struct {
	Id       int `json:"log_id"`
	Type     int `json:"log_type"`
	OrderId  int `json:"order_id"`
	ObjectId int `json:"object_id"`
	Date     int `json:"date"`
}

func (order *Order) HasInvoiceNip() bool {
	return len(order.InvoiceNip) > 10
}

func (order *Order) GetClientFullName() string {
	// From Invoice
	if len(order.InvoiceCompanyName) > 0 {
		return order.InvoiceCompanyName
	} else if len(order.InvoiceClientName) > 0 {
		return order.InvoiceClientName
	}

	// From Client
	if len(order.ClientCompanyName) > 0 {
		return order.ClientCompanyName
	} else if len(order.ClientName) > 0 {
		return order.ClientName
	}

	return ""
}
