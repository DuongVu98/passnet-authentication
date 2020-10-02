package midlewares

import (
	"github.com/DuongVu98/passnet-authentication/src/main/config/handles"
	"github.com/labstack/echo/v4"
	"reflect"
)

func GetBeanMiddlewareProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		handles.Push(reflect.TypeOf(handles.GetBeanRequest{}))
		return next(c)
	}
}
