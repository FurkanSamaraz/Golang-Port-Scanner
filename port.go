package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var ip string
	var baslaPort int
	var bitisPort int
	var host int
	fmt.Println("Ip adresini giriniz: ")
	fmt.Scan(&ip)
	fmt.Println("Baslangıç Portunu giriniz: ")
	fmt.Scan(&baslaPort)
	fmt.Println("Bitiş Portunu giriniz: ")
	fmt.Scan(&bitisPort)
	for host := baslaPort; host < bitisPort; host++ {
		hostIP := fmt.Sprintf("%s:%d", ip, host)

		_, err := net.DialTimeout("tcp", hostIP, 10*time.Millisecond)

		if err != nil {

		} else {
			fmt.Printf("Açık port %d\n", host)
		}
	}
	fmt.Println("Tarama Tamamlandı.")
	f, _ := os.OpenFile("mC:/Users/samar/OneDrive/Masaüstü/port.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()
	_, err := f.WriteString(string(host))

	if err != nil {
		fmt.Println(err)
	}

}
