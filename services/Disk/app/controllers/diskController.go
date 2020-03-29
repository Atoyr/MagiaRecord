package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/atoyr/MagiaRecord/services/Disk/app/models"
	util "github.com/atoyr/MagiaRecord/services/http"
)

type diskController struct {
	util.APIResourceBase
}

func init() {
	http.Handle("/Disk", util.APIResourceHandler(diskController{}))
}

func (c diskController) Get(url string, queries url.Values, body io.Reader) (util.APIStatus, interface{}) {
	disks, err := models.GetDiskAll()
	if err != nil {
		log.Println(err)
		return util.Fail(http.StatusNotFound, fmt.Sprintln(err)), nil
	}
	return util.Success(http.StatusOK), disks
}
