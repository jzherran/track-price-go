package internal

// DataSource interface for all type to implement for store and retrieve data
type DataSource interface {

	// Read it's the function to read stored data from data store
	Read(dataRange string) [][]interface{}

	// Write it's the function to write data to store from data store
	Write(dataRange string) [][]interface{}
}
