package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/Type/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type typeController struct {
	util.APIResourceBase
}

func init() {
	http.Handle("/MagicalGirl", util.APIResourceHandler(typeController{}))
}

func (c typeController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	types, err := models.GetTypeAll()
	if err != nil {
		log.Println(err)
		return util.Fail(http.StatusNotFound, fmt.Sprintln(err)), nil
	}
	return util.Success(http.StatusOK), types
}
