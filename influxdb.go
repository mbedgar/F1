package main

import (
//	"log"
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	db       = "f1"
	username = "admin"
	password = "admin"
	addr     = "influxdb:8089"
)

// Point is a struct with data for sending to InfluxDB
type Point struct {
	tp *TelemetryPack
	t  time.Time
}

func influxDBSender(ch chan Point) {
	ConnectUDP:
	c, err := client.NewUDPClient(client.UDPConfig{
		Addr: addr,
	})
	if err != nil {
		fmt.Println(err)
//		log.Fatal(err)
		time.Sleep(2 * time.Second)
		goto ConnectUDP
	}
	defer c.Close()

	for p := range ch {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database: db,
		})
		if err != nil {
			fmt.Println(err)
//			log.Fatal(err)
		}
		pt, err := client.NewPoint("TelemetryPack", nil, p.tp.ToMap(), p.t)
		if err != nil {
			fmt.Println(err)
//			log.Fatal(err)
		}
		bp.AddPoint(pt)
		if err := c.Write(bp); err != nil {
			fmt.Println(err)
//			log.Fatal(err)
		}

	}

}
