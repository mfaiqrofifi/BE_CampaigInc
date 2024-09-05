package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"messages"`
	Data    interface{} `json:"data"`
}
