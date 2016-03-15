package middlewares

import (
	"github.com/ANPez/gogsi/interfaces"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
)

type DBConnection struct {
	*mgo.Session
}

func (dbc *DBConnection) Clone() interfaces.DBConnection {
	copy := dbc.Session.Copy()
	copy.SetMode(mgo.Monotonic, true)
	return &DBConnection{copy}
}

func (dbc *DBConnection) DB(name string) interfaces.Database {
	return &DB{dbc.Session.DB(name)}
}

type DB struct {
	*mgo.Database
}

func (db *DB) C(name string) interfaces.DBCollection {
	return &DBCollection{db.Database.C(name)}
}

type DBCollection struct {
	*mgo.Collection
}

func (dbc *DBCollection) AggregationIterator(query interface{}) interfaces.DBResultIterator {
	return DBResultIterator{dbc.Collection.Pipe(query).Iter()}
}

func (dbc *DBCollection) AggregationOne(query, result interface{}) error {
	return dbc.Collection.Pipe(query).One(result)
}

func (dbc *DBCollection) Find(query interface{}) interfaces.DBQuery {
	return &DBQuery{dbc.Collection.Find(query)}
}

func (dbc *DBCollection) Insert(object interface{}) error {
	return dbc.Collection.Insert(object)
}

// FindOne finds the first object matched by the query and stores it into result.
// Returns whether the object has been found or not and an error.
func (dbc *DBCollection) FindOne(query interface{}, result interface{}) (bool, error) {
	err := dbc.Collection.Find(query).One(result)
	if nil != err {
		if mgo.ErrNotFound == err {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

type DBQuery struct {
	*mgo.Query
}

func (dbq *DBQuery) Iterator() interfaces.DBResultIterator {
	return DBResultIterator{dbq.Query.Iter()}
}

type DBResultIterator struct {
	*mgo.Iter
}

func NewDatabaseMiddleware(host, dbname string) gin.HandlerFunc {
	session, err := mgo.Dial(host)
	if nil != err {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		sess := session.Copy()
		sess.SetMode(mgo.Monotonic, true)

		c.Set("mongodb_connection", &DBConnection{sess})
		c.Set("mongodb_db", &DB{sess.DB(dbname)})

		c.Next()

		sess.Close()
	}
}
