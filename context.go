package contexterr

type ContextError struct {
	message string
	fields  Fields
}

func New(message string, fields Fields) ContextError {
	e := ContextError{}
	e.message = message
	e.fields = fields
	return e
}

func From(ce ContextError, fields Fields) {
	e := ContextError{}
	e.message = ce.message
	e.fields = ce.fields
	for k, v := range fields {
		e.fields[k] = v
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
