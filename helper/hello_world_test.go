package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Aldi",
			request: "Aldi",
		},
		{
			name:    "Syah",
			request: "Syah",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Aldi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aldi")
		}
	})
	b.Run("Syah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syah")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Aldi")
	}
}

func BenchmarkHelloWorldSyah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syah")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Aldi",
			request:  "Aldi",
			expected: "Hello Aldi",
		},
		{
			name:     "Syah",
			request:  "Syah",
			expected: "Hello Syah",
		},
		{
			name:     "Putra",
			request:  "Putra",
			expected: "Hello Putra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Aldi", func(t *testing.T) {
		result := HelloWorld("Aldi")
		require.Equal(t, "Hello Aldi", result, "Result must be 'Hello Aldi'")
	})
	t.Run("Syah", func(t *testing.T) {
		result := HelloWorld("Syah")
		require.Equal(t, "Hello Syah", result, "Result must be 'Hello Syah'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	m.Run() // eksekusi semua unit test

	// after
	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Aldi")
	require.Equal(t, "Hello Aldi", result, "Result must be 'Hello Aldi'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Aldi")
	require.Equal(t, "Hello Aldi", result, "Result must be 'Hello Aldi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Aldi")
	assert.Equal(t, "Hello Aldi", result, "Result must be 'Hello Aldi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAldi(t *testing.T) {
	result := HelloWorld("Aldi")
	if result != "Hello Aldi" {
		// error
		t.Error("Result must be 'Hello Aldi'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldSyah(t *testing.T) {
	result := HelloWorld("Syah")
	if result != "Hello Syah" {
		// error
		t.Fatal("Result must be 'Hello Syah'")
	}

	fmt.Println("TestHelloSyah Done")
}
