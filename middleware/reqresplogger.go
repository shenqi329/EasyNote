package middleware

import (
	"fmt"
	"github.com/labstack/echo"
)

func ReqRespLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			//req := c.Request()
			//resp := c.Response()

			if err = next(c); err != nil {
				c.Error(err)
			}
			fmt.Println("ReqRespLogger")
			return err
		}
	}
}
