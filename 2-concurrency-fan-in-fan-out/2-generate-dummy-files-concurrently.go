package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 10000
const contentLength = 15000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

type FileInfo struct {
	Index       int
	Filename    string
	WorkerIndex int
	Err         error
}

func main() {
	log.Println("start")
	start := time.Now()

	generatesFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
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

	// pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic (creating files)
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0

	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.Filename, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				Filename: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)

	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		// dispacth N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {

				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {
					// do the jobs
					filePath := filepath.Join(tempPath, job.Filename)
					content := randomString(contentLength)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.Filename, "file generation")

					// construct the job's result, and send it to `chanOut`
					chanOut <- FileInfo{
						Filename:    job.Filename,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}

				// if `chanIn` is closed, and the remaining jobs are finished,
				// only then we mark the worker as complete.

				// ngabarin kalo goroutine ini udah done
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and then all workers are done,
	// because right after that - we need to close the `chanOut` channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
