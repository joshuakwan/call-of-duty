package db

import (
	"reflect"

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
		session, err = mgo.Dial("mongodb://127.0.0.1")
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
	c := GetDBSession().DB("test").C(collectionName)
	return c.Insert(object)
}

// RemoveDBObject removes a DB object
func RemoveDBObject(colName string, selector interface{}) error {
	col := GetDBSession().DB("test").C(colName)
	err := col.Remove(selector)
	return err
}

// FindDBObjects finds a DB object
func FindDBObjects(colName string, query interface{}, result interface{}) error {
	col := GetDBSession().DB("test").C(colName)
	err := col.Find(query).All(result)
	return err
}

// UpdateDBObject updates a DB object
func UpdateDBObject(colName string, selector interface{}, update interface{}) error {
	col := GetDBSession().DB("test").C(colName)
	err := col.Update(selector, update)
	return err
}

func main() {

	// session, err := mgo.Dial("mongodb://127.0.0.1")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()

	// // Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

	// c := session.DB("test").C("people")
	// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)
}
