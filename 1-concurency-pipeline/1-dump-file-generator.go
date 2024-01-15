package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFile = 5000
const contentLength = 15000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	generatesFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	// prefer use unixNano karena unix() return valuenya sama ketika processornya cepet
	// jadi isi filenya bisa beda2
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func generatesFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%v.txt", i))
		content := randomString(contentLength)
		err := os.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("error writing file::", filename)
		}
		if i%100 == 0 && i > 0 {
			log.Printf("%v file created", i)
		}
	}

	log.Printf("%v of total file created", totalFile)
}
