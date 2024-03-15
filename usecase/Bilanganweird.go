package usecase

import (
	"basic/models"
	"fmt"
)

func Bilanganweird(N int) models.Weird {
	if N%2 == 0 {
		if N >= 2 && N <= 5 {
			fmt.Println("not weird")
		} else if N >= 6 && N <= 20 {
			fmt.Println("weird")
		} else {
			fmt.Println("not weird")
		}
	} else {
		fmt.Println("weird")
	}
	return models.Weird{}
}
