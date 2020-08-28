package regression

import (
	"fmt"
	"math"
)

// LinerRegression xの値と線形回帰weightのパラメータを受け取って予測値を返す
func LinerRegression(x float64, w []float64, degree int) (float64, error) {
	if len(w) != degree {
		return 0, fmt.Errorf("error:引数が次元と一致しません")
	}
	y := 0.0
	for i := 0; i < degree; i++ {
		y += math.Pow(x, float64(i)) * w[i]
	}
	return y, nil
}
