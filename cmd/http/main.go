package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

// func main() {
// 	r := echo.New()

// 	r.GET("/", func(ctx echo.Context) error {
// 		// Method .String()
// 		// data := "Hello from /index"
// 		// return ctx.String(http.StatusOK, data)

// 		// Method .HTML()
// 		// data := "Hello from /index"
// 		// return ctx.HTML(http.StatusOK, data)

// 		// Method .Redirect()
// 		// return ctx.Redirect(http.StatusTemporaryRedirect, "/")

// 		// Method .JSON()
// 		data := M{"Message": "Hello", "Counter": 2}
// 		return ctx.JSON(http.StatusOK, data)
// 	})

// 	// Parsing Query String
// 	r.GET("/page1", func(ctx echo.Context) error {
// 		name := ctx.QueryParam("name")
// 		data := fmt.Sprintf("Hello %s", name)

// 		return ctx.String(http.StatusOK, data)
// 	})

// 	// Parsing URL Path Param
// 	r.GET("/page2/:name", func(ctx echo.Context) error {
// 		name := ctx.Param("name")
// 		data := fmt.Sprintf("Hello %s", name)

// 		return ctx.String(http.StatusOK, data)
// 	})

// 	// Parsing URL Path Param dan Setelahnya
// 	r.GET("/page3/:name/*", func(ctx echo.Context) error {
// 		name := ctx.Param("name")
// 		message := ctx.Param("*")

// 		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

// 		return ctx.String(http.StatusOK, data)
// 	})

// 	// Parsing Form Data
// 	r.POST("/page4", func(ctx echo.Context) error {
// 		name := ctx.FormValue("name")
// 		message := ctx.FormValue("message")

// 		data := fmt.Sprintf(
// 			"Hello %s, I have message for you: %s",
// 			name,
// 			strings.Replace(message, "/", "", 1),
// 		)

// 		return ctx.String(http.StatusOK, data)
// 	})

// 	r.Start(":9000")
// }

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

func main() {
	r := echo.New()

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	r.Start(":9000")
}
