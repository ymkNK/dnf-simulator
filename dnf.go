package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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

		if UpgradeOnce() {
			currentLevel++
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

var r = rand.New(rand.NewSource(time.Now().Unix()))

func UpgradeOnce() bool {

	// get one num in [0,100)
	num := r.Intn(100)

	// if num < 22, upgrade success
	if num < 22 {
		return true
	}

	return false
}
