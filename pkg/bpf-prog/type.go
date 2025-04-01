package bpfprog

import "github.com/cilium/ebpf"

type ProgramInfo struct {
	Type        string
	ProgramInfo *ebpf.ProgramInfo
}
