package pkg

import (
	"log"
	"math"
)

// NodDel は近接する頂点を削除する
func DelNode(nod int, XY [][]float64, ext []float64, deg []float64) (nodz int,
	cordz [][]float64, extz []float64, degz []float64) {

	// 削除する頂点のリストを作成する
	var delLst []int
	var deltmp int

	// 頂点間の距離を算出して近接する頂点は削除する
	// １番目の頂点と最後の頂点に対する処理
	chkX := XY[0][0] - XY[nod-1][0]
	chkY := XY[0][1] - XY[nod-1][1]
	dist1 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	log.Println("dist1", dist1)
	// １番目の頂点と最後の頂点が近接している場合
	if dist1 < 0.5 {
		// 終点を削除する場合
		if math.Abs(deg[0]-90) < 5 && math.Abs(deg[nod-1]-90) > 5 {
			// 前方頂点が直角、後方頂点が直角ではないので後方頂点を削除する
			deltmp = nod - 1
			log.Println("削除ノードリストE0", deltmp) // Ctrl+/
			// １番目から最後の１つ前までの頂点に対する処理
			for i := 1; i < nod-2; i++ {
				// 次の頂点までの距離を求める
				next := i + 1
				chkX = XY[next][0] - XY[i][0]
				chkY = XY[next][1] - XY[i][1]
				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
				log.Println("dist2", dist2)
				if dist2 < 0.5 {
					// log.Println("chkX=", chkX)
					// log.Println("chkY=", chkY)
					if math.Abs(deg[next]-90) < 5 && math.Abs(deg[i]-90) > 5 {
						// 前方頂点が直角、後方頂点が直角ではないので後方頂点を削除する
						delLst = append(delLst, i)
						log.Println("削除ノードリストE1", delLst) // Ctrl+/
					} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) < 5 {
						// 後方頂点が直角、前方頂点が直角ではないので前方頂点を削除する
						delLst = append(delLst, next)
						log.Println("削除ノードリストE2", delLst) // Ctrl+/
					} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) > 5 {
						// 前方，後方頂点が共に直角ではないので，前方頂点の座標値を後方頂点との平均にする
						XY[next][0] = (XY[next][0] + XY[i][0]) / 2
						XY[next][1] = (XY[next][1] + XY[i][1]) / 2
						delLst = append(delLst, i)
						log.Println("削除ノードリストE3", delLst) // Ctrl+/
					}
				}
			}

			// 起点削除する場合
		} else if math.Abs(deg[0]-90) > 5 && math.Abs(deg[nod-1]-90) < 5 {
			// 後方頂点が直角、前方頂点が直角ではないので前方頂点を削除する
			delLst = append(delLst, 0)
			log.Println("削除ノードリストS0", delLst) // Ctrl+/
			// ２番目以降から最後の頂点に対する処理
			for i := 1; i < nod-1; i++ {
				// 次の頂点までの距離を求める
				next := i + 1
				chkX = XY[next][0] - XY[i][0]
				chkY = XY[next][1] - XY[i][1]
				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
				log.Println("dist2", dist2)
				if dist2 < 0.5 {
					// log.Println("chkX=", chkX)
					// log.Println("chkY=", chkY)
					if math.Abs(deg[next]-90) < 5 && math.Abs(deg[i]-90) > 5 {
						// 前方頂点が直角、後方頂点が直角ではないので後方頂点を削除する
						delLst = append(delLst, i)
						log.Println("削除ノードリストS1", delLst) // Ctrl+/
					} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) < 5 {
						// 後方頂点が直角、前方頂点が直角ではないので前方頂点を削除する
						delLst = append(delLst, next)
						log.Println("削除ノードリストS2", delLst) // Ctrl+/
					} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) > 5 {
						// 前方，後方頂点が共に直角ではないので，前方頂点の座標値を後方頂点との平均にする
						XY[next][0] = (XY[next][0] + XY[i][0]) / 2
						XY[next][1] = (XY[next][1] + XY[i][1]) / 2
						delLst = append(delLst, i)
						log.Println("削除ノードリストS3", delLst) // Ctrl+/
					}
				}
			}

			// 起点を平均化して終点を削除する
		} else if math.Abs(deg[0]-90) > 5 && math.Abs(deg[nod-1]-90) > 5 {
			// 前方，後方頂点が共に直角ではないので，前方頂点の座標値を後方頂点との平均にする
			XY[0][0] = (XY[0][0] + XY[nod-1][0]) / 2
			XY[0][1] = (XY[0][1] + XY[nod-1][1]) / 2
			deltmp = nod - 1
			log.Println("削除ノードリストM0", deltmp) // Ctrl+/
			// ２番目以降から最後の頂点に対する処理
			for i := 1; i < nod-2; i++ {
				// 次の頂点までの距離を求める
				next := i + 1
				chkX = XY[next][0] - XY[i][0]
				chkY = XY[next][1] - XY[i][1]
				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
				log.Println("dist2", dist2)
				if dist2 < 0.5 {
					// log.Println("chkX=", chkX)
					// log.Println("chkY=", chkY)
					if math.Abs(deg[next]-90) < 5 && math.Abs(deg[i]-90) > 5 {
						// 前方頂点が直角、後方頂点が直角ではないので後方頂点を削除する
						delLst = append(delLst, i)
						log.Println("削除ノードリストM1", delLst) // Ctrl+/
					} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) < 5 {
						// 後方頂点が直角、前方頂点が直角ではないので前方頂点を削除する
						delLst = append(delLst, next)
						log.Println("削除ノードリストM2", delLst) // Ctrl+/
						// } else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) > 5 {
					} else {
						// 前方，後方頂点が共に直角ではない、もしくは共に直角なので，前方頂点の座標値を後方頂点との平均にする
						XY[next][0] = (XY[next][0] + XY[i][0]) / 2
						XY[next][1] = (XY[next][1] + XY[i][1]) / 2
						delLst = append(delLst, i)
						log.Println("削除ノードリストM3", delLst) // Ctrl+/
					}
				}
			}
		}

		// １番目の頂点以降で頂点間の距離が近接している場合
	} else {
		// ２番目以降から最後の頂点に対する処理
		for i := 0; i < nod-1; i++ {
			// 次の頂点までの距離を求める
			next := i + 1
			chkX = XY[next][0] - XY[i][0]
			chkY = XY[next][1] - XY[i][1]
			dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
			log.Println("dist2", i, dist2)
			if dist2 < 0.5 {
				// log.Println("chkX=", chkX)
				// log.Println("chkY=", chkY)
				if math.Abs(deg[next]-90) < 5 && math.Abs(deg[i]-90) > 5 {
					// 前方頂点が直角、後方頂点が直角ではないので後方頂点を削除する
					delLst = append(delLst, i)
					log.Println("削除ノードリスト1", delLst) // Ctrl+/
				} else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) < 5 {
					// 後方頂点が直角、前方頂点が直角ではないので前方頂点を削除する
					delLst = append(delLst, next)
					log.Println("削除ノードリスト2", delLst) // Ctrl+/
					// } else if math.Abs(deg[next]-90) > 5 && math.Abs(deg[i]-90) > 5 {
				} else {
					// 前方，後方頂点が共に直角ではない、もしくは共に直角なので，前方頂点の座標値を後方頂点との平均にする
					XY[next][0] = (XY[next][0] + XY[i][0]) / 2
					XY[next][1] = (XY[next][1] + XY[i][1]) / 2
					delLst = append(delLst, i)
					log.Println("削除ノードリスト3", delLst) // Ctrl+/
				}
			}
		}
	}

	/* 	// 頂点間の距離を算出して近接する頂点は削除する
	   	// １番目の頂点と最後の頂点に対する処理
	   	chkX := XY[0][0] - XY[nod-1][0]
	   	chkY := XY[0][1] - XY[nod-1][1]
	   	dist1 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	   	log.Println("dist1", dist1)
	   	if dist1 < 0.9 {
	   		if ext[nod-1] > 0 {
	   			// 後方が右回りなので削除する
	   			deltmp = nod - 1
	   			log.Println("削除ノードリスト0", delLst) // Ctrl+/
	   			// １番目から最後の１つ前までの頂点に対する処理
	   			for i := 1; i < nod-2; i++ {
	   				// 次の頂点までの距離を求める
	   				next := i + 1
	   				chkX = XY[next][0] - XY[i][0]
	   				chkY = XY[next][1] - XY[i][1]
	   				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	   				log.Println("dist2", dist2)
	   				if dist2 < 0.9 {
	   					// log.Println("chkX=", chkX)
	   					// log.Println("chkY=", chkY)
	   					if ext[i] > 0 {
	   						// 後方が右回りなので削除する
	   						delLst = append(delLst, i)
	   						log.Println("削除ノードリスト1", delLst) // Ctrl+/
	   					} else if ext[next] > 0 {
	   						// 後方が左回り，前方が右回りなので前方を削除する
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト2", delLst) // Ctrl+/
	   					} else {
	   						// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	   						XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	   						XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト3", delLst) // Ctrl+/
	   					}
	   				}
	   			}

	   		} else if ext[0] > 0 {
	   			// 後方が左回り，前方が右回りなので前方を削除する
	   			delLst = append(delLst, 0)
	   			log.Println("削除ノードリスト0", delLst) // Ctrl+/
	   			// ２番目以降から最後の頂点に対する処理
	   			for i := 1; i < nod-1; i++ {
	   				// 次の頂点までの距離を求める
	   				next := i + 1
	   				chkX = XY[next][0] - XY[i][0]
	   				chkY = XY[next][1] - XY[i][1]
	   				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	   				log.Println("dist2", dist2)
	   				if dist2 < 0.9 {
	   					// log.Println("chkX=", chkX)
	   					// log.Println("chkY=", chkY)
	   					if ext[i] > 0 {
	   						// 後方が右回りなので削除する
	   						delLst = append(delLst, i)
	   						log.Println("削除ノードリスト1", delLst) // Ctrl+/
	   					} else if ext[next] > 0 {
	   						// 後方が左回り，前方が右回りなので前方を削除する
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト2", delLst) // Ctrl+/
	   					} else {
	   						// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	   						XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	   						XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト3", delLst) // Ctrl+/
	   					}
	   				}
	   			}

	   		} else {
	   			// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	   			XY[nod-1][0] = (XY[0][0] + XY[nod-1][0]) / 2
	   			XY[nod-1][1] = (XY[0][1] + XY[nod-1][1]) / 2
	   			delLst = append(delLst, 0)
	   			log.Println("削除ノードリスト0", delLst) // Ctrl+/
	   			// ２番目以降から最後の頂点に対する処理
	   			for i := 1; i < nod-1; i++ {
	   				// 次の頂点までの距離を求める
	   				next := i + 1
	   				chkX = XY[next][0] - XY[i][0]
	   				chkY = XY[next][1] - XY[i][1]
	   				dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	   				log.Println("dist2", dist2)
	   				if dist2 < 0.9 {
	   					// log.Println("chkX=", chkX)
	   					// log.Println("chkY=", chkY)
	   					if ext[i] > 0 {
	   						// 後方が右回りなので削除する
	   						delLst = append(delLst, i)
	   						log.Println("削除ノードリスト1", delLst) // Ctrl+/
	   					} else if ext[next] > 0 {
	   						// 後方が左回り，前方が右回りなので前方を削除する
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト2", delLst) // Ctrl+/
	   					} else {
	   						// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	   						XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	   						XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	   						delLst = append(delLst, next)
	   						log.Println("削除ノードリスト3", delLst) // Ctrl+/
	   					}
	   				}
	   			}
	   		}

	   	} else {
	   		// ２番目以降から最後の頂点に対する処理
	   		for i := 0; i < nod-1; i++ {
	   			// 次の頂点までの距離を求める
	   			next := i + 1
	   			chkX = XY[next][0] - XY[i][0]
	   			chkY = XY[next][1] - XY[i][1]
	   			dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	   			log.Println("dist2", dist2)
	   			if dist2 < 0.9 {
	   				// log.Println("chkX=", chkX)
	   				// log.Println("chkY=", chkY)
	   				if ext[i] > 0 {
	   					// 後方が右回りなので削除する
	   					delLst = append(delLst, i)
	   					log.Println("削除ノードリスト1", delLst) // Ctrl+/
	   				} else if ext[next] > 0 {
	   					// 後方が左回り，前方が右回りなので前方を削除する
	   					delLst = append(delLst, next)
	   					log.Println("削除ノードリスト2", delLst) // Ctrl+/
	   				} else {
	   					// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	   					XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	   					XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	   					delLst = append(delLst, next)
	   					log.Println("削除ノードリスト3", delLst) // Ctrl+/
	   				}
	   			}
	   		}
	   	} */

	// for i := 1; i < nod-1; i++ {
	// 	// 次の頂点までの距離を求める
	// 	next := i + 1
	// 	chkX = XY[next][0] - XY[i][0]
	// 	chkY = XY[next][1] - XY[i][1]
	// 	dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	// 	log.Println("dist2", dist2)
	// 	if dist2 < 0.9 {
	// 		// log.Println("chkX=", chkX
	// 		// log.Println("chkY=", chkY)
	// 		if ext[i] > 0 {
	// 			// 後方が右回りなので削除する
	// 			delLst = append(delLst, i)
	// 			log.Println("削除ノードリスト1", delLst) // Ctrl+/
	// 		} else if ext[next] > 0 {
	// 			// 後方が左回り，前方が右回りなので前方を削除する
	// 			delLst = append(delLst, next)
	// 			log.Println("削除ノードリスト2", delLst) // Ctrl+/
	// 		} else {
	// 			// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	// 			XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	// 			XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	// 			delLst = append(delLst, next)
	// 			log.Println("削除ノードリスト3", delLst) // Ctrl+/
	// 		}
	// 	}
	// }

	// for i := 1; i < nod-1; i++ {
	// 	// 次の頂点までの距離を求める
	// 	next := i + 1
	// 	chkX = XY[next][0] - XY[i][0]
	// 	chkY = XY[next][1] - XY[i][1]
	// 	dist2 := math.Sqrt(math.Pow(chkX, 2) + math.Pow(chkY, 2))
	// 	log.Println("dist2", dist2)
	// 	if dist2 < 0.9 {
	// 		// log.Println("chkX=", chkX)
	// 		// log.Println("chkY=", chkY)
	// 		if ext[i] > 0 {
	// 			// 後方が右回りなので削除する
	// 			delLst = append(delLst, i)
	// 			log.Println("削除ノードリスト1", delLst) // Ctrl+/
	// 		} else if ext[next] > 0 {
	// 			// 後方が左回り，前方が右回りなので前方を削除する
	// 			delLst = append(delLst, next)
	// 			log.Println("削除ノードリスト2", delLst) // Ctrl+/
	// 		} else {
	// 			// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
	// 			XY[i][0] = (XY[next][0] + XY[i][0]) / 2
	// 			XY[i][1] = (XY[next][1] + XY[i][1]) / 2
	// 			delLst = append(delLst, next)
	// 			log.Println("削除ノードリスト3", delLst) // Ctrl+/
	// 		}
	// 	}
	// }

	if deltmp > 0 {
		delLst = append(delLst, deltmp)
	}

	// 近接する頂点を削除する
	delCnt := len(delLst)
	if delCnt != 0 {
		inCnt := 0
		for i := 0; i < delCnt; i++ {
			log.Println("削除するノード", delLst[i]) // Ctrl+/
			XY = append(XY[:delLst[i]-inCnt], XY[delLst[i]+1-inCnt:]...)
			// ext = append(ext[:delLst[i]-inCnt], ext[delLst[i]+1-inCnt:]...)
			// deg = append(deg[:delLst[i]-inCnt], deg[delLst[i]+1-inCnt:]...)
			inCnt++
		}
	}
	nodz = nod - delCnt
	cordz = XY

	// Trivert を使って外戚と内角を計算する
	extz, degz, _ = TriVert(nodz, cordz)

	return nodz, cordz, extz, degz
}
