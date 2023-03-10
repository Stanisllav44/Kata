package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	Println("Input:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	err := incorrect(text)
	if err != nil {
		Println("Ошибка ввода:", err)
	}

	check := character(text)
	if check == false {
		arab(text)
	} else {
		rim(text)
	}
}

func incorrect(s string) error {
	replacer := strings.NewReplacer("+", "", "-", "", "*", "", "/", "",
		"I", "", "II", "", "III", "", "IV", "", "V", "", "VI", "", "VII", "", "VIII", "", "IX", "", "X", "",
		"10", "", "9", "", "8", "", "7", "", "6", "", "5", "", "4", "", "3", "", "2", "", "1", "")
	total := replacer.Replace(s)
	if total != "" {
		return Errorf("\n-> некорректный ввод")
	}
	return nil
}

func character(s string) bool {
	if strings.ContainsAny(s, "I") == true {
		return true
	} else if strings.ContainsAny(s, "V") == true {
		return true
	} else if strings.ContainsAny(s, "X") == true {
		return true
	} else {
		return false
	}
}


func toNum(s string) int {
	switch s {
	case "I":
		toInt := 1
		return toInt
	case "II":
		toInt := 2
		return toInt
	case "III":
		toInt := 3
		return toInt
	case "IV":
		toInt := 4
		return toInt
	case "V":
		toInt := 5
		return toInt
	case "VI":
		toInt := 6
		return toInt
	case "VII":
		toInt := 7
		return toInt
	case "VIII":
		toInt := 8
		return toInt
	case "IX":
		toInt := 9
		return toInt
	case "X":
		toInt := 10
		return toInt
	default:
		toInt, _ := strconv.Atoi(s)
		return toInt
	}
}


func add(a, b int) (int, error) {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		return 0, Errorf("\n-> число не может выходить за диапазон 1...10")
	}
	return a + b, nil
}


func sub(a, b int) (int, error) {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		return 0, Errorf("\n-> число не может выходить за диапазон 1...10")
	}
	if a <= b {
		return 0, Errorf("\n-> результатом работы калькулятора с римскими числами могут быть только положительные числа")
	}
	return a - b, nil
}


func mul(a, b int) (int, error) {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		return 0, Errorf("\n-> число не может выходить за диапазон 1...10")
	}
	return a * b, nil
}


func div(a, b int) (int, error) {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		return 0, Errorf("\n-> число не может выходить за диапазон 1...10")
	}
	return a / b, nil
}

func arab(s string) {
	for _, n := range s {
		switch string(n) {
		case "+":
			buf := strings.Split(s, "+")
			a := toNum(buf[0])
			b := toNum(buf[1])
			c, err := add(a, b)
			if err != nil {
				Println("Ошибка ввода:", err)
			} else {
				Println("Output:\n", c)
			}
		case "-":
			buf := strings.Split(s, "-")
			a := toNum(buf[0])
			b := toNum(buf[1])
			c, err := sub(a, b)
			if err != nil {
				Println("Ошибка ввода:", err)
			} else {
				Println("Output:\n", c)
			}
		case "/":
			buf := strings.Split(s, "/")
			a := toNum(buf[0])
			b := toNum(buf[1])
			c, err := div(a, b)
			if err != nil {
				Println("Ошибка ввода:", err)
			} else {
				Println("Output:\n", c)
			}
		case "*":
			buf := strings.Split(s, "*")
			err := check(buf[0], buf[1])
			if err != nil {
				Println("Ошибка ввода", err)
				break
			}
			a := toNum(buf[0])
			b := toNum(buf[1])
			c, err := mul(a, b)
			Println("Output:\n", c)
		default:

		}
	}

}
func rim(s string) {
	for _, n := range s {
		switch string(n) {
		case "+":
			buf := strings.Split(s, "+")
			err := check(buf[0], buf[1])
			if err != nil {
				Println("Ошибка ввода", err)
				break
			}
			a := toNum(buf[0])
			b := toNum(buf[1])
			bufc, err := add(a, b)
			if err != nil {
				Println("Ошибка ввода", err)
			} else {
				c := toArab(bufc)
				Println("Output:\n", c)
			}
		case "-":
			buf := strings.Split(s, "-")
			err := check(buf[0], buf[1])
			if err != nil {
				Println("Ошибка ввода", err)
				break
			}
			a := toNum(buf[0])
			b := toNum(buf[1])
			bufc, err := sub(a, b)
			if err != nil {
				Println("Ошибка ввода", err)
			} else {
				c := toArab(bufc)
				Println("Output:\n", c)
			}
		case "/":
			buf := strings.Split(s, "/")
			err := check(buf[0], buf[1])
			if err != nil {
				Println("Ошибка ввода", err)
				break
			}
			a := toNum(buf[0])
			b := toNum(buf[1])
			bufc, err := div(a, b)
			if err != nil {
				Println("Ошибка ввода", err)
			} else {
				c := toArab(bufc)
				Println("Output:\n", c)
			}
		case "*":
			buf := strings.Split(s, "*")
			err := check(buf[0], buf[1])
			if err != nil {
				Println("Ошибка ввода", err)
				break
			}
			a := toNum(buf[0])
			b := toNum(buf[1])
			bufc, err := mul(a, b)
			if err != nil {
				Println("Ошибка ввода", err)
			} else {
				c := toArab(bufc)
				Println("Output:\n", c)
			}
		default:

		}
	}
}
func check(s1, s2 string) error {
	var a, b bool
	if strings.ContainsAny(s1, "I") == true ||
		strings.ContainsAny(s1, "V") == true ||
		strings.ContainsAny(s1, "X") == true {
		a = true
	} else {
		a = false
	}
	if strings.ContainsAny(s2, "I") == true ||
		strings.ContainsAny(s2, "V") == true ||
		strings.ContainsAny(s2, "X") == true {
		b = true
	} else {
		b = false
	}
	if a == true && b == false {
		return Errorf("\n-> калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	} else if a == false && b == true {
		return Errorf("\n-> калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}
	return nil
}

func toArab(num int) string {
	base := "I"
	for i := 1; i < num; i++ {
		base += "I"
	}
	base = strings.ReplaceAll(base, "IIIII", "V")
	base = strings.ReplaceAll(base, "VV", "X")
	base = strings.ReplaceAll(base, "XXXXX", "L")
	base = strings.ReplaceAll(base, "LL", "C")

	base = strings.ReplaceAll(base, "LXXXX", "XC")
	base = strings.ReplaceAll(base, "XXXX", "XL")
	base = strings.ReplaceAll(base, "VIIII", "IX")
	base = strings.ReplaceAll(base, "IIII", "IV")
	return base
}
