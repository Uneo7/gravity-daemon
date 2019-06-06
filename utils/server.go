package utils

import (
	Cfg "gravity-daemon/config"
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

	Server Cfg.ServerConfig
	Game   Cfg.GameConfig
}

func (server *Server) Load(uid string, sid string) {

	path := filepath.Join(config.Daemon.Root, uid, sid)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	serverConfig := Cfg.LoadServerConfig(sid)
	gameConfig := Cfg.LoadGameConfig(serverConfig.Game.Id)

	if serverConfig.Sid == "" || gameConfig.Name == "" {
		return
	}

	server.Uid = uid
	server.Sid = sid
	server.Path = path
	server.Server = serverConfig
	server.Game = gameConfig
}

func (server *Server) GetPID() {

	cmd := exec.Command("/usr/bin/tmux", "list-panes", "-t"+server.Sid, "-F '#{pane_pid}'")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return
	}

	spid := CleanupString(string(output))
	pid, err := strconv.Atoi(spid)

	if err != nil {
		return
	}

	server.Pid = pid
}

func (server *Server) Start() bool {

	logger := "|& tee output.log"
	runString := parseStartCommand(*server)

	cmd := exec.Command("/usr/bin/tmux", "new", "-d", "-s"+server.Sid, runString+logger)
	cmd.Dir = server.Path

	_, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	server.GetPID()

	if server.Pid == 0 {
		return false
	}

	return true
}

func (server Server) Stop() bool {

	cmd := exec.Command("/usr/bin/tmux", "send-keys", "-t"+server.Sid, server.Game.Commands.Stop, "ENTER")
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
