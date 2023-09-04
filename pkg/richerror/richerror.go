package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation string
	warpError error
	message   string
	kind      Kind
	meta      map[string]interface{}
}

func New(operation string) RichError {
	return RichError{operation: operation}
}
func (r RichError) WhitWarpError(err error) RichError {
	r.warpError = err
	return r
}
func (r RichError) WhitMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WhitKind(kind Kind) RichError {
	r.kind = kind
	return r
}
func (r RichError) WhitMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}
func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}
	re, ok := r.warpError.(RichError)
	if !ok {
		return 0
	}
	return re.Kind()
}
func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}
	re, ok := r.warpError.(RichError)
	if !ok {
		return r.warpError.Error()
	}
	return re.Message()
}

func (r RichError) Error() string {
	return r.message

}

/*
func New(args ...interface{}) RichError {
	r:=RichError{}
	for _,arg := range args {
		switch arg.(type) {
		case Op:
			r.Operation = arg.(Op)
		case string:
			r.Message = arg.(string)
		case error:
			r.WarpError=arg.(error)
		case Kind:
			r.Kind = arg.(Kind)
		case map[string]interface{}:
			r.Meta=arg.(map[string]interface{})




		}
	}

}*/
