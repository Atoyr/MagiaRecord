package controllers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/MagicalGirl/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type magicalGirlController struct {
	util.APIResourceBase
}

func (c magicalGirlController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	return util.Success(http.StatusOK), models.MagicalGirl{}
}
