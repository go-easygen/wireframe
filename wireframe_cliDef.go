////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Myself <me@mine.org> (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// wireframe

type rootT struct {
	cli.Helper
	Self      *rootT      `cli:"c,config" usage:"config file\n" json:"-" parser:"jsoncfg" dft:"wireframe_cfg.json"`
	Host      string      `cli:"H,host" usage:"host addr" dft:"$HOST"`
	Port      int         `cli:"p,port" usage:"listening port\n"`
	Demo      string      `cli:"long" usage:"Now can use the \\n to arrange parameters in groups\n\t\t\tJust like what is showing here, even with extreme long usage text that can spread across multiple lines\n" dft:"$Demo"`
	Daemonize bool        `cli:"D,daemonize" usage:"daemonize the service"`
	Verbose   cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity)\n"`
}

var root = &cli.Command{
	Name: "wireframe",
	Desc: "wire framing\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2021, Myself <me@mine.org>",
	Text:   "Tool to showcase wire-framing command line app fast prototype",
	Global: true,
	Argv:   func() interface{} { t := new(rootT); t.Self = t; return t },
	Fn:     Wireframe,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Self	*rootT
//  	Host	string
//  	Port	int
//  	Demo	string
//  	Daemonize	bool
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "wireframe"
//          version   = "0.1.0"
//          date = "2021-02-14"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle)
//  	if err := cli.Root(root,
//  		cli.Tree(putDef),
//  		cli.Tree(getDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

// Wireframe - main dispatcher dumb handler
//  func Wireframe(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// put

//  func putCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*putT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Host, Opts.Port, Opts.Demo, Opts.Daemonize, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Host, rootArgv.Port, rootArgv.Demo, rootArgv.Daemonize, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei,
//  	//return nil
//  	return DoPut()
//  }
//
// DoPut implements the business logic of command `put`
//  func DoPut() error {
//  	fmt.Fprintf(os.Stderr, "Upload into service\n")
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Myself <me@mine.org>\n\n")
//  	// err := ...
//  	// clis.WarnOn("Doing Put", err)
//  	// or,
//  	// clis.AbortOn("Doing Put", err)
//  	return nil
//  }

type putT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"The file to upload from (mandatory)"`
}

var putDef = &cli.Command{
	Name: "put",
	Desc: "Upload into service",
	Text: "Usage:\n  wireframe put -i /tmp/f",
	Argv: func() interface{} { return new(putT) },
	Fn:   putCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// get

//  func getCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*getT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Host, Opts.Port, Opts.Demo, Opts.Daemonize, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Host, rootArgv.Port, rootArgv.Demo, rootArgv.Daemonize, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei, argv.Fileo,
//  	//return nil
//  	return DoGet()
//  }
//
// DoGet implements the business logic of command `get`
//  func DoGet() error {
//  	fmt.Fprintf(os.Stderr, "Get from the service\n")
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Myself <me@mine.org>\n\n")
//  	// err := ...
//  	// clis.WarnOn("Doing Get", err)
//  	// or,
//  	// clis.AbortOn("Doing Get", err)
//  	return nil
//  }

type getT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"The file to upload from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"The output file (default: some file)"`
}

var getDef = &cli.Command{
	Name: "get",
	Desc: "Get from the service",
	Text: "Usage:\n  wireframe get -o /tmp/f mandatory some more args",
	Argv: func() interface{} { return new(getT) },
	Fn:   getCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}
