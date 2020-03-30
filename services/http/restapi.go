package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	Status  string `json:"status"`
	Message string `json:"message"`
}

type apienvelope struct {
	Header   apiheader
	Response interface{}
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

func APIResourceHandler(resource APIResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(r.Body, b)
		r.Body = ioutil.NopCloser(b)
		defer r.Body.Close()

		e := r.ParseForm()
		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var status APIStatus
		var data interface{}

		switch r.Method {
		case http.MethodOptions:
			status, data = resource.Options(r.URL.Path, r.Form, reader)
		case http.MethodGet:
			status, data = resource.Get(r.URL.Path, r.Form, reader)
		case http.MethodPost:
			status, data = resource.Post(r.URL.Path, r.Form, reader)
		case http.MethodPut:
			status, data = resource.Put(r.URL.Path, r.Form, reader)
		case http.MethodPatch:
			status, data = resource.Patch(r.URL.Path, r.Form, reader)
		case http.MethodDelete:
			status, data = resource.Delete(r.URL.Path, r.Form, reader)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// Return API response
		var content []byte

		if !status.success {
			content, e = json.Marshal(apienvelope{
				Header: apiheader{Status: "fail", Message: status.message},
			})
		} else {
			content, e = json.Marshal(apienvelope{
				Header:   apiheader{Status: "success"},
				Response: data,
			})
		}
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)
		w.Write(content)
	}
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
