package httperrors

type CustomError struct {
	Text   string `json:"messge"`
	Status int    `json:"code"`
}

func (c CustomError) Error() string {
	return c.Text
}

type ClientSideError struct {
	CustomError
}

type ServerSideError struct {
	CustomError
}

type BadRequestError struct {
	ClientSideError
}

type UnauthorizedError struct {
	ClientSideError
}

type NotFoundError struct {
	ClientSideError
}

type MethodNotAllowedError struct {
	ClientSideError
}

type InternalServerError struct {
	ServerSideError
}

func NewBadRequestError(text string) BadRequestError {
	return BadRequestError{
		ClientSideError: ClientSideError{
			CustomError: CustomError{
				Status: 400,
				Text:   text,
			},
		},
	}
}

func NewUnauthorizedError(text string) UnauthorizedError {
	return UnauthorizedError{
		ClientSideError: ClientSideError{
			CustomError: CustomError{
				Status: 401,
				Text:   text,
			},
		},
	}
}

func NewNotFoundError(text string) NotFoundError {
	return NotFoundError{
		ClientSideError: ClientSideError{
			CustomError: CustomError{
				Status: 404,
				Text:   text,
			},
		},
	}
}

func NewMethodNotAllowedError(text string) MethodNotAllowedError {
	return MethodNotAllowedError{
		ClientSideError: ClientSideError{
			CustomError: CustomError{
				Status: 405,
				Text:   text,
			},
		},
	}
}

func NewInternalServerError(text string) InternalServerError {
	return InternalServerError{
		ServerSideError: ServerSideError{
			CustomError: CustomError{
				Status: 500,
				Text:   text,
			},
		},
	}
}
