package db

import (
	"reflect"

	"github.com/joshuakwan/call-of-duty/config"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

type Person struct {
	Name  string
	Phone string
}

func GetDBSession() *mgo.Session {
	var err error
	if session == nil {
		session, err = mgo.Dial(config.GetMongoURL())
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
	return session
}

// CreateDBObject creates a new mongo object
func CreateDBObject(object interface{}) error {
	collectionName := reflect.ValueOf(object).Elem().Type().String()
	c := GetDBSession().DB(config.GetDatabaseName()).C(collectionName)
	return c.Insert(object)
}

// RemoveDBObject removes a DB object
func RemoveDBObject(colName string, selector interface{}) error {
	col := GetDBSession().DB(config.GetDatabaseName()).C(colName)
	err := col.Remove(selector)
	return err
}

// FindDBObjects finds a DB object
func FindDBObjects(colName string, query interface{}, result interface{}) error {
	col := GetDBSession().DB(config.GetDatabaseName()).C(colName)
	err := col.Find(query).All(result)
	return err
}

// UpdateDBObject updates a DB object
func UpdateDBObject(colName string, selector interface{}, update interface{}) error {
	col := GetDBSession().DB(config.GetDatabaseName()).C(colName)
	err := col.Update(selector, update)
	return err
}
