package fsync

import (
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

func (fsync *fsync) Copy() error {
	src := fsync.src
	dst := fsync.dst

	copyTree(src, dst)

	return nil
}

func copyTree(src, dst string) error {

	files, err := os.ReadDir(src)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		srcPath := filepath.Join(src, f.Name())
		dstPath := filepath.Join(dst, f.Name())

		fInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		// name := fInfo.Name()
		// modTime := fInfo.ModTime()
		// size := fInfo.Size()
		// mode := fInfo.Mode()

		// hash := hash(fmt.Sprintf(name, modTime, size, mode))

		// fsync.cache[hash] = [2]bool{true, false}

		switch fInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := makeDir(dstPath, 0755); err != nil {
				return err
			}
			if err := copyTree(srcPath, dstPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := copySymLink(srcPath, dstPath); err != nil {
				return err
			}
		default:
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}

	}

	return nil

}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)

	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func makeDir(dir string, perm os.FileMode) error {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func copySymLink(src, dst string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}
	return os.Symlink(link, dst)
}

// func hash(str string) string {
// 	h := md5.New()
// 	io.WriteString(h, str)
// 	return fmt.Sprintf("%x", h.Sum(nil))
// }

func Hello() {
	fmt.Println("Hello")
}

// hello()
