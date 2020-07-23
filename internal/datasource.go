package internal

// Datasource interface for all type to implement for store and retrieve data
type Datasource interface {
	Read(dataRange string) [][]interface{}
}
