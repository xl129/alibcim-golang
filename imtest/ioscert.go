package imtest

import (
	"ali/im"
	"fmt"
)

func IOSCertSandBoxSet() {
	u := im.NewIOSCertProductionSet()
	fmt.Println(u)
}

func IOSCertProductionSet() {
	u := im.NewIOSCertSandBoxSet()
	fmt.Println(u)
}