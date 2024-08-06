package apperrors

type AppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (err *AppError) Error() string {
	return err.Err.Error()
}

func (err *AppError) Unwrap() error {
	return err.Err
}
