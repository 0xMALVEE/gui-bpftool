package bpfprog

import (
	"fmt"
	"strings"

	"github.com/cilium/ebpf"
)

func GetAllBpfProgList() []ebpf.ProgramID {
	var startId ebpf.ProgramID = 0
	ebpfProgList := []ebpf.ProgramID{}

	for {
		progId, err := ebpf.ProgramGetNextID(startId)
		if err != nil {
			fmt.Println("Error getting next bpf program: ", err.Error())
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

func GetProgListWithInfo(programName string) ([]ProgramInfo, error) {
	progList := GetAllBpfProgList()
	porgListInfo := []ProgramInfo{}
	for _, id := range progList {
		info, err := GetProgInfo(id)
		if err != nil {
			return nil, err
		}
		if !strings.Contains(info.Name, programName) {
			continue
		}
		porgListInfo = append(porgListInfo, ProgramInfo{Type: GetProgramTypeString(info.Type), ProgramInfo: info})
	}
	return porgListInfo, nil
}
