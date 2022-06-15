package main

// ポイント情報構造体
type PointInfo struct {
	Point          string // ポインt増減
	Apply_date     string // 反映日
	Type           string // 利用/獲得
	Description    string // 利用内容
	Use_date       string // 利用日
	Limited        string // 期間・用途限定
	Effective_date string // 有効期限
}

func (p *PointInfo) clear() {
	p.Point = ""
	p.Apply_date = ""
	p.Type = ""
	p.Description = ""
	p.Use_date = ""
	p.Limited = ""
	p.Effective_date = ""
}

func (p *PointInfo) header() {
	p.Point = "ポイント"
	p.Apply_date = "反映日"
	p.Type = "種別"
	p.Description = "内容"
	p.Use_date = "利用日"
	p.Limited = "期間・用途限定"
	p.Effective_date = "有効期限"
}

func (p *PointInfo) toString() string {
	line := p.Point + ","
	line += p.Apply_date + ","
	line += p.Type + ","
	line += p.Description + ","
	line += p.Limited + ","
	line += p.Use_date + ","
	line += p.Effective_date
	return line
}
