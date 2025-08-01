package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// generateSecret 生成1~100的随机秘密数字（仅负责生成）
func generateSecret() int {
	//注意：实际项目中建议只在程序启动时设置一次种子（如main函数中）
	//这里为了函数独立性，演示如何在函数内设置种子（但肯能影响随机性）
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1 //生成1~100的随机数
}
func playRound() {
	secret := generateSecret() //调用函数生成的秘密数字
	attempts := 0              // 记录猜测次数

	fmt.Println("新的一局开始!我已经想好了1~100之间的数字,请开始猜测吧!")

	for {
		fmt.Print("请输入你的猜测(1~100): ")
		var guess int
		//读取用户输入并检查是否为有效数字（关键错误处理）
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("输入无效，请输入一个整数!")
			//清空输入缓冲区（避免错误输入卡住后续输入）
			_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
			continue // 跳过本次循环，重新输入
		}
		attempts++ //每次有效输入后增加尝试次数

		//判断猜测结果并给出提示
		if guess < 1 || guess > 100 {
			fmt.Println("输入不合法,请输入1~100之间的数字!")
		} else if guess < secret {
			fmt.Println("太小了!")
		} else if guess > secret {
			fmt.Println("太大了!")
		} else {
			// 猜中时输出结果并结束当前局
			fmt.Printf("恭喜你，猜对了！答案就是%d,你一共猜了%d次。", secret, attempts)
			return // 结束当前局循环
		}

	}
}
func main() {
	fmt.Println("===猜数字小游戏===")
	// 初始化随机数种子（仅需执行一次，保证全局随机性）
	// 注意：如果generateSecret函数内也设置了种子，这里可以注释掉
	// rand.Seed(time.Now().UnixNano())

	// 执行一局游戏（可扩展为多局循环）
	playRound()

	// 可选：添加多局游戏支持（用户输入y继续，否则退出）
	/*
		var again string
		fmt.Print("再玩一局？(y/n)：")
		fmt.Scan(&again)
		if again == "y" || again == "Y" {
			playRound()
		}
	*/
}
