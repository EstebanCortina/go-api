package exceptions

type ServerException struct {
	Name string
	Body string
	Err  error
}
