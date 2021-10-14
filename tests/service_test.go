package main

import (
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var service_url = "http://localhost:8081"

func Test_Echo_Body(t *testing.T) {
	data_json := `{"int_field":123,"string_field":"the value"}`

	req, err := http.NewRequest("POST", service_url+`/echo/body`, strings.NewReader(data_json))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), data_json)
}

func Test_Echo_Query_Params(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/echo/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	q.Add("int_query", "123")
	q.Add("string_query", "the value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), `{"int_field":123,"string_field":"the value"}`)
}

func Test_Echo_Query_Params_Missing(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/echo/query`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Echo_Header_Params(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/echo/header`, nil)
	assert.NilError(t, err)
	h := req.Header
	h.Add("Int-Header", "123")
	h.Add("String-Header", "the value")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), `{"int_field":123,"string_field":"the value"}`)
}

func Test_Echo_Url_Params(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/echo/url_params/123/value`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), `{"int_field":123,"string_field":"value"}`)
}

func Test_Check_Response_Empty(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/empty`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Query_Params(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	q.Add("p_string", "the string 1")
	q.Add("p_string_opt", "the string 2")
	q.Add("p_string_array", "str1")
	q.Add("p_string_array", "str2")
	q.Add("p_date", "2020-01-01")
	q.Add("p_date_array", "2020-01-01")
	q.Add("p_date_array", "2020-01-02")
	q.Add("p_datetime", "2019-11-30T17:45:55")
	q.Add("p_int", "123")
	q.Add("p_long", "1234")
	q.Add("p_decimal", "12345")
	q.Add("p_enum", "SECOND_CHOICE")
	q.Add("p_string_defaulted", "value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Query_Params_Missing_Required_Param(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	//q.Add("p_string", "the string 1") <- this is required param and it's not provided
	q.Add("p_string_opt", "the string 2")
	q.Add("p_string_array", "str1")
	q.Add("p_string_array", "str2")
	q.Add("p_date", "2020-01-01")
	q.Add("p_date_array", "2020-01-01")
	q.Add("p_date_array", "2020-01-02")
	q.Add("p_datetime", "2019-11-30T17:45:55")
	q.Add("p_int", "123")
	q.Add("p_long", "1234")
	q.Add("p_decimal", "12345")
	q.Add("p_enum", "SECOND_CHOICE")
	q.Add("p_string_defaulted", "value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 400)
}

func Test_Check_Query_Params_Missing_Optional_Param(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	q.Add("p_string", "the string 1")
	//q.Add("p_string_opt", "the string 2") <- this is optional param and it's not provided
	q.Add("p_string_array", "str1")
	q.Add("p_string_array", "str2")
	q.Add("p_date", "2020-01-01")
	q.Add("p_date_array", "2020-01-01")
	q.Add("p_date_array", "2020-01-02")
	q.Add("p_datetime", "2019-11-30T17:45:55")
	q.Add("p_int", "123")
	q.Add("p_long", "1234")
	q.Add("p_decimal", "12345")
	q.Add("p_enum", "SECOND_CHOICE")
	q.Add("p_string_defaulted", "value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Query_Params_Missing_Defaulted_Param(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/query`, nil)
	assert.NilError(t, err)
	q := req.URL.Query()
	q.Add("p_string", "the string 1")
	q.Add("p_string_opt", "the string 2")
	q.Add("p_string_array", "str1")
	q.Add("p_string_array", "str2")
	q.Add("p_date", "2020-01-01")
	q.Add("p_date_array", "2020-01-01")
	q.Add("p_date_array", "2020-01-02")
	q.Add("p_datetime", "2019-11-30T17:45:55")
	q.Add("p_int", "123")
	q.Add("p_long", "1234")
	q.Add("p_decimal", "12345")
	q.Add("p_enum", "SECOND_CHOICE")
	q.Add("p_string_defaulted", "value")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Url_Params(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/url_params/123/value/1.23/true/123e4567-e89b-12d3-a456-426655440000/1.23/2019-11-30/SECOND_CHOICE`, nil)
	assert.NilError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
}

func Test_Check_Response_Forbidden(t *testing.T) {
	req, err := http.NewRequest("GET", service_url+`/check/forbidden`, nil)
	assert.NilError(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 403)
}

func Test_Echo_Body_V2(t *testing.T) {
	data_json := `{"bool_field":true,"string_field":"the value"}`

	req, err := http.NewRequest("POST", service_url+`/v2/echo/body`, strings.NewReader(data_json))
	assert.NilError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)
	err = resp.Body.Close()
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, strings.TrimSuffix(string(body), "\n"), data_json)
}
