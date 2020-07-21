package domain

//IdentifiableError represents an error identifiable by its Code
type IdentifiableError interface {
	GetCode() int
}

// GenericError holds the info about an error occurred during the application execution
type GenericError struct {
	Code int `json:"code"`

	Message string `json:"message"`
}

func (e GenericError) Error() string {
	return e.Message
}

//GetCode identifies an error by its Code
func (e GenericError) GetCode() int {
	return e.Code
}

//AlreadyExistsError represents an specialized Already Exists Error
type AlreadyExistsError struct {
	GenericError
}

//AlreadyExists builds an specialized Already Exists Error
func AlreadyExists(message string) AlreadyExistsError {
	alreadyExists := AlreadyExistsError{}
	alreadyExists.Code = 409
	alreadyExists.Message = message
	return alreadyExists
}

//InternalError builds an specialized Internal Error
func InternalError(message string) GenericError {
	internalError := GenericError{}
	internalError.Code = 500
	internalError.Message = message
	return internalError
}

//ConstraintViolationError represents an specialized Constraint Violation Error
type ConstraintViolationError struct {
	GenericError
}

//ConstraintViolation builds an specialized Constraint Violation Error
func ConstraintViolation(message string) ConstraintViolationError {
	constraintViolation := ConstraintViolationError{}
	constraintViolation.Code = 400
	constraintViolation.Message = message
	return constraintViolation
}

//NotFoundError represents an specialized Not Found Error
type NotFoundError struct {
	GenericError
}

//NotFound builds an specialized Not Found Error
func NotFound(message string) NotFoundError {
	notFound := NotFoundError{}
	notFound.Code = 404
	notFound.Message = message
	return notFound
}
