package dayan_test

import (
	"fmt"
	"testing"

	"github.com/sillydong/suangua/dayan"
)

func TestDayan(t *testing.T) {
	d := dayan.New()
	d.Run()
	fmt.Printf("爻 %v\n", d.Yaos())
	bengua := d.Gua(false)
	fmt.Println("本卦")
	fmt.Printf("%s %s\n", bengua.ShangGua, bengua.ShangGua.Name())
	fmt.Printf("%s %s\n", bengua.XiaGua, bengua.XiaGua.Name())
	fmt.Printf("%s\n", bengua.String())
	fmt.Printf("%+v\n", bengua.Name())
	biangua := d.Gua(true)
	fmt.Println("变卦")
	fmt.Printf("%s %s\n", biangua.ShangGua, biangua.ShangGua.Name())
	fmt.Printf("%s %s\n", biangua.XiaGua, biangua.XiaGua.Name())
	fmt.Printf("%s\n", biangua.String())
	fmt.Printf("%+v\n", biangua.Name())
}
