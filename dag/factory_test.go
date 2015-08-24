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
	f := NewDAGFactory()
	d, err := f.CreateDAGFromFile("./test.dot")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}

func TestDAGFactoryCreateDAGFromBytes(t *testing.T) {
	data, err := ioutil.ReadFile("./test.dot")
	if err != nil {
		t.Error(err)
		return
	}
	f := NewDAGFactory()
	d, err := f.CreateDAGFromBytes(data)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d)
}
