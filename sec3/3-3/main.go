package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func makeZipfile(src io.Reader) int64 {
	// zip ファイル作成
	file, err := os.Create("result.txt.gz")
	if err != nil {
		panic(err)
	}
	// zip ファイルに対する書き込み用の構造体 (io.Writer ではない)
	zipWriter := zip.NewWriter(file)

	// txt ファイルをインターフェースとして zip ファイルへ書き込む io.Writer
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	// ソースの内容を zip ファイルの writer にコピー
	writeSize, err := io.Copy(writer, src)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer zipWriter.Close()
	return writeSize
}

func main() {
	// srcfile, err := os.Open("rand.txt")
	// if err != nil {
	// 	panic(err)
	// }
	hoge := strings.NewReader(strings.Repeat("hoge", 3000))
	makeZipfile(hoge)
	// srcfile.Close()
}
