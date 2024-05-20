// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package dict_test

import (
	"os"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam/dict"
)

var testDicts = []string{
	"./testdata/base.xml",
	"./testdata/credit_control.xml",
	"./testdata/network_access_server.xml",
	"./testdata/tgpp_ro_rf.xml",
	"./testdata/tgpp_s6a.xml",
	"./testdata/tgpp_s6c.xml",
	"./testdata/tgpp_sh.xml",
	"./testdata/tgpp_sgd_gdd.xml",
	"./testdata/tgpp_swx.xml"}

func TestNewParser(t *testing.T) {
	for _, newDict := range testDicts {
		p, err := dict.NewParser(newDict)
		if err != nil {
			t.Fatalf("Error Creating Parser from %s: %s", newDict, err)
		}
		t.Log(p)
	}
}

func TestLoadFile(t *testing.T) {
	for _, newDict := range testDicts {
		p, _ := dict.NewParser()
		if err := p.LoadFile(newDict); err != nil {
			t.Fatalf("Error Loading %s: %s", newDict, err)
		}
	}
}

func TestLoad(t *testing.T) {
	for _, newDict := range testDicts {
		f, err := os.Open(newDict)
		if err != nil {
			t.Fatalf("Error Opening %s: %s", newDict, err)
		}
		defer f.Close()
		p, _ := dict.NewParser()
		if err = p.Load(f); err != nil {
			t.Fatalf("Error Loading Parsing %s: %s", newDict, err)
		}
	}
}
