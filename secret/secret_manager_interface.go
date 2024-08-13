package secret

// SecretManagerInterface defines an abstract interface for APIs that allow to
// retrieve secrets from different srouces.
type SecretManagerInterface interface {
	Fetch(string) (string, error)
}
