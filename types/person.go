package types

//go:generate ../main -package=types -name=p -type=Person
type Person struct {
	Name string
	Age int
}
