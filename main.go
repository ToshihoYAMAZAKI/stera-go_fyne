package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	window *fyne.Window
)

func main() {
	// ログファイルを新規作成、追記、書き込み専用、パーミッションは読むだけ
	file, err := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	// アプリとウインドウの初期化
	myApp := app.New()
	myWindow := myApp.NewWindow("３次元都市モデル生成プログラム")

	// 各画面（ページ）のコンテンツを作成
	page1 := container.NewVBox(
		widget.NewLabel("建物モデルの設定"),
		widget.NewButton("建物モデルを生成", func() {
			log.Println("建物モデルを生成ボタンがクリックされました")
		}),
	)

	page2 := container.NewVBox(
		widget.NewLabel("屋根設定の設定"),
		widget.NewButton("屋根を生成", func() {
			log.Println("屋根を生成ボタンがクリックされました")
		}),
	)

	page3 := container.NewVBox(
		widget.NewLabel("用途地域の設定"),
		widget.NewButton("用途地域を生成", func() {
			log.Println("用途地域を生成ボタンがクリックされました")
		}),
	)

	page4 := container.NewVBox(
		widget.NewLabel("地形モデルの設定"),
		widget.NewButton("地形モデルを生成", func() {
			log.Println("地形モデルを生成ボタンがクリックされました")
		}),
	)

	// タブの定義
	tabs := container.NewAppTabs(
		container.NewTabItem("建物モデル", page1),
		container.NewTabItem("屋根設定", page2),
		container.NewTabItem("用途地域", page3),
		container.NewTabItem("地形モデル", page4),
	)

	tabs.SetTabLocation(container.TabLocationLeading) // タブの位置を左側に設定

	// ウインドウに表示して実行
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(640, 480))
	myWindow.ShowAndRun()
}
