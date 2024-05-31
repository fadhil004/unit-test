package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T)  {
	t.Run("Fadhil", func(t *testing.T)  {
		result := HelloWorld("Fadhil")
		require.Equal(t, "Hello Fadhil", result, "Result must be 'Hello Fadhil'")
	})
	t.Run("Farraz", func(t *testing.T)  {
		result := HelloWorld("Farraz")
		require.Equal(t, "Hello Farraz", result, "Result must be 'Hello Farraz'")
	})
}

func TestMain(m *testing.M)  {
	//before
	fmt.Println("Sebelum Unit Test")
	m.Run()
	//after
	fmt.Println("Sesudah Unit Test")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

	result := HelloWorld("Fadhil")
	require.Equal(t, "Hello Fadhil", result, "Result must be 'Hello Fadhil'")
}

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