package cmd

import (
	"errors"
	"flag"

	"github.com/CameronGorrie/sc"
)

type Free struct {
	groupId int
	nodeId  int
	client  *sc.Client
}

func (f *Free) Run(args []string) error {
	if len(args[1:]) == 0 {
		return errors.New("no arguments provided to free ")
	}

	fs := flag.NewFlagSet("free", flag.ContinueOnError)
	fs.IntVar(&f.groupId, "gid", 0, "group id")
	fs.IntVar(&f.nodeId, "id", 0, "node id")

	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	if f.groupId != 0 {
		if err := f.client.FreeAll(int32(f.groupId)); err != nil {
			return err
		}
	}

	if f.nodeId != 0 {
		if err := f.client.NodeFree(int32(f.nodeId)); err != nil {
			return err
		}
	}

	return nil
}
