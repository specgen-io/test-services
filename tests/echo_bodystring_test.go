package main

import (
	"net/http"
	"strings"
	"testing"
)

const ECHO_BODY_STRING = `some text`

func Test_EchoBodyString(t *testing.T) {
	req, _ := http.NewRequest("POST", serviceUrl+`/echo/body_string`, strings.NewReader(ECHO_BODY_STRING))
	req.Header.Add("Content-Type", "text/plain")

	assertTextResponse(t, req, 200, ECHO_BODY_STRING)
}

func Test_EchoBodyString_Request_ContentType_Charset(t *testing.T) {
	req, _ := http.NewRequest("POST", serviceUrl+`/echo/body_string`, strings.NewReader(ECHO_BODY_STRING))
	req.Header.Add("Content-Type", "text/plain;charset=utf-8")

	assertTextResponse(t, req, 200, ECHO_BODY_STRING)
}

func Test_EchoBodyString_Request_ContentType_Empty(t *testing.T) {
	req, _ := http.NewRequest("POST", serviceUrl+`/echo/body_string`, strings.NewReader(ECHO_BODY_STRING))
	req.Header.Add("Content-Type", "")

	if check(ERRORS) {
		assertJsonResponse(t, req, 400, "", map[string]interface{}{
			"$.location":       "header",
			"$.message":        "Failed to parse header",
			"$.errors[0].code": "missing",
			"$.errors[0].path": "Content-Type",
		})
	} else {
		assertResponseStatus(t, req, 400)
	}
}

func Test_EchoBodyString_Request_ContentType_Missing(t *testing.T) {
	req, _ := http.NewRequest("POST", serviceUrl+`/echo/body_string`, strings.NewReader(ECHO_BODY_STRING))

	if check(ERRORS) {
		assertJsonResponse(t, req, 400, "", map[string]interface{}{
			"$.location":       "header",
			"$.message":        "Failed to parse header",
			"$.errors[0].code": "missing",
			"$.errors[0].path": "Content-Type",
		})
	} else {
		assertResponseStatus(t, req, 400)
	}
}
