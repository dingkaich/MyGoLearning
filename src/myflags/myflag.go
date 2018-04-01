package myflags

import (
	"flag"
	"fmt"
)

func MyflagsMain() {
	port := flag.Int64("port", 6060, "listen port")
	listen := flag.String("listen", "127.0.0.1", "listen ip")
	flag.Parse()
	fmt.Println(*port, *listen)
	fmt.Println(flag.Arg(1))
}
