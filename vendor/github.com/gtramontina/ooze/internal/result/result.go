package result

type Result[Type any] interface {
	seal() string
	String() string
	IsOk() bool
}

func Ok[Type any](value Type) Result[Type] {
	return ok[Type]{value}
}

func Err[Type any](errorMessage string) Result[Type] {
	return err[Type]{errorMessage}
}
