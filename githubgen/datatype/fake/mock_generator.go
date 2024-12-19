// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"go.opentelemetry.io/build-tools/githubgen/datatype"
	"sync"
)

// MockGenerator is a mock implementation of datatype.Generator.
//
//	func TestSomethingThatUsesGenerator(t *testing.T) {
//
//		// make and configure a mocked datatype.Generator
//		mockedGenerator := &MockGenerator{
//			GenerateFunc: func(data datatype.GithubData) error {
//				panic("mock out the Generate method")
//			},
//		}
//
//		// use mockedGenerator in code that requires datatype.Generator
//		// and then make assertions.
//
//	}
type MockGenerator struct {
	// GenerateFunc mocks the Generate method.
	GenerateFunc func(data datatype.GithubData) error

	// calls tracks calls to the methods.
	calls struct {
		// Generate holds details about calls to the Generate method.
		Generate []struct {
			// Data is the data argument value.
			Data datatype.GithubData
		}
	}
	lockGenerate sync.RWMutex
}

// Generate calls GenerateFunc.
func (mock *MockGenerator) Generate(data datatype.GithubData) error {
	if mock.GenerateFunc == nil {
		panic("MockGenerator.GenerateFunc: method is nil but Generator.Generate was just called")
	}
	callInfo := struct {
		Data datatype.GithubData
	}{
		Data: data,
	}
	mock.lockGenerate.Lock()
	mock.calls.Generate = append(mock.calls.Generate, callInfo)
	mock.lockGenerate.Unlock()
	return mock.GenerateFunc(data)
}

// GenerateCalls gets all the calls that were made to Generate.
// Check the length with:
//
//	len(mockedGenerator.GenerateCalls())
func (mock *MockGenerator) GenerateCalls() []struct {
	Data datatype.GithubData
} {
	var calls []struct {
		Data datatype.GithubData
	}
	mock.lockGenerate.RLock()
	calls = mock.calls.Generate
	mock.lockGenerate.RUnlock()
	return calls
}
