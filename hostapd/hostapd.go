package hostapd

import (
  "log"
  "os"
  "os/exec"
  "io/ioutil"
  "strconv"
  "strings"
  "fmt"
)

const (
  PID_FILE = "/var/run/hostapd.pid"
  CFG_FILE = "/etc/hostapd.conf"
)

func IsRunning() bool {
  pid := GetPID()
  if pid < 0 { return false }

  _, err := os.FindProcess( pid )
  if err != nil {
    return false
  }
  return true
}

func GetPID() int {
  data,err := ioutil.ReadFile(PID_FILE)
  if err != nil {
    fmt.Println("GetPID: cannot read from", PID_FILE)
    return -1
  }

  pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
  if err != nil {
    fmt.Printf("GetPID: cannot convert %s to an integer", string(data))
    return -1
  }
  return pid
}

func Start() bool {
  fmt.Println("HostAPD: starting")
  cmd := exec.Command("/usr/sbin/hostapd", "-B", "-P", PID_FILE, CFG_FILE)

  err := cmd.Start()
  if err != nil {
      log.Fatal(err)
  }
  log.Printf("HostAPD Start: Waiting for command to finish...")
  err = cmd.Wait()
  log.Printf("HostAPD Start: Command finished with error: %v", err)
  return true
}

func Stop() bool {
  fmt.Println("HostAPD Stop: HostAPD: stopping")
  pid := GetPID()
  fmt.Println("HostAPD Stop: HostAPD's pid=", pid)
  if pid < 0 { return false }

  process,err := os.FindProcess( pid )
  if err != nil {
    return false
  }
  err = process.Signal(os.Interrupt)
  if err != nil {
    return false
  }
  process.Wait()
  return true
}
