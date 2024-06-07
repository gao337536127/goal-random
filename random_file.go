package GoalRandom

import (
	"io"
	"os"
)

var randomFile *os.File
var byteCache = make(chan byte, 30)

func init() {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		panic(err.Error())
	}
	randomFile = file

	go func() {
		for {
			randomFile.Seek(0, io.SeekStart)
			bytes := make([]byte, 10)
			_, err := randomFile.Read(bytes)
			if err != nil {
				panic(err.Error())
			}
			for _, b := range bytes {
				byteCache <- b
			}
		}
	}()
}
