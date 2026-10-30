package booklooker

type MediaType int

const (
	Books      MediaType = 0
	Movies     MediaType = 1
	Music      MediaType = 2
	AudioBooks MediaType = 3
	Games      MediaType = 4
)

// DataType is the kind of an upload of an article file.
type DataType int

const (
	// DataUpdate adds, changes and deletes the articles of the file.
	DataUpdate DataType = 0
	// DataReplace replaces all articles with the articles of the file.
	DataReplace DataType = 1
	// DataDelete deletes the articles of the file.
	DataDelete DataType = 2
)

type FileType string

const (
	ArticleFile FileType = "article"
	// ImageFile has to be a ZIP archive.
	ImageFile FileType = "pic"
)

// Encoding of an article file, without one the encoding of the last upload is used and CP-1252 on the first.
type Encoding string

const (
	CP437    Encoding = "IBMPC/CR"
	MacRoman Encoding = "macintosh"
	UTF8     Encoding = "UTF-8"
)

// ArticleField is returned by ArticleList, ISBN and EAN require a media type.
type ArticleField string

const (
	FieldOrderNo ArticleField = "orderNo"
	FieldISBN    ArticleField = "isbn"
	FieldEAN     ArticleField = "ean"
)

type ArticleStatus string

const (
	ArticleActive  ArticleStatus = "ACTIVE"
	ArticleSold    ArticleStatus = "SOLD"
	ArticleDeleted ArticleStatus = "DELETED"
)

type FileStatus string

const (
	FileReceived     FileStatus = "FILE_RECEIVED"
	UploadInProgress FileStatus = "UPLOAD_IN_PROGRESS"
	UploadDone       FileStatus = "UPLOAD_DONE"
)

type MessageType string

const (
	PaymentInformation MessageType = "PAYMENT_INFORMATION"
	PaymentReminder    MessageType = "PAYMENT_REMINDER"
	// ShippingNotice can only be sent once per order.
	ShippingNotice MessageType = "SHIPPING_NOTICE"
)

// OrderStatus only organizes the orders, CancelOrder and CancelOrderItem are needed to cancel without commission.
type OrderStatus string

// The statuses up to Canceled belong to sales, the following ones to purchases, Canceled to both.
const (
	WaitingForPayment        OrderStatus = "WAITING_FOR_PAYMENT"
	ReadyForShipment         OrderStatus = "READY_FOR_SHIPMENT"
	ShippedWaitingForPayment OrderStatus = "SHIPPED_WAITING_FOR_PAYMENT"
	ShippedAndPaid           OrderStatus = "SHIPPED_AND_PAID"
	BuyerNoReaction          OrderStatus = "BUYER_NO_REACTION"
	Canceled                 OrderStatus = "CANCELED"
	ToBePaid                 OrderStatus = "TO_BE_PAID"
	WaitingForPaymentInfo    OrderStatus = "WAITING_FOR_PAYMENT_INFO"
	WaitingForShipment       OrderStatus = "WAITING_FOR_SHIPMENT"
	PaidWaitingForShipment   OrderStatus = "PAID_WAITING_FOR_SHIPMENT"
	ReceivedAndPaid          OrderStatus = "RECEIVED_AND_PAID"
	VendorNoReaction         OrderStatus = "VENDOR_NO_REACTION"
)

type PaymentMethod int

const (
	BankTransfer            PaymentMethod = 1
	Invoice                 PaymentMethod = 2
	DirectDebit             PaymentMethod = 3
	CreditCard              PaymentMethod = 4
	CashOnDelivery          PaymentMethod = 5
	PayPal                  PaymentMethod = 6
	Skrill                  PaymentMethod = 8
	PickupAndCash           PaymentMethod = 9
	InvoicePrepaymentOption PaymentMethod = 11
	PayPalREST              PaymentMethod = 12
)

type Rating int

const (
	// RatingNeverAgain and RatingProblematic require a comment of at least 15 letters.
	RatingNeverAgain  Rating = 1
	RatingProblematic Rating = 2
	RatingOK          Rating = 3
	RatingPerfect     Rating = 4
)

// Account is the seller account of the API key.
type Account struct {
	SellerID    int    `json:"sellerId"`
	Username    string `json:"username"`
	ReadyToSell bool   `json:"readyToSell"`
	// Missing lists why the account cannot sell, e.g. ACCOUNT_NOT_VERIFIED or SHIPPING_COSTS_MISSING.
	Missing []string `json:"missing"`
}
