package types

//go:generate ../main  -package=types -name=s -type=Student -pointer
type Student struct {
	Name    string
	Age     int
	College string
}
