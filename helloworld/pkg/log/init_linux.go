//go:build linux
// +build linux

package log

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	if defaultLogger == nil {
		InitLog()
	}

	// 信号监听
	go func() {
		if runtime.GOOS != "windows" {
			for {
				c := make(chan os.Signal)
				signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)
				s := <-c

				switch s {
				case syscall.SIGUSR1: //kill -10 pid
					SetLevelString("debug") //切换日志级别DEBUG
					fmt.Println(`get usr1 signal and change log level to "debug":`, s)
				case syscall.SIGUSR2: //kill -12 pid
					SetLevelString("error") //切换日志级别ERROR
					fmt.Println(`get usr2 signal and change log level to "error":`, s)
				}
			}
		}
	}()
}
