package interfaces

type Database interface {
	C(string) DBCollection
}
