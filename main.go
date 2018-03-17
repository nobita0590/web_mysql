package main

import (
	"gopkg.in/kataras/iris.v8"
	"gopkg.in/kataras/iris.v8/websocket"
	"github.com/nobita0590/web_mysql/modules"
	//"github.com/nobita0590/crm_service/models"
	"fmt"
	"github.com/nobita0590/web_mysql/key"
	"github.com/nobita0590/web_mysql/periodic"
	"github.com/nobita0590/web_mysql/config"
)

func main() {
	app := iris.Default()
	/*
	Setup websocket
	*/
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)
	app.Get("/ws_gate", ws.Handler())
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})

	periodic.Init()
	/*
	Setup db connection
	*/


	/*
	Register route
	*/

	modules.BindRoute(app)

	app.StaticWeb("/public", config.FilePath + "/public")

	// Method:   GET
	// Resource: http://localhost:8080/
	key.Init()

	app.Run(iris.Addr(":" + config.Port),iris.WithoutVersionChecker,iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner:
		// c.Emit("chat", msg)
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}
