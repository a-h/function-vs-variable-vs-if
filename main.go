package main

import "io"

func Plain(w io.Writer, devMode bool) {
	w.Write([]byte("prod"))
}

func PlainVar(w io.Writer, devMode bool) {
	prod := "prod"
	w.Write([]byte(prod))
}

const prod = "prod"

func PlainConst(w io.Writer, devMode bool) {
	w.Write([]byte(prod))
}

func If(w io.Writer, devMode bool) {
	if devMode {
		w.Write([]byte("dev"))
	} else {
		w.Write([]byte("prod"))
	}
}

func Var(w io.Writer, devMode bool) {
	s := "prod"
	if devMode {
		s = "dev"
	}
	w.Write([]byte(s))
}

func Or(w io.Writer, devMode bool) {
	w.Write([]byte(get(devMode, "dev", "prod")))
}

func get(devMode bool, a, b string) string {
	if devMode {
		return a
	}
	return b
}
