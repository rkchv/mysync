package fsync

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type fsync struct {
	dst   string
	src   string
	cache map[string][2]bool
}

func New(src, dst string) *fsync {
	return &fsync{
		dst:   dst,
		src:   src,
		cache: make(map[string][2]bool),
	}
}

func (fsync *fsync) Start() error {

	filesSrc, err := os.ReadDir(fsync.src)
	if err != nil {
		log.Fatal(err)
	}

	filesDst, err := os.ReadDir(fsync.dst)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesSrc {
		src := filepath.Join(fsync.src, file.Name())

		fileInfo, err := os.Stat(src)
		if err != nil {
			return err
		}

		name := fileInfo.Name()
		modTime := fileInfo.ModTime()
		size := fileInfo.Size()
		mode := fileInfo.Mode()

		str := hash(fmt.Sprintf(name, modTime, size, mode))

		fsync.cache[str] = [2]bool{true, false}

	}

	for _, file := range filesDst {
		src := filepath.Join(fsync.dst, file.Name())

		fileInfo, err := os.Stat(src)
		if err != nil {
			return err
		}

		name := fileInfo.Name()
		modTime := fileInfo.ModTime()
		size := fileInfo.Size()
		mode := fileInfo.Mode()

		str := hash(fmt.Sprintf(name, modTime, size, mode))

		_, ok := fsync.cache[str]

		if ok {
			fmt.Println("in cache")
		}

		// if exist {
		// 	ok := os.SameFile(fileInfo, cached)
		// 	fmt.Println(ok)
		// }

	}

	return fmt.Errorf("my error")
}

// func (fsync *fsync) Scan() error {
// 	return fmt.Errorf("my Scan error")
// }

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

func hash(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Hello() {
	fmt.Println("Hello")
}

// hello()
