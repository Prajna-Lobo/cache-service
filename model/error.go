package model

type Error struct {
	Code    string `json:"ErrorCode"`
	Message string `json:"Errormessage"`
}

func GetBadRequestError() Error {
	return Error{
		Code:    "ERR_BAD_REQUEST",
		Message: "bad request",
	}
}

func GetInternalServerError() Error {
	return Error{
		Code:    "ERR_INTERNAL_SERVER",
		Message: "internal server error",
	}
}
