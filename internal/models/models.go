package models

type Employee struct {
	ID     int     `bson:"_id"`
	Name   string  `bson:"name"`
	Sex    string  `bson:"sex"`
	Age    int     `bson:"age"`
	Salary float64 `bson:"salary"`
}

type UpdateEmployee struct {
	Name   string  `bson:"name"`
	Sex    string  `bson:"sex"`
	Age    int     `bson:"age"`
	Salary float64 `bson:"salary"`
}

type EmployeeGroup struct {
	ID             int        `bson:"_id"`
	GroupTitle     string     `bson:"group_title"`
	EmployeeNumber int        `bson:"employee_number"`
	EmployeeList   []Employee `bson:"employee_list"`
}
