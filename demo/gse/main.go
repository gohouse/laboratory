package main

import (
	"fmt"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg gse.Segmenter
	posSeg pos.Segmenter

	new = gse.New("zh,testdata/test_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func init() {
	// 加载默认字典
	seg.LoadDict()
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")
}

func cut() {

	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	hmm = new.CutAll(`	tb := "超市一般是指商品开放陈列、顾客自我选购、排队收银结算，以经营生鲜食品水果、日杂用品为主的商店。一种消费者自助选购、统一收银结算的零售企业。在中国，超级市场被引入于1978年，当时称作自选商场。超级市场一般经销食品和日用品为主，其特点主要是"`)
	fmt.Println("cut all: ", hmm)
}

func main() {
	cut()

	segCut()
}

func posAndTrim(cut []string) {
	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	posSeg.WithGse(seg)
	po := posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func cutPos() {
	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))

	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}

func segCut() {

	// 分词文本
	tb := "超市一般是指商品开放陈列、顾客自我选购、排队收银结算，以经营生鲜食品水果、日杂用品为主的商店"

	// 处理分词结果
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb))
	fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb,true))

	segments := seg.Segment([]byte(tb))
	// 处理分词结果
	fmt.Println(gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	fmt.Println(gse.ToString(segments1, true))
}