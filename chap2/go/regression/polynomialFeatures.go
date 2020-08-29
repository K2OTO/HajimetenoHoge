package regression

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// PolynomialFeatures 多項式近似関数:
// 最小二乗法を用いている変数名はご愛敬
// wikiを詳細
// https://ja.wikipedia.org/wiki/%E5%A4%9A%E9%A0%85%E5%BC%8F%E5%9B%9E%E5%B8%B0
func PolynomialFeatures(x, y []float64, degree int) ([]float64, error) {
	Y := mat64.NewDense(len(y), 1, y)
	var xData []float64
	for i := 0; i < len(x); i++ {
		for j := 0; j < degree; j++ {
			xData = append(xData, math.Pow(x[i], float64(j)))
		}
	}
	X := mat64.NewDense(len(x), degree, xData)
	X1 := mat64.NewDense(degree, degree, nil)
	X3 := mat64.NewDense(degree, degree, nil)
	X4 := mat64.NewDense(degree, 1, nil)
	B := mat64.NewDense(degree, 1, nil)
	Xt := X.T()
	X1.Mul(Xt, X)
	X3.Inverse(X1)
	X4.Mul(Xt, Y)
	B.Mul(X3, X4)
	var ans []float64
	for i := 0; i < degree; i++ {
		ans = append(ans, B.RawRowView(i)[0])
	}
	return ans, nil
}
