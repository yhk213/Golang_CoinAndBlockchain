package main

import (
	"github.com/yhk213/nomadcoin/cli"
	"github.com/yhk213/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
