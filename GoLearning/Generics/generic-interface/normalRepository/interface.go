package normalRepository

type Requester interface {
	GetResponse() Response
	MockResponse() Response
}
