package apperror

type FieldError struct {
	Field 	string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	Code 			uint		  `json:"code"` // Strict checking only for devs
	Message 	 	string		  `json:"message"`
	FieldErrors		[]FieldError  `json:"field_errors"`
	originalerror 	error

}

func New(code uint, err error) *AppError{
	return &AppError{
		Code: 	 	   code,
		Message: 	   err.Error(),
		originalerror: err,
	}
}

func (appError *AppError) Error() string {
	return appError.originalerror.Error()
}