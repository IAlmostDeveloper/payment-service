package service

import (
	"fmt"
	entities "payment-service/Entities"
	"strconv"
	"strings"
)

func SimpleLuhnCheck(cardNumber string) bool {
	if len(cardNumber) != 16 {
		return false
	}
	a := strings.Split(cardNumber, "")
	sum := 0
	for i, s := range a {
		num, _ := strconv.Atoi(s)
		if i%2 == 0 {
			if 2*num > 9 {
				sum += 2*num - 9
			} else {
				sum += 2 * num
			}
		} else {
			sum += num
		}
	}
	fmt.Println(sum)
	return sum%10 == 0 && sum > 0
}

func ValidateCard(cardData entities.CardData) bool {
	return SimpleLuhnCheck(cardData.Number)
}
