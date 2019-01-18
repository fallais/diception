package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestOMG tests the OMG function.
func TestOMG(t *testing.T) {
	solutions := map[string]string{
		"Martine dirige le pays !":                 "rige",
		"J'suis chaud pour une bonne dichotomie !": "chotomie",
		"Il a une légère dyslexie le bonhomme !":   "slexie",
		"dithyrambique":                            "thyrambique",
		"dite":                                     "te",
		"dis":                                      "",
	}

	for key, value := range solutions {
		// Create a ResponseRecorder
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(OMG)

		// Prepare the URL
		var reqURL *url.URL
		reqURL, err := url.Parse("/omg")
		if err != nil {
			t.Fatal("Error while parsing the URL")
		}
		parameters := url.Values{}
		parameters.Add("s", key)
		reqURL.RawQuery = parameters.Encode()

		// Create the request
		req, err := http.NewRequest("GET", reqURL.String(), nil)
		if err != nil {
			t.Fatal("Error while creating the request")
		}

		// Do the request
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		if rr.Body.String() != value {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), value)
		}
	}
}
