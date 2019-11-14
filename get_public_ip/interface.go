package main

import "net"
import "fmt"
import "strings"
import(
	"net/http"
	"os"
	"io/ioutil"
)

func Ips() (map[string]string, error) {

    ips :=  make(map[string]string)

    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }

    for _, i := range interfaces {
        byName, err := net.InterfaceByName(i.Name)
        if err != nil {
            return nil, err
		}
		
		//fmt.Println(byName)

		fmt.Println(i.Name, ":")
		fmt.Println(byName.Addrs())
		
		addresses, err := byName.Addrs()
        for _, v := range addresses {
			//fmt.Println(v.Network(), ":",  v.String())
            ips[byName.Name] = v.String()
        }
    }
    return ips, nil
}

func getPublicIp() (string, error) {
	interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, v := range interfaces {
		if v.Name == "以太网" {
			iface, err := net.InterfaceByName(v.Name)
			if err != nil {
				return "", err
			}

			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, iter := range addrs {
				ips := strings.Split(iter.String(), "/")
				ip := ips[0]
				fmt.Println(ip)
				if len(ip) <= len("123.123.123.123") {
					if ip != "192.168.0.77" {
						return ip, nil
					}
				}
			}
		}
	}

	return "", nil
}

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("--------------")

	b, err := ioutil.ReadAll(resp.Body)
	//b := make([]byte, 100)
	//n, err := resp.Body.Read(b)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	

	return string(b[0:])
}

func main() {
	Ips()
	public_ip, _ := getPublicIp() 
	if public_ip != "" {
		fmt.Println("public ip :", public_ip)
	}

	fmt.Println(get_external())
 }