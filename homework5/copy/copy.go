package copy

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func copy(src string, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Init() {

	//go run main.go -src="C:\test\tst.txt" -dst="C:\test1\tst2.txt"

	var src, dst string
	flag.StringVar(&src, "src", "", "src")
	flag.StringVar(&dst, "dst", "", "dst")

	flag.Parse()

	fmt.Println("src", src)
	fmt.Println("dst", dst)

	nBytes, err := copy(src, dst)
	if err != nil {
		fmt.Println("copyError", err)
	}

	fmt.Println("copyNBytes", nBytes)

}
