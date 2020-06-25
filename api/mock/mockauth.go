// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-authorisation/auth"
	"github.com/ONSdigital/dp-recipe-api/api"
)

var (
	lockAuthHandlerMockRequire sync.RWMutex
)

// Ensure, that AuthHandlerMock does implement AuthHandler.
// If this is not the case, regenerate this file with moq.
var _ api.AuthHandler = &AuthHandlerMock{}

// AuthHandlerMock is a mock implementation of api.AuthHandler.
//
//     func TestSomethingThatUsesAuthHandler(t *testing.T) {
//
//         // make and configure a mocked api.AuthHandler
//         mockedAuthHandler := &AuthHandlerMock{
//             RequireFunc: func(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc {
// 	               panic("mock out the Require method")
//             },
//         }
//
//         // use mockedAuthHandler in code that requires api.AuthHandler
//         // and then make assertions.
//
//     }
type AuthHandlerMock struct {
	// RequireFunc mocks the Require method.
	RequireFunc func(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc

	// calls tracks calls to the methods.
	calls struct {
		// Require holds details about calls to the Require method.
		Require []struct {
			// Required is the required argument value.
			Required auth.Permissions
			// Handler is the handler argument value.
			Handler http.HandlerFunc
		}
	}
}

// Require calls RequireFunc.
func (mock *AuthHandlerMock) Require(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc {
	if mock.RequireFunc == nil {
		panic("AuthHandlerMock.RequireFunc: method is nil but AuthHandler.Require was just called")
	}
	callInfo := struct {
		Required auth.Permissions
		Handler  http.HandlerFunc
	}{
		Required: required,
		Handler:  handler,
	}
	lockAuthHandlerMockRequire.Lock()
	mock.calls.Require = append(mock.calls.Require, callInfo)
	lockAuthHandlerMockRequire.Unlock()
	return mock.RequireFunc(required, handler)
}

// RequireCalls gets all the calls that were made to Require.
// Check the length with:
//     len(mockedAuthHandler.RequireCalls())
func (mock *AuthHandlerMock) RequireCalls() []struct {
	Required auth.Permissions
	Handler  http.HandlerFunc
} {
	var calls []struct {
		Required auth.Permissions
		Handler  http.HandlerFunc
	}
	lockAuthHandlerMockRequire.RLock()
	calls = mock.calls.Require
	lockAuthHandlerMockRequire.RUnlock()
	return calls
}
