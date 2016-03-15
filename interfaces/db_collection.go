package interfaces

type DBCollection interface {
	AggregationIterator(interface{}) DBResultIterator
	AggregationOne(interface{}, interface{}) error
	Find(interface{}) DBQuery
	FindOne(interface{}, interface{}) (bool, error)
	Insert(object interface{}) error
}
