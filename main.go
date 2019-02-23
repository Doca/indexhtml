package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
	checkDir()
}
func copyFile(){
	srcFile, err := os.Open("/mountready/index.html")
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create("/www/html/kubertestvolume/index.html") // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	check(err)

	err = destFile.Sync()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}


func checkDir(){
	if _, err := os.Stat("/www/html/kubertestvolume/"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Waiting...")
			delaySecond(10)

		}
	}
	copyFile()
}

func main() {
	checkDir()

}