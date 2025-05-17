package errors

func SomethingWentWrong() Error {
	return Error{
		Status: 500,
		Error: ErrorMdl{
			Code:    "something_went_wrong",
			Message: "Please try again after sometime.",
		},
	}
}

func BadRequst() Error {
	return Error{
		Status: 400,
		Error: ErrorMdl{
			Code:    "bad_request",
			Message: "The request could not be processed due to invalid input. Please check your request parameters and try again.",
		},
	}
}

func BadRequstMessage(msg any) Error {
	return Error{
		Status: 400,
		Error: ErrorMdl{
			Code:    "bad_request",
			Message: msg,
		},
	}
}
