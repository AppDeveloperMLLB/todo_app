package apperrors

// AppError - エラー情報
type AppError struct {
	ErrCode `json:"err_code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (err *AppError) Error() string {
	return err.Err.Error()
}

func (err *AppError) Unwrap() error {
	return err.Err
}
