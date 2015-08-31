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
 * @file factory_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Aug 19 00:45:32 2015
 *
 **/

package dag

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDAGFactoryCreateDAGFromFile(t *testing.T) {
	f := NewFactory()
	d, err := f.CreateDAGFromFile("../example/wordcount/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}

func TestDAGFactoryCreateDAGFromBytes(t *testing.T) {
	data, err := ioutil.ReadFile("../example/wordcount/wordcount.dot")
	if err != nil {
		t.Error(err)
		return
	}
	f := NewFactory()
	d, err := f.CreateDAGFromBytes(data)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}