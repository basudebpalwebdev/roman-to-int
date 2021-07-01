package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var DataMap map[byte]int

type Roman struct {
	RomanValue string `json:"roman_value"`
}

func main() {
	app := fiber.New()
	defer app.Listen(":3000")
	app.Post("/romantoint", func(c *fiber.Ctx) error {
		var roman Roman
		err := c.BodyParser(&roman)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(RomanToInt(roman.RomanValue))
	})
}

func RomanToInt(roman string) int {

	intVal := 0
	populateDataMap(&DataMap)
	temp := []byte{}
	roman = strings.ToUpper(roman)
	for i := 0; i < len(roman); i++ {
		c := roman[i]
		if len(temp) == 0 {
			temp = append(temp, c)
		} else if DataMap[temp[len(temp)-1]] >= DataMap[c] {
			intVal += ToDigit(temp)
			temp = []byte{c}
		} else {
			temp = append(temp, c)
		}
	}
	if len(temp) > 0 {
		intVal += ToDigit(temp)
	}
	return intVal
}

func ToDigit(roman []byte) int {
	digit := DataMap[roman[len(roman)-1]]
	for i := len(roman) - 2; i >= 0; i-- {
		digit -= DataMap[roman[i]]
	}
	return digit
}

func populateDataMap(dataMap *map[byte]int) {
	*dataMap = make(map[byte]int, 7)
	(*dataMap)['I'] = 1
	(*dataMap)['V'] = 5
	(*dataMap)['X'] = 10
	(*dataMap)['L'] = 50
	(*dataMap)['C'] = 100
	(*dataMap)['D'] = 500
	(*dataMap)['M'] = 1000
}
