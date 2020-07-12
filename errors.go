package baselinker

const (
	ErrorNoCode             = "ERROR_NO_CODE"
	ErrorCodeAccountBlocked = "ERROR_USER_ACCOUNT_BLOCKED"
)

type Error interface {
	Error() string
	CodeError() string
}

func NewSimpleError(err error) *BaseResponse {
	return &BaseResponse{
		Status:       "ERROR",
		ErrorMessage: err.Error(),
		ErrorCode:    ErrorNoCode,
	}
}
