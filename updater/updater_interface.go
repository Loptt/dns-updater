package updater

// UpdaterInterface provices an abstract interface to Update DDNS records.
type UpdaterInterface interface {
	Update() error
}
