package code

var (
	Success = &Err{
		Code:    200,
		Message: "success",
	}
	ERROR = &Err{
		Code:    500,
		Message: "Error",
	}
)

type Err struct {
	Code    int
	Message string
}

func (err Err) Error() string {
	return err.Message
}

func DecodeErr(err *Err) (int, string) {
	if err == nil {
		return Success.Code, Success.Message
	}

	return err.Code, err.Message
}
