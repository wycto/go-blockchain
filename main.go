package main

import "wycto/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("weiyi")
	bc.SendData("我爱北京天安门，天安门上太阳升，伟大领袖毛主席，指引我们向前进")
	bc.Print()
}
