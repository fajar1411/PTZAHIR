package helper

import (
	"net/http"
	"strings"
)

type Responsive struct {
	Status  string `json:"status"`
	Massage string `json:"masssage"`
	Data    any    `json:"data"`
}
type ResponsivePage struct {
	Status  string `json:"status"`
	Massage string `json:"masssage"`
	Data    any    `json:"data"`
	Page    int    `json:"page"`
}

type ResponsFail struct {
	Status  string `json:"status"`
	Massage string `json:"masssage"`
}

func PesanSuksesHelper(msg string) map[string]any {
	return map[string]any{
		"Status": "Berhasil",
		"MSG":    msg,
	}
}

func PesanDataBerhasilHelper(data Responsive) map[string]any {

	respon := map[string]any{
		"status":  data.Status,
		"massage": data.Massage,
		"data":    data.Data,
	}

	return respon
}
func FailedResponse(msg string) map[string]any {
	return map[string]any{

		"message": msg,
	}
}

func PesanGagalHelper(msg string) (int, map[string]any) {
	var code int
	resp := map[string]interface{}{}
	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "please upload the"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "duplicated"):
		code = http.StatusConflict
	case strings.Contains(msg, "syntax"):
		code = http.StatusNotFound
		resp["message"] = "not found"
	case strings.Contains(msg, "input invalid"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "input value"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validation"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		resp["message"] = "failed to unmarshal json"
		code = http.StatusBadRequest
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "denied"):
		code = http.StatusUnauthorized
	case strings.Contains(msg, "jwt"):
		msg = "access is denied due to invalid credential"
		code = http.StatusUnauthorized
	case strings.Contains(msg, "Unauthorized"):
		code = http.StatusUnauthorized
	case strings.Contains(msg, "empty"):
		code = http.StatusBadRequest
	}

	return code, resp
}

type PaginationResponse struct {
	Page        int `json:"page"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
}

type WithPagination struct {
	Pagination PaginationResponse `json:"pagination"`
	Data       interface{}        `json:"data"`
	Message    string             `json:"message"`
}
