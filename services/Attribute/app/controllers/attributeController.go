package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/Attribute/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type attributeController struct {
	util.APIResourceBase
}

func init() {
	http.Handle("/Attribute", util.APIResourceHandler(typeController{}))
}

func (c typeController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	attributes, err := models.GetAttributeAll()
	if err != nil {
		log.Println(err)
		return util.Fail(http.StatusNotFound, fmt.Sprintln(err)), nil
	}
	return util.Success(http.StatusOK), attributes
}
