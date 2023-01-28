package main

import (
        "os/exec"
)

func main() {
        // Ağ arayüzünün adı
        interfaceName := "wlo1"

        // Yeni MAC adresi
        newMAC := "00:11:22:33:44:55"

        
        _, err := exec.Command("sudo", "ifconfig", interfaceName, "down").Output()
        if err != nil {
                panic(err)
        }
        _, err = exec.Command("sudo", "ifconfig", interfaceName, "hw", "ether", newMAC).Output()
        if err != nil {
                panic(err)
        }
        _, err = exec.Command("sudo", "ifconfig", interfaceName, "up").Output()
        if err != nil {
                panic(err)
        }
}
