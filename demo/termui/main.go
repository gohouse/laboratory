// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	gauge()
}
func gauge()  {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	g0 := widgets.NewGauge()
	g0.Title = "Slim Gauge"
	g0.SetRect(20, 20, 30, 30)
	g0.Percent = 75
	g0.BarColor = ui.ColorRed
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorCyan

	g2 := widgets.NewGauge()
	g2.Title = "Slim Gauge"
	g2.SetRect(0, 3, 50, 6)
	g2.Percent = 60
	g2.BarColor = ui.ColorYellow
	g2.LabelStyle = ui.NewStyle(ui.ColorBlue)
	g2.BorderStyle.Fg = ui.ColorWhite

	g1 := widgets.NewGauge()
	g1.Title = "Big Gauge"
	g1.SetRect(0, 6, 50, 11)
	g1.Percent = 30
	g1.BarColor = ui.ColorGreen
	g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	g1.TitleStyle.Fg = ui.ColorMagenta
	g1.BorderStyle.Fg = ui.ColorWhite

	g3 := widgets.NewGauge()
	g3.Title = "Gauge with custom label"
	g3.SetRect(0, 11, 50, 14)
	g3.Percent = 50
	g3.Label = fmt.Sprintf("%v%% (100MBs free)", g3.Percent)

	g4 := widgets.NewGauge()
	g4.Title = "Gauge"
	g4.SetRect(0, 14, 50, 17)
	g4.Percent = 50
	g4.Label = "Gauge with custom highlighted label"
	g4.BarColor = ui.ColorGreen
	g4.LabelStyle = ui.NewStyle(ui.ColorYellow)

	ui.Render(g0, g1, g2, g3, g4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
func canvas()  {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	ui.Render(c)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
func barchart()  {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bc := widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
func img() {
	var images []image.Image
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			log.Fatalf("failed to fetch image: %v", err)
		}
		image, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatalf("failed to decode fetched image: %v", err)
		}
		images = append(images, image)
	}
	if len(images) == 0 {
		image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(GOPHER_IMAGE)))
		if err != nil {
			log.Fatalf("failed to decode gopher image: %v", err)
		}
		images = append(images, image)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	img := widgets.NewImage(nil)
	img.SetRect(0, 0, 40, 20)
	index := 0
	render := func() {
		img.Image = images[index]
		if !img.Monochrome {
			img.Title = fmt.Sprintf("Color %d/%d", index+1, len(images))
		} else if !img.MonochromeInvert {
			img.Title = fmt.Sprintf("Monochrome(%d) %d/%d", img.MonochromeThreshold, index+1, len(images))
		} else {
			img.Title = fmt.Sprintf("InverseMonochrome(%d) %d/%d", img.MonochromeThreshold, index+1, len(images))
		}
		ui.Render(img)
	}
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Left>", "h":
			index = (index + len(images) - 1) % len(images)
		case "<Right>", "l":
			index = (index + 1) % len(images)
		case "<Up>", "k":
			img.MonochromeThreshold++
		case "<Down>", "j":
			img.MonochromeThreshold--
		case "<Enter>":
			img.Monochrome = !img.Monochrome
		case "<Tab>":
			img.MonochromeInvert = !img.MonochromeInvert
		}
		render()
	}
}

const GOPHER_IMAGE = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`
const GOPHER_IMAGE2 = `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`
