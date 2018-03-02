package contexterr

type ContextError struct {
	message string
	fields  Fields
}

func New(message string) ContextError {
	e := ContextError{}
	e.message = message
	e.fields = make(Fields)
	return e
}

func WithFields(message string, fields Fields) ContextError {
	e := New(message)
	if fields != nil {
		e.fields = fields
	}
	return e
}

func FromError(e error, f Fields) {
	ce := ContextError{}
	ce.message = e.Error()

	if _, ok := e.(FieldError); ok {
		ce.fields = e.(FieldError).Fields()
	}

	// Merge the provided fields into this
	for k, v := range f {
		ce.fields[k] = v
	}
}

func (e ContextError) Error() string {
	return e.message
}

func (e ContextError) Fields() Fields {
	return e.fields
}

func (e ContextError) WithFields(fields Fields) FieldError {
	for k, v := range fields {
		e.fields[k] = v
	}
	return e
}
