package contexterr

type ContextError struct {
	message string
	fields  Fields
}

func Errorf(message string, fields Fields) ContextError {
	e := ContextError{}
	e.message = message
	e.fields = fields
	return e
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
