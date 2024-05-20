// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatype_test

import (
	"bytes"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

func TestFloat64(t *testing.T) {
	n := datatype.Float64(3.1415926535)
	b := []byte{0x40, 0x09, 0x21, 0xfb, 0x54, 0x41, 0x17, 0x44}
	if x := n.Serialize(); !bytes.Equal(b, x) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, x)
	}
	if n.Len() != 8 {
		t.Fatalf("Unexpected len. Want 8, have %d", n.Len())
	}
	if n.Padding() != 0 {
		t.Fatalf("Unexpected padding. Want 0, have %d", n.Padding())
	}
	if n.Type() != datatype.Float64Type {
		t.Fatalf("Unexpected type. Want %d, have %d",
			datatype.Float64Type, n.Type())
	}
	if len(n.String()) == 0 {
		t.Fatalf("Unexpected empty string")
	}
}

func TestDecodeFloat64(t *testing.T) {
	b := []byte{0x40, 0x09, 0x21, 0xfb, 0x54, 0x41, 0x17, 0x44}
	n, err := datatype.DecodeFloat64(b)
	if err != nil {
		t.Fatal(err)
	}
	if v := n.(datatype.Float64); v != 3.1415926535 {
		t.Fatalf("Unexpected value. Want 3.1415926535, have %0.4f", v)
	}
}

func BenchmarkFloat64(b *testing.B) {
	v := datatype.Float64(3.1415926535)
	for n := 0; n < b.N; n++ {
		v.Serialize()
	}
}

func BenchmarkDecodeFloat64(b *testing.B) {
	v := []byte{0x40, 0x09, 0x21, 0xfb, 0x54, 0x41, 0x17, 0x44}
	for n := 0; n < b.N; n++ {
		datatype.DecodeFloat64(v)
	}
}
