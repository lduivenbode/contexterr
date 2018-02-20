package contexterr

type Fields map[string]interface{}

type FieldError interface {
	Fields() Fields
	WithFields(fields Fields) FieldError
}
