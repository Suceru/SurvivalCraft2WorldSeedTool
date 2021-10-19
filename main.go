package main

import "fmt"
import "math"

func HandleSeed(seed string) int {
	var num, num2 = 0, 1
	for c := 0; c < len(seed); c++ {
		num += int(seed[c]) * num2
		num2 += 29
	}

	return num
}

func HandleNum(num int, Mixer int, Isletter bool) string {
	var mixer = int(math.Abs(float64(Mixer))) % 127
	if mixer == 0 {
		return ""
	}
	var seed = ""
	var j = 0
	for i := 1; num > 0; i += 29 {
		num -= i * mixer
		j = i
	}
	for k := j; k > 0; k -= 29 {
		seed = string(rune(mixer+(num/k))) + seed
		switch Isletter {
		case true:
			var chr = rune(mixer + (num / k))
			if chr <= 64 || (chr >= 91 && chr <= 96) || chr >= 123 {
				return ""
			}
			break
		default:
			if (mixer+(num/k)) <= 32 || (mixer+(num/k)) == 127 {
				return ""
			}
			break
		}

		num %= k
	}
	if seed != "" {
		return seed
	} else {
		return ""
	}
}

func main() {
	fmt.Println("Code written by: Suceru")
	fmt.Println("Email: fidax@qq.com")
	fmt.Println("Qq: 1246454212")
	fmt.Println("Version: 082410ZA")
	fmt.Println("=========================")
	fmt.Println("Please enter the function number:[1]Num to Seed. [2]Seed to Num. [3]Show All Seed. [4]Show Help.")
	var n int
	fmt.Print("FunctionNum:")
	fmt.Scanln(&n)
	switch n {
	case 1:
		fmt.Println("Please enter the seed number!!!")
		var num int
		fmt.Print("SeedNum:")
		fmt.Scanln(&num)
		for i := 127; i > 0; i-- {
			switch HandleNum(num, i, true) {
			case "":
				break
			default:
				switch num == HandleSeed(HandleNum(num, i, false)) {
				case true:
					fmt.Println("Seed Contents in [Example]!!!")
					fmt.Println("[" + HandleNum(num, i, false) + "]")
					fmt.Println("=========================")
					fmt.Print("Press enter to exit!")
					fmt.Scanln()
					return
				default:
					fmt.Println("ERROR")
					break
				}

				break
			}

		}
		for i := 127; i > 0; i-- {
			switch HandleNum(num, i, false) {
			case "":
				break
			default:
				switch num == HandleSeed(HandleNum(num, i, false)) {
				case true:
					fmt.Println("Seed Contents in [Example]!!!")
					fmt.Println("[" + HandleNum(num, i, false) + "]")
					fmt.Println("=========================")
					fmt.Print("Press enter to exit!")
					fmt.Scanln()
					return
				default:
					fmt.Println("ERROR")
					break
				}

				break
			}

		}
		fmt.Println("Seed Number is ERROR!!!")
		fmt.Println("=========================")
		fmt.Print("Press enter to exit!")
		fmt.Scanln()
		break

	case 2:
		fmt.Println("Please enter the seed!!!")
		var numx string
		fmt.Print("SeedContent:")
		fmt.Scanln(&numx)
		fmt.Printf("Seed Number is:[%d]\n", HandleSeed(numx))
		fmt.Println("=========================")
		fmt.Print("Press enter to exit!")
		fmt.Scanln()
		return
	case 3:
		fmt.Println("Please enter the seed number!!!")
		var numz int
		fmt.Print("SeedNum:")
		fmt.Scanln(&numz)

		var numnull bool = true
		for i := 127; i > 0; i-- {
			switch HandleNum(numz, i, false) {
			case "":
				break
			default:
				switch numz == HandleSeed(HandleNum(numz, i, false)) {
				case true:
					if numnull == true {
						fmt.Println("Seed Contents in [Example]!!!")
					}
					fmt.Println("[" + HandleNum(numz, i, false) + "]")
					numnull = false
				default:
					fmt.Println("ERROR")
					break
				}

				break
			}

		}
		if numnull == true {
			fmt.Println("Seed Number is ERROR!!!")
		}
		fmt.Println("=========================")
		fmt.Print("Press enter to exit!")
		fmt.Scanln()
		break
	case 4:

		fmt.Println("[Help] You can find the value of \"WorldSeed\" in Project.xml by opening the .scworld file and use it to get the SeedNum.")
		fmt.Println("=========================")
		fmt.Print("Press enter to exit!")
		fmt.Scanln()
	default:
		fmt.Println("Function is ERROR")
		fmt.Println("=========================")
		fmt.Print("Press enter to exit!")
		fmt.Scanln()
		break
	}

}
