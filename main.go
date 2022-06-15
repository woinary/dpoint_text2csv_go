package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数処理
	infile := flag.String("infile", "", "input dPoint text filename")
	outfile := flag.String("outfile", "", "output CSV filename (override)")
	encode := flag.String("encode", "", "output encode (sjis or others)")
	flag.Parse()
	// fmt.Printf("file: %s\n", *file)//DEBUG

	// ファイルからの読み込み
	data, err := read_file(*infile)
	if err != nil {
		fmt.Println("file read error: " + err.Error())
		os.Exit(1)
	}

	// 出力
	if err := write_file(*outfile, *encode, data); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
