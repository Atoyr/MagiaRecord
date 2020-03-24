package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/MagicalGirl/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type magicalGirlController struct {
	util.APIResourceBase
}

func init() {
	http.Handle("/MagicalGirl", util.APIResourceHandler(magicalGirlController{}))
}

func (c magicalGirlController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	magicalGirls, err := models.GetMagicalGirlAll()
	if err != nil {
		log.Println(err)
		return util.Fail(http.StatusNotFound, fmt.Sprintln(err)), nil
	}
	return util.Success(http.StatusOK), magicalGirls
}
