package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
)

func main() {
	p, _ := process.NewProcess(int32(os.Getpid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	cpu, _ := p.CPUPercent()
	fmt.Printf("parent pid: %d name: %s cmd: %s cpu: %.3f\n", p.Pid, name, cmd, cpu)
}
