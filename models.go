package baselinker

type intArray []int

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
