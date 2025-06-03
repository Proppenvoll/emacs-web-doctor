package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAssertRequest(t *testing.T) {
	testCases := []struct {
		name   string
		input  http.Request
		output *requestAssertion
	}{
		{
			name: "asserts the post method",
			input: http.Request{
				Method: http.MethodPost,
				Header: map[string][]string{
					"Content-Type": {"application/x-www-form-urlencoded"},
				},
			},
			output: nil,
		},
		{
			name: "generates an output on a non post method",
			input: http.Request{
				Method: "whatever",
				Header: map[string][]string{
					"Content-Type": {"application/x-www-form-urlencoded"},
				},
			},
			output: &requestAssertion{
				text:       http.StatusText(http.StatusMethodNotAllowed),
				httpStatus: http.StatusMethodNotAllowed,
				header:     requestAssertionHeader{"Allow", "POST"},
			},
		},
		{
			name: "asserts the content type application/x-www-form-urlencoded",
			input: http.Request{
				Method: http.MethodPost,
				Header: map[string][]string{
					"Content-Type": {"application/x-www-form-urlencoded"},
				},
			},
			output: nil,
		},
		{
			name: "generates an output for an unsupported content type",
			input: http.Request{
				Method: http.MethodPost,
				Header: map[string][]string{
					"Content-Type": {"whatever"},
				},
			},
			output: &requestAssertion{
				text:       http.StatusText(http.StatusUnsupportedMediaType),
				httpStatus: http.StatusUnsupportedMediaType,
				header: requestAssertionHeader{
					"Accept-Post",
					"application/x-www-form-urlencoded",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := assertRequest(testCase.input)
			if !reflect.DeepEqual(result, testCase.output) {
				t.Errorf("got: %v, want: %v", result, testCase.output)
			}
		})
	}
}
