package utils

import (
	"context"
	"fmt"
	"os"
	"path"
)

func UpdatePIDFile(ctx context.Context, pidFile string) {
	var (
		runDir string = path.Dir(pidFile)
		f      *os.File
		pid    int = os.Getpid()
		err    error
	)

	if _, err = os.Stat(runDir); os.IsNotExist(err) {
		if err = os.Mkdir(runDir, 0755); err != nil {
			return
		}
	}

	f, err = os.OpenFile(pidFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	fmt.Fprint(f, pid)
	f.Close()
}
