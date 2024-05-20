// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatype_test

import (
	"bytes"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

func TestFloat32(t *testing.T) {
	n := datatype.Float32(3.1415)
	b := []byte{0x40, 0x49, 0x0e, 0x56}
	if x := n.Serialize(); !bytes.Equal(b, x) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, x)
	}
	if n.Len() != 4 {
		t.Fatalf("Unexpected len. Want 4, have %d", n.Len())
	}
	if n.Padding() != 0 {
		t.Fatalf("Unexpected padding. Want 0, have %d", n.Padding())
	}
	if n.Type() != datatype.Float32Type {
		t.Fatalf("Unexpected type. Want %d, have %d",
			datatype.Float32Type, n.Type())
	}
	if len(n.String()) == 0 {
		t.Fatalf("Unexpected empty string")
	}
}

func TestDecodeFloat32(t *testing.T) {
	b := []byte{0x40, 0x49, 0x0e, 0x56}
	n, err := datatype.DecodeFloat32(b)
	if err != nil {
		t.Fatal(err)
	}
	if v := n.(datatype.Float32); v != 3.1415 {
		t.Fatalf("Unexpected value. Want 3.1414, have %0.4f", v)
	}
}

func BenchmarkFloat32(b *testing.B) {
	v := datatype.Float32(3.1415)
	for n := 0; n < b.N; n++ {
		v.Serialize()
	}
}

func BenchmarkDecodeFloat32(b *testing.B) {
	v := []byte{0x40, 0x49, 0x0e, 0x56}
	for n := 0; n < b.N; n++ {
		datatype.DecodeFloat32(v)
	}
}
