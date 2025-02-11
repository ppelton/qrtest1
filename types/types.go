package types

type ContentProvider struct {
	Name    string
	Path    string
	Provide func() ([]byte, error)
}
