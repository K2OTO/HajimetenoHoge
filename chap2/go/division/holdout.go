package division

import (
	"fmt"
	"log"
	"math/rand"
	"pattern-recognaize/regression"
)

// HoldoutMethod divide data
// 手元のデータを二つに分割して、片方をトレーニングに使用し、片方をテストに使う手法
func HoldoutMethod(x, y []float64, per float64, valueFunc func([]float64, []float64) (float64, error), degree int) ([]float64, []float64, error) {
	if per >= 1 {
		return nil, nil, fmt.Errorf("per %f is not bigger than 1", per)
	}
	// 配列xのシャッフルを行なっている
	destX := make([]float64, len(x))
	destY := make([]float64, len(y))
	perm := rand.Perm(len(x))
	for i, v := range perm {
		destX[i] = x[v]
		destY[i] = y[v]
	}
	fmt.Println("destX, destY")
	fmt.Println(destX, destY)
	t := float64(len(x)) * per
	// テストデータとトレーニングデータに分割している
	// 全体のper%文を先頭から取り出し学習データに残りをテストデータに使用する
	trainDataX := make([]float64, int(t))
	testDataX := make([]float64, len(x)-int(t))
	trainDataY := make([]float64, int(t))
	testDataY := make([]float64, len(x)-int(t))
	for i := 0; i < int(t); i++ {
		trainDataX[i] = destX[i]
		trainDataY[i] = destY[i]
	}
	for i := int(t); i < len(x); i++ {
		testDataX[i-int(t)] = destX[i]
		testDataY[i-int(t)] = destY[i]
	}

	// 評価関数で得られた値を保存する
	var trainScoreList []float64
	var testScoreList []float64

	// 1次の近似からdegree次の近似まで求める
	for i := 1; i <= degree; i++ {
		var predYTest []float64
		var predYTrain []float64
		// 関数のweightを求める
		weight, err := regression.PolynomialFeatures(trainDataX, trainDataY, i)
		if err != nil {
			log.Fatalf("%v\n", err.Error())
		}
		// 学習データに対して予測値を出す
		for j := 0; j < len(trainDataX); j++ {
			value, err := regression.LinerRegression(trainDataX[j], weight, i)
			if err != nil {
				log.Fatalf("%v\n", err.Error())
			}
			predYTrain = append(predYTrain, value)
		}
		// テストデータに対して予測値を出す
		for j := 0; j < len(testDataX); j++ {
			value, err := regression.LinerRegression(testDataX[j], weight, i)
			if err != nil {
				log.Fatalf("%v\n", err.Error())
			}
			predYTest = append(predYTest, value)
		}
		// 評価関数を用いて値の良さを評価する
		// 評価された値をscoreListに保存する
		tr, _ := valueFunc(predYTrain, trainDataY)
		te, _ := valueFunc(predYTest, testDataY)
		trainScoreList = append(trainScoreList, tr)
		testScoreList = append(testScoreList, te)
	}
	return trainScoreList, testScoreList, nil
}
