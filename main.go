package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting the collector")
	ResolveDB:
	laddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20777")
	if err != nil {
		fmt.Println("Cannot resolve Influxdb. Retrying...")
		time.Sleep(1 * time.Second)
		goto ResolveDB
	}
	ConnectDB:
	con, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("InfluxDB unavailable. Retrying...")
		time.Sleep(1 * time.Second)
		goto ConnectDB
	}
	defer con.Close()
	fmt.Println("Collector started")
	buf := make([]byte, 1289)

	ch := make(chan Point, 1000)
	for i := 0; i < 5; i++ {
		go influxDBSender(ch)
	}

	for {
		_, err := con.Read(buf)
		if err != nil {
			fmt.Println(err)
			//log.Fatal(err)
		}
		tp, err := NewTelemetryPack(buf)
		if err != nil {
			fmt.Println(err)
			//log.Fatal(err)
		}
		p := Point{tp: tp, t: time.Now()}
		ch <- p

	}
}
