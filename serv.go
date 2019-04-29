package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvisnews/basedef"
)

func startServ() {
	fmt.Printf("jarvisnews start...\n")
	fmt.Printf("jarvisnews version is %v \n", jarvisnewsbasedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/jarvisnode.yaml")
	if err != nil {
		fmt.Printf("load jarvisnode.yaml fail!\n")

		return
	}

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	// dtd, err := dtdata.NewDTData("./cfg/config.yaml")
	// if err != nil {
	// 	fmt.Printf("NewDTData %v", err)

	// 	return
	// }

	// pprof
	jarviscore.InitPprof(cfg)

	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		fmt.Printf("jarviscore.NewNode fail! %v \n", err)

		return
	}

	node.SetNodeTypeInfo(jarvisnewsbasedef.JARVISNODETYPE, jarvisnewsbasedef.VERSION)

	// go dtd.Start(context.Background(), node)
	node.Start(context.Background())

	fmt.Printf("jarvisnews end.\n")
}
