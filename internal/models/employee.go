package models

type Employee struct {
	ID     int     `bson:"_id"`
	Name   string  `bson:"name"`
	Sex    string  `bson:"sex"`
	Age    int     `bson:"age"`
	Salary float64 `bson:"salary"`
}
