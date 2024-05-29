package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Fadhil")
	assert.Equal(t, "Hello Fadhil", result, "Result must be 'Hello Fadhil'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Fadhil")
	require.Equal(t, "Hello Fadhil", result, "Result must be 'Hello Fadhil'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Fadhil")
	if result != "Hello Fadhil" {
		//unit test failed
		t.Error("Result must be 'Hello Fadhil'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldFaraz(t *testing.T)  {
	result := HelloWorld("Faraz")
	if result != "Hello Faraz" {
		//unit test failed
		t.Fatal("Result must be 'Hello Faraz'")
	}

	fmt.Println("TestHelloWorldFaraz Done")
}