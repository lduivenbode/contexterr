package contexterr

type Fields map[string]interface{}

type FieldError interface {
	Error() string
	Fields() Fields
	WithFields(fields Fields) FieldError
}
