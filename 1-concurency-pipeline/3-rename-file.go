package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

type FileInfo struct {
	FilePath  string
	Content   []byte
	Sum       string
	IsRenamed bool
}

func main() {
	log.Println("start")
	start := time.Now()

	// pipeline1: loop all file and read it
	chanFileContent := readFiles()

	// pipeline2: get md5 sum
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2)

	// pipeline3: rename
	chanRename1 := rename(chanFileSum)
	chanRename2 := rename(chanFileSum)
	chanRename := mergeChanFileInfo(chanRename1, chanRename2)

	// hitung
	counterTotal := 0
	counterRenamed := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "second")
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// send data to channel
			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil
		})
		if err != nil {
			log.Println("ERROR::", err.Error())
		}

		close(chanOut)
	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))

			chanOut <- fileInfo
		}
		close(chanOut)
	}()
	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	// karena variadic
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			defer wg.Done()
			for eachDataChan := range eachChan {
				chanOut <- eachDataChan
			}
		}(eachChan)
	}
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil

			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}
