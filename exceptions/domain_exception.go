package exceptions

type DomainException struct {
	Message string
}

func (e *DomainException) Error() string {
	return e.Message
}
