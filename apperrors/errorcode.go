package apperrors

type ErrCode string

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}

const (
	// Unknown - 予期せぬエラー
	Unknown ErrCode = "U000"
	// InsertDataField - データの挿入に失敗
	InsertDataField ErrCode = "S001"
	// GetDataFailed - データの取得に失敗
	GetDataFailed ErrCode = "S002"
	// NAData - データが存在しない
	NAData ErrCode = "S003"
	// NoTargetData - 対象データが存在しない
	NoTargetData ErrCode = "S004"
	// UpdateDataFailed - データの更新に失敗
	UpdateDataFailed ErrCode = "S005"
	// DeleteDataFailed - データの削除に失敗
	DeleteDataFailed ErrCode = "S006"
	// Forbidden - 権限がない
	Forbidden ErrCode = "S007"
	// ReqBodyDecodeFailed - リクエストボディのデコードに失敗
	ReqBodyDecodeFailed ErrCode = "R001"
	// BadParam - パラメータが不正
	BadParam ErrCode = "R002"
	// Unauthorized - 認証エラー
	Unauthorized ErrCode = "A001"
	// InvalidHeader - ヘッダが不正
	InvalidHeader ErrCode = "A002"
	// CreateValidatorFailed - Validatorの作成に失敗
	CreateValidatorFailed   ErrCode = "A003"
	NotMatchUser            ErrCode = "A004"
	InvalidOauthState       ErrCode = "A005"
	OauthConfExchangeFailed ErrCode = "A006"
	NoIdToken               ErrCode = "A007"
	ClientGetFailed         ErrCode = "A008"
)
