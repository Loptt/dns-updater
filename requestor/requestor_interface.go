package requestor

// RequestorInterface provides an abstract interface to requestor object.
// Requestors implement the business logic to connect to a remote service
// and post an action to them.
type RequestorInterface interface {
	Request(string) (string, error)
}
