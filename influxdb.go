package main

import (
	"log"
	"time"
	"fmt"

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
	
	for try := 1;;try++ {
		c, err := client.NewUDPClient(client.UDPConfig{
			Addr: addr,
		})
	
		if err != nil {
			fmt.Println("InfluxDB not found retrying...")
			time.Sleep(5 * time.Second)
		} else if try > 12 {
			log.Fatal(err)
		} else if err == nil{
			fmt.Println("Connected to InfluxDB...")
			break
		}
	}
	defer c.Close()

	for p := range ch {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database: db,
		})
		if err != nil {
			log.Fatal(err)
		}
		pt, err := client.NewPoint("TelemetryPack", nil, p.tp.ToMap(), p.t)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

	}

}
