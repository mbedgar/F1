A fork of rafaelreinert/F1 so I could see if I could dockerise it and then run this with influxdb and grafana all in containers. Primary driver is to have it up on the wall in my work Christmas party where we have a sim rig set up on a time trial in F1 2017.

I have included the json for my dashboard in repo as well as a starting point for anyone who want to have a look.

# F1
A Telemetry System For F1 2017 Using Grafana and InfluxDB

For datails see rafaelreinert's Post: [Link](https://medium.com/@rafaelreinert/building-my-own-telemetry-system-for-f1-2017-game-using-golang-influxdb-and-grafana-48dedbd2cdc1)
# Build
``` sh
git clone https://github.com/mbedgar/F1.git
cd F1
go get github.com/influxdata/influxdb1-client/v2
go build
./F1
```
# Docker
```docker run -d --rm --name grafana -p 3000:3000 grafana/grafana
docker run -d --rm --name influxdb -p 8089:8089/udp -p 8086:8086 -e INFLUXDB_DB=f1 -e INFLUXDB_UDP_ENABLED=true -v C:\Docker\Volumes\influx\config:\etc\influxdb -v C:\Docker\Volumes\influx\lib:/var/lib/influxdb  influxdb
```
not sure if this is required
```
invoke-webrequest -Method POST -Uri http://localhost:8089/query -Body "q=CREATE DATABASE f1"
```
```
docker run -d --rm --name F1GO -p 20777:20777/udp csmax/f1go
```
