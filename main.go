package main

import (
	"fmt"
)

/*
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
*/

/*type Tomato struct {
	Variety string
	Weight  int
}

func (t Tomato) Pick() (string, error) {
	if t.Weight < 100 {
		return "", errors.New("Tomato is too small to pick")
	}
	return "Picked a " + t.Variety + " tomato weighing " + fmt.Sprint(t.Weight) + " grams.", nil
}
*/

func main() {
	fmt.Println("any type")

	var inventary []any

	inventary = append(inventary, "Test string")
	inventary = append(inventary, 450)
	inventary = append(inventary, true)

	fmt.Println("---Analize inventary items---")

	for _, item := range inventary {

		switch v := item.(type) {
		case string:
			fmt.Printf("Item is a string: %s\n", v)
		case int:
			fmt.Printf("Item is an int: %d\n", v)
		case bool:
			fmt.Printf("Item is a bool: %t\n", v)
		default:
			fmt.Printf("Item is of unknown type")
		}
	}

	/*
		greenTomato := Tomato{
			Variety: "Green Zebra",
			Weight:  90,
		}

		message, err := greenTomato.Pick()

		if err != nil {
			fmt.Printf("Увага: %s\n", err)
		}
		fmt.Println(message)

		/*
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
	*/

}
