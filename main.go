package main

import "io"

func Plain(w io.Writer, devMode bool) {
	w.Write([]byte("prod"))
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
