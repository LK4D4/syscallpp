package syscallpp

//go:generate go run ppgen/gen.go ppgen/genlinux.go

type ArgType uint

const (
	ARG_INT ArgType = iota
	ARG_STR
	ARG_PTR
)
