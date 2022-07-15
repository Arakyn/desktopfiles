package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(6 + 7)
	time := time.Now()
	fmt.Println(time.Format("Mon 02-01-2006"))
}
