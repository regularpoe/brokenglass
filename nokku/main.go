package main

import (
    "fmt"
    "log"
    "net"
    "os/exec"
)

func main() {
    listener, err := net.Listen("tcp4", ":4404")
    if err != nil {
        log.Fatalf("Error creating TCP listener: %s", err)
    }
    defer listener.Close()

    fmt.Println("Listening on port 4404...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %s", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 100)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Printf("Error reading data: %s", err)
        return
    }

    if checkPacket(buffer) {
        fmt.Println("Valid packet received!")
        executeCommand()
    } else {
        fmt.Println("Invalid packet")
    }
}

func checkPacket(data []byte) bool {
    for i := 0; i < len(data)-1; i++ {
        if data[i] != 0 {
            return false
        }
    }
    return data[len(data)-1] == 1
}

func executeCommand() {
    cmd := "ls -lah"
    fmt.Printf("Executing command: %s\n", cmd)

    out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

    fmt.Printf("Output:\n%s\n", out)

    if err != nil {
        log.Printf("Error executing command: %s", err)
    }
}

