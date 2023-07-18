package router

import (
	"os"
	"shift-app-go-api/controller"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, sc controller.IShiftController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	s := e.Group("/shifts")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	})) //ミドルウェア適用
	s.GET("", sc.GetAllShifts)
	s.GET("/:shiftId", sc.GetShiftById)
	s.POST("", sc.CreateShift)
	s.PUT("/:shiftId", sc.UpdateShift)
	s.DELETE("/:shiftId", sc.DeleteShift)
	return e
}
