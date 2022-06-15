package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// テキストファイルの読み込み
func read_file(file string) (string, error) {
	var fp *os.File
	if file == "" {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(file)
		if err != nil {
			return "", err
		}
		defer fp.Close()
	}

	sc := bufio.NewScanner(fp)
	item := PointInfo{}
	data := ""
	count := 0
	description_flag := false
	for sc.Scan() {
		line := sc.Text()
		count += 1

		// 空行はスキップ
		if line == "" {
			continue
		}

		// ポイント増減
		re_point_io := regexp.MustCompile(`^([+-]\d+)P$`)
		if re_point_io.MatchString(line) {
			res_point_io := re_point_io.ReplaceAllString(line, "$1")
			if count != 1 {
				// 既存のデータを出力しておく
				data += item.toString() + "\n"

				// 構造体を一旦クリア
				item.clear()

			} else {
				// ヘッダ出力
				item.header()
				data += item.toString() + "\n"
				item.clear()
			}

			// 新しく値を入れる
			item.Point = res_point_io
			continue
		}

		// 利用・獲得
		re_point_type := regexp.MustCompile(`^(利用|獲得)$`)
		if re_point_type.MatchString(line) {
			item.Type = line
			description_flag = true
			continue
		}

		// 反映日
		re_apply_date := regexp.MustCompile(`^反映日：(\d+/\d*/\d+)$`)
		if re_apply_date.MatchString(line) {
			item.Apply_date = re_apply_date.ReplaceAllString(line, "$1")
		}

		// 内容
		if description_flag {
			line = strings.Replace(line, "　", " ", -1)
			// 全角スペースと余分な前後のスペースを削除
			item.Description = strings.TrimSpace(line)
			description_flag = false
			continue
		}

		// 期間・用途限定
		if line == "期間・用途限定" {
			item.Limited = line
			continue
		}

		// 利用日
		re_use_date := regexp.MustCompile(`^利用日：(\d+/\d*/\d+).+`)
		if re_use_date.MatchString(line) {
			item.Use_date = re_use_date.ReplaceAllString(line, "$1")
		}

		// 有効期限
		re_effective_date := regexp.MustCompile(`.*有効期限：(\d+/\d*/\d+)`)
		if re_effective_date.MatchString(line) {
			item.Effective_date = re_effective_date.ReplaceAllString(line, "$1")
		}

	}
	data += item.toString() + "\n"

	return data, nil
}
