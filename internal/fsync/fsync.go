package fsync

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type fsync struct {
	dst string
	src string
}

func New(src, dst string) *fsync {
	return &fsync{
		dst: dst,
		src: src,
	}
}

func (fsync *fsync) Start() error {

	l1, err := os.ReadDir(fsync.src)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l1)

	return fmt.Errorf("my error")
}

func (fsync *fsync) CopyNext() error {
	files, err := os.ReadDir(fsync.src)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		src := filepath.Join(fsync.src, file.Name())
		// destPath := filepath.Join(fsync.dst, file.Name())

		fileInfo, err := os.Stat(src)
		if err != nil {
			return err
		}

		fmt.Println(fileInfo)
	}

	return fmt.Errorf("my error")
}

func Hello() {
	fmt.Println("Hello")
}

// hello()
