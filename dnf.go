package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	r             = rand.New(rand.NewSource(time.Now().Unix()))
	successNumMap = map[int]int{
		1:  1000,
		2:  1000,
		3:  1000,
		4:  1000,
		5:  800,
		6:  700,
		7:  600,
		8:  700,
		9:  600,
		10: 500,
		11: 400,
		12: 300,
	}
)

const highLevelSuccessNum = 200

func main() {
	itemCnt := 0                  // 胚子数量
	upgradeCnt := 0               // 增幅次数
	succCnt := 0                  // 成功次数
	failCnt := 0                  // 失败次数
	maxContinuousFailCnt := 0     // 最大连续失败数
	startLevel := 12              // 胚子等级
	targetLevel := 18             // 目标等级
	currentLevel := startLevel    // 起始等级
	currentContinuousFailCnt := 0 // 当前连续失败数

	for {
		if currentLevel == targetLevel {
			break
		}

		if currentLevel == startLevel {
			itemCnt++
		}

		upgradeCnt++

		newLevel, success := UpgradeOnce(currentLevel)
		if success {
			currentLevel = newLevel
			succCnt++
			maxContinuousFailCnt = int(math.Max(float64(currentContinuousFailCnt), float64(maxContinuousFailCnt)))

			currentContinuousFailCnt = 0

			continue
		}

		currentLevel = startLevel
		failCnt++
		currentContinuousFailCnt++
	}

	fmt.Printf("胚子数量: %d\n"+
		"增幅次数: %d\n"+
		"成功次数: %d\n"+
		"失败次数: %d\n"+
		"最大连败次数: %d\n"+
		"成功率: %f",
		itemCnt,
		upgradeCnt,
		succCnt,
		failCnt,
		maxContinuousFailCnt,
		float32(succCnt)/float32(upgradeCnt)*100)
}

func UpgradeOnce(currentLevel int, opts ...Option) (int, bool) {
	successNum := successNumMap[currentLevel+1]
	successNum = int(math.Max(float64(successNum), float64(highLevelSuccessNum)))

	for _, opt := range opts {
		successNum += opt()
	}

	// get one num in [0,1000)
	num := r.Intn(1000)

	// if num < successNum, upgrade success
	if num < successNum {
		return currentLevel + 1, true
	}

	return currentLevel, false
}

type Option func() int
