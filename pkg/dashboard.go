package pkg

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

// keybinding

func KeyDashboard(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, dquit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, searchBuffer); err != nil {
		return err
	}
	if err := g.SetKeybinding("searchBuffer", gocui.KeyEnter, gocui.ModNone, delsearchBuffer); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}
	return nil
}

func dashboard(g *gocui.Gui) error {
	origin.maxX, origin.maxY = g.Size()
	if len(origin.Cluster) > 0 {
		num := len(origin.Cluster)
		len := origin.maxX / num
		for n, x := range origin.Cluster {
			var endX int
			startX := n * len
			startY := 0
			if n == num-1 {
				endX = origin.maxX - 1
			} else {
				endX = (n+1)*len - 1
			}
			// endX = (n + 1) * len
			endY := origin.maxY/4 - 1

			err := StatusTable(g, startX, startY, endX, endY, x)
			if err != nil {
				return err
			}
		}
	}

	if len(origin.ServiceConfig) > 0 {
		num := len(origin.Cluster)
		len := origin.maxX / num
		for n, x := range origin.ServiceConfig {
			var endX int
			startX := n * len
			startY := origin.maxY / 4
			if n == num-1 {
				endX = origin.maxX - 1
			} else {
				endX = (n+1)*len - 1
			}
			// endX = (n + 1) * len
			endY := origin.maxY/2 - 1

			err := StatusTable(g, startX, startY, endX, endY, x)
			if err != nil {
				return err
			}
		}
	}

	// if v, err := g.SetView("bottom", 0, maxY/2, maxX/2-1, maxY-1); err != nil {
	// 	if err != gocui.ErrUnknownView {
	// 		return err
	// 	}
	// 	v.Title = "工作负载状态"
	// 	v.Wrap = true
	// 	// v.Highlight = true
	// 	// v.Autoscroll = true
	// 	v.SelBgColor = gocui.ColorGreen
	// 	v.SelFgColor = gocui.ColorBlack
	// 	// v.Editable = true
	// 	// fmt.Fprintf(v, time.Now().Format("2006-01-02 15:04:05"))
	// 	// uri = fmt.Sprintf("/a%s", time.Now().Format("150405"))
	// 	fmt.Fprintln(v, fmt.Sprintf("URL => 0.0.0.0:%s <= \nPATH: => %s <=", "9999", "/tmp"))
	// 	fmt.Fprintln(v, origin.Cluster)
	// 	if _, err := setCurrentViewOnTop(g, "bottom"); err != nil {
	// 		return err
	// 	}
	// }

	// if err := WorkLoadTable(g, 0, maxY/2, maxX/2-1, maxY-1); err != nil {
	if err := WorkLoadTable(g, 0, origin.maxY/2, origin.maxX-1, origin.maxY*3/4-1); err != nil {
		return err
	}

	// if v, err := g.SetView("pod", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
	// 	if err != gocui.ErrUnknownView {
	// 		return err
	// 	}
	// 	v.Title = "Pod"
	// 	v.Wrap = true
	// 	// v.Highlight = true
	// 	// v.Autoscroll = true
	// 	v.SelBgColor = gocui.ColorGreen
	// 	v.SelFgColor = gocui.ColorBlack
	// 	// v.Editable = true
	// 	// fmt.Fprintf(v, time.Now().Format("2006-01-02 15:04:05"))
	// 	// uri = fmt.Sprintf("/a%s", time.Now().Format("150405"))
	// 	fmt.Fprintln(v, fmt.Sprintf("URL => 0.0.0.0:%s <= \nPATH: => %s <=", "9999", "/tmp"))
	// 	fmt.Fprintln(v, origin.Cluster)
	// }

	// if err := PodsTable(g, maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
	if err := PodsTable(g, 0, origin.maxY*3/4, origin.maxX-1, origin.maxY-1); err != nil {
		return err
	}

	return nil
}

func RefreshWorkLoad(g *gocui.Gui, startx, starty, endx, endy int) error {
	if v, err := origin.Gui.View("bottom"); err != nil {
		return err
	} else {
		v.Clear()
		// v.Autoscroll = true
		v.Wrap = true
		v.Highlight = true
		v.SetCursor(startx+2, starty+1)
		// v.Editable = true
		num := 0
		tableNow := NewTable(endx - startx - 1)

		// tableNow.AddCol("ID").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Type").SetColor("red").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("NAME").SetColor("red").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Namespace").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Tags").SetColor("red").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Ready").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Images").SetColor("red").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Time").SetColor("red").SetTextAlign(TextRight).SetBgColor("black")
		tableNow.CalColumnWidths()
		for _, value := range origin.PodControllers {
			num++
			if num == 1 {
				tableNow.FprintHeader(v)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			tt := NewCol()
			tt.Data = fmt.Sprintf("*%s", value.Type)
			tt.TextAlign = TextLeft
			tt.Color = "yellow"
			tableNow.AddRow(0, tt)

			name := NewCol()
			name.Data = fmt.Sprintf("%s", value.Name)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(1, name)

			ns := NewCol()
			ns.Data = fmt.Sprintf("%s", value.Namespace)
			ns.TextAlign = TextCenter
			ns.Color = "yellow"
			tableNow.AddRow(2, ns)

			ttags := ""
			for k, v := range value.Tags {
				ttags += fmt.Sprintf("%s:%s ", k, v)
			}
			Tags := NewCol()
			Tags.Data = fmt.Sprintf("%s", ttags)
			Tags.TextAlign = TextLeft
			Tags.Color = "yellow"
			tableNow.AddRow(3, Tags)

			rd := NewCol()
			rd.Data = fmt.Sprintf("%s", value.ContainerGroup)
			rd.TextAlign = TextCenter
			rd.Color = "yellow"
			tableNow.AddRow(4, rd)

			image := NewCol()
			image.Data = fmt.Sprintf("%s", value.Images)
			image.TextAlign = TextLeft
			image.Color = "yellow"
			tableNow.AddRow(5, image)

			time := NewCol()
			time.Data = fmt.Sprintf("%s", strings.Replace(value.Time, "\n", "", -1))
			time.TextAlign = TextRight
			time.Color = "yellow"
			tableNow.AddRow(6, time)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v)
	}
	return nil
}

func WorkLoadTable(g *gocui.Gui, startx, starty, endx, endy int) error {
	if v, err := g.SetView("bottom", startx, starty, endx, endy); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = fmt.Sprintf("WorkLoad/%d", len(origin.PodControllers))
		v.Highlight = true
		v.Editable = true
		// v.Wrap = true
		// v.MoveCursor(startx+2, starty+3, true)
		v.SetCursor(startx+2, starty+1)

		v.Clear()
		num := 0
		tableNow := NewTable(endx - startx - 1)

		// tableNow.AddCol("ID").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Type").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("NAME").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Namespace").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Tags").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Ready").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Images").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Time").SetColor("dgreen").SetTextAlign(TextRight).SetBgColor("black")
		tableNow.CalColumnWidths()
		for _, value := range origin.PodControllers {
			num++
			if num == 1 {
				tableNow.FprintHeader(v)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			tt := NewCol()
			tt.Data = fmt.Sprintf("*%s", value.Type)
			tt.TextAlign = TextLeft
			tt.Color = "yellow"
			tableNow.AddRow(0, tt)

			name := NewCol()
			name.Data = fmt.Sprintf("%s", value.Name)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(1, name)

			ns := NewCol()
			ns.Data = fmt.Sprintf("%s", value.Namespace)
			ns.TextAlign = TextCenter
			ns.Color = "yellow"
			tableNow.AddRow(2, ns)

			ttags := ""
			for k, v := range value.Tags {
				ttags += fmt.Sprintf("%s:%s ", k, v)
			}
			Tags := NewCol()
			Tags.Data = fmt.Sprintf("%s", ttags)
			Tags.TextAlign = TextLeft
			Tags.Color = "yellow"
			tableNow.AddRow(3, Tags)

			rd := NewCol()
			rd.Data = fmt.Sprintf("%s", value.ContainerGroup)
			rd.TextAlign = TextCenter
			rd.Color = "yellow"
			tableNow.AddRow(4, rd)

			image := NewCol()
			image.Data = fmt.Sprintf("%s", value.Images)
			image.TextAlign = TextLeft
			image.Color = "yellow"
			tableNow.AddRow(5, image)

			time := NewCol()
			time.Data = fmt.Sprintf("%s", value.Time)
			time.TextAlign = TextRight
			time.Color = "yellow"
			tableNow.AddRow(6, time)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v)

		if _, err := setCurrentViewOnTop(g, "bottom"); err != nil {
			return err
		}
	}
	return nil
}

func RefreshPods(g *gocui.Gui, startx, starty, endx, endy int) error {
	if v1, err := origin.Gui.View("Pod"); err != nil {
		return err
	} else {
		v1.Clear()
		v1.Wrap = true
		v1.Highlight = true
		v1.Editable = true
		v1.SetCursor(startx+2, starty+1)

		num := 0
		tableNow := NewTable(origin.maxX - 1)

		// tableNow.AddCol("ID").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("NAME").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Namespace").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Node").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Ready").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Restarts").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Time").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.CalColumnWidths()

		for _, value := range origin.Pods {
			num++
			if num == 1 {
				tableNow.FprintHeader(v1)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			name := NewCol()
			name.Data = fmt.Sprintf("*%s", value.Name)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(0, name)

			ns := NewCol()
			ns.Data = fmt.Sprintf("%s", value.Namespace)
			ns.TextAlign = TextCenter
			ns.Color = "yellow"
			tableNow.AddRow(1, ns)

			node := NewCol()
			node.Data = fmt.Sprintf("%s", value.Node)
			node.TextAlign = TextCenter
			node.Color = "yellow"
			tableNow.AddRow(2, node)

			rd := NewCol()
			rd.Data = fmt.Sprintf("%s", value.Ready)
			rd.TextAlign = TextCenter
			rd.Color = "yellow"
			tableNow.AddRow(3, rd)

			rs := NewCol()
			rs.Data = fmt.Sprintf("%s", value.Restarts)
			rs.TextAlign = TextCenter
			rs.Color = "yellow"
			tableNow.AddRow(4, rs)

			time := NewCol()
			time.Data = fmt.Sprintf("%s", value.Time)
			time.TextAlign = TextLeft
			time.Color = "yellow"
			tableNow.AddRow(5, time)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v1)
	}

	if v, err := origin.Gui.View("pod"); err != nil {
		return err
	} else {
		v.Clear()
		// v.Autoscroll = true
		v.Wrap = true
		v.Highlight = true
		v.Editable = true
		// v.Wrap = true
		// v.MoveCursor(startx, endy, false)
		v.SetCursor(startx+2, starty+1)

		num := 0
		tableNow := NewTable(endx - startx - 1)

		// tableNow.AddCol("ID").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("NAME").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Namespace").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Node").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Ready").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Restarts").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Time").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.CalColumnWidths()

		for _, value := range origin.Pods {
			num++
			if num == 1 {
				tableNow.FprintHeader(v)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			name := NewCol()
			name.Data = fmt.Sprintf("*%s", value.Name)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(0, name)

			ns := NewCol()
			ns.Data = fmt.Sprintf("%s", value.Namespace)
			ns.TextAlign = TextCenter
			ns.Color = "yellow"
			tableNow.AddRow(1, ns)

			node := NewCol()
			node.Data = fmt.Sprintf("%s", value.Node)
			node.TextAlign = TextCenter
			node.Color = "yellow"
			tableNow.AddRow(2, node)

			rd := NewCol()
			rd.Data = fmt.Sprintf("%s", value.Ready)
			rd.TextAlign = TextCenter
			rd.Color = "yellow"
			tableNow.AddRow(3, rd)

			rs := NewCol()
			rs.Data = fmt.Sprintf("%s", value.Restarts)
			rs.TextAlign = TextCenter
			rs.Color = "yellow"
			tableNow.AddRow(4, rs)

			time := NewCol()
			time.Data = fmt.Sprintf("%s", value.Time)
			time.TextAlign = TextLeft
			time.Color = "yellow"
			tableNow.AddRow(5, time)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v)
	}
	return nil
}

func PodsTable(g *gocui.Gui, startx, starty, endx, endy int) error {
	if v, err := g.SetView("pod", startx, starty, endx, endy); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = fmt.Sprintf("Pod/%d", len(origin.Pods))
		v.Highlight = true
		v.Editable = true
		// v.Wrap = true
		// v.MoveCursor(startx, endy, false)
		v.SetCursor(startx+2, starty+1)

		num := 0
		tableNow := NewTable(endx - startx - 1)

		// tableNow.AddCol("ID").SetColor("red").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("NAME").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.AddCol("Namespace").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Node").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Ready").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Restarts").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.AddCol("Time").SetColor("dgreen").SetTextAlign(TextLeft).SetBgColor("black")
		tableNow.CalColumnWidths()

		for _, value := range origin.Pods {
			num++
			if num == 1 {
				tableNow.FprintHeader(v)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			name := NewCol()
			name.Data = fmt.Sprintf("*%s", value.Name)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(0, name)

			ns := NewCol()
			ns.Data = fmt.Sprintf("%s", value.Namespace)
			ns.TextAlign = TextCenter
			ns.Color = "yellow"
			tableNow.AddRow(1, ns)

			node := NewCol()
			node.Data = fmt.Sprintf("%s", value.Node)
			node.TextAlign = TextCenter
			node.Color = "yellow"
			tableNow.AddRow(2, node)

			rd := NewCol()
			rd.Data = fmt.Sprintf("%s", value.Ready)
			rd.TextAlign = TextCenter
			rd.Color = "yellow"
			tableNow.AddRow(3, rd)

			rs := NewCol()
			rs.Data = fmt.Sprintf("%s", value.Restarts)
			rs.TextAlign = TextCenter
			rs.Color = "yellow"
			tableNow.AddRow(4, rs)

			time := NewCol()
			time.Data = fmt.Sprintf("%s", value.Time)
			time.TextAlign = TextLeft
			time.Color = "yellow"
			tableNow.AddRow(5, time)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v)
	}
	return nil
}

func StatusTable(g *gocui.Gui, startx, starty, endx, endy int, data ClusterStatus) error {
	if v, err := g.SetView(data.Title, startx, starty, endx, endy); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = fmt.Sprintf("%s/%d", data.Title, data.Count)
		v.Highlight = true
		v.Editable = true
		// v.Wrap = true
		// v.MoveCursor(startx+2, starty+1, false)
		v.SetCursor(startx+2, starty+1)

		num := 0
		tableNow := NewTable(endx - startx)

		tableNow.AddCol("NAME").SetColor("dgreen").SetTextAlign(TextCenter).SetBgColor("black")
		tableNow.CalColumnWidths()

		for _, value := range data.Data {
			num++
			if num == 1 {
				tableNow.FprintHeader(v)
			}

			// id := NewCol()
			// id.Data = fmt.Sprintf("%d", num)
			// id.TextAlign = TextCenter
			// id.Color = "yellow"
			// tableNow.AddRow(0, id)

			name := NewCol()
			name.Data = fmt.Sprintf("*%s", value)
			name.TextAlign = TextLeft
			name.Color = "yellow"
			tableNow.AddRow(0, name)

			// fmt.Fprintln(w, s)
		}
		tableNow.Fprint(v)
	}
	return nil
}
