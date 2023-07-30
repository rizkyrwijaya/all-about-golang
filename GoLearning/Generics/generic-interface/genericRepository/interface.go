package genericrepository

type Requester[T any] interface {
	GetResponse() T
	MockResponse() T
}
