# dポイントサイトの履歴からCSVを生成する

## 概要

dポイントサイトをコピペしたテキストファイルから、CSVファイルを生成します。

dポイントのポイント履歴はエクスポートできるようになっておらず、また、表示内容も一度に表示されないといった使い勝手が悪いものです。
本当はそこからスクレイピングして取得したいですが、認証を通した上で自動的に表示内容を追加するのが面倒ですので、そこは手動で対応した上で、そこからコピペしたベタテキストファイルを読み込んで、CSVファイルを生成します。

## 使い方

dポイントサイトからコピペした情報を標準入力から受け取り、CSV化したものを標準出力に書き出します。

オプション指定をすることで、入出力ファイル名を指定できます。

Excelで操作するために文字コードをShiftJISに変更できます。

### オプション

-infile `filename`
  指定したファイルを読み込みます。

-outfile `filename`
  指定したファイルに書き出します。

-encode `sjis`
  sjisを指定するとShiftJISで書き出します。それ以外のものは全て標準のUTF8で書き出します。

