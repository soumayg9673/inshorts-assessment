package errors

type ErrorMdl struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Status int      `json:"status"`
	Error  ErrorMdl `json:"error"`
}
