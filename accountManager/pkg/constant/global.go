package constant

type JWTClaim string
type StatusTransaction string
type PaymentAccount string

const (
	DefaultPage    = 1
	DefaultPerPage = 10
)

const (
	ASC  = "asc"
	DESC = "desc"
)

const (
	USERIDKEY JWTClaim = "user_id"
	ROLEKEY   JWTClaim = "role"
)

const (
	Debit   PaymentAccount = "debit"
	Credit  PaymentAccount = "credit"
	Loan    PaymentAccount = "loan"
	Default PaymentAccount = "default"
)

const (
	Pending StatusTransaction = "pending"
	Success StatusTransaction = "success"
	Failed  StatusTransaction = "failed"
)
