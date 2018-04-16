# nexus-scripting

nexus-scripting provides an go api and a command line interface for the scripting api Sonatype Nexus 3

## Requirements

* [Go](https://golang.org/) >= 1.10
* [Dep](https://golang.github.io/dep/)

## Testing

* build
```bash
make
```
* start nexus
```bash
docker-compose up -d
```
* configure nexus-scripting
```bash
export NEXUS_URL="http://localhost:8081"
export NEXUS_USER="admin"
export NEXUS_PASSWORD="admin123"
```
* execute sample script
```bash
./target/nexus-scripting execute sample.groovy
./target/nexus-scripting execute --payload 'Tricia McMillian' sample.groovy
```

