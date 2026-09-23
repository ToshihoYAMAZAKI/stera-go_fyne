package internal

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
	"stera/pkg"
	"strings"
)

// SlopeRoof は傾斜屋根プログラム用の構造体の定義
type SlopeRoof struct {
	ID    string
	Fid   string
	Elv   float64
	Story int
	Type  string
	Area  string
	List  [][]float64
}

// Polygon は多角柱プログラム用の構造体の定義
type Polygon struct {
	ID    string
	Fid   string
	Elv   float64
	Story int
	Roof  float64
	Area  string
	List  [][]float64
}

// SquarePolyは普通建物のデータを傾斜屋根を掛けるために四角形の集まりに分割する
func SquarePoly() {
	// ログファイルを新規作成，追記，書き込み専用，パーミションは読むだけ
	file, err := os.OpenFile("hutsu.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	// 普通建物のファイルを開く
	fp, er := os.Open("data/hutsu_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer fp.Close()
	// log.Println("ファイルポインタ", fp)

	// 構造体のフィールド
	var id string
	var fid string
	var elv float64
	var st int
	var rfh float64
	var area string
	var cords [][]float64

	// 四角形データ（構造体）のスライスを作成する
	rectList := []*SlopeRoof{}

	// 多角形データ（構造体）のスライスを作成する
	polyList := []*Polygon{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		jStr := scanner.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		jStr = strings.TrimRight(jStr, ",")

		// MultiPolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [ [") {
			jStr = strings.Replace(jStr, "[ [ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ] ]", "] ]", 1)
		}
		// PolygonをLineStringに置換する
		if strings.Contains(jStr, "[ [ [") {
			jStr = strings.Replace(jStr, "[ [ [", "[ [", 1)
			jStr = strings.Replace(jStr, "] ] ]", "] ]", 1)
		}

		// JSONファイルからデータを読み出す
		id, fid, elv, rfh, st, area, _, _, cords = pkg.ParseJSON(jStr)
		// 階層データがない場合の処理
		// log.Println("st =", st)
		if st == 0 {
			st = 2
		}
		// 屋上高さのデータがある場合の処理
		// log.Println("rfh =", rfh)
		if rfh != 0.0 {
			st = int((rfh - elv) / 3.3)
		}

		// 配列の長さを取得する
		l := (len(cords))
		log.Println("l =", l)
		// 頂点データ数をチェックする，２以下の場合は処理を中止して次の行に進む
		if l <= 2 {
			// 頂点数が２以下なのでモデリングしない
			// 当該処理を飛ばして次の処理に移る
			continue
		}
		// 配列の要素を取得する
		// log.Println("配列の要素", cordnts[0][2][0])
		// 平面直角座標系のX軸は真北に向かう値が正，Y軸は真東に向かう値が正
		// QGISに読み込むことによりXY座標に変換される
		log.Println("ID", id)       // Ctrl+/
		log.Println("2次元配列", cords) // Ctrl+/

		// 傾斜屋根でモデリングする
		var slR bool
		slR = true
		// 閉じた図形かどうかを判断し頂点数を求める
		chkCls := pkg.ChkClose(l, cords)
		var numV int
		if chkCls {
			numV = l - 1
		} else {
			numV = l
			// 閉じていない図形としての頂点数
			// log.Println("閉じていない図形の頂点数", numV) // Ctrl+/
		}
		// log.Println("numV =", numV)

		// 頂点数が5以上（四角形）の場合は多角形の箱モデルでモデリングする
		if numV > 4 {
			slR = false
		}

		// 四角形の場合の処理
		// 傾斜屋根モデリングのプログラムに処理を渡す
		if slR {
			// 外積を計算して右向き・左向きを求める
			// 内積を計算して内角の角度を求める
			// 時計回りかどうか判断する
			extLst, degLst, rotate := pkg.TriVert(numV, cords)

			// 時計周りは反時計回りに修正する
			var nCords [][]float64
			var nExt, nDeg []float64
			if rotate == "cw" {
				nCords, nExt, nDeg = pkg.CwRev(numV, cords, extLst, degLst)
			} else if rotate == "ccw" {
				nCords = cords[:numV]
				nExt = extLst[:numV]
				nDeg = degLst[:numV]
			}

			// 内角が約180度の頂点を削除する
			// 対象とする内角の削除はflattenvert.goで行う
			// nod2, cord2, ext2, deg2 := pkg.FlatVert(nodz, cordz, extz, degz)
			nodz, cordz, extz, degz := pkg.FlatVert(numV, nCords, nExt, nDeg)

			// 近接している頂点を削除する
			// 頂点間の距離の計算はnododel.goで行う
			// nodz, cordz, extz, degz := pkg.DelNode(numV, nCords, nExt, nDeg)
			// nod2, cord2, ext2, deg2 := pkg.DelNode(nodz, cordz, extz, degz)
			nodx, cordx, extx, degx := pkg.DelNode(nodz, cordz, extz, degz)
			// var cord1 [][]float64
			// for i := 0; i < nodz; i++ {
			// 	cord1 = append(cord1, cordz[i])
			// }

			// 近接している頂点を削除した結果、内角が約180度となった頂点を削除する
			nod2, cord2, ext2, deg2 := pkg.FlatVert(nodx, cordx, extx, degx)

			// // 頂点数が4未満の場合は多角形の箱モデルでモデリングする
			// if nod2 < 4 {
			// 	slR = false
			// }

			// // 内角条件を設定し，満たさない内角がある場合は，四角形分割に適さない
			// var errang []int
			// for d := range deg2 {
			// 	if deg2[d] < 75.0 || deg2[d] > 105.0 {
			// 		errang = append(errang, d)
			// 		// TODO:三角メッシュの分割プログラムに渡す
			// 		// 外積の計算は済んでいるのでTriMeshDivに渡せるか？
			// 		// slR = false
			// 		log.Println("内角条件を満たさない")
			// 	}
			// }
			// log.Println("errang", errang)

			// 四角形分割のために多角形から凹頂点のL点を抽出する
			// Ｎ角形  内角数：N=2x,x=N/2，凹角数：L=x-2=N/2-2
			lcnt := nod2/2 - 2
			// L点の座標リストを作成する
			// 頂点並びのL点・R点の辞書を作成する
			// L点とR点をリストおよび辞書に振り分ける
			lL, _, order, lrPtn, lrIdx := pkg.Lexicogra(nod2, cord2, ext2)
			// log.Println("lL", lL)
			// log.Println("len(lL)", len(lL))
			log.Println("order", order)
			log.Println("lrPtn", lrPtn)
			log.Println("lrIdx", lrIdx)

			// L点と凹角数が一致しない場合は傾斜屋根でモデリングしない
			if nod2%2 == 0 {
				if lcnt != len(lL) {
					slR = false
					log.Println("L点と凹角数が一致しない")
					log.Println("slR=", slR)
				}
			}
			// log.Println("slR", slR)

			// "頂点座標"cord2と"LR並び"orderを使って多角形の分割を行う
			// 四角形分割ができない場合，三角メッシュ分割を行う
			// 三角メッシュの分割プログラムでは，L点を基準としたLR並びでパターン分けする
			// まずLRRLを捜す，次いでLRLを探す，その後LRRを捜す
			// L点が無くなったら任意のR点で扇形分割を行う
			// 三角メッシュ分割プログラムの呼び出しは，普通建物においては例外処理となる

			if nod2 == 4 {
				// 四角形に切妻屋根を掛ける
				rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: cord2}
				rectList = append(rectList, &rect0)

			} else if nod2 == 5 {
				// ５角形に変形の寄棟屋根を掛ける
				// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
				cord5, yane := pkg.PentaNode(deg2, cord2)

				rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
				rectList = append(rectList, &rect0)

			} else if nod2 == 6 {
				// ６角形の四角形分割
				_, rect1L, rect2L := pkg.HexaDiv(cord2, order)
				if rect1L == nil || rect2L == nil {
					log.Println("6角形を四角形分割できない\n", id, elv, cord2)
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				} else {
					id1 := id + "_01"
					fid1 := fid + "_01"
					rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
					rectList = append(rectList, &rect1)
					id2 := id + "_02"
					fid2 := fid + "_02"
					rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
					rectList = append(rectList, &rect2)
				}

			} else if nod2 == 7 {
				// ７角形を３つに分割して片流れ屋根を掛ける
				rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(lrPtn, deg2, cord2, order)

				if !chk7 {
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				} else if chk7 {
					id3 := id + "_03"
					fid3 := fid + "_03"
					rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
					rectList = append(rectList, &rect3)
					log.Println("rect3=", rect3)
					id1 := id + "_01"
					fid1 := fid + "_01"
					rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
					rectList = append(rectList, &rect1)
					log.Println("rect1=", rect1)
					id2 := id + "_02"
					fid2 := fid + "_02"
					rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
					rectList = append(rectList, &rect2)
					log.Println("rect2=", rect2)
				}

			} else if nod2 == 8 {
				// ８角形の四角形分割
				_, rect1L, rect2L, rect3L, story, yane := pkg.OctaDiv(cord2, order, lrPtn, lrIdx)
				if rect1L == nil || rect2L == nil || rect3L == nil {
					log.Println("8角形を四角形分割できない\n", id, elv, cord2)
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				} else {
					id1 := id + "_01"
					fid1 := fid + "_01"
					rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
					rectList = append(rectList, &rect1)
					id2 := id + "_02"
					fid2 := fid + "_02"
					rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
					rectList = append(rectList, &rect2)
					id3 := id + "_03"
					fid3 := fid + "_03"
					rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
					rectList = append(rectList, &rect3)
				}

			} else if nod2 == 10 {
				// 10角形の四角形分割
				rect1L, rect2L, rect3L, rect4L, story, yane := pkg.DecaDiv(cord2, order, lrPtn, lrIdx)
				if rect1L == nil || rect2L == nil || rect3L == nil || rect4L == nil {
					log.Println("10角形を四角形分割できない\n", id, elv, cord2)
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				} else {
					// oct1Lは８角形の四角形分割プログラムに渡されて四角形分割される
					id1 := id + "_01"
					fid1 := fid + "_01"
					rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
					rectList = append(rectList, &rect1)
					id2 := id + "_02"
					fid2 := fid + "_02"
					rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
					rectList = append(rectList, &rect2)
					id3 := id + "_03"
					fid3 := fid + "_03"
					rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
					rectList = append(rectList, &rect3)
					id4 := id + "_04"
					fid4 := fid + "_04"
					rect4 := SlopeRoof{ID: id4, Fid: fid4, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect4L}
					rectList = append(rectList, &rect4)
					// log.Println("rectList=", rectList)
				}
				log.Println("story=", story)
				log.Println("yane=", yane)

			} else if nod2 == 12 {
				// 12角形の四角形分割
				rect1L, rect2L, rect3L, rect4L, rect5L, story, yane := pkg.DodecaDiv(cord2, order, lrPtn, lrIdx)
				if rect1L == nil || rect2L == nil || rect3L == nil || rect4L == nil || rect5L == nil {
					log.Println("12角形を四角形分割できない\n", id, elv, cord2)
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				} else {
					// deca1Lは10角形の四角形分割プログラムに渡されて四角形分割される
					id1 := id + "_01"
					fid1 := fid + "_01"
					rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
					rectList = append(rectList, &rect1)
					id2 := id + "_02"
					fid2 := fid + "_02"
					rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
					rectList = append(rectList, &rect2)
					id3 := id + "_03"
					fid3 := fid + "_03"
					rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
					rectList = append(rectList, &rect3)
					id4 := id + "_04"
					fid4 := fid + "_04"
					rect4 := SlopeRoof{ID: id4, Fid: fid4, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect4L}
					rectList = append(rectList, &rect4)
					id5 := id + "_05"
					fid5 := fid + "_05"
					rect5 := SlopeRoof{ID: id5, Fid: fid5, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect5L}
					rectList = append(rectList, &rect5)
					// log.Println("rectList=", rectList)
				}
				log.Println("story=", story)
				log.Println("yane=", yane)

			} else if nod2 <= 13 {
				// 頂点数が奇数の場合，5角形を1つ分離して，残りの多角形を四角形分割する
				// 角度が最も大きく両側の頂点の角度が90度に近い頂点を選択する
				slice1, slice2, result13 := pkg.OddPoly(lrPtn, lrIdx, deg2, cord2, order)

				if result13 {
					// slice1 の内角を求め直して５角形に屋根を掛ける
					num1 := len(slice1)
					if num1 == 4 {
						// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
						chk := pkg.AnglChk(slice1)
						if !chk {
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
							rectList = append(rectList, &rect0)
						} else if chk {
							// 四角形に切妻屋根を掛ける
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
							rectList = append(rectList, &rect0)
						}

					} else if num1 == 5 {
						// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
						_, deg5, _ := pkg.TriVert(num1, slice1)
						// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
						cord5, yane := pkg.PentaNode(deg5, slice1)

						rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
						rectList = append(rectList, &rect0)
					}

					// slice2 の頂点数に応じて四角形分割を行う
					num2 := len(slice2)
					// slice2 に適用するには新しいextとdegが必要
					// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
					extmp, newdeg, _ := pkg.TriVert(num2, slice2)
					// FlatVert で内角が約180°の頂点を削除する
					newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
					// L点，R点の辞書を作り直す
					_, _, neworder, newPtn, newIdx := pkg.Lexicogra(newnum, newslice, newext)
					log.Println("newdeg=", newdeg)
					log.Println("newnum=", newnum)
					log.Println("newslice=", newslice)
					log.Println("newext=", newext)
					log.Println("neworder=", neworder)
					log.Println("newPtn=", newPtn)
					log.Println("newIdx=", newIdx)

					// newslice の頂点数に応じて屋根を掛ける
					if len(newslice) == 4 {
						// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
						chk := pkg.AnglChk(newslice)
						if !chk {
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: newslice}
							rectList = append(rectList, &rect0)
						} else if chk {
							// 四角形に切妻屋根を掛ける
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: newslice}
							rectList = append(rectList, &rect0)
						}

					} else if len(newslice) == 5 {
						// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
						_, deg5, _ := pkg.TriVert(newnum, newslice)
						// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
						cord5, yane := pkg.PentaNode(deg5, newslice)

						rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
						rectList = append(rectList, &rect0)

					} else if len(newslice) == 6 {
						// ６角形の四角形分割
						_, rect1L, rect2L := pkg.HexaDiv(newslice, neworder)
						if rect1L == nil || rect2L == nil {
							log.Println("6角形を四角形分割できない\n", id, elv, newslice)
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						} else {
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
						}

					} else if len(newslice) == 7 {
						// ７角形を３つに分割して片流れ屋根を掛ける
						rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(newPtn, newdeg, newslice, neworder)

						if !chk7 {
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
							polyList = append(polyList, &poly)
						} else if chk7 {
							id3 := id + "_03"
							fid3 := fid + "_03"
							rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
							rectList = append(rectList, &rect3)
							log.Println("rect3=", rect3)
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							log.Println("rect1=", rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
							log.Println("rect2=", rect2)
						}

					} else if len(newslice) == 8 {
						// ８角形の四角形分割
						_, rect1L, rect2L, rect3L, story, yane := pkg.OctaDiv(newslice, neworder, newPtn, newIdx)
						if rect1L == nil || rect2L == nil || rect3L == nil {
							log.Println("8角形を四角形分割できない\n", id, elv, newslice)
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						} else {
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
							id3 := id + "_03"
							fid3 := fid + "_03"
							rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
							rectList = append(rectList, &rect3)
						}

					} else if len(newslice) == 9 {
						// ９角形の四角形分割
						slice1, slice2, result9 := pkg.OddPoly(newPtn, newIdx, newdeg, newslice, neworder)

						if result9 {
							// slice1に頂点数に応じて屋根を掛ける
							num1 := len(slice1)
							if len(slice1) == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(slice1)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								}

							} else if num1 == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(num1, slice1)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, slice1)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)
							}

							// slice2 の頂点数に応じて四角形分割を行う
							num2 := len(slice2)
							// slice2 に適用するには新しいextとdegが必要
							// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
							extmp, newdeg, _ := pkg.TriVert(num2, slice2)
							// FlatVert で内角が約180°の頂点を削除する
							newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
							// L点，R点の辞書を作り直す
							_, _, neworder, newPtn, _ := pkg.Lexicogra(newnum, newslice, newext)

							// newslice の頂点数に応じて屋根を掛ける
							if len(newslice) == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(newslice)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								}

							} else if len(newslice) == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(newnum, newslice)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, newslice)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)

							} else if len(newslice) == 6 {
								// ６角形の四角形分割
								_, rect1L, rect2L := pkg.HexaDiv(newslice, neworder)
								if rect1L == nil || rect2L == nil {
									log.Println("6角形を四角形分割できない\n", id, elv, newslice)
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								} else {
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
								}
							} else if len(newslice) == 7 {
								// ７角形の四角形分割
								rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(newPtn, newdeg, newslice, neworder)

								if !chk7 {
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
									polyList = append(polyList, &poly)
								} else if chk7 {
									id3 := id + "_03"
									fid3 := fid + "_03"
									rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
									rectList = append(rectList, &rect3)
									log.Println("rect3=", rect3)
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									log.Println("rect1=", rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
									log.Println("rect2=", rect2)
								}
							}
						}
						if !result9 {
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						}

					} else if len(newslice) == 10 {
						// 10角形の四角形分割
						rect1L, rect2L, rect3L, rect4L, story, yane := pkg.DecaDiv(cord2, order, lrPtn, lrIdx)
						if rect1L == nil || rect2L == nil || rect3L == nil || rect4L == nil {
							log.Println("10角形を四角形分割できない\n", id, elv, cord2)
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
							polyList = append(polyList, &poly)
						} else {
							// oct1Lは８角形の四角形分割プログラムに渡されて四角形分割される
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
							id3 := id + "_03"
							fid3 := fid + "_03"
							rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
							rectList = append(rectList, &rect3)
							id4 := id + "_04"
							fid4 := fid + "_04"
							rect4 := SlopeRoof{ID: id4, Fid: fid4, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect4L}
							rectList = append(rectList, &rect4)
							// log.Println("rectList=", rectList)
						}
						log.Println("story=", story)
						log.Println("yane=", yane)

					} else if len(newslice) == 11 {
						// 11角形の四角形分割
						slice1, slice2, result11 := pkg.OddPoly(newPtn, newIdx, newdeg, newslice, neworder)

						if result11 {
							// slice1に頂点数に応じて屋根を掛ける
							num1 := len(slice1)
							if len(slice1) == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(slice1)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								}

							} else if num1 == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(num1, slice1)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, slice1)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)
							}

							// slice2 の頂点数に応じて四角形分割を行う
							num2 := len(slice2)
							// slice2 に適用するには新しいextとdegが必要
							// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
							extmp, newdeg, _ := pkg.TriVert(num2, slice2)
							// FlatVert で内角が約180°の頂点を削除する
							newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
							// L点，R点の辞書を作り直す
							_, _, neworder, newPtn, _ := pkg.Lexicogra(newnum, newslice, newext)

							// newslice の頂点数に応じて屋根を掛ける
							if len(newslice) == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(newslice)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								}

							} else if len(newslice) == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(newnum, newslice)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, newslice)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)

							} else if len(newslice) == 6 {
								// ６角形の四角形分割
								_, rect1L, rect2L := pkg.HexaDiv(newslice, neworder)
								if rect1L == nil || rect2L == nil {
									log.Println("6角形を四角形分割できない\n", id, elv, newslice)
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								} else {
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
								}

							} else if len(newslice) == 7 {
								// ７角形の四角形分割
								rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(newPtn, newdeg, newslice, neworder)

								if !chk7 {
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
									polyList = append(polyList, &poly)
								} else if chk7 {
									id3 := id + "_03"
									fid3 := fid + "_03"
									rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
									rectList = append(rectList, &rect3)
									log.Println("rect3=", rect3)
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									log.Println("rect1=", rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
									log.Println("rect2=", rect2)
								}

							} else if len(newslice) == 8 {
								// ８角形の四角形分割
								_, rect1L, rect2L, rect3L, story, yane := pkg.OctaDiv(newslice, neworder, newPtn, newIdx)
								if rect1L == nil || rect2L == nil || rect3L == nil {
									log.Println("8角形を四角形分割できない\n", id, elv, newslice)
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								} else {
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
									id3 := id + "_03"
									fid3 := fid + "_03"
									rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
									rectList = append(rectList, &rect3)
								}

							} else if len(newslice) == 9 {
								// ９角形の四角形分割
								slice1, slice2, result9 := pkg.OddPoly(newPtn, newIdx, newdeg, newslice, neworder)

								if result9 {
									// slice1に頂点数に応じて屋根を掛ける
									num1 := len(slice1)
									if len(slice1) == 4 {
										// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
										chk := pkg.AnglChk(slice1)
										if !chk {
											rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
											rectList = append(rectList, &rect0)
										} else if chk {
											// 四角形に切妻屋根を掛ける
											rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
											rectList = append(rectList, &rect0)
										}

									} else if num1 == 5 {
										// slice1 の内角を求め直して５角形に屋根を掛ける
										// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
										_, deg5, _ := pkg.TriVert(num1, slice1)
										// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
										cord5, yane := pkg.PentaNode(deg5, slice1)

										rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
										rectList = append(rectList, &rect0)
									}

									// slice2 の頂点数に応じて四角形分割を行う
									num2 := len(slice2)
									// slice2 に適用するには新しいextとdegが必要
									// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
									extmp, newdeg, _ := pkg.TriVert(num2, slice2)
									// FlatVert で内角が約180°の頂点を削除する
									newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
									// L点，R点の辞書を作り直す
									_, _, neworder, newPtn, _ := pkg.Lexicogra(newnum, newslice, newext)

									// newslice の頂点数に応じて屋根を掛ける
									if len(newslice) == 4 {
										// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
										chk := pkg.AnglChk(newslice)
										if !chk {
											rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: newslice}
											rectList = append(rectList, &rect0)
										} else if chk {
											// 四角形に切妻屋根を掛ける
											rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: newslice}
											rectList = append(rectList, &rect0)
										}

									} else if len(newslice) == 5 {
										// slice1 の内角を求め直して５角形に屋根を掛ける
										// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
										_, deg5, _ := pkg.TriVert(newnum, newslice)
										// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
										cord5, yane := pkg.PentaNode(deg5, newslice)

										rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
										rectList = append(rectList, &rect0)

									} else if len(newslice) == 6 {
										// ６角形の四角形分割
										_, rect1L, rect2L := pkg.HexaDiv(newslice, neworder)
										if rect1L == nil || rect2L == nil {
											log.Println("6角形を四角形分割できない\n", id, elv, newslice)
											poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
											polyList = append(polyList, &poly)
										} else {
											id1 := id + "_01"
											fid1 := fid + "_01"
											rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
											rectList = append(rectList, &rect1)
											id2 := id + "_02"
											fid2 := fid + "_02"
											rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
											rectList = append(rectList, &rect2)
										}

									} else if len(newslice) == 7 {
										// ７角形の四角形分割
										rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(newPtn, newdeg, newslice, neworder)

										if !chk7 {
											poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
											polyList = append(polyList, &poly)
										} else if chk7 {
											id3 := id + "_03"
											fid3 := fid + "_03"
											rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
											rectList = append(rectList, &rect3)
											log.Println("rect3=", rect3)
											id1 := id + "_01"
											fid1 := fid + "_01"
											rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
											rectList = append(rectList, &rect1)
											log.Println("rect1=", rect1)
											id2 := id + "_02"
											fid2 := fid + "_02"
											rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
											rectList = append(rectList, &rect2)
											log.Println("rect2=", rect2)
										}
									}
								}
								if !result9 {
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								}
							}
						}
						if !result11 {
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						}
					}
				}
				if !result13 {
					// slR = false
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				}

			} else if nod2 == 14 {
				// 14角形から四角形を分割して12角形以下の分割プログラムに渡す
				slice1, slice2, result14 := pkg.FortenPoly(lrPtn, lrIdx, deg2, cord2, order)

				if result14 {
					// slice1 の頂点数に応じて屋根を掛ける
					if len(slice1) == 4 {
						// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
						chk := pkg.AnglChk(slice1)
						if !chk {
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
							rectList = append(rectList, &rect0)
						} else if chk {
							// 四角形に切妻屋根を掛ける
							rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
							rectList = append(rectList, &rect0)
						}

					} else if len(slice1) == 5 {
						// slice1 の内角を求め直して５角形に屋根を掛ける
						// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
						_, deg5, _ := pkg.TriVert(len(slice1), slice1)
						// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
						cord5, yane := pkg.PentaNode(deg5, slice1)

						rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
						rectList = append(rectList, &rect0)
					}

					// 分割した多角形の頂点の内角を確認して分割プログラムに渡す
					// FlatVert で内角が約180°の頂点を削除する
					// slice2 に適用するには新しいextとdegが必要
					num2 := len(slice2)
					// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
					extmp, newdeg, _ := pkg.TriVert(num2, slice2)
					// FlatVert で内角が約180°の頂点を削除する
					newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
					// L点，R点の辞書を作り直す
					_, _, neworder, newPtn, newIdx := pkg.Lexicogra(newnum, newslice, newext)
					log.Println("newdeg=", newdeg)
					log.Println("newnum=", newnum)
					log.Println("newslice=", newslice)
					log.Println("newext=", newext)
					log.Println("neworder=", neworder)
					log.Println("newPtn=", newPtn)
					log.Println("newIdx=", newIdx)

					// newslice の頂点数に応じて屋根を掛ける
					if len(newslice) == 12 {
						// 12角形の四角形分割
						rect1L, rect2L, rect3L, rect4L, rect5L, story, yane := pkg.DodecaDiv(newslice, neworder, newPtn, newIdx)
						if rect1L == nil || rect2L == nil || rect3L == nil || rect4L == nil || rect5L == nil {
							log.Println("12角形を四角形分割できない\n", id, elv, newslice)
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						} else {
							// deca1Lは10角形の四角形分割プログラムに渡されて四角形分割される
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
							id3 := id + "_03"
							fid3 := fid + "_03"
							rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
							rectList = append(rectList, &rect3)
							id4 := id + "_04"
							fid4 := fid + "_04"
							rect4 := SlopeRoof{ID: id4, Fid: fid4, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect4L}
							rectList = append(rectList, &rect4)
							id5 := id + "_05"
							fid5 := fid + "_05"
							rect5 := SlopeRoof{ID: id5, Fid: fid5, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect5L}
							rectList = append(rectList, &rect5)
							// log.Println("rectList=", rectList)
						}
						log.Println("story=", story)
						log.Println("yane=", yane)

					} else if len(newslice) == 11 {
						// 11角形の四角形分割
						slice1, slice2, result11 := pkg.OddPoly(newPtn, newIdx, newdeg, newslice, neworder)

						if result11 {
							// slice1 の頂点数に応じて屋根を掛ける
							num1 := len(slice1)
							if num1 == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(slice1)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: slice1}
									rectList = append(rectList, &rect0)
								}

							} else if num1 == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(num1, slice1)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, slice1)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)
							}

							// FlatVert で内角が約180°の頂点を削除する
							// slice2 に適用するには新しいextとdegが必要
							num2 := len(slice2)
							// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
							extmp, newdeg, _ := pkg.TriVert(num2, slice2)
							// FlatVert で内角が約180°の頂点を削除する
							newnum, newslice, newext, _ := pkg.FlatVert(num2, slice2, extmp, newdeg)
							// L点，R点の辞書を作り直す
							_, _, neworder, newPtn, newIdx := pkg.Lexicogra(newnum, newslice, newext)
							log.Println("newdeg=", newdeg)
							log.Println("newnum=", newnum)
							log.Println("newslice=", newslice)
							log.Println("newext=", newext)
							log.Println("neworder=", neworder)
							log.Println("newPtn=", newPtn)
							log.Println("newIdx=", newIdx)

							// newslice の頂点数に応じて屋根を掛ける
							if len(newslice) == 4 {
								// 内角を確認して，内角が条件を満たさない場合は傾斜屋根を掛けない
								chk := pkg.AnglChk(newslice)
								if !chk {
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "flat", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								} else if chk {
									// 四角形に切妻屋根を掛ける
									rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: st, Type: "kiri", Area: area, List: newslice}
									rectList = append(rectList, &rect0)
								}

							} else if len(newslice) == 5 {
								// slice1 の内角を求め直して５角形に屋根を掛ける
								// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
								_, deg5, _ := pkg.TriVert(newnum, newslice)
								// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
								cord5, yane := pkg.PentaNode(deg5, newslice)

								rect0 := SlopeRoof{ID: id, Fid: fid, Elv: elv, Story: 2, Type: yane, Area: area, List: cord5}
								rectList = append(rectList, &rect0)

							} else if len(newslice) == 6 {
								_, rect1L, rect2L := pkg.HexaDiv(newslice, neworder)
								if rect1L == nil || rect2L == nil {
									log.Println("6角形を四角形分割できない\n", id, elv, newslice)
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								} else {
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: st, Type: "kiri", Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
								}

							} else if len(newslice) == 7 {
								// ７角形を３つに分割して片流れ屋根を掛ける
								rect1L, rect2L, rect3L, type1L, type2L, type3L, story, chk7 := pkg.HeptaDiv(newPtn, newdeg, newslice, neworder)

								if !chk7 {
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
									polyList = append(polyList, &poly)
								} else if chk7 {
									id3 := id + "_03"
									fid3 := fid + "_03"
									rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: type3L, Area: area, List: rect3L}
									rectList = append(rectList, &rect3)
									log.Println("rect3=", rect3)
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: type1L, Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									log.Println("rect1=", rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: type2L, Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
									log.Println("rect2=", rect2)
								}

							} else if len(newslice) == 8 {
								_, rect1L, rect2L, rect3L, story, yane := pkg.OctaDiv(newslice, neworder, newPtn, newIdx)
								if rect1L == nil || rect2L == nil || rect3L == nil {
									log.Println("8角形を四角形分割できない\n", id, elv, newslice)
									poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
									polyList = append(polyList, &poly)
								} else {
									id1 := id + "_01"
									fid1 := fid + "_01"
									rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
									rectList = append(rectList, &rect1)
									id2 := id + "_02"
									fid2 := fid + "_02"
									rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
									rectList = append(rectList, &rect2)
									id3 := id + "_03"
									fid3 := fid + "_03"
									rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
									rectList = append(rectList, &rect3)
								}
							}
						}
						if !result11 {
							// TODO:
							// slR = false
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						}

					} else if len(newslice) == 10 {
						rect1L, rect2L, rect3L, rect4L, story, yane := pkg.DecaDiv(newslice, neworder, newPtn, newIdx)
						if rect1L == nil || rect2L == nil || rect3L == nil || rect4L == nil {
							log.Println("10角形を四角形分割できない\n", id, elv, newslice)
							poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: newslice}
							polyList = append(polyList, &poly)
						} else {
							// oct1Lは８角形の四角形分割プログラムに渡されて四角形分割される
							id1 := id + "_01"
							fid1 := fid + "_01"
							rect1 := SlopeRoof{ID: id1, Fid: fid1, Elv: elv, Story: story[0], Type: yane[0], Area: area, List: rect1L}
							rectList = append(rectList, &rect1)
							id2 := id + "_02"
							fid2 := fid + "_02"
							rect2 := SlopeRoof{ID: id2, Fid: fid2, Elv: elv, Story: story[1], Type: yane[1], Area: area, List: rect2L}
							rectList = append(rectList, &rect2)
							id3 := id + "_03"
							fid3 := fid + "_03"
							rect3 := SlopeRoof{ID: id3, Fid: fid3, Elv: elv, Story: story[2], Type: yane[2], Area: area, List: rect3L}
							rectList = append(rectList, &rect3)
							id4 := id + "_04"
							fid4 := fid + "_04"
							rect4 := SlopeRoof{ID: id4, Fid: fid4, Elv: elv, Story: 1, Type: "kata", Area: area, List: rect4L}
							rectList = append(rectList, &rect4)
							// log.Println("rectList=", rectList)
						}
						log.Println("story=", story)
						log.Println("yane=", yane)
					}
				}

				if !result14 {
					// TODO:
					// slR = false
					poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
					polyList = append(polyList, &poly)
				}

			} else {
				// 傾斜屋根モデリングできない多角形は三角メッシュ分割する
				// 四角形分割ができなかった場合，三角メッシュに分割される
				// ポリゴンリストに追加する
				// log.Println("nod2=", nod2)
				// log.Println("deg2=", deg2)
				// log.Println("lrPtn=", lrPtn)
				// log.Println("lrIdx=", lrIdx)
				// log.Println("order=", order)
				// slR = false
				poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cord2}
				polyList = append(polyList, &poly)
			}

		} else if !slR {
			// 建物本体と屋根を分離してモデリングする
			poly := Polygon{ID: id, Fid: fid, Elv: elv, Story: st, Roof: rfh, Area: area, List: cords}
			polyList = append(polyList, &poly)
			// 内角条件を満たさない頂点を無視して屋根を掛ける
			// TODO:
			// if errang != nil {
			// 	i := 0
			// 	for e := 0; e < len(errang); e++ {
			// 		log.Println("e=", e)
			// 		j := errang[e-i]
			// 		cord2 = cord2[:j+copy(cord2[j:], cord2[j+1:])]
			// 		i = i + 1
			// 	}
			// }
			// log.Println("cord2=", cord2)
			// log.Println("len(cord2)=", len(cord2))
		}

		if er = scanner.Err(); er != nil {
			// エラー処理
			break
		}
	}

	// log.Println(rectList)
	// p90 := &rectList
	// log.Printf("pointer:%p\n", p90)
	// log.Println("rectList", unsafe.Sizeof(rectList))

	// log.Println(polyList)
	// p99 := &polyList
	// log.Printf("pointer:%p\n", p99)
	// log.Println("polyList", unsafe.Sizeof(polyList))

	// 四角形データファイルの作成
	fr, err := os.Create("./rectfile.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()
	// エンコーダーの作成
	rEncoder := gob.NewEncoder(fr)
	// エンコード
	if err := rEncoder.Encode(rectList); err != nil {
		log.Fatal(err)
	}

	// 多角形データファイルの作成
	// 三角メッシュ分割は多角柱モデリングで行う
	// ここでは多角形データを出力する
	fp, erp := os.Create("./polyfile.gob")
	if erp != nil {
		log.Fatal(erp)
	}
	defer fp.Close()
	// エンコーダーの作成
	pEncoder := gob.NewEncoder(fp)
	// エンコード
	if erp := pEncoder.Encode(polyList); erp != nil {
		log.Fatal(erp)
	}

	// セーブファイルから復元
	flr, err := os.Open("./rectfile.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer flr.Close()
	var rectList2 = make([]SlopeRoof, 0)
	// デコーダーの作成
	rDecoder := gob.NewDecoder(flr)
	// デコード
	if err := rDecoder.Decode(&rectList2); err != nil {
		log.Fatal("decode error:", err)
	}
	log.Println("rectList")

	for index, v := range rectList2 {
		log.Println(index, ":", v.Elv, v.ID, v.List)
	}

	// セーブファイルから復元
	flp, erp := os.Open("./polyfile.gob")
	if erp != nil {
		log.Fatal(erp)
	}
	defer flp.Close()
	var polyList2 = make([]Polygon, 0)
	// デコーダーの作成
	pDecoder := gob.NewDecoder(flp)
	// デコード
	if erp := pDecoder.Decode(&polyList2); erp != nil {
		log.Fatal("decode error:", erp)
	}
	log.Println("polyList")

	for index, v := range polyList2 {
		log.Println(index, ":", v.Elv, v.ID, v.List)
	}
}
