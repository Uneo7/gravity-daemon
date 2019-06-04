package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Server struct {
	Uid  string
	Sid  string
	Path string
	Pid  int
}

func (server Server) Load(uid string, sid string) Server {

	path := filepath.Join(config.Daemon.Root, uid, sid)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Server{}
	}

	return Server{
		Uid:  uid,
		Sid:  sid,
		Path: path,
	}
}

func (server Server) GetPID() Server {

	cmd := exec.Command("/usr/bin/tmux", "list-panes", "-t"+server.Sid, "-F '#{pane_pid}'")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return server
	}

	spid := CleanupString(string(output))
	pid, err := strconv.Atoi(spid)

	if err != nil {
		return server
	}

	server.Pid = pid

	return server
}

func (server Server) Start() Server {

	runString := "java -jar ./server.jar 2> >(sed 's/^/\\[Error\\] /' >&1) | tee ./output.log"
	cmd := exec.Command("/usr/bin/tmux", "new", "-d", "-s"+server.Sid, runString)
	cmd.Dir = server.Path

	_, err := cmd.CombinedOutput()

	if err != nil {
		return server
	}

	return server.GetPID()
}

func (server Server) Stop() bool {

	cmd := exec.Command("/usr/bin/tmux", "send-keys", "-t"+server.Sid, "stop", "ENTER")
	cmd.Dir = server.Path

	_, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	return true
}

func (server Server) Command(command string) bool {

	cmd := exec.Command("/usr/bin/tmux", "send-keys", "-t"+server.Sid, command, "ENTER")
	cmd.Dir = server.Path

	_, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	return true
}

func (server Server) Kill() bool {

	cmd := exec.Command("/usr/bin/tmux", "kill-session", "-t"+server.Sid)
	cmd.Dir = server.Path

	_, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	return true
}
