package main

import (
	"fmt"
	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	"core.google.com/p/go.net/context"
)


func main() {
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &core.BuildCfg {
		Repo: r,
		Onlne: true,
	}

	nd, err := core.NewNode(ctx, cfg)

	if err != nil {
		panic(err)
	}

	list, err :== corenet.Listen(nd, "/app/whyusleeping")
	if err != nil {
		panic(err)
	}

	fmt.Printf("I am peer: %s\n", nd.Identity.Pretty())

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return 
		}

		defer con.Close()

		fmt.Fprintln(con, "Hello! This is whyrusleepings awesome ipfs service")
		fmt.Prinf("Connection from %s \n", con.Conn().RemotePeer())
	}
}
