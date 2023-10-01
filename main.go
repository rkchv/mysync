package main

import (
	fsync "github.com/rkchv/mysync/internal/fsync"
)

func main() {
	f1 := "/Users/roman/Desktop/1"
	f2 := "/Users/roman/Desktop/2"

	syncher := fsync.New(f1, f2)
	syncher.Start()
	// syncher.CopyNext()
}
