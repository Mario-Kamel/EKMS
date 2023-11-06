package cerrors

type InvalidIDError struct {
	Method  string
	Service string
	Err     error
}

func (e *InvalidIDError) Error() string {
	return e.Err.Error()
}

func (e *InvalidIDError) Unwrap() error {
	return e.Err
}

func (e *InvalidIDError) Log() string {
	return e.Service + " " + e.Method + ": " + e.Error()
}

func NewInvalidIDError(method, service string, err error) *InvalidIDError {
	return &InvalidIDError{
		Method:  method,
		Service: service,
		Err:     err,
	}
}

type InternalServerError struct {
	Method  string
	Service string
	Err     error
}

func (e *InternalServerError) Error() string {
	return e.Err.Error()
}

func (e *InternalServerError) Unwrap() error {
	return e.Err
}

func (e *InternalServerError) Log() string {
	return e.Service + " " + e.Method + ": " + e.Error()
}

func NewInternalServerError(method, service string, err error) *InternalServerError {
	return &InternalServerError{
		Method:  method,
		Service: service,
		Err:     err,
	}
}
