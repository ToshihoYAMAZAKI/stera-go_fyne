package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// TerrXMLはLandXML形式で地形モデルを作成する
func TerrXML(filename string, x_matrix, y_matrix, z_matrix [][]float64, x_dot, y_dot int, z_max, z_min float64) {
	// LandXML形式で出力するためのファイルを開く
	// f, err := os.Create("data/outputgeo.xml")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// LandXMLファイルの出力開始
	// ファイル出力（ヘッダー部分の書き出し）
	xml := "<LandXML xmlns=\"http://www.landxml.org/schema/LandXML-1.2\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" version=\"1.2\" date=\"2008-07-29\" time=\"10:39:13\" xsi:schemaLocation=\"http://www.landxml.org/schema/LandXML-1.2 http://www.landxml.org/schema/LandXML-1.2/LandXML-1.2.xsd\">\n"
	data := []byte(xml)
	cnt, err := f.Write(data)
	if err != nil {
		log.Println(err)
		log.Println("fail to write file")
	}
	fmt.Printf("write %d bytes\n", cnt)

	// LandXMLデータを作成したアプリケーション情報
	f.WriteString("\t<Application name=\"Stera_Dem2Geo_Go\" version=\"0.1.0 \" manufacturer=\"Urban Planning Cyber Studio\" manufacturerURL=\"https://lisa9500jp.wixsite.com/upcs\">\n")
	f.WriteString("\t\t<Author createdBy=\"Toshio Yamazaki\"/>\n")
	f.WriteString("\t</Application>\n")

	// 座標参照系の情報
	f.WriteString("\t<CoordinateSystem name=\"CRS1\" horizontalDatum=\"JGD2000\" verticalDatum=\"O.P\" horizontalCoordinateSystemName=\"9(X,Y)\" desc=\"第9系\">\n")
	f.WriteString("\t\t<Feature>\n")
	f.WriteString("\t\t\t<Property label=\"differTP\" value=\"-1.3000\"/>\n")
	f.WriteString("\t\t</Feature>\n")
	f.WriteString("\t</CoordinateSystem>\n")

	// LandXMLで利用する単位の設定
	f.WriteString("\t<Units>\n")
	f.WriteString("\t\t<Metric areaUnit=\"squareMeter\" linearUnit=\"meter\" volumeUnit=\"cubicMeter\" temperatureUnit=\"celsius\" pressureUnit=\"HPA\" />\n")
	f.WriteString("\t</Units>\n")

	// TIN（不等辺三角形網）で地形の３次元形状を表現
	// idの設定
	id := 1
	// Rowの設定
	row := y_dot - 1
	// Columnの設定
	col := x_dot - 1

	// 要素種別サーフェス定義
	f.WriteString("\t<Surfaces>\n")
	f.WriteString("\t\t<Surface name=\"現況地形\" desc=\"ExistingGround\">\n")
	def := "\t\t\t<Definition surfType=\"TIN\" elevMin=\"" + strconv.FormatFloat(z_min*0.0254, 'f', -1, 64) + "\" elevMax=\"" + strconv.FormatFloat(z_max*0.0254, 'f', -1, 64) + "\">\n"
	f.WriteString(def)

	// 点集合
	f.WriteString("\t\t\t\t<Pnts>\n")
	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			pnt := "\t\t\t\t\t<P id=\"" + strconv.Itoa(id) + "\">" + strconv.FormatFloat(x_matrix[i][j]*0.0254, 'f', -1, 64) + " " + strconv.FormatFloat(y_matrix[i][j]*0.0254, 'f', -1, 64) + " " + strconv.FormatFloat(z_matrix[i][j]*0.0254, 'f', -1, 64) + "</P>\n"
			f.WriteString(pnt)
			id = id + 1
		}
	}
	f.WriteString("\t\t\t\t</Pnts>\n")

	// 点（X座標，Y座標，Z座標）
	f.WriteString("\t\t\t\t<Faces>\n")
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fce1 := "\t\t\t\t\t<F>" + strconv.Itoa(x_dot*(i-1)+j) + " " + strconv.Itoa(x_dot*(i-1)+j+1) + " " + strconv.Itoa(x_dot*i+j) + "</F>\n"
			f.WriteString(fce1)
			fce2 := "\t\t\t\t\t<F>" + strconv.Itoa(x_dot*(i-1)+j+1) + " " + strconv.Itoa(x_dot*i+j) + " " + strconv.Itoa(x_dot*i+j+1) + "</F>\n"
			f.WriteString(fce2)
		}
	}

	f.WriteString("\t\t\t\t</Faces>\n")
	f.WriteString("\t\t\t</Definition>\n")
	f.WriteString("\t\t</Surface>\n")
	f.WriteString("\t</Surfaces>\n")

	// LandXMLファイルの出力終了
	f.WriteString("</LandXML>")
}
