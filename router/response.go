package router

const (
	ResponseStatusOK                  = 200
	ResponseStatusNotFound            = 404
	ResponseStatusInternalServerError = 500
)

const (
	ResponseOkMessage                  = "OK"
	ResponseNotFoundMessage            = "Not Found"
	ResponseInternalServerErrorMessage = "Internal Server Error"
)

// Response is the response object for the API.
type Response struct {
	// Status is the status code of the response.
	Status int `json:"status"`
	// Message is the message of the response.
	Message string `json:"message"`
	// Data is the data of the response.
	Data interface{} `json:"data"`
}
