package helper

type Responsive struct {
	Status  string `json:"status"`
	Massage string `json:"masssage"`
	Data    any    `json:"data"`
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
		"status":  "failed",
		"message": msg,
	}
}
