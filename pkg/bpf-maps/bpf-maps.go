package bpfmaps

import (
	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/cilium/ebpf"
)

func getAllMapIDsOfProgram(id ebpf.ProgramID) ([]ebpf.MapID, error) {
	info, err := bpfprog.GetProgInfo(id)
	mapIds := []ebpf.MapID{}
	if err != nil {
		return nil, err
	}
	ids, _ := info.MapIDs()
	mapIds = append(mapIds, ids...)
	return mapIds, nil
}

func GetAllMaps(id ebpf.ProgramID) ([]*ebpf.Map, error) {
	mapIds, err := getAllMapIDsOfProgram(id)
	if err != nil {
		return nil, err
	}
	maps := []*ebpf.Map{}
	for _, mapId := range mapIds {
		ebpfMap, err := ebpf.NewMapFromID(mapId)
		if err != nil {
			return nil, err
		}
		maps = append(maps, ebpfMap)
	}
	return maps, nil
}
