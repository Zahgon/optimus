package resource

const (
	Bigquery Store = "bigquery"
)

// Store represents the type of datasource, resource corresponds to
type Store string

func (s Store) String() string { _ = "STUB: not implemented"; return "" }

func FromStringToStore(name string) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}
