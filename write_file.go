package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func write_file(file, encode, data string) error {
	var fp *os.File
	if file == "" {
		fp = os.Stdout
	} else {
		var err error
		fp, err = os.Create(file)
		if err != nil {
			return err
		}
		defer fp.Close()
	}

	if strings.ToLower(encode) == "sjis" {
		// Excel用にShiftJISで出力する
		localized_data, _, err := transform.Bytes(japanese.ShiftJIS.NewEncoder(), []byte(data))
		if err != nil {
			return err
		}
		fmt.Fprintln(fp, string(localized_data))
	} else {
		// UTF-8で出力する
		fmt.Fprintln(fp, data)
	}
	return nil
}
