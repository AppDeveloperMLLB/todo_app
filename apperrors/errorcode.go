package apperrors

type ErrCode string

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}

const (
	// Unknown - 予期せぬエラー
	Unknown                 ErrCode = "U000"
	InsertDataField         ErrCode = "S001"
	GetDataFailed           ErrCode = "S002"
	NAData                  ErrCode = "S003"
	NoTargetData            ErrCode = "S0004"
	UpdateDataFailed        ErrCode = "S005"
	ReqBodyDecodeFailed     ErrCode = "R001"
	BadParam                ErrCode = "R002"
	Unauthorized            ErrCode = "A001"
	InvalidHeader           ErrCode = "A002"
	CreateValidatorFailed   ErrCode = "A003"
	NotMatchUser            ErrCode = "A004"
	InvalidOauthState       ErrCode = "A005"
	OauthConfExchangeFailed ErrCode = "A006"
	NoIdToken               ErrCode = "A007"
	ClientGetFailed         ErrCode = "A008"
)
