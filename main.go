package main

import (
	"image/color"
	"log"
	"os"
	"path/filepath"
	"stera_ui/internal"
	"stera_ui/pkg"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	window    *fyne.Window
	filename  string
	yanedtail internal.Yanedtail
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

	prgentry := widget.NewMultiLineEntry()
	prgentry.Wrapping = fyne.TextWrapWord
	card := widget.NewCard("進行状況", "",
		container.NewVBox(
			prgentry,
		),
	)
	prgentry.SetText("「開く」ボタンを押して，基盤地図情報(国土地理院）からQGIS2.18で作成した建物標高.geojsonファイルを読み込ませて下さい")

	entry1 := widget.NewEntry()
	entry1.SetText("ファイル名")
	label1 := widget.NewLabel("建物データ数")
	entry2 := widget.NewEntry()
	button1 := widget.NewButton("開く", func() {
		log.Println("開くボタンがクリックされました")
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("ファイルを開く際にエラーが発生しました:", err)
				return
			}
			if uc == nil {
				log.Println("ファイルの選択がキャンセルされました")
				return
			}
			fn := uc.URI().Path()
			log.Println("選択されたファイル:", fn)

			filename = fn
			// ファイル名
			fname := filepath.Base(fn)
			entry1.SetText(fname)
			// データ件数
			fl, er := os.Open(fn)
			if er != nil {
				log.Println("ファイルを読み込む際にエラーが発生しました:", er)
				return
			}
			defer fl.Close()

			_, l, _, _ := pkg.FileCount(fl)
			entry2.SetText(strconv.Itoa(l - 7))
		}, myWindow)
		prgentry.SetText("建物標高.geojsonファイルが読み込まれました．「ファイル分割」ボタンを押して、建物データを普通建物・堅ろう建物・無璧舎建物に分割して下さい．")
	})
	qbtn1 := widget.NewButton("？", func() {
		log.Println("Q1ボタンがクリックされました")
		dialog.ShowInformation("オープンファイルダイアログ", "QGIS 2.18で作成した建物標高.geojsonファイルを選択してください.", myWindow)
	})

	label2 := widget.NewLabel("普通建物")
	label3 := widget.NewLabel("堅ろう建物")
	label4 := widget.NewLabel("無璧舎建物")
	entry3 := widget.NewEntry()
	entry4 := widget.NewEntry()
	entry5 := widget.NewEntry()

	button2 := widget.NewButton("ファイル分割", func() {
		log.Println("SplitFileButton Clicked")
		internal.DivideLine(filename)

		wfl1, er := os.Open("data/hutsu_list.txt")
		_, l1, _, _ := pkg.FileCount(wfl1)
		entry3.SetText(strconv.Itoa(l1))
		if er != nil {
			log.Fatal(er)
		}
		defer wfl1.Close()

		wfl2, er := os.Open("data/kenrou_list.txt")
		_, l2, _, _ := pkg.FileCount(wfl2)
		entry4.SetText(strconv.Itoa(l2))
		if er != nil {
			log.Fatal(er)
		}
		defer wfl2.Close()

		wfl3, er := os.Open("data/other_list.txt")
		_, l3, _, _ := pkg.FileCount(wfl3)
		entry5.SetText(strconv.Itoa(l3))
		if er != nil {
			log.Fatal(er)
		}
		defer wfl3.Close()

		prgentry.SetText("建物データが普通建物・堅ろう建物・無璧舎建物に分割されました．普通建物に傾斜屋根をかける場合は四角形分割を行なって下さい．普通建物に傾斜屋根をかけない場合はファイル出力に進んでください．なお，COLLADAファイル以外で出力する場合は三角メッシュに分割して下さい。")
	})
	qbtn2 := widget.NewButton("？", func() {
		log.Println("Q2ボタンがクリックされました")
		dialog.ShowInformation("geojsonファイルの分割", "読み込んだファイルを建物種別に分割してください.", myWindow)
	})
	qbtn3 := widget.NewButton("？", func() {
		log.Println("Q3ボタンがクリックされました")
		dialog.ShowInformation("傾斜屋根を持つ低層木造建物", "屋根設定タブで屋根種別ごとに屋根勾配・軒の出等を設定してください.", myWindow)
	})
	qbtn4 := widget.NewButton("？", func() {
		log.Println("Q4ボタンがクリックされました")
		dialog.ShowInformation("普通建物の四角形分割", "普通建物に傾斜屋根を掛けるために四角形に分割してください.", myWindow)
	})
	qbtn5 := widget.NewButton("？", func() {
		log.Println("Q5ボタンがクリックされました")
		dialog.ShowInformation("陸屋根の非木造建物", "COLLADAファイル以外に出力する場合は三角メッシュに分割する必要があります.", myWindow)
	})
	qbtn6 := widget.NewButton("？", func() {
		log.Println("Q6ボタンがクリックされました")
		dialog.ShowInformation("陸屋根の無璧舎建物", "COLLADAファイル以外に出力する場合は三角メッシュに分割する必要があります.", myWindow)
	})
	qbtn7 := widget.NewButton("？", func() {
		log.Println("Q7ボタンがクリックされました")
		dialog.ShowInformation("多角形の３角メッシュ分割", "COLLADAファイル以外で出力する場合は堅ろう建物と無璧舎建物を三角メッシュに分割してください.", myWindow)
	})
	sbtn1 := widget.NewButton("四角形分割", func() {
		log.Println("HutsuBuildingsButton Clicked")
		internal.SquarePoly()
		prgentry.SetText("普通建物に四角形分割が行われました．ファイル出力に進んでください．なお，COLLADAファイル以外で出力する場合は三角メッシュに分割して下さい。")
	})
	sbtn2 := widget.NewButton("３角メッシュ分割", func() {
		log.Println("TriangleMeshDevideButton Clicked")
		internal.TriMeshDiv()
		prgentry.SetText("三角メッシュ分割が行われました．ファイル出力に進んでください．出力するファイル形式を選択して下さい。")
	})

	label5 := widget.NewLabel("出力ファイル形式")
	rb := widget.NewRadioGroup([]string{"COLLADA ファイル", "CityGML ファイル"}, func(selected string) {
		log.Println("選択された値:", selected)
	})
	qbtn8 := widget.NewButton("？", func() {
		log.Println("Q8ボタンがクリックされました")
		dialog.ShowInformation("出力ファイルのフォーマット選択", "出力するファイルのファイル形式を選択してください.", myWindow)
	})
	entry6 := widget.NewEntry()

	button3 := widget.NewButton("ファイル出力", func() {
		log.Println("OutputButton Clicked")
		dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				log.Println("ファイルを開く際にエラーが発生しました:", err)
				return
			}
			if uc == nil {
				log.Println("ファイルの選択がキャンセルされました")
				return
			}
			fn := uc.URI().Path()
			log.Println("選択されたファイル:", fn)

			// ファイル名
			fname := filepath.Base(fn)
			entry6.SetText(fname)
			log.Println(fname)

			// 出力するファイル形式の選択
			selectedFormat := rb.Selected
			if selectedFormat == "COLLADA ファイル" {
				// internal.BuildDAE(fname, yanedtail)
				log.Println("COLLADA Building Clicked")
			} else if selectedFormat == "CityGML ファイル" {
				// internal.BuildCity(fname, yanedtail)
				log.Println("CityGML Building Clicked")
			} else {
				log.Println("出力するファイル形式が選択されていません")
				dialog.ShowInformation("出力ファイルのフォーマット選択", "出力するファイルのファイル形式を選択してください.", myWindow)
			}
			prgentry.SetText("建物モデルの３次元データが指定されたファイル形式で出力されました．")
		}, myWindow)
	})
	qbtn9 := widget.NewButton("？", func() {
		log.Println("Q9ボタンがクリックされました")
		dialog.ShowInformation("セーブファイルダイアログ", "作成する３次元建物モデルの出力先を指定してください.", myWindow)
	})

	page1 := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("建物モデルの設定"),
		widget.NewButton("建物モデルを生成", func() {
			log.Println("建物モデルを生成ボタンがクリックされました")
		}),
		container.New(
			layout.NewHBoxLayout(),
			container.New(
				layout.NewVBoxLayout(),
				container.New(
					layout.NewGridLayout(2),
					button1,
					entry1,
					label1,
					entry2,
				),
				container.New(
					layout.NewGridLayout(3),
					button2,
					layout.NewSpacer(),
					layout.NewSpacer(),
					layout.NewSpacer(),
					label2,
					entry3,
					layout.NewSpacer(),
					sbtn1,
					layout.NewSpacer(),
					layout.NewSpacer(),
					label3,
					entry4,
					layout.NewSpacer(),
					label4,
					entry5,
					layout.NewSpacer(),
					sbtn2,
					layout.NewSpacer(),
				),
			),
			container.NewVBox(
				qbtn1,
				layout.NewSpacer(),
				qbtn2,
				qbtn3,
				qbtn4,
				qbtn5,
				qbtn6,
				qbtn7,
			),
			container.New(
				layout.NewVBoxLayout(),
				container.New(
					layout.NewGridLayout(2),
					label5,
					rb,
				),
				container.New(
					layout.NewGridLayout(2),
					button3,
					entry6,
				),
			),
			container.NewVBox(
				qbtn8,
				layout.NewSpacer(),
				qbtn9,
				layout.NewSpacer(),
				layout.NewSpacer(),
				layout.NewSpacer(),
				layout.NewSpacer(),
				layout.NewSpacer(),
				layout.NewSpacer(),
			),
		),
		container.NewVBox(
			card,
		),
	)

	nlabel1 := widget.NewLabel("軒の長さ")
	// 軒の長さのリスト
	noki_opt := []string{"300 mm", "450 mm", "600 mm", "900 mm", "1200 mm"}
	// Selectウィジェットの作成
	noki1 := widget.NewSelect(noki_opt, func(selected string) {
		log.Println("軒の長さ:", selected)
	})
	// デフォルトで選択する項目を設定
	noki1.SetSelected("450 mm")
	// 軒の長さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// ケラバ長さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kirinoki = noki(noki1.Selected)
	log.Println(yanedtail.Kirinoki)

	klabel1 := widget.NewLabel("ケラバ長さ")
	// ケラバ長さのリスト
	kera_opt := []string{"150 mm", "225 mm", "300 mm", "450 mm", "600 mm"}
	// Selectウィジェットの作成
	kera1 := widget.NewSelect(kera_opt, func(selected string) {
		log.Println("ケラバ長さ:", selected)
	})
	// デフォルトで選択する項目を設定
	kera1.SetSelected("300 mm")
	// ケラバ長さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根勾配等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kirikera = kera(kera1.Selected)
	log.Println(yanedtail.Kirikera)

	ilabel1 := widget.NewLabel("屋根勾配")
	// 屋根勾配のリスト
	incl_opt := []string{"３寸勾配", "3.5寸勾配", "４寸勾配", "4.5寸勾配", "５寸勾配", "5.5寸勾配", "６寸勾配"}
	// Selectウィジェットの作成
	incl1 := widget.NewSelect(incl_opt, func(selected string) {
		log.Println("屋根勾配:", selected)
	})
	// デフォルトで選択する項目を設定
	incl1.SetSelected("4.5寸勾配")
	// 屋根勾配のコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根厚さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kiriincl = incl(incl1.Selected)
	log.Println(yanedtail.Kiriincl)

	rlabel1 := widget.NewLabel("屋根厚さ")
	// 屋根厚さのリスト
	roof_opt := []string{"瓦屋根  110 mm", "スレート屋根  75 mm", "トタン屋根  70 mm"}
	// Selectウィジェットの作成
	roof1 := widget.NewSelect(roof_opt, func(selected string) {
		log.Println("屋根厚さ:", selected)
	})
	// デフォルトで選択する項目を設定
	roof1.SetSelected("瓦屋根  110 mm")
	// 屋根厚さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根勾配等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kiriroof = roof(roof1.Selected)
	log.Println(yanedtail.Kiriroof)

	card1 := widget.NewCard("切妻屋根", "",
		container.NewVBox(
			// widget.NewCheck("有効にする", nil),
			nlabel1,
			noki1,
			klabel1,
			kera1,
			ilabel1,
			incl1,
			rlabel1,
			roof1,
		),
	)

	nlabel2 := widget.NewLabel("軒の長さ")
	// 軒の長さのリスト
	noki2_opt := []string{"300 mm", "450 mm", "600 mm", "900 mm"}
	// Selectウィジェットの作成
	noki2 := widget.NewSelect(noki2_opt, func(selected string) {
		log.Println("軒の長さ:", selected)
	})
	// デフォルトで選択する項目を設定
	noki2.SetSelected("450 mm")
	// 軒の長さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// ケラバ長さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Yosenoki = noki(noki2.Selected)
	log.Println(yanedtail.Yosenoki)

	ilabel2 := widget.NewLabel("屋根勾配")
	// 屋根勾配のリスト
	incl2_opt := []string{"３寸勾配", "3.5寸勾配", "４寸勾配", "4.5寸勾配", "５寸勾配", "5.5寸勾配", "６寸勾配"}
	// Selectウィジェットの作成
	incl2 := widget.NewSelect(incl2_opt, func(selected string) {
		log.Println("屋根勾配:", selected)
	})
	// デフォルトで選択する項目を設定
	incl2.SetSelected("4.5寸勾配")
	// 屋根勾配のコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根厚さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Yoseincl = incl(incl2.Selected)
	log.Println(yanedtail.Yoseincl)

	rlabel2 := widget.NewLabel("屋根厚さ")
	// 屋根厚さのリスト
	roof2_opt := []string{"瓦屋根  110 mm", "スレート屋根  75 mm", "トタン屋根  70 mm"}
	// Selectウィジェットの作成
	roof2 := widget.NewSelect(roof2_opt, func(selected string) {
		log.Println("屋根厚さ:", selected)
	})
	// デフォルトで選択する項目を設定
	roof2.SetSelected("瓦屋根  110 mm")
	// 屋根厚さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根勾配等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Yoseroof = roof(roof2.Selected)
	log.Println(yanedtail.Yoseroof)

	card2 := widget.NewCard("寄棟屋根", "",
		container.NewVBox(
			// widget.NewCheck("有効にする", nil),
			nlabel2,
			noki2,
			layout.NewSpacer(),
			layout.NewSpacer(),
			ilabel2,
			incl2,
			rlabel2,
			roof2,
		),
	)

	nlabel3 := widget.NewLabel("軒の長さ")
	// 軒の長さのリスト
	noki3_opt := []string{"300 mm", "450 mm", "600 mm", "900 mm", "1200 mm"}
	// Selectウィジェットの作成
	noki3 := widget.NewSelect(noki3_opt, func(selected string) {
		log.Println("軒の長さ:", selected)
	})
	// デフォルトで選択する項目を設定
	noki3.SetSelected("450 mm")
	// 軒の長さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// ケラバ長さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Katanoki = noki(noki3.Selected)
	log.Println(yanedtail.Katanoki)

	klabel3 := widget.NewLabel("ケラバ長さ")
	// ケラバ長さのリスト
	kera3_opt := []string{"150 mm", "225 mm", "300 mm", "450 mm", "600 mm"}
	// Selectウィジェットの作成
	kera3 := widget.NewSelect(kera3_opt, func(selected string) {
		log.Println("ケラバ長さ:", selected)
	})
	// デフォルトで選択する項目を設定
	kera3.SetSelected("300 mm")
	// ケラバ長さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根勾配等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Katakera = kera(kera3.Selected)
	log.Println(yanedtail.Katakera)

	ilabel3 := widget.NewLabel("屋根勾配")
	// 屋根勾配のリスト
	incl3_opt := []string{"３寸勾配", "3.5寸勾配", "４寸勾配", "4.5寸勾配", "５寸勾配", "5.5寸勾配", "６寸勾配"}
	// Selectウィジェットの作成
	incl3 := widget.NewSelect(incl3_opt, func(selected string) {
		log.Println("屋根勾配:", selected)
	})
	// デフォルトで選択する項目を設定
	incl3.SetSelected("4.5寸勾配")
	// 屋根勾配のコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根厚さ等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kataincl = incl(incl3.Selected)
	log.Println(yanedtail.Kataincl)

	rlabel3 := widget.NewLabel("屋根厚さ")
	// 屋根厚さのリスト
	roof3_opt := []string{"瓦屋根  110 mm", "スレート屋根  75 mm", "トタン屋根  70 mm"}
	// Selectウィジェットの作成
	roof3 := widget.NewSelect(roof3_opt, func(selected string) {
		log.Println("屋根厚さ:", selected)
	})
	// デフォルトで選択する項目を設定
	roof3.SetSelected("瓦屋根  110 mm")
	// 屋根厚さのコンボボックスの選択結果に応じて軒の長さの数値を関数で設定する
	// 屋根勾配等の数値とともに構造体にしてbuildcollada.goに渡す
	yanedtail.Kataroof = roof(roof3.Selected)
	log.Println(yanedtail.Kataroof)

	card3 := widget.NewCard("片流れ屋根", "",
		container.NewVBox(
			// widget.NewCheck("有効にする", nil),
			nlabel3,
			noki3,
			klabel3,
			kera3,
			ilabel3,
			incl3,
			rlabel3,
			roof3,
		),
	)

	rf_btn1 := widget.NewButton("屋根種別", func() {
		log.Println("屋根種別ボタンがクリックされました")
		dialog.ShowInformation("主な傾斜屋根の形", "屋根頂部の棟から両側に流れる形状が切妻屋根．中央から四方に傾斜面がある形状が寄棟屋根．一方向だけに傾斜がついているのが片流れ屋根．平らな屋根が陸屋根．", myWindow)
	})
	rf_btn2 := widget.NewButton("軒の長さ", func() {
		log.Println("軒の長さボタンがクリックされました")
		dialog.ShowInformation("軒の出（庇の長さ）", "軒は屋根が外壁・窓より突き出ている部分で雨や日差しから建物を守る．建築面積に影響するため900mm以下とする場合が多いが，北海道では1200mmが推奨されている．", myWindow)
	})
	rf_btn3 := widget.NewButton("ケラバ長さ", func() {
		log.Println("ケラバ長さボタンがクリックされました")
		dialog.ShowInformation("ケラバの出（ケラバ長さ）", "切妻屋根の棟の両端部の名称．軒には雨どいがあるが，ケラバには雨どいがない．軒の長さの半分程度が理想的とされている．", myWindow)
	})
	rf_btn4 := widget.NewButton("屋根勾配", func() {
		log.Println("屋根勾配ボタンがクリックされました")
		dialog.ShowInformation("屋根勾配（傾斜面の傾き）", "屋根面の傾斜の大きさを「寸」という単位で表す．４寸勾配は21.8度であり１寸当たりで約5〜6度変化する．瓦屋根の場合４寸勾配が最低勾配である．", myWindow)
	})
	rf_btn5 := widget.NewButton("屋根厚さ", func() {
		log.Println("屋根厚さボタンがクリックされました")
		dialog.ShowInformation("屋根材と屋根下地の厚み", "瓦の厚さは約20mm，スレートは約5mm，金属材料は1mm未満．屋根材の下地の野地板が9〜12mm，屋根材と野地板を支える垂木は瓦の場合は75mm，金属屋根は60mm．", myWindow)
	})

	page2 := container.NewVBox(
		widget.NewLabel("屋根タイプの設定"),
		widget.NewButton("屋根を生成", func() {
			log.Println("屋根を生成ボタンがクリックされました")
		}),
		container.NewGridWithColumns(3,
			card1,
			card2,
			card3,
		),
		container.NewGridWithColumns(5,
			rf_btn1,
			rf_btn2,
			rf_btn3,
			rf_btn4,
			rf_btn5,
		),
	)

	cbox := widget.NewCheck("建物モデルを用途地域別に着色する場合はチェックを付ける", func(checked bool) {
		log.Println("用途地域別に着色するボタンがクリックされました")
		internal.Checked = checked
		if checked {
			log.Println("Checkbox is checked")
		} else {
			log.Println("Checkbox is unchecked")
		}
	})

	label_1 := widget.NewLabel("第一種低層住居専用地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_1 := canvas.NewRectangle(color.RGBA{R: 0, G: 165, B: 104, A: 255})
	color_1.SetMinSize(fyne.NewSize(60, 36))
	label1r := widget.NewLabel("R")
	label1g := widget.NewLabel("G")
	label1b := widget.NewLabel("B")
	entry1r := widget.NewEntry()
	entry1r.SetText("0")
	entry1g := widget.NewEntry()
	entry1g.SetText("165")
	entry1b := widget.NewEntry()
	entry1b.SetText("104")
	qbtn_1 := widget.NewButton("？", func() {
		log.Println("第一種低層住居専用地域ボタンがクリックされました")
		dialog.ShowInformation("第一種低層住居専用地域", "第一種低層住居専用地域は低層住宅の良好な住環境を守るための地域．（床面積の合計が）50m²までの住居を兼ねた一定条件の店舗や，小規模な公共施設，小中学校，診療所などを建てることができる．", myWindow)
	})

	label_2 := widget.NewLabel("第二種低層住居専用地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_2 := canvas.NewRectangle(color.RGBA{R: 119, G: 183, B: 158, A: 255})
	color_2.SetMinSize(fyne.NewSize(60, 36))
	label2r := widget.NewLabel("R")
	label2g := widget.NewLabel("G")
	label2b := widget.NewLabel("B")
	entry2r := widget.NewEntry()
	entry2r.SetText("119")
	entry2g := widget.NewEntry()
	entry2g.SetText("183")
	entry2b := widget.NewEntry()
	entry2b.SetText("158")
	qbtn_2 := widget.NewButton("？", func() {
		log.Println("第二種低層住居専用地域ボタンがクリックされました")
		dialog.ShowInformation("第二種低層住居専用地域", "第二種低層住居専用地域は主に低層住宅の良好な住環境を守るための地域．150m²までの一定条件の店舗等が建てられる．", myWindow)
	})

	label_3 := widget.NewLabel("第一種中高層住居専用地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_3 := canvas.NewRectangle(color.RGBA{R: 80, G: 175, B: 108, A: 255})
	color_3.SetMinSize(fyne.NewSize(60, 36))
	label3r := widget.NewLabel("R")
	label3g := widget.NewLabel("G")
	label3b := widget.NewLabel("B")
	entry3r := widget.NewEntry()
	entry3r.SetText("80")
	entry3g := widget.NewEntry()
	entry3g.SetText("175")
	entry3b := widget.NewEntry()
	entry3b.SetText("108")
	qbtn_3 := widget.NewButton("？", func() {
		log.Println("第一種中高層住居専用地域ボタンがクリックされました")
		dialog.ShowInformation("第一種中高層住居専用地域", "第一種中高層住居専用地域は中高層住宅の良好な住環境を守るための地域．500m²までの一定条件の店舗等が建てられる．中規模な公共施設，病院・大学なども建てられる．", myWindow)
	})

	label_4 := widget.NewLabel("第二種中高層住居専用地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_4 := canvas.NewRectangle(color.RGBA{R: 190, G: 205, B: 0, A: 255})
	color_4.SetMinSize(fyne.NewSize(60, 36))
	label4r := widget.NewLabel("R")
	label4g := widget.NewLabel("G")
	label4b := widget.NewLabel("B")
	entry4r := widget.NewEntry()
	entry4r.SetText("190")
	entry4g := widget.NewEntry()
	entry4g.SetText("205")
	entry4b := widget.NewEntry()
	entry4b.SetText("0")
	qbtn_4 := widget.NewButton("？", func() {
		log.Println("第二種中高層住居専用地域ボタンがクリックされました")
		dialog.ShowInformation("第二種中高層住居専用地域", "第二種中高層住居専用地域は主に中高層住宅の良好な住環境を守るための地域．1500m²までの一定条件の店舗や事務所等が建てられる．", myWindow)
	})

	label_5 := widget.NewLabel("第一種住居地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_5 := canvas.NewRectangle(color.RGBA{R: 0, G: 239, B: 68, A: 255})
	color_5.SetMinSize(fyne.NewSize(60, 36))
	label5r := widget.NewLabel("R")
	label5g := widget.NewLabel("G")
	label5b := widget.NewLabel("B")
	entry5r := widget.NewEntry()
	entry5r.SetText("0")
	entry5g := widget.NewEntry()
	entry5g.SetText("239")
	entry5b := widget.NewEntry()
	entry5b.SetText("68")
	qbtn_5 := widget.NewButton("？", func() {
		log.Println("第一種住居地域ボタンがクリックされました")
		dialog.ShowInformation("第一種住居地域", "第一種住居地域は住居の環境を保護するための地域．3000m²までの一定条件の店舗・事務所・ホテル等や，環境影響の小さいごく小規模な工場が建てられる．", myWindow)
	})

	label_6 := widget.NewLabel("第二種住居地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_6 := canvas.NewRectangle(color.RGBA{R: 249, G: 178, B: 0, A: 255})
	color_6.SetMinSize(fyne.NewSize(60, 36))
	label6r := widget.NewLabel("R")
	label6g := widget.NewLabel("G")
	label6b := widget.NewLabel("B")
	entry6r := widget.NewEntry()
	entry6r.SetText("249")
	entry6g := widget.NewEntry()
	entry6g.SetText("178")
	entry6b := widget.NewEntry()
	entry6b.SetText("0")
	qbtn_6 := widget.NewButton("？", func() {
		log.Println("第二種住居地域ボタンがクリックされました")
		dialog.ShowInformation("第二種住居地域", "第二種住居地域は主に住居の環境を保護するための地域．10000m²までの一定条件の店舗・事務所・ホテル・パチンコ屋・カラオケボックス等や，環境影響の小さいごく小規模な工場が建てられる．", myWindow)
	})

	label_7 := widget.NewLabel("準住居地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_7 := canvas.NewRectangle(color.RGBA{R: 238, G: 127, B: 0, A: 255})
	color_7.SetMinSize(fyne.NewSize(60, 36))
	label7r := widget.NewLabel("R")
	label7g := widget.NewLabel("G")
	label7b := widget.NewLabel("B")
	entry7r := widget.NewEntry()
	entry7r.SetText("238")
	entry7g := widget.NewEntry()
	entry7g.SetText("127")
	entry7b := widget.NewEntry()
	entry7b.SetText("0")
	qbtn_7 := widget.NewButton("？", func() {
		log.Println("準住居地域ボタンがクリックされました")
		dialog.ShowInformation("準住居地域", "準住居地域は道路の沿道等において，自動車関連施設などと，住居が調和した環境を保護するための地域．10000m²までの一定条件の店舗・事務所・ホテル・パチンコ屋・カラオケボックス等や，小規模の映画館，車庫・倉庫，環境影響の小さいごく小規模な工場も建てられる．", myWindow)
	})

	label_8 := widget.NewLabel("近隣商業地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_8 := canvas.NewRectangle(color.RGBA{R: 240, G: 145, B: 174, A: 255})
	color_8.SetMinSize(fyne.NewSize(60, 36))
	label8r := widget.NewLabel("R")
	label8g := widget.NewLabel("G")
	label8b := widget.NewLabel("B")
	entry8r := widget.NewEntry()
	entry8r.SetText("240")
	entry8g := widget.NewEntry()
	entry8g.SetText("145")
	entry8b := widget.NewEntry()
	entry8b.SetText("174")
	qbtn_8 := widget.NewButton("？", func() {
		log.Println("近隣商業地域ボタンがクリックされました")
		dialog.ShowInformation("近隣商業地域", "近隣商業地域は近隣の住民が日用品の買物をする店舗等の，業務の利便の増進を図る地域．ほとんどの商業施設・事務所のほか，住宅・店舗・ホテル・パチンコ屋・カラオケボックス等のほか，映画館，車庫・倉庫，小規模の工場も建てられる．延べ床面積規制が無いため，場合によっては中規模以上の建築物が建つ．", myWindow)
	})

	label_9 := widget.NewLabel("商業地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_9 := canvas.NewRectangle(color.RGBA{R: 232, G: 88, B: 133, A: 255})
	color_9.SetMinSize(fyne.NewSize(60, 36))
	label9r := widget.NewLabel("R")
	label9g := widget.NewLabel("G")
	label9b := widget.NewLabel("B")
	entry9r := widget.NewEntry()
	entry9r.SetText("232")
	entry9g := widget.NewEntry()
	entry9g.SetText("88")
	entry9b := widget.NewEntry()
	entry9b.SetText("133")
	qbtn_9 := widget.NewButton("？", func() {
		log.Println("商業地域ボタンがクリックされました")
		dialog.ShowInformation("商業地域", "商業地域は主に商業等の業務の利便の増進を図る地域．ほとんどの商業施設・事務所，住宅・店舗・ホテル・パチンコ屋・カラオケボックス等，映画館，車庫・倉庫，小規模の工場のほか，広義の風俗営業および性風俗関連特殊営業関係の施設も建てられる．延べ床面積規制が無く，容積率限度も相当高いため，高層ビル群も建てられる．", myWindow)
	})

	label_10 := widget.NewLabel("準工業地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_10 := canvas.NewRectangle(color.RGBA{R: 209, G: 189, B: 217, A: 255})
	color_10.SetMinSize(fyne.NewSize(60, 36))
	label10r := widget.NewLabel("R")
	label10g := widget.NewLabel("G")
	label10b := widget.NewLabel("B")
	entry10r := widget.NewEntry()
	entry10r.SetText("289")
	entry10g := widget.NewEntry()
	entry10g.SetText("189")
	entry10b := widget.NewEntry()
	entry10b.SetText("217")
	qbtn_10 := widget.NewButton("？", func() {
		log.Println("準工業地域ボタンがクリックされました")
		dialog.ShowInformation("準工業地域", "準工業地域は主に軽工業の工場等，環境悪化の恐れのない工場の利便を図る地域．住宅や商店も建てることができる．ただし，危険性・環境悪化のおそれが大きい花火工場や石油コンビナートなどは建設できない．", myWindow)
	})

	label_11 := widget.NewLabel("工業地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_11 := canvas.NewRectangle(color.RGBA{R: 191, G: 226, B: 231, A: 255})
	color_11.SetMinSize(fyne.NewSize(60, 36))
	label11r := widget.NewLabel("R")
	label11g := widget.NewLabel("G")
	label11b := widget.NewLabel("B")
	entry11r := widget.NewEntry()
	entry11r.SetText("191")
	entry11g := widget.NewEntry()
	entry11g.SetText("226")
	entry11b := widget.NewEntry()
	entry11b.SetText("231")
	qbtn_11 := widget.NewButton("？", func() {
		log.Println("工業地域ボタンがクリックされました")
		dialog.ShowInformation("工業地域", "工業地域は主に工業の業務の利便の増進を図る地域．どんな工場でも建てられる．住宅・店舗は建てられる．学校・病院・ホテル等は建てられない．", myWindow)
	})

	label_12 := widget.NewLabel("工業専用地域")
	// 色の四角形を作成し、サイズを60x36に固定する
	color_12 := canvas.NewRectangle(color.RGBA{R: 80, G: 107, B: 173, A: 255})
	color_12.SetMinSize(fyne.NewSize(60, 36))
	label12r := widget.NewLabel("R")
	label12g := widget.NewLabel("G")
	label12b := widget.NewLabel("B")
	entry12r := widget.NewEntry()
	entry12r.SetText("80")
	entry12g := widget.NewEntry()
	entry12g.SetText("107")
	entry12b := widget.NewEntry()
	entry12b.SetText("173")
	qbtn_12 := widget.NewButton("？", func() {
		log.Println("工業専用地域ボタンがクリックされました")
		dialog.ShowInformation("工業専用地域", "工業専用地域は工業の業務の利便の増進を図る地域．どんな工場でも建てられる．住宅・物品販売店舗・飲食店・学校・病院・ホテル等は建てられない．福祉施設（老人ホームなど）も不可．住宅が建設できない唯一の用途地域でもある．", myWindow)
	})

	cardy := widget.NewCard("用途地域", "",
		container.NewHBox(
			container.NewVBox(
				label_1,
				label_2,
				label_3,
				label_4,
				label_5,
				label_6,
				label_7,
				label_8,
				label_9,
				label_10,
				label_11,
				label_12,
			),
			container.NewVBox(
				color_1,
				color_2,
				color_3,
				color_4,
				color_5,
				color_6,
				color_7,
				color_8,
				color_9,
				color_10,
				color_11,
				color_12,
			),
			container.NewVBox(
				label1r,
				label2r,
				label3r,
				label4r,
				label5r,
				label6r,
				label7r,
				label8r,
				label9r,
				label10r,
				label11r,
				label12r,
			),
			container.NewGridWrap(fyne.Size{Width: 144, Height: 36},
				entry1r,
				entry2r,
				entry3r,
				entry4r,
				entry5r,
				entry6r,
				entry7r,
				entry8r,
				entry9r,
				entry10r,
				entry11r,
				entry12r,
			),
			container.NewVBox(
				label1g,
				label2g,
				label3g,
				label4g,
				label5g,
				label6g,
				label7g,
				label8g,
				label9g,
				label10g,
				label11g,
				label12g,
			),
			container.NewGridWrap(fyne.Size{Width: 144, Height: 36},
				entry1g,
				entry2g,
				entry3g,
				entry4g,
				entry5g,
				entry6g,
				entry7g,
				entry8g,
				entry9g,
				entry10g,
				entry11g,
				entry12g,
			),
			container.NewVBox(
				label1b,
				label2b,
				label3b,
				label4b,
				label5b,
				label6b,
				label7b,
				label8b,
				label9b,
				label10b,
				label11b,
				label12b,
			),
			container.NewGridWrap(fyne.Size{Width: 144, Height: 36},
				entry1b,
				entry2b,
				entry3b,
				entry4b,
				entry5b,
				entry6b,
				entry7b,
				entry8b,
				entry9b,
				entry10b,
				entry11b,
				entry12b,
			),
			container.NewVBox(
				qbtn_1,
				qbtn_2,
				qbtn_3,
				qbtn_4,
				qbtn_5,
				qbtn_6,
				qbtn_7,
				qbtn_8,
				qbtn_9,
				qbtn_10,
				qbtn_11,
				qbtn_12,
			),
		),
	)

	page3 := container.NewVBox(
		widget.NewLabel("用途地域の設定"),
		widget.NewButton("用途地域を生成", func() {
			log.Println("用途地域を生成ボタンがクリックされました")
		}),
		cbox,
		cardy,
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

	// tabs.SetTabLocation(container.TabLocationLeading) // タブの位置を左側に設定

	// ウインドウに表示して実行
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 480))
	myWindow.ShowAndRun()
}

func noki(nkid string) (nksize float64) {
	if nkid == "300 mm" {
		nksize = 0.3
	} else if nkid == "450 mm" {
		nksize = 0.45
	} else if nkid == "600 mm" {
		nksize = 0.6
	} else if nkid == "900 mm" {
		nksize = 0.9
	} else if nkid == "1200 mm" {
		nksize = 1.2
	}
	return nksize
}

func kera(krid string) (krsize float64) {
	if krid == "150 mm" {
		krsize = 0.15
	} else if krid == "225 mm" {
		krsize = 0.225
	} else if krid == "300 mm" {
		krsize = 0.3
	} else if krid == "450 mm" {
		krsize = 0.45
	} else if krid == "600 mm" {
		krsize = 0.6
	}
	return krsize
}

func incl(clid string) (clratio float64) {
	if clid == "３寸勾配" {
		clratio = 0.3
	} else if clid == "3.5寸勾配" {
		clratio = 0.35
	} else if clid == "４寸勾配" {
		clratio = 0.4
	} else if clid == "4.5寸勾配" {
		clratio = 0.45
	} else if clid == "５寸勾配" {
		clratio = 0.5
	} else if clid == "5.5寸勾配" {
		clratio = 0.55
	} else if clid == "６寸勾配" {
		clratio = 0.6
	}
	return clratio
}

func roof(rfid string) (rfsize float64) {
	if rfid == "瓦屋根  110 mm" {
		rfsize = 0.11
	} else if rfid == "スレート屋根  75 mm" {
		rfsize = 0.075
	} else if rfid == "トタン屋根  70 mm" {
		rfsize = 0.07
	}
	return rfsize
}
