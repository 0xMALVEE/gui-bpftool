package bpfprog

import (
	"fmt"

	"github.com/cilium/ebpf"
)

func GetAllBpfProgList() []ebpf.ProgramID {
	var startId ebpf.ProgramID = 0
	ebpfProgList := []ebpf.ProgramID{}

	for {
		progId, err := ebpf.ProgramGetNextID(startId)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		ebpfProgList = append(ebpfProgList, progId)
		startId = progId
	}
	return ebpfProgList
}

func GetProgInfo(id ebpf.ProgramID) (*ebpf.ProgramInfo, error) {
	prog, err := ebpf.NewProgramFromID(id)
	if err != nil {
		return nil, err
	}
	progInfo, err := prog.Info()
	return progInfo, err
}
