package internal

import (
	"github.com/cilium/ebpf"
)

const (
	ProgramListView uint32 = 0
	ProgramMapsView uint32 = 1
)

var (
	CurrentView       uint32 = ProgramListView
	SelectedProgramID ebpf.ProgramID
)
