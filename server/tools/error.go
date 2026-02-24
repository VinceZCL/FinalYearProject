package tools

type Xerror struct {
	Code    int
	Msg     string
	Details string
}

func (e *Xerror) Error() string {
	return e.Details
}

func new(code int, msg string, details string) *Xerror {
	return &Xerror{
		Code:    code,
		Msg:     msg,
		Details: details,
	}
}

func ErrBadRequest(details string) *Xerror {
	return new(400, "bad request", details)
}

func ErrUnauthorized(details string) *Xerror {
	return new(401, "authorization", details)
}

func ErrForbidden(details string) *Xerror {
	return new(403, "invalid operation", details)
}

func ErrNotFound(msg string, details string) *Xerror {
	return new(404, msg, details)
}

func ErrInternal(msg string, details string) *Xerror {
	return new(500, msg, details)
}
