package main

import (
	cmds "github.com/indraniel/weaver/commands"
)

const version = "0.0.3"

func main() {
	cmds.Execute(version)
}
