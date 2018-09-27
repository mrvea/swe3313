package handler

import "testing"

// All of these test should be run with Async = false

func testLoggedOutAttempt(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 403
	rt.User = nil
	rt.H = UserGet
	rt.Method = "GET"
	rt.Pattern = "/users/:id"
	rt.URL = "/users/1"
	rt.Async = false

	rt.Run(t)
}
