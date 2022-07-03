package main

import (
	"marusya/cmd/commands"
	"marusya/internal/extlib"

	"github.com/alecthomas/kong"
)

const (
	serviceName    = "namaztime"
	serviceVersion = "0.1"
)

var cli struct {
	Server commands.Server `kong:"cmd,help:'Run server'"`

	EnvFile extlib.ENVFileConfig `kong:"optional,name=env-file,default=.env,help='Path to .env file'"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Name(serviceName),
		kong.Vars{
			"serviceName":    serviceName,
			"serviceVersion": serviceVersion,
		},
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run())
}
