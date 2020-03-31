package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/Connect/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type connectController struct {
	util.APIResourceBase
}

func init() {
	http.Handle("/Connect", util.APIResourceHandler(connectController{}))
}

func (c connectController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	var connects []models.Connect
	var err error

	if len(queries) == 0 {
		connects, err = models.GetConnectAll()
		if err != nil {
			log.Println(err)
			return util.Fail(http.StatusInternalServerError, fmt.Sprintln(err)), nil
		}
	}
	return util.Success(http.StatusOK), connects
}
