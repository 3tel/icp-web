// Copyright 2024 3tel. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/urfave/cli/v2"
)

func stringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
