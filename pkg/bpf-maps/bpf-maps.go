package bpfmaps

import (
	bpfprog "gui-bpftool/pkg/bpf-prog"

	"github.com/cilium/ebpf"
)

func GetAllMapsOnProgram(id ebpf.ProgramID) ([]ebpf.MapID, error) {
	info, err := bpfprog.GetProgInfo(id)
	mapIds := []ebpf.MapID{}
	if err != nil {
		return nil, err
	}
	ids, _ := info.MapIDs()
	mapIds = append(mapIds, ids...)
	return mapIds, nil
}
