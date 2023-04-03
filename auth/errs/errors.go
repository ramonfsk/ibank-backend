package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e *AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
