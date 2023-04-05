package mongo

import (
	"fmt"
	"os/exec"

	"github.com/slayer/autorestart"
	"github.com/webmakom-com/saiStorage/config"
)

type Server struct {
	Config config.Configuration
}

func NewMongoServer(c config.Configuration) Server {
	return Server{
		Config: c,
	}
}

func (m Server) Start() {
	autorestart.WatchFilename = "/usr/bin/mongod"
	autorestart.StartWatcher()

	startMongoCmd := exec.Command("/usr/bin/mongod")
	err := startMongoCmd.Start()

	if err != nil {
		fmt.Println("Mongo has been failed to start:", err)
		return
	}

	fmt.Println("Mongo has been started. PID:", startMongoCmd.Process.Pid)
}
