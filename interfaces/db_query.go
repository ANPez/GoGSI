package interfaces

type DBQuery interface {
	Distinct(string, interface{}) error
	Iterator() DBResultIterator
}
