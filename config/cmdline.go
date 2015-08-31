/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file cmdline.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 23 20:59:30 2014
 *
 **/

package config

import (
	"flag"
	//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

var (
	Help               bool
	Verbose            bool
	WorkPath           string
	EntryFile          string
	NameNode           string
	HadoopStreamingJar string
)

func InitFlags() {
	flag.BoolVar(&Help, "help", false, "Print help message")
	flag.BoolVar(&Help, "h", false, "Print help message")
	flag.BoolVar(&Verbose, "verbose", false, "Use verbose output")
	flag.BoolVar(&Verbose, "v", false, "Use verbose output")
	flag.StringVar(&WorkPath, "path", "./", "Working path")
	flag.StringVar(&WorkPath, "p", "./", "Working path")
	flag.StringVar(&EntryFile, "flow", "", "Entry of the flow")
	flag.StringVar(&EntryFile, "f", "", "Entry of the flow")
	flag.StringVar(&NameNode, "namenode", "", "Hadoop name node url, namenode:8020 eg.")
	flag.StringVar(&HadoopStreamingJar, "jar", "", "Hadoop streaming jar")
}

func Parse() {
	flag.Parse()
}

//===================================================================
// Private
//===================================================================