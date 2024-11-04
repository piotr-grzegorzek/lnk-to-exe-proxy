package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: lnk-to-exe-proxy.exe <path_to_lnk_file>")
        os.Exit(1)
    }

    lnkPath := os.Args[1]

    // Start the command in detached mode to avoid any subprocess blocking
    cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", lnkPath)
    cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}

    err := cmd.Start()
    if err != nil {
        fmt.Printf("Failed to open %s: %v\n", lnkPath, err)
        os.Exit(1)
    }

    fmt.Println("Shortcut launched successfully")
    os.Exit(0) // Explicit exit to prevent hanging
}
