package code

type SliceError struct {
	detail string
	code   SliceCode
}

func NewSliceError(sc SliceCode) SliceError {
	return SliceError{
		detail: sc.String(),
		code:   sc,
	}
}

func NewSliceErrorWithMsg(sc SliceCode, msg string) SliceError {
	return SliceError{
		detail: msg,
		code:   sc,
	}

}

// Error Error()
func (s SliceError) Error() string {
	return s.detail
}
