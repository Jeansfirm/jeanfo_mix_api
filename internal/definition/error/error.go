package error_definition

type BaseError struct {
	Code int
	Msg  string
}

type BadRequestError struct {
	BaseError
}

func (be *BaseError) Error() string {
	return be.Msg
}
