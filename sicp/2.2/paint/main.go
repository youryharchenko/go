//+build js,wasm

package main

import (
	"fmt"
	"log"
	"strconv"
	"syscall/js"
	"time"

	"github.com/youryharchenko/go/kit"
)

var glob *kit.Global

func init() {
	glob = kit.Env.G
}

// App -
type App struct {
	Console *kit.Console
	Body    *kit.Element
	Canvas  *kit.Element
	Ctx     *kit.Context2d
	Ticker  *time.Ticker
	Stop    chan bool
	LogCh   chan []interface{}
}

func main() {

	c := make(chan bool)

	app := &App{
		Console: glob.Console(),
		Body:    glob.Document().Body(),
		Canvas:  glob.Document().GetElementByID("c"),
		Ticker:  time.NewTicker(time.Millisecond * 30),
		Stop:    make(chan bool, 0),
		LogCh:   make(chan []interface{}, 100),
	}

	app.Console.Log("Hello!")

	ctx := app.Canvas.GetContext("2d")
	if ctx == nil {
		p := glob.Document().CreateElement("p")
		p.SetInnerHTML("Context2d does not supprted")
		app.Body.AppendChild(p)
		return
	}
	//log.Println(ctx)
	app.Ctx = ctx

	go logger(app)
	go painter(app, draw)

	resize(app)

	glob.AddEventListener("resize", js.FuncOf(func(this js.Value, args []js.Value) (r interface{}) {
		//app.LogCh <- []interface{}{"resize", this, args}
		resize(app)
		return
	}))

	<-c

}

func draw(ctx *kit.Context2d, w1 int, h1 int) {
	d1, d2 := 5, 3
	ctx.FillRect(0, 0, w1, h1)
	w2, h2 := w1-w1/d1, h1-h1/d1
	x2, y2 := w1/d1/2, h1/d1/2
	ctx.ClearRect(x2, y2, w2, h2)
	w3, h3 := w1-w1/d2, h1-h1/d2
	x3, y3 := w1/d2/2, h1/d2/2
	ctx.StrokeRect(x3, y3, w3, h3)
}

func resize(app *App) {
	app.Ticker.Stop()
	app.Stop <- true

	bw := app.Body.ClientWidth()
	bh := app.Body.ClientHeight()

	app.Canvas.Set("width", bw)
	app.Canvas.Set("height", bh)

	app.Canvas.Style().Set("width", strconv.Itoa(bw)+"px")
	app.Canvas.Style().Set("height", strconv.Itoa(bh)+"px")
	app.Canvas.Style().Set("position", "absolute")

	app.Ticker = time.NewTicker(time.Millisecond * 30)
	go painter(app, draw)
}

func painter(app *App, f func(ctx *kit.Context2d, w int, h int)) {
	for {
		select {
		case <-app.Ticker.C:
			cw := app.Canvas.ClientWidth()
			ch := app.Canvas.ClientHeight()
			//app.LogCh <- []interface{}{"tick", cw, ch}
			f(app.Ctx, cw, ch)
		case <-app.Stop:
			return
		}
	}
}

func logger(app *App) {
	for args := range app.LogCh {
		msg := append([]interface{}{time.Now().Format("2006-01-02 15:04:05")}, args...)
		app.Console.Log(fmt.Sprintln(msg...))
	}
}

func printMessage(this js.Value, inputs []js.Value) (r interface{}) {
	message := inputs[0].String()
	log.Println(message)
	return
}
