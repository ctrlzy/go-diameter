// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datatype_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

func TestUnsigned64(t *testing.T) {
	n := datatype.Unsigned64(math.MaxUint64)
	b := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if x := n.Serialize(); !bytes.Equal(b, x) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, x)
	}
	if n.Len() != 8 {
		t.Fatalf("Unexpected len. Want 8, have %d", n.Len())
	}
	if n.Padding() != 0 {
		t.Fatalf("Unexpected padding. Want 0, have %d", n.Padding())
	}
	if n.Type() != datatype.Unsigned64Type {
		t.Fatalf("Unexpected type. Want %d, have %d",
			datatype.Unsigned64Type, n.Type())
	}
	if len(n.String()) == 0 {
		t.Fatalf("Unexpected empty string")
	}
}

func TestDecodeUnsigned64(t *testing.T) {
	b := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	n, err := datatype.DecodeUnsigned64(b)
	if err != nil {
		t.Fatal(err)
	}
	z := uint64(math.MaxUint64)
	if uint64(n.(datatype.Unsigned64)) != z {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", z, n)
	}
}

func BenchmarkUnsigned64(b *testing.B) {
	v := datatype.Unsigned64(math.MaxUint64)
	for n := 0; n < b.N; n++ {
		v.Serialize()
	}
}

func BenchmarkDecodeUnsigned64(b *testing.B) {
	v := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	for n := 0; n < b.N; n++ {
		datatype.DecodeUnsigned64(v)
	}
}
