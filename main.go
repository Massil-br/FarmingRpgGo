package main

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const InitialWidth int32 = 1280
const InitialHeight int32 = 720

var number1 = 10
var number2 = 20

var isGameRunning = false

type Human struct {
	name string
	hp   int
}

func (human *Human) SetName(name string) (string, error) {
	if len(name) > 10 {
		human.name = "error"
		return  human.name,errors.New("error lenght too high")
	}

	if human.hp <= 0 {
		return  human.name,errors.New("error hp too low")

	}

	human.name = name
	return human.name, nil
}

func main() {
	if number1 < number2 {
		fmt.Println("number2 is higher")
	}
	number1 = 5 + number2
	fmt.Println(number1)
	if number1 > number2 {
		fmt.Println("number1 is higher")
	}
	imed := Human{name: "imed", hp: 10}
	fmt.Println(imed.name)

	data, err := imed.SetName("aehmaheahemahmehamaz")
	if err != nil {
		fmt.Println("Eror while using setname", err)
	}
	fmt.Println(data)
	fmt.Println(imed.name)

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(InitialWidth, InitialHeight, "FarmingRpgGo")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(rl.GetRenderWidth())/2, int32(rl.GetRenderHeight())/2, 100, 100, rl.Blue)

		rl.EndDrawing()
	}

	defer rl.CloseWindow()

}
