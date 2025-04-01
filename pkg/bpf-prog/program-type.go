package bpfprog

import (
	"fmt"

	"github.com/cilium/ebpf"
)

func GetProgramTypeString(typeID ebpf.ProgramType) string {
	// Convert the uint32 to ProgramType for comparison
	progType := ebpf.ProgramType(typeID)

	switch progType {
	case ebpf.UnspecifiedProgram:
		return "UnspecifiedProgram"
	case ebpf.SocketFilter:
		return "SocketFilter"
	case ebpf.Kprobe:
		return "Kprobe"
	case ebpf.SchedCLS:
		return "SchedCLS"
	case ebpf.SchedACT:
		return "SchedACT"
	case ebpf.TracePoint:
		return "TracePoint"
	case ebpf.XDP:
		return "XDP"
	case ebpf.PerfEvent:
		return "PerfEvent"
	case ebpf.CGroupSKB:
		return "CGroupSKB"
	case ebpf.CGroupSock:
		return "CGroupSock"
	case ebpf.LWTIn:
		return "LWTIn"
	case ebpf.LWTOut:
		return "LWTOut"
	case ebpf.LWTXmit:
		return "LWTXmit"
	case ebpf.SockOps:
		return "SockOps"
	case ebpf.SkSKB:
		return "SkSKB"
	case ebpf.CGroupDevice:
		return "CGroupDevice"
	case ebpf.SkMsg:
		return "SkMsg"
	case ebpf.RawTracepoint:
		return "RawTracepoint"
	case ebpf.CGroupSockAddr:
		return "CGroupSockAddr"
	case ebpf.LWTSeg6Local:
		return "LWTSeg6Local"
	case ebpf.LircMode2:
		return "LircMode2"
	case ebpf.SkReuseport:
		return "SkReuseport"
	case ebpf.FlowDissector:
		return "FlowDissector"
	case ebpf.CGroupSysctl:
		return "CGroupSysctl"
	case ebpf.RawTracepointWritable:
		return "RawTracepointWritable"
	case ebpf.CGroupSockopt:
		return "CGroupSockopt"
	case ebpf.Tracing:
		return "Tracing"
	case ebpf.StructOps:
		return "StructOps"
	case ebpf.Extension:
		return "Extension"
	case ebpf.LSM:
		return "LSM"
	case ebpf.SkLookup:
		return "SkLookup"
	case ebpf.Syscall:
		return "Syscall"
	case ebpf.Netfilter:
		return "Netfilter"
	default:
		return fmt.Sprintf("UnknownProgramType(%d)", typeID)
	}
}
