package imtest

import (
	"ali/im"
	"fmt"
)

func IOSCertSandBoxSet() {
	u := im.IOSCertSandBoxSet{Paras: make(map[string]string)}
	fmt.Println(u)
}

func IOSCertProductionSet() {
	u := im.IOSCertProductionSet{Paras: make(map[string]string)}
	fmt.Println(u)
}