package interfaces

type DBResultIterator interface {
	Next(interface{}) bool
	Close() error
	Err() error
}
