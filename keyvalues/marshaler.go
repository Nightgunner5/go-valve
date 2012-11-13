package keyvalues

type Marshaler interface {
	ToKeyValues() *KeyValues
	FromKeyValues(*KeyValues) error
}
