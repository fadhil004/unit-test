package helper

import "testing"

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Fadhil")
	if result != "Hello Fadhil" {
		//unit test failed
		panic("Result is not 'Hello Fadhil'")
	}
}