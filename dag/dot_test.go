/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file dot_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Aug 17 21:17:34 2015
 *
 **/

package dag

import (
	"fmt"
	"testing"
)

func TestDotLoaderLoadBytes(t *testing.T) {
	d := NewDotLoader()
	g, err := d.LoadBytes([]byte(`
digraph wordcount_example {
  wordcount1 [
    name="wordcount1"
  ]
  wordcount2 [
    name="wordcount2"
  ]
  wordcount1 -> wordcount2 -> wordcount3
  wordcount1 -> wordcount3
}
`))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(g)
}

func TestDotLoaderLoadFile(t *testing.T) {
	d := NewDotLoader()
	g, err := d.LoadFile("./test.dot")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(g)
}
