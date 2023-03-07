package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Create a new server
	server := echo.New()

	server.GET("/", hello)
	server.POST("/login", login)

	server.Logger.Fatal(server.Start(":8080"))

}

func hello(c echo.Context) error {
	x := &X{
		Text: "Hello World!",
	}
	return c.JSON(http.StatusOK, x)
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	c.Bind(req)

	log.Printf("req data: %+v", req)

	if req.Username != "admin" || req.Password != "123" {
		return c.String(http.StatusUnauthorized, "username/password invalid")
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		Token: "77777",
	})
}

type LoginRequest struct {
	Username string `json:"username" form: "username" query:"username"`
	Password string `json:"password" form: "password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}
