package common

import "fmt"

type Gua struct {
	X YingYang
	Y YingYang
	Z YingYang
}

func (g Gua) String() string {
	return fmt.Sprintf("%s,%s,%s", g.X, g.Y, g.Z)
}

func (g Gua) Name() string {
	return bagua[g]
}

var (
	Qian  Gua = Gua{Yang, Yang, Yang}
	Kan       = Gua{Ying, Yang, Ying}
	Gen       = Gua{Yang, Ying, Ying}
	Zheng     = Gua{Ying, Ying, Yang}
	Xun       = Gua{Yang, Yang, Ying}
	Li        = Gua{Yang, Ying, Yang}
	Kun       = Gua{Ying, Ying, Ying}
	Dui       = Gua{Ying, Yang, Yang}
)

var bagua = map[Gua]string{
	Qian:  "乾（天）",
	Kan:   "坎（水）",
	Gen:   "艮（山）",
	Zheng: "震（雷）",
	Xun:   "巽（风）",
	Li:    "离（火）",
	Kun:   "坤（地）",
	Dui:   "兑（泽）",
}

type Guas struct {
	ShangGua Gua
	XiaGua   Gua
}

func (g Guas) String() string {
	return fmt.Sprintf("%s,%s", g.ShangGua.String(), g.XiaGua.String())
}

func (g Guas) Name() string {
	return liushisigua[g]
}

var liushisigua = map[Guas]string{
	{Qian, Qian}:   "乾为天，刚健如龙，过刚则悔",
	{Kan, Qian}:    "水天需，饮食宴乐，待机而动",
	{Gen, Qian}:    "山天大畜，储存资源，养精蓄锐",
	{Zheng, Qian}:  "雷天大壮，野马奔腾，扣锁缰绳",
	{Xun, Qian}:    "风天小畜，密云不雨，雨过天晴",
	{Li, Qian}:     "火天大有，阳光普照，万事如意",
	{Kun, Qian}:    "地天泰，三阳开泰，万事亨通",
	{Dui, Qian}:    "泽天决，老虎啸月，果敢勇决",
	{Qian, Kan}:    "天水讼，争讼无益，难达目的",
	{Kan, Kan}:     "坎为水，陷入漩涡，烦恼不安",
	{Gen, Kan}:     "山水蒙，求知极佳，寻物其难",
	{Zheng, Kan}:   "雷水解，迎刃而解，宽大为怀",
	{Xun, Kan}:     "风水涣，离人起航，一帆风顺",
	{Li, Kan}:      "火水未济，时运不济，有待来兹",
	{Kun, Kan}:     "地水师，容名畜众，需要支持",
	{Dui, Kan}:     "泽水困，困苦贫乏，志不得伸",
	{Qian, Gen}:    "天山遁，逃避现实，没落衰退",
	{Kan, Gen}:     "水山蹇，危险难行，不宜妄动",
	{Gen, Gen}:     "艮为山，屹立不移，停止不前",
	{Zheng, Gen}:   "雷山小过，行为过度，竞见相背",
	{Xun, Gen}:     "风山渐，鸿鸟齐飞，有女于归",
	{Li, Gen}:      "火山旅，艰苦旅程，小亨贞吉",
	{Kun, Gen}:     "地山谦，谦和诚实，集多益寡",
	{Dui, Gen}:     "泽山咸，二气感应，得心应手",
	{Qian, Zheng}:  "天雷无妄，听其自然，顺天从命",
	{Kan, Zheng}:   "水雷屯，暂时郁结，等待新机",
	{Gen, Zheng}:   "山雷颐，颐养合作，当心病祸",
	{Zheng, Zheng}: "震为雷，雷声隆隆，震惊百里",
	{Xun, Zheng}:   "风雷益，损上益下，利人利己",
	{Li, Zheng}:    "火雷噬嗑，积极行动，克服困难",
	{Kun, Zheng}:   "地雷复，春回大地，重新开始",
	{Dui, Zheng}:   "泽雷随，由强变弱，随从机动",
	{Qian, Xun}:    "天风姤，不期而遇，结婚不宜",
	{Kan, Xun}:     "水风井，无丧无得，往来井井",
	{Gen, Xun}:     "山风蛊，内部不安，危机潜伏",
	{Zheng, Xun}:   "雷风恒，保持常态，平淡不变",
	{Xun, Xun}:     "巽为风，随风飘浮，往来不定",
	{Li, Xun}:      "火风鼎，鼎力合作，大亨养贤",
	{Kun, Xun}:     "地风升，种子发芽，柔以时升",
	{Dui, Xun}:     "泽风大过，负担过重，本未衰弱",
	{Qian, Li}:     "天火同人，同心同力，与人交往",
	{Kan, Li}:      "水火既济，功成名就，和谐顺心",
	{Gen, Li}:      "山火贲，绚丽晚霞，好景不常",
	{Zheng, Li}:    "雷火丰，成果丰硕，运势转衰",
	{Xun, Li}:      "风火家人，正家定位，温暖在内",
	{Li, Li}:       "离为火，光明艳丽，炎烈不前",
	{Kun, Li}:      "地火明夷，落日余晖，休养生息",
	{Dui, Li}:      "泽火革，新陈代谢，己需转机",
	{Qian, Kun}:    "天地否，闭塞不通，运气不佳",
	{Kan, Kun}:     "水地比，水田丰满，亲比竞争",
	{Gen, Kun}:     "山地剥，外强中空，危在旦夕",
	{Zheng, Kun}:   "雷地豫，细心策划，从长计议",
	{Xun, Kun}:     "风地观，万事难行，检讨反省",
	{Li, Kun}:      "火地晋，如日东升，充满活力",
	{Kun, Kun}:     "坤为地，柔顺包容，缺积极性",
	{Dui, Kun}:     "泽地萃，人文荟萃，欣欣向荣",
	{Qian, Dui}:    "天泽履，跟从别人，不可冒进",
	{Kan, Dui}:     "水泽节，循序渐进，适可而止",
	{Gen, Dui}:     "山泽损，损下益上，奉献储畜",
	{Zheng, Dui}:   "雷泽归妹，少女出嫁，运势稍待",
	{Xun, Dui}:     "风泽中孚，忠心诚信，成功在望",
	{Li, Dui}:      "火泽睽，事与愿违，小事尚吉",
	{Kun, Dui}:     "地泽临，随机应变，与时推移",
	{Dui, Dui}:     "兑为泽，爽朗喜悦，丽泽多言",
}
