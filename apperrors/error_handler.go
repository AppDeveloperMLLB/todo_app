package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/AppDeveloperMLLB/todo_app/common"
)

// AppError - エラー情報
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceId := common.GetTraceId(req.Context())
	log.Printf("[%d]error: %s\n", traceId, appErr)

	var statusCode int

	switch appErr.ErrCode {
	// 指定されたデータがない場合
	case NAData:
		statusCode = http.StatusNotFound
	// 更新対象のデータがない場合
	// JSONでコードに失敗した場合
	// クエリ、パスパラメータが不正の場合
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case Unauthorized, InvalidHeader, CreateValidatorFailed:
		statusCode = http.StatusUnauthorized
	case NotMatchUser, Forbidden:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
