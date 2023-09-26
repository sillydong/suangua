package common

import "fmt"

type YingYang string

const (
	Yang YingYang = "——"
	Ying YingYang = "--"
)

type Yao uint

func (x Yao) YingYang(bian bool) YingYang {
	if bian {
		x = x.Bian()
	}
	switch x {
	case LaoYang, ShaoYang:
		return Yang
	case LaoYing, ShaoYing:
		return Ying
	}
	return ""
}

func (x Yao) String() string {
	return fmt.Sprintf("%d", x)
}

func (x Yao) Bian() Yao {
	switch x {
	case LaoYang: // 9
		return ShaoYing // 8
	case LaoYing: // 6
		return ShaoYang // 7
	default:
		return x
	}
}

const (
	LaoYang  Yao = 9
	ShaoYing Yao = 8
	ShaoYang Yao = 7
	LaoYing  Yao = 6
)
