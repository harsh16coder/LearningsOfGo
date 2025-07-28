package acceptancetest

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const (
	baseBinName = "temp-testbinary"
)

func LaunchTestProgram(port string) (cleanup func(), sendInterrupt func() error, err error) {
	binname, err := buildBinary()
	if err != nil {
		return nil, nil, err
	}
	sendInterrupt, Kill, err := runServer(binname, port)
	cleanup = func() {
		if Kill != nil {
			Kill()
		}
		os.Remove(binname)
	}
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return cleanup, sendInterrupt, nil
}

func buildBinary() (string, error) {
	binName := randomString(10) + "-" + baseBinName
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		return "", fmt.Errorf("cannot build tool %s: %s", binName, err)
	}
	return binName, nil
}

func runServer(binname string, port string) (sendInterrupt func() error, KILL func(), err error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	cmdPath := filepath.Join(dir, binname)
	cmd := exec.Command(cmdPath)
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("Cannot run the temp converter: %s", err)
	}
	KILL = func() {
		cmd.Process.Kill()
	}

	sendInterrupt = func() error {
		return cmd.Process.Signal(syscall.SIGTERM)
	}

	err = waitForServerListening(port)
	return
}

func waitForServerListening(port string) error {
	for i := 0; i < 30; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", port))
		if conn != nil {
			conn.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("nothing seems to be listening on localhost:%s", port)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
