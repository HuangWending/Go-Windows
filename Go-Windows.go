package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("黄文定的窗口", 200, 100, false)
		label := ui.NewLabel("Hello, HuangWending")
		window.SetChild(label)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		fmt.Printf("发生错误：%v", err)
	}
}
