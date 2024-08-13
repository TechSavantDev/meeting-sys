package route

type MethodType uint8

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
)
