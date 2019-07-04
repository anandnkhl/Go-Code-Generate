package types

//go:generate ../main -package=types -name=e -type=Employee -pointer
type Employee struct {
	Name   string
	Age    int
	Salary float64
}
