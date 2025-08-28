package response

type BadRequestError struct {
	Error string `json:"error" example:"Bad Request"`
}

type NotFoundError struct {
	Error string `json:"error" example:"Not Found"`
}

type InternalServerError struct {
	Error string `json:"error" example:"Internal Server Error"`
}
