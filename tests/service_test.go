package main

import (
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var serviceUrl = "http://localhost:8081"

func Test_Echo_Body_String(t *testing.T) {
	dataText := "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"

	req, err := http.NewRequest("POST", serviceUrl+`/echo/body_string`, strings.NewReader(dataText))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(body), dataText)
}

func Test_Echo_Body(t *testing.T) {
	dataJson := `{"int_field":123,"string_field":"the value"}`

	req, err := http.NewRequest("POST", serviceUrl+`/echo/body`, strings.NewReader(dataJson))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), dataJson)
}

func Test_Echo_Body_Bad_Request(t *testing.T) {
	dataJson := `{"int_field":"the string","string_field":"the value"}`

	req, err := http.NewRequest("POST", serviceUrl+`/echo/body`, strings.NewReader(dataJson))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Echo_Query_Params(t *testing.T) {
	dataJson := `{"int_field":123,"long_field":12345,"float_field":1.23,"double_field":12.345,"decimal_field":12345,"bool_field":true,"string_field":"the value","string_opt_field":"the value","string_defaulted_field":"value","string_array_field":["the str1","the str2"],"uuid_field":"123e4567-e89b-12d3-a456-426655440000","date_field":"2020-01-01","date_array_field":["2020-01-01","2020-01-02"],"datetime_field":"2019-11-30T17:45:55","enum_field":"SECOND_CHOICE"}`

	req, err := http.NewRequest("GET", serviceUrl+`/echo/query`, nil)
	assert.NilError(t, err)

	q := req.URL.Query()
	q.Add("int_query", "123")
	q.Add("long_query", "12345")
	q.Add("float_query", "1.23")
	q.Add("double_query", "12.345")
	q.Add("decimal_query", "12345")
	q.Add("bool_query", "true")
	q.Add("string_query", "the value")
	q.Add("string_opt_query", "the value")
	q.Add("string_defaulted_query", "value")
	q.Add("string_array_query", "the str1")
	q.Add("string_array_query", "the str2")
	q.Add("uuid_query", "123e4567-e89b-12d3-a456-426655440000")
	q.Add("date_query", "2020-01-01")
	q.Add("date_array_query", "2020-01-01")
	q.Add("date_array_query", "2020-01-02")
	q.Add("datetime_query", "2019-11-30T17:45:55")
	q.Add("enum_query", "SECOND_CHOICE")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), dataJson)
}

func Test_Echo_Query_Params_Bad_Request(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/echo/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	q.Add("int_query", "the value")
	q.Add("string_query", "the value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Echo_Query_Params_Missing(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/echo/query`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Echo_Header_Params(t *testing.T) {
	dataJson := `{"int_field":123,"long_field":12345,"float_field":1.23,"double_field":12.345,"decimal_field":12345,"bool_field":true,"string_field":"the value","string_opt_field":"the value","string_defaulted_field":"value","string_array_field":["the str1","the str2"],"uuid_field":"123e4567-e89b-12d3-a456-426655440000","date_field":"2020-01-01","date_array_field":["2020-01-01","2020-01-02"],"datetime_field":"2019-11-30T17:45:55","enum_field":"SECOND_CHOICE"}`

	req, err := http.NewRequest("GET", serviceUrl+`/echo/header`, nil)
	assert.NilError(t, err)
	h := req.Header
	h.Add("Int-Header", "123")
	h.Add("Long-Header", "12345")
	h.Add("Float-Header", "1.23")
	h.Add("Double-Header", "12.345")
	h.Add("Decimal-Header", "12345")
	h.Add("Bool-Header", "true")
	h.Add("String-Header", "the value")
	h.Add("String-Opt-Header", "the value")
	h.Add("String-Defaulted-Header", "value")
	h.Add("String-Array-Header", "the str1")
	h.Add("String-Array-Header", "the str2")
	h.Add("Uuid-Header", "123e4567-e89b-12d3-a456-426655440000")
	h.Add("Date-Header", "2020-01-01")
	h.Add("Date-Array-Header", "2020-01-01")
	h.Add("Date-Array-Header", "2020-01-02")
	h.Add("Datetime-Header", "2019-11-30T17:45:55")
	h.Add("Enum-Header", "SECOND_CHOICE")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), dataJson)
}

func Test_Echo_Header_Params_Bad_Request(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/echo/header`, nil)
	assert.NilError(t, err)
	h := req.Header
	h.Add("Int-Header", "the value")
	h.Add("String-Header", "the value")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Echo_Url_Params(t *testing.T) {
	dataJson := `{"int_field":123,"long_field":12345,"float_field":1.23,"double_field":12.345,"decimal_field":12345,"bool_field":true,"string_field":"the value","uuid_field":"123e4567-e89b-12d3-a456-426655440000","date_field":"2020-01-01","datetime_field":"2019-11-30T17:45:55","enum_field":"SECOND_CHOICE"}`

	req, err := http.NewRequest("GET", serviceUrl+`/echo/url_params/123/12345/1.23/12.345/12345/true/the value/123e4567-e89b-12d3-a456-426655440000/2020-01-01/2019-11-30T17:45:55/SECOND_CHOICE`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), dataJson)
}

func Test_Echo_Url_Params_Bad_Request(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/echo/url_params/value/12345/1.23/12.345/12345/true/the value/123e4567-e89b-12d3-a456-426655440000/2020-01-01/2019-11-30T17:45:55/SECOND_CHOICE`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Check_Response_Empty(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/check/empty`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Url_Params(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/check/url_params/123/first_static_part/value/second_static_part`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Response_Forbidden(t *testing.T) {
	req, err := http.NewRequest("GET", serviceUrl+`/check/forbidden`, nil)
	assert.NilError(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 403)
}

func Test_Echo_Body_V2(t *testing.T) {
	dataJson := `{"bool_field":true,"string_field":"the value"}`

	req, err := http.NewRequest("POST", serviceUrl+`/v2/echo/body`, strings.NewReader(dataJson))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), dataJson)
}
