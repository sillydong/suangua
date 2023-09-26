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
	Qian:  "乾（天）1",
	Kan:   "坎（水）6",
	Gen:   "艮（山）7",
	Zheng: "震（雷）4",
	Xun:   "巽（风）5",
	Li:    "离（火）3",
	Kun:   "坤（地）8",
	Dui:   "兑（泽）2",
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
	{Qian, Qian}:   "1. 乾为天，刚健如龙，过刚则悔",
	{Kan, Qian}:    "5. 水天需，饮食宴乐，待机而动",
	{Gen, Qian}:    "26. 山天大畜，储存资源，养精蓄锐",
	{Zheng, Qian}:  "34. 雷天大壮，野马奔腾，扣锁缰绳",
	{Xun, Qian}:    "9. 风天小畜，密云不雨，雨过天晴",
	{Li, Qian}:     "14. 火天大有，阳光普照，万事如意",
	{Kun, Qian}:    "11. 地天泰，三阳开泰，万事亨通",
	{Dui, Qian}:    "43. 泽天决，老虎啸月，果敢勇决",
	{Qian, Kan}:    "6. 天水讼，争讼无益，难达目的",
	{Kan, Kan}:     "29. 坎为水，陷入漩涡，烦恼不安",
	{Gen, Kan}:     "4. 山水蒙，求知极佳，寻物其难",
	{Zheng, Kan}:   "40. 雷水解，迎刃而解，宽大为怀",
	{Xun, Kan}:     "59. 风水涣，离人起航，一帆风顺",
	{Li, Kan}:      "64. 火水未济，时运不济，有待来兹",
	{Kun, Kan}:     "7. 地水师，容名畜众，需要支持",
	{Dui, Kan}:     "47. 泽水困，困苦贫乏，志不得伸",
	{Qian, Gen}:    "33. 天山遁，逃避现实，没落衰退",
	{Kan, Gen}:     "39. 水山蹇，危险难行，不宜妄动",
	{Gen, Gen}:     "52. 艮为山，屹立不移，停止不前",
	{Zheng, Gen}:   "62. 雷山小过，行为过度，竞见相背",
	{Xun, Gen}:     "53. 风山渐，鸿鸟齐飞，有女于归",
	{Li, Gen}:      "56. 火山旅，艰苦旅程，小亨贞吉",
	{Kun, Gen}:     "15. 地山谦，谦和诚实，集多益寡",
	{Dui, Gen}:     "31. 泽山咸，二气感应，得心应手",
	{Qian, Zheng}:  "25. 天雷无妄，听其自然，顺天从命",
	{Kan, Zheng}:   "3. 水雷屯，暂时郁结，等待新机",
	{Gen, Zheng}:   "27. 山雷颐，颐养合作，当心病祸",
	{Zheng, Zheng}: "51. 震为雷，雷声隆隆，震惊百里",
	{Xun, Zheng}:   "42. 风雷益，损上益下，利人利己",
	{Li, Zheng}:    "21. 火雷噬嗑，积极行动，克服困难",
	{Kun, Zheng}:   "24. 地雷复，春回大地，重新开始",
	{Dui, Zheng}:   "17. 泽雷随，由强变弱，随从机动",
	{Qian, Xun}:    "44. 天风姤，不期而遇，结婚不宜",
	{Kan, Xun}:     "48. 水风井，无丧无得，往来井井",
	{Gen, Xun}:     "18. 山风蛊，内部不安，危机潜伏",
	{Zheng, Xun}:   "32. 雷风恒，保持常态，平淡不变",
	{Xun, Xun}:     "57. 巽为风，随风飘浮，往来不定",
	{Li, Xun}:      "50. 火风鼎，鼎力合作，大亨养贤",
	{Kun, Xun}:     "46. 地风升，种子发芽，柔以时升",
	{Dui, Xun}:     "28. 泽风大过，负担过重，本未衰弱",
	{Qian, Li}:     "13. 天火同人，同心同力，与人交往",
	{Kan, Li}:      "63. 水火既济，功成名就，和谐顺心",
	{Gen, Li}:      "22. 山火贲，绚丽晚霞，好景不常",
	{Zheng, Li}:    "55. 雷火丰，成果丰硕，运势转衰",
	{Xun, Li}:      "37. 风火家人，正家定位，温暖在内",
	{Li, Li}:       "30. 离为火，光明艳丽，炎烈不前",
	{Kun, Li}:      "36. 地火明夷，落日余晖，休养生息",
	{Dui, Li}:      "49. 泽火革，新陈代谢，己需转机",
	{Qian, Kun}:    "12. 天地否，闭塞不通，运气不佳",
	{Kan, Kun}:     "8. 水地比，水田丰满，亲比竞争",
	{Gen, Kun}:     "23. 山地剥，外强中空，危在旦夕",
	{Zheng, Kun}:   "16. 雷地豫，细心策划，从长计议",
	{Xun, Kun}:     "20. 风地观，万事难行，检讨反省",
	{Li, Kun}:      "35. 火地晋，如日东升，充满活力",
	{Kun, Kun}:     "2. 坤为地，柔顺包容，缺积极性",
	{Dui, Kun}:     "45. 泽地萃，人文荟萃，欣欣向荣",
	{Qian, Dui}:    "10. 天泽履，跟从别人，不可冒进",
	{Kan, Dui}:     "60. 水泽节，循序渐进，适可而止",
	{Gen, Dui}:     "41. 山泽损，损下益上，奉献储畜",
	{Zheng, Dui}:   "54. 雷泽归妹，少女出嫁，运势稍待",
	{Xun, Dui}:     "61. 风泽中孚，忠心诚信，成功在望",
	{Li, Dui}:      "38. 火泽睽，事与愿违，小事尚吉",
	{Kun, Dui}:     "19. 地泽临，随机应变，与时推移",
	{Dui, Dui}:     "58. 兑为泽，爽朗喜悦，丽泽多言",
}
