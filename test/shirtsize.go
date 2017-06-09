package test

type ShirtSize byte

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
	_
)

const (
	foo = iota
)

type bar string

const baz bar = ""
