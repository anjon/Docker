package main

import (
	"fmt"
	"os"
	"runtime"
)

func main()  {
	hostname, _ :=os.Hostname()
	fmt.Printf("I am using %s hosting relese party on %s/%s \n", hostname, runtime.GOOS, runtime.GOARCH)
}
