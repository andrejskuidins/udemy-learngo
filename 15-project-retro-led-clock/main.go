package main
import (
    "fmt"
    "time"
)

func main()  {
	digits := [5][17]string{
		{"###", "## ", "###", "###", "# #", "###", "###", "###", "###", "###", "   ", "   ", "###", "#  ", "## ", "# #", " # ",},
		{"# #", " # ", "  #", "  #", "# #", "#  ", "#  ", "  #", "# #", "# #", " # ", "   ", "# #", "#  ", "# #", "###", " # ",},
		{"# #", " # ", "###", "###", "###", "###", "###", "  #", "###", "###", "   ", "   ", "###", "#  ", "## ", "# #", " # ",},
		{"# #", " # ", "#  ", "  #", "  #", "  #", "# #", "  #", "# #", "  #", " # ", "   ", "# #", "#  ", "# #", "# #", "   ",},
		{"###", "###", "###", "###", "  #", "###", "###", "  #", "###", "###", "   ", "   ", "# #", "###", "# #", "# #", " # ",},
	}

	for {
		fmt.Print("\033[H\033[2J")
		t := time.Now()

		hourStr := fmt.Sprintf("%02d", t.Hour())
		f1 := int(hourStr[0] - '0')
		f2 := int(hourStr[1] - '0')

		minuteStr := fmt.Sprintf("%02d", t.Minute())
		f3 := int(minuteStr[0] - '0')
		f4 := int(minuteStr[1] - '0')

		secondStr := fmt.Sprintf("%02d", t.Second())
		f5 := int(secondStr[0] - '0')
		f6 := int(secondStr[1] - '0')

		for i := 0; i < len(digits); i++ {
			if f6 % 10 == 0 {
				fmt.Println(digits[i][12], digits[i][13], digits[i][12], digits[i][14], digits[i][15], digits[i][16])
			} else if f6 % 2 == 0 {
				fmt.Println(digits[i][f1], digits[i][f2], digits[i][10], digits[i][f3], digits[i][f4], digits[i][10], digits[i][f5], digits[i][f6])
			} else {
				fmt.Println(digits[i][f1], digits[i][f2], digits[i][11], digits[i][f3], digits[i][f4], digits[i][11], digits[i][f5], digits[i][f6])
			}
		}
		time.Sleep(time.Second)
	}
}