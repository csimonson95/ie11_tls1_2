package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

type TLS_Settings struct {
	key_path     string
	machine_type string
}

func Console(show bool) {
	var getWin = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	var showWin = syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	hwnd, _, _ := getWin.Call()
	if hwnd == 0 {
		return
	}
	if show {
		var SW_RESTORE uintptr = 9
		showWin.Call(hwnd, SW_RESTORE)
	} else {
		var SW_HIDE uintptr = 0
		showWin.Call(hwnd, SW_HIDE)
	}
}

func ConsoleExit() {
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key
}

func (t TLS_Settings) Create_Key_Paths() {
	required_paths := [2]string{`\Client`, `\Server`}
	for _, k := range required_paths {
		d_path := t.key_path + k
		k, key_existed, err := registry.CreateKey(registry.LOCAL_MACHINE,
			d_path,
			registry.ALL_ACCESS)
		if err != nil {
			log.Fatal("Failed to CreateKey key =", d_path, " Error = ", err)
		}
		key_modified := !key_existed
		fmt.Println("Registry key:", d_path, "was modified:", key_modified)
		defer k.Close()
	}
}

func (t *TLS_Settings) Update_Key_Paths(requested_key string, requested_value uint32) {
	required_paths := []string{}
	if t.machine_type == "Server" {
		required_paths = append(required_paths, `\Client`, `\Server`)
	} else {
		required_paths = append(required_paths, `\Client`)
	}

	for _, p := range required_paths {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE,
			t.key_path+p,
			registry.ALL_ACCESS)
		if err != nil {
			log.Fatal("Failed to OpenKey =", t.key_path+p, "Error = ", err)
		}
		fmt.Println("Modifying the requested path:", t.key_path+p,
			"\nSetting requested_key:", requested_key, "to requested_value:", requested_value)
		defer k.Close()

		err = k.SetDWordValue(requested_key, requested_value)
		if err != nil {
			log.Fatal("Failed to set DWordValue requested_value =", requested_value, "Error = ", err)
		}
	}
}

func main() {
	Console(true)
	defer Console(true)
	TLS := &TLS_Settings{`SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.2`, "Server"}
	TLS.Create_Key_Paths()
	TLS.Update_Key_Paths("DisabledByDefault", 0)
	TLS.Update_Key_Paths("Enabled", 1)
	fmt.Println("You can now close your Console window.")
	time.Sleep(24 * time.Hour)
}
