package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/yhk213/nomadcoin/explorer"
	"github.com/yhk213/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("please use the following flags: \n\n")
	fmt.Printf("-port:		Set the Port of the server \n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest' \n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set Port of the Server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
