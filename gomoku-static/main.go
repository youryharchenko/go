package main

import (
	"fmt"
	"log"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const name = "gomoku"

type ResProvider struct {
	app.ResourceProvider
	pack   string
	static string
	wasm   string
}

func (res ResProvider) Package() string {
	return res.pack
}

func (res ResProvider) Static() string {
	return res.static
}

func (res ResProvider) AppWASM() string {
	return res.wasm
}

type gomoku struct {
	app.Compo
}

func (g *gomoku) Render() app.UI {
	return app.Div().
		Body(
			app.Div().
				Style("display", "flex").
				Style("justify-content", "center").
				Body(
					app.H1().Text("Gomoku game"),
				),
			app.Div().
				Style("display", "flex").
				Style("justify-content", "center").
				Body(
					&desk{n: 0, who: 1},
				),
		)
}

type desk struct {
	app.Compo
	center *position
	all    []*position
	n      int
	steps  []Step
	who    int
	play   bool
}

func (d *desk) Render() app.UI {
	rows := make([]string, 15)
	cols := make([]string, 15)
	for i := 0; i < 15; i += 1 {
		rows[i] = fmt.Sprintf("r%02d", i+1)
		cols[i] = fmt.Sprintf("c%02d", i+1)
	}
	return app.Table().ID("desk").
		Style("border-spacing", "0px 0px").
		Style("display", "block").
		Style("border-bottom", "0px").
		Style("border-spacing", "0").
		Style("margin", "0px").
		Style("border", "0px").
		Style("padding", "0px").
		Body(
			app.Div().
				Style("display", "flex").
				Style("justify-content", "center").
				Body(
					app.Button().Text("New: Black").OnClick(func(ctx app.Context, e app.Event) {
						for _, p := range d.all {
							p.s = 0
							p.Update()
						}
						d.n = 1
						d.steps = []Step{{x: 7, y: 7}}
						d.center.s = 1
						d.play = true
						d.center.Update()

						st, stat := calcStep(d.steps)

						fmt.Println("stat", stat)
						if len(st) > len(d.steps) {
							x := st[len(st)-1].x
							y := st[len(st)-1].y
							for _, cp := range d.all {
								if cp.x == x && cp.y == y && cp.s == 0 {
									d.n += 1
									if d.n%2 == 1 {
										cp.s = 1
									} else {
										cp.s = 2
									}
									cp.Update()
									d.steps = append(d.steps, st[len(st)-1])
								}
							}
						}
					}),
					app.Button().Text("New: White").OnClick(func(ctx app.Context, e app.Event) {
						for _, p := range d.all {
							p.s = 0
							p.Update()
						}
						d.n = 1
						d.steps = []Step{{x: 7, y: 7}}
						d.center.s = 1
						d.play = true
						d.center.Update()
					}),
				),
			app.Range(rows).Slice(func(i int) app.UI {
				return app.Tr().ID(rows[i]).
					Style("height", "40px").
					Style("display", "block").
					Body(
						app.Range(cols).Slice(func(j int) app.UI {
							p := &position{desk: d, id: rows[i] + cols[j], x: j, y: i, s: 0}
							if i == 7 && j == 7 {
								d.center = p
							}
							d.all = append(d.all, p)
							return p
						}),
					)
			}),
		)
}

type position struct {
	app.Compo
	desk *desk
	id   string
	s    int
	x    int
	y    int
}

func (p *position) Render() app.UI {
	return app.Td().ID(p.id).
		//Style("display", "block").
		Style("height", "40px").
		Style("width", "40px").
		Style("margin", "0px").
		Style("border", "0px").
		Style("padding", "0px").
		Body(
			app.If(p.s == 0,
				app.Img().Src("static/img/Empty.png").
					Style("height", "40px").
					Style("width", "40px"),
			).ElseIf(p.s == 1,
				app.Img().Src("static/img/Black.png").
					Style("height", "40px").
					Style("width", "40px"),
			).ElseIf(p.s == 2,
				app.Img().Src("static/img/White.png").
					Style("height", "40px").
					Style("width", "40px"),
			),
		).
		OnClick(p.onClick)
}

func (p *position) onClick(ctx app.Context, e app.Event) {

	if p.desk.play && p.s == 0 {
		fmt.Println("onClick is called", p.x, p.y, p.s, p.desk.n)
		p.desk.n += 1
		if p.desk.n%2 == 1 {
			p.s = 1
		} else {
			p.s = 2
		}
		p.Update()
		p.desk.steps = append(p.desk.steps, Step{x: p.x, y: p.y})

		st, stat := calcStep(p.desk.steps)

		fmt.Println("stat", stat)
		if len(st) > len(p.desk.steps) {
			x := st[len(st)-1].x
			y := st[len(st)-1].y
			for _, cp := range p.desk.all {
				if cp.x == x && cp.y == y && cp.s == 0 {
					p.desk.n += 1
					if p.desk.n%2 == 1 {
						cp.s = 1
					} else {
						cp.s = 2
					}
					cp.Update()
					p.desk.steps = append(p.desk.steps, st[len(st)-1])
				}
			}
		}
		if stat != "play" {
			p.desk.play = false
		}
	}
}

func main() {

	app.Route("/", &gomoku{})

	app.RunWhenOnBrowser()

	res := &ResProvider{
		pack:   fmt.Sprintf("/%s/", name),
		static: fmt.Sprintf("/%s/static/", name),
		wasm:   fmt.Sprintf("/%s/web/app.wasm", name),
	}

	err := app.GenerateStaticWebsite("./", &app.Handler{
		Name:        "Gomoku",
		Description: "Gomoku game",
		Resources:   res,
		Icon: app.Icon{
			Default: "static/img/Gomoku.png", // Specify default favicon.
		},
	})
	if err != nil {
		log.Fatal(err)
	}

}
