package main

import (
	"fmt"

	"github.com/hysios/gonginx/config"
	"github.com/hysios/gonginx/dumper"
	"github.com/hysios/gonginx/parser"
)

func main() {
	p := parser.NewStringParser(`http{
	upstream my_backend{
		server 127.0.0.1:443;
		server 127.0.0.2:443 backup;
	}
	}`)

	conf, err := p.Parse()
	if err != nil {
		panic(err)
	}
	upstreams := conf.FindUpstreams()

	upstreams[0].AddServer(&config.UpstreamServer{
		Address: "127.0.0.1:443",
		Parameters: map[string]string{
			"weight": "5",
		},
		Flags: []string{"down"},
	})

	fmt.Println(dumper.DumpBlock(conf.Block, dumper.IndentedStyle))

}
