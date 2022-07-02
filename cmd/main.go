package main

import (
	"marusya/cmd/commands"

	"github.com/alecthomas/kong"
)

const (
	serviceName    = "marusya"
	serviceVersion = "0.1"
)

var cli struct {
	Server commands.Server `kong:"cmd,help:'Run server'"`

	EnvFile commands.ENVFileConfig `kong:"optional,name=env-file,default=.env,help='Path to .env file'"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Name(serviceName),
		kong.Vars{
			"serviceName": serviceName,
		},
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run())
}
