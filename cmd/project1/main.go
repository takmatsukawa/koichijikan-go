package main

import (
	"fmt"
)

type character struct {
	hp    int
	maxHp int
	mp    int
	mapMP int
	name  string
	aa    string
}

var (
	player = character{
		hp: 15, maxHp: 15, mp: 15, mapMP: 15, name: "ゆうしゃ",
	}
	slime = character{
		hp: 3, maxHp: 3, mp: 0, mapMP: 0, name: "スライム", aa: "／・Д・＼\n〜〜〜〜〜",
	}
	monster = character{}
)

func init() {

}

func drawBattleScreen() {
	fmt.Println(player.name)
	fmt.Printf("HP:%d/%d MP:%d/%d\n", player.hp, player.maxHp, player.mp, player.mapMP)
	fmt.Println()

	fmt.Print(monster.aa)
	fmt.Printf("HP:%d/%d\n", monster.hp, monster.maxHp)
	fmt.Println()
}

func battle(currentMonster character) {
	monster = currentMonster
	drawBattleScreen()
}

func main() {
	battle(slime)
}
