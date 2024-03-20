package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

var Version = "(dev)" // set at compile time

var args struct {
	Oculus []string `arg:"-o,--oculus" help:"Oculus Username(s), up to 4" placeholder:"USERNAME"`
	Pico   []string `arg:"-p,--pico" help:"Pico Username(s), up to 4" placeholder:"USERNAME"`
	Vive   []string `arg:"-v,--vive" help:"Viveport ID(s), up to 4" placeholder:"ID"`
}

func parseArgs() error {
	p, err := arg.NewParser(arg.Config{
		IgnoreEnv: true,
	}, &args)
	if err != nil {
		return err
	}
	err = p.Parse(flags())

	var buff bytes.Buffer

	switch {
	case err == arg.ErrHelp:
		p.WriteHelp(&buff)
	case err != nil:
		p.WriteUsage(&buff)
		fmt.Fprintln(&buff, "\nError:", err.Error())
	default:
		if args.Oculus == nil && args.Pico == nil && args.Vive == nil {
			p.WriteHelp(&buff)
			fmt.Fprintln(&buff, "\nError:", "at least one account must be specified")
		}
		if len(args.Oculus) > 4 || len(args.Pico) > 4 || len(args.Vive) > 4 {
			p.WriteUsage(&buff)
			fmt.Fprintln(&buff, "\nError:", "only 4 usernames per platform supported")
		}
	}

	if buff.Len() > 0 {
		return errors.New(buff.String())
	}
	return nil
}

func flags() []string {
	if len(os.Args) == 0 {
		return nil
	}
	return os.Args[1:]
}
