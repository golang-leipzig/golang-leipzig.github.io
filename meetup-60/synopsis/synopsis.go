package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	configPath = flag.String("c", "synopsis.conf", "config file with one command per line")
	timeout    = flag.Duration("t", 60*time.Second, "max recording duration")
	startDelay = flag.Duration("d", 2*time.Second, "delay before injecting commands")
	layout     = flag.String("l", "tiled", "tmux layout: tiled, even-horizontal, even-vertical, main-horizontal, main-vertical")
	waitFull   = flag.Bool("w", false, "wait the full timeout instead of exiting when all commands finish")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	// allow tmux attach to work even when run from inside an existing tmux
	os.Unsetenv("TMUX")

	cmds, err := readCommands(*configPath)
	if err != nil {
		log.Fatalf("read %s: %v", *configPath, err)
	}
	if len(cmds) == 0 {
		log.Fatalf("no commands in %s", *configPath)
	}
	log.Printf("synopsis: %d commands, timeout %s", len(cmds), *timeout)

	session := fmt.Sprintf("synopsis-%d", os.Getpid())
	// keep temp file on the same filesystem as the destination so rename works
	tempPath := fmt.Sprintf(".%s.cast", session)

	if err := setupTmux(session, len(cmds), *layout); err != nil {
		log.Fatalf("tmux setup: %v", err)
	}

	rec := exec.Command("asciinema", "rec", "--overwrite", tempPath,
		"-c", "tmux attach -t "+session)
	rec.Stdin = os.Stdin
	rec.Stdout = os.Stdout
	rec.Stderr = os.Stderr
	if err := rec.Start(); err != nil {
		killSession(session)
		log.Fatalf("start asciinema: %v", err)
	}

	go func() {
		time.Sleep(*startDelay)
		for i, c := range cmds {
			// without -w, exec replaces the pane's shell so the pane closes
			// when the command exits; with -w, the shell remains and we run
			// the full timeout
			line := c
			if !*waitFull {
				line = "exec " + c
			}
			if err := sendKeys(session, i, line); err != nil {
				fmt.Fprintf(os.Stderr, "send-keys pane %d: %v\n", i, err)
			}
		}
	}()

	done := make(chan error, 1)
	go func() { done <- rec.Wait() }()

	select {
	case <-time.After(*timeout):
		log.Println("timeout reached, closing session")
		killSession(session)
		<-done
	case <-done:
		killSession(session)
	}

	final := fmt.Sprintf("synopsis-recording-%s.cast", time.Now().Format("20060102-150405"))
	if err := os.Rename(tempPath, final); err != nil {
		log.Fatalf("rename %s: %v", tempPath, err)
	}
	log.Println("saved:", final)
}

func readCommands(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cmds []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		cmds = append(cmds, line)
	}
	return cmds, s.Err()
}

// paneShell is run inside each tmux pane; --norc avoids ~/.bashrc clobbering PS1.
const paneShell = `PS1='$ ' exec bash --norc`

func setupTmux(session string, n int, layout string) error {
	if err := exec.Command("tmux", "new-session", "-d", "-s", session, paneShell).Run(); err != nil {
		return fmt.Errorf("new-session: %w", err)
	}
	for i := 1; i < n; i++ {
		if err := exec.Command("tmux", "split-window", "-t", session, paneShell).Run(); err != nil {
			return fmt.Errorf("split-window: %w", err)
		}
		// rebalance after every split so we don't run out of space
		if err := exec.Command("tmux", "select-layout", "-t", session, layout).Run(); err != nil {
			return fmt.Errorf("select-layout %q: %w", layout, err)
		}
	}
	return nil
}

func sendKeys(session string, pane int, text string) error {
	target := fmt.Sprintf("%s.%d", session, pane)
	return exec.Command("tmux", "send-keys", "-t", target, text, "Enter").Run()
}

func killSession(session string) error {
	return exec.Command("tmux", "kill-session", "-t", session).Run()
}
