package handler

import (
	"database/sql"
	"net/http"

	"github.com/atoyr/MagiaRecord/bff/models"
)

func GetAttributesHandler() HandlerFunc {
	return func(c *Context) error {
		var err error
		var apierr APIError
		db, err := sql.Open(c.DriverName, c.DataSourceName)
		if err != nil {
			apierr.Code = 900
			apierr.Message = "Database Server Not Found"
			c.JSON(http.StatusInternalServerError, apierr)
			return err
		}
		defer db.Close()
		if err = db.Ping(); err != nil {
			apierr.Code = 900
			apierr.Message = "Database Server Not Found"
			c.JSON(http.StatusInternalServerError, apierr)
			return err
		}
		as, err := models.GetAttributes(db)
		if err != nil {
			apierr.Code = 100
			apierr.Message = "Query Error"
			c.JSON(http.StatusInternalServerError, apierr)
			return err
		}
		return c.JSON(http.StatusOK, as)
	}
}
