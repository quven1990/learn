package main

import (
	"fmt"
	"time"
)

func main() {
	MakeCake()

}
func MakeCake() {
	i := 0
	for i < 10 {
		go func() {
			fmt.Print("z")
			fmt.Print(i)
			fmt.Print("a")
		}()
		i++
	}
	time.Sleep(1 * time.Second)
}
