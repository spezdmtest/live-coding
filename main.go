package main

import "fmt"

type Switcher interface {
	TurnOn() string
	TurnOff() string
}

type Charger interface {
	Charge() string
}

type SmartDevice interface {
	Switcher
	Charger
}

type Phone struct {
	Brand string
	Model string
}

func (p Phone) TurnOn() string {
	return p.Brand + " " + p.Model + " is turned on."
}

func (p Phone) TurnOff() string {
	return p.Brand + " " + p.Model + " is turned off."
}

func (p Phone) Charge() string {
	return p.Brand + " " + p.Model + " is charging."
}

func main() {
	//fmt.Println("Проєкт live-coding успішно запущено!")
	fmt.Println("Embedded Interfaces")

	myPhone := Phone{
		Brand: "Oukitel",
		Model: "WP19",
	}

	var device SmartDevice = myPhone

	fmt.Println("--- Test SmartDevice Interface ---")

	fmt.Println(device.TurnOn())
	fmt.Println(device.Charge())
	fmt.Println(device.TurnOff())
}
