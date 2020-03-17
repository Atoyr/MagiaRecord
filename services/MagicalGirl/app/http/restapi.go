package http

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	options = "OPTIONS"
	get     = "GET"
	post    = "POST"
	put     = "PUT"
	patch   = "PATCH"
	delete  = "DELETE"
)

// APIStatus represents API's result status
type APIStatus struct {
	success bool
	code    int
	message string
}

// APIResource represents RESTful API Interfaces
type APIResource interface {
	Options(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Put(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Patch(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
	Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{})
}

type apiheader struct {
	Stastus string `json:"status"`
	Message string `json:"message"`
}

// APIResourceBase is defined for composition
type APIResourceBase struct{}

func (APIResourceBase) Options(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Get(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Post(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Put(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Patch(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func (APIResourceBase) Delete(url string, queries url.Values, body io.Reader) (APIStatus, interface{}) {
	return FailSimple(http.StatusMethodNotAllowed), nil
}

func Success(code int) APIStatus {
	return APIStatus{success: true, code: code, message: ""}
}

func Fail(code int, message string) APIStatus {
	return APIStatus{success: false, code: code, message: message}
}

func FailSimple(code int) APIStatus {
	return APIStatus{success: false, code: code, message: strconv.Itoa(code) + " " + http.StatusText(code)}
}
