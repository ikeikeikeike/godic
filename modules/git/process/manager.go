package process

import (
	"fmt"
	"os/exec"
	"time"
)

type Process struct {
	Pid         int64
	Description string
	Start       time.Time
	Cmd         *exec.Cmd
}

var (
	curPid    int64 = 1
	Processes []*Process
)

func Add(desc string, cmd *exec.Cmd) int64 {
	pid := curPid
	Processes = append(Processes, &Process{
		Pid:         pid,
		Description: desc,
		Start:       time.Now(),
		Cmd:         cmd,
	})
	curPid++
	return pid
}

func Remove(pid int64) {
	for i, proc := range Processes {
		if proc.Pid == pid {
			Processes = append(Processes[:i], Processes[i+1:]...)
			return
		}
	}
}

func Kill(pid int64) error {
	for i, proc := range Processes {
		if proc.Pid == pid {
			if proc.Cmd.Process != nil && proc.Cmd.ProcessState != nil && !proc.Cmd.ProcessState.Exited() {
				if err := proc.Cmd.Process.Kill(); err != nil {
					return fmt.Errorf("fail to kill process(%d/%s): %v", proc.Pid, proc.Description, err)
				}
			}
			Processes = append(Processes[:i], Processes[i+1:]...)
			return nil
		}
	}
	return nil
}
