package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	itemCnt := 0       // 胚子数量
	upgradeCnt := 0    // 增幅次数
	succCnt := 0       // 成功次数
	failCnt := 0       // 失败次数
	currentLevel := 12 // 胚子等级
	targetLevel := 18  // 目标等级

	for {
		if currentLevel == targetLevel {
			break
		}

		if currentLevel == 12 {
			itemCnt++
		}

		upgradeCnt++

		if UpgradeOnce() {
			currentLevel++
			succCnt++
			continue
		}

		currentLevel = 12
		failCnt++
	}

	fmt.Printf("item cnt: %d\nupgrade cnt: %d\nsucc cnt: %d\nfail cnt: %d\nsucc rate: %f",
		itemCnt, upgradeCnt, succCnt, failCnt, float32(succCnt)/float32(upgradeCnt)*100)
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
