package main

import (
	"bytes"
	"testing"
)

func BenchmarkPlain(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Plain(w, false)
		w.Reset()
	}
}

func BenchmarkPlainVar(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		PlainVar(w, false)
		w.Reset()
	}
}

func BenchmarkPlainConst(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		PlainConst(w, false)
		w.Reset()
	}
}

func BenchmarkIfDev(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		If(w, true)
		w.Reset()
	}
}

func BenchmarkIfProd(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		If(w, false)
		w.Reset()
	}
}

func BenchmarkVarDev(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Var(w, true)
		w.Reset()
	}
}

func BenchmarkVarProd(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Var(w, false)
		w.Reset()
	}
}

func BenchmarkOrDev(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Or(w, true)
		w.Reset()
	}
}

func BenchmarkOrProd(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Or(w, false)
		w.Reset()
	}
}

func BenchmarkSlice(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		Slice(w)
		w.Reset()
	}
}
