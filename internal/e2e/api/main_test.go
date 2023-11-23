package api

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

var pkg = "vanjiii/kook-book-server"

func TestMain(m *testing.M) {
	binpath := gobuild()

	defer os.Remove(binpath)

	curdir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("failed to get current working directory, err: %+v", err))
	}

	cmd := exec.Command(binpath, "serve")
	cmd.Dir = curdir

	// NOTE: Inherit environment of the host/container running the binary,
	// to make sure we're not isolating factors.
	// cmd.Env = os.Environ()

	// NOTE: Don't need stdin.
	cmd.Stdin = nil
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	defer func(cmd *exec.Cmd) {
		_ = cmd.Process.Signal(syscall.SIGTERM)
		_ = cmd.Wait()
	}(cmd)

	os.Exit(m.Run())
}

func gobuild() string {
	tmpfile, err := os.CreateTemp("", "e2e_tests_kookbook")
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	tmpfilename := tmpfile.Name()

	var binaryPath string

	if os.Getenv("GOROOT") == "" {
		binaryPath = "go"
	} else {
		binaryPath = filepath.Join(os.Getenv("GOROOT"), "bin", "go")
	}

	args := []string{
		"build",
		"-v",
		"-o",
		tmpfilename,
		pkg,
	}

	cmd := exec.Command(binaryPath, args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to build executable: %s", err)
	}

	return tmpfilename
}
