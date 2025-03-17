package main

import "math/rand"

var MIN = 0
var MAX = 94

func main() {
	// random number, not secure:
	randomNumber := random(0, 10)
	println(randomNumber)

	randomString := randomString(20)
	println(randomString)

}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}
