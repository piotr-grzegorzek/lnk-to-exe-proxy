package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: lnk-to-exe-proxy.exe <path_to_lnk_file>")
        os.Exit(1)
    }

    lnkPath := os.Args[1]

    // Use rundll32 to open the .lnk file
    cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", lnkPath)
    err := cmd.Start()
    if err != nil {
        fmt.Printf("Failed to open %s: %v\n", lnkPath, err)
        os.Exit(1)
    }

    fmt.Println("Shortcut launched successfully")
}
