package errors

type ErrorMdl struct {
	Code    string `json:"code"`
	Message any    `json:"message"`
}

type Error struct {
	Status int      `json:"status"`
	Error  ErrorMdl `json:"error"`
}
