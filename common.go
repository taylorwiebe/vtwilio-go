package vtwilio

// Method Request methods
type Method string

const (
	// GET method
	GET Method = "GET"
	// POST method
	POST Method = "POST"
)

func (m Method) String() string {
	return string(m)
}
