package dayan

import (
	"github.com/sillydong/suangua/common"
)

type Dayan struct {
	start uint
	yaos  []common.Yao
	one   common.Yao
	two   common.Yao
	three common.Yao
	four  common.Yao
	five  common.Yao
	six   common.Yao
}

func New() Dayan {
	return Dayan{start: 50}
}

func (d *Dayan) Run() {
	result := make([]common.Yao, 6)
	for i := 0; i < 6; i++ {
		result[i] = d.siying()
	}

	// fmt.Printf("%+v\n", result)
	d.yaos = result
	d.one = result[0]
	d.two = result[1]
	d.three = result[2]
	d.four = result[3]
	d.five = result[4]
	d.six = result[5]
}

func (d Dayan) siying() common.Yao {
	extra := uint(1)         //太极
	for i := 0; i < 3; i++ { //三演
		left := d.start - extra
		tian := common.Rand(left) //分二
		ren := uint(1)            //挂一
		di := left - tian - 1
		tianji := tian % 4 //折四
		if tianji == 0 {
			tianji = 4
		}
		diji := di % 4
		if diji == 0 {
			diji = 4
		}
		ji := ren + tianji + diji //归奇
		// fmt.Printf("tian: %d, tianji: %d, di: %d, diji: %d, ren: %d, ji: %d\n", tian, tianji, di, diji, ren, ji)
		extra = extra + ji
	}
	zu := (d.start - extra) / 4
	// fmt.Printf("%d\n", zu)

	return common.Yao(zu)
}

func (d Dayan) Yaos() []common.Yao {
	return d.yaos
}

func (d Dayan) ShangGua(bian bool) common.Gua {
	return common.Gua{X: d.six.YingYang(bian), Y: d.five.YingYang(bian), Z: d.four.YingYang(bian)}
}

func (d Dayan) XiaGua(bian bool) common.Gua {
	return common.Gua{X: d.three.YingYang(bian), Y: d.two.YingYang(bian), Z: d.one.YingYang(bian)}
}

func (d Dayan) Gua(bian bool) common.Guas {
	return common.Guas{
		ShangGua: d.ShangGua(bian),
		XiaGua:   d.XiaGua(bian),
	}
}
