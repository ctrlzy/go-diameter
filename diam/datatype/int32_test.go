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

func TestInteger32(t *testing.T) {
	n := datatype.Integer32(math.MaxInt32)
	b := []byte{0x7f, 0xff, 0xff, 0xff}
	if x := n.Serialize(); !bytes.Equal(b, x) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, x)
	}
	if n.Len() != 4 {
		t.Fatalf("Unexpected len. Want 4, have %d", n.Len())
	}
	if n.Padding() != 0 {
		t.Fatalf("Unexpected padding. Want 0, have %d", n.Padding())
	}
	if n.Type() != datatype.Integer32Type {
		t.Fatalf("Unexpected type. Want %d, have %d",
			datatype.Integer32Type, n.Type())
	}
	if len(n.String()) == 0 {
		t.Fatalf("Unexpected empty string")
	}
}

func TestNegativeInteger32(t *testing.T) {
	n := datatype.Integer32(math.MinInt32)
	b := []byte{0x80, 0, 0, 0}
	if x := n.Serialize(); !bytes.Equal(b, x) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, x)
	}
}

func TestDecodeInteger32(t *testing.T) {
	b := []byte{0x7f, 0xff, 0xff, 0xff}
	n, err := datatype.DecodeInteger32(b)
	if err != nil {
		t.Fatal(err)
	}
	z := int32(math.MaxInt32)
	if int32(n.(datatype.Integer32)) != z {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", z, n)
	}
}

func TestDecodeNegativeInteger32(t *testing.T) {
	b := []byte{0x80, 0, 0, 0}
	n, err := datatype.DecodeInteger32(b)
	if err != nil {
		t.Fatal(err)
	}
	z := int32(math.MinInt32)
	if int32(n.(datatype.Integer32)) != z {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", z, n)
	}
}

func BenchmarkInteger32(b *testing.B) {
	v := datatype.Integer32(math.MaxInt32)
	for n := 0; n < b.N; n++ {
		v.Serialize()
	}
}

func BenchmarkDecodeInteger32(b *testing.B) {
	v := []byte{0x7f, 0xff, 0xff, 0xff}
	for n := 0; n < b.N; n++ {
		datatype.DecodeInteger32(v)
	}
}
