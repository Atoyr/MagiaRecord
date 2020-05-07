package handler

import (
	"github.com/labstack/echo"
)

type HandlerFunc func(*Context) error

type Context struct {
	echo.Context
	DriverName     string
	DataSourceName string
}
