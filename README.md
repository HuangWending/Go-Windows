# Go-Windows
Go语言创建窗口
package main

这一行指定了这个文件所属的包。在Go语言中，`package main` 表示这是一个可执行程序。

import (
	"fmt"
	"github.com/andlabs/ui"
)

这些 `import` 语句导入了程序所需的包。`fmt` 包提供了格式化和输出功能，`github.com/andlabs/ui` 包提供了用户界面的功能。

func main() {
	err := ui.Main(func() {
		// 程序逻辑
	})
	if err != nil {
		fmt.Printf("发生错误：%v", err)
	}
}

这是程序的入口点，main 函数。ui.Main 函数是一个阻塞调用，用于启动用户界面的主循环。

window := ui.NewWindow("黄文定的窗口", 200, 100, false)

这一行创建了一个新的窗口对象。NewWindow 函数接受窗口标题、宽度、高度和是否可调整大小等参数，并返回一个窗口对象。

label := ui.NewLabel("Hello, HuangWending")

这一行创建了一个标签对象。NewLabel 函数接受一个字符串参数，并返回一个带有指定文本内容的标签对象。

window.SetChild(label)

这一行将标签对象设置为窗口的子控件，以便在窗口中显示标签内容。

window.OnClosing(func(*ui.Window) bool {
	ui.Quit()
	return true
})
这一段代码设置了窗口关闭事件的处理函数。当用户尝试关闭窗口时，将调用此函数。在这个处理函数中，我们调用 `ui.Quit()` 函数来退出应用程序，并返回 `true` 以允许窗口关闭。
window.Show()

这一行显示窗口，使其可见。
最后，我们在 err 变量中捕获可能的错误，并使用 fmt.Printf函数打印错误信息。
