package interfaces

type DBConnection interface {
	Clone() DBConnection
	DB(string) Database
	Close()
}
