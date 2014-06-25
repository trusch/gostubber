package main

import (
	"./stubber"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var stage = flag.Int("stage", 1, "stage 1 (prepare) or 2 (test)")

func main() {
	flag.Parse()
	if *stage == 1 {
		err := exec.Command("go", "run", "stubGenerator.go", "-in", "./stubber/stubber.go", "-name", "stubGen", "-out", "./stubber").Run()
		if err != nil {
			log.Fatal(err)
		}
		beautify()
		os.Exit(0)
	}
	if *stage == 2 {
		data, err := stubber.Get("stubGen")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(data))
		os.Exit(0)
	}
	os.Exit(1)
}

func beautify() {
	err := exec.Command("go", "fmt", "./stubber").Run()
	if err != nil {
		log.Fatal(err)
	}
}
