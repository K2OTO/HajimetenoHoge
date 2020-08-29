package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"pattern-recognaize/division"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 図の生成
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	//label
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 補助線
	p.Add(plotter.NewGrid())

	//各クラスのサンプル
	const n = 100

	// サンプルデータの作成
	var x []float64
	var y []float64
	for i := 0; i < n; i++ {
		// x座標をランダムに生成する
		x = append(x, rand.Float64()*math.Pi*2)
		// sinの値+正規分布によるノイズ
		y = append(y, math.Sin(x[i])+normRandom())
	}

	// 散布図の作成
	plot1, err := plotter.NewScatter(randomPoints(n, x, y))
	if err != nil {
		panic(err)
	}
	//色を指定する
	plot1.GlyphStyle.Color = color.RGBA{R: 155, B: 255, A: 55}
	p.Add(plot1)
	p.Legend.Add("ランダム生成データ", plot1)
	// 座標範囲
	p.X.Min = 0
	p.X.Max = 6
	p.Y.Min = -1
	p.Y.Max = 1
	// plot.pngに保存
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}

	// val1, val2, _ := divisions.holdoutMethod(x, y, 0.8, MSE, 10)
	val1, val2, err := division.HoldoutMethod(x, y, 0.8, MSE, 10)

	// インスタンスを生成
	plot2, err := plot.New()
	if err != nil {
		panic(err)
	}

	// 表示項目の設定
	plot2.Title.Text = "holdoutMethod"
	plot2.X.Label.Text = "degree"
	plot2.Y.Label.Text = "error"
	points := make(plotter.XYs, len(val1))

	for i := 0; i < len(val1); i++ {
		points[i].X = float64(i)
		points[i].Y = val1[i]
	}

	points2 := make(plotter.XYs, len(val2))

	for i := 0; i < len(val1); i++ {
		points2[i].X = float64(i)
		points2[i].Y = val2[i]
	}

	// グラフを描画
	err = plotutil.AddLinePoints(plot2, "train", points, "test", points2)
	if err != nil {
		panic(err)
	}

	// 描画結果を保存
	// "5*vg.Inch" の数値を変更すれば，保存する画像のサイズを調整できます．
	if err := plot2.Save(5*vg.Inch, 5*vg.Inch, "holdout.png"); err != nil {
		panic(err)
	}
}

// MSE 近似の良さの評価手法
// 平均二乗誤差 mean square error
func MSE(y, t []float64) (float64, error) {
	if len(y) != len(t) {
		return 0, fmt.Errorf("size %d is not equal to size %d", len(y), len(t))
	}
	sum := 0.0
	for i := 0; i < len(y); i++ {
		sum += math.Pow(y[i]-t[i], 2)
	}
	sum /= float64(len(y))
	return sum, nil
}

//正規乱数の作成
func normRandom() float64 {
	// 偏差0.1
	// ※グラフが似たのでよしとする
	deviation := 0.1
	rand.Seed(time.Now().UnixNano())
	return rand.NormFloat64() * deviation
}

// ptsに保存している
func randomPoints(n int, x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

// func matPrint(X mat64.Matrix) {
// 	fa := mat64.Formatted(X, mat64.Prefix(""), mat64.Squeeze())
// 	fmt.Printf("%v\n", fa)
// }
