# SwarMon

## Prerequisites
TODO

## Build script for Linux systems
```
env GOOS=linux GOARCH=amd64 go build
```

## Usage
TODO

## Project tree
```
├── main.go
├── go.mod
├── go.sum
├── README.md
├── cmd
│   ├── config
│   │   ├── config.go
│   │   └── types.go
│   ├── docker
│   │   ├── containers.go
│   │   ├── nodes.go
│   │   └── services.go
│   ├── page
│   │   ├── containerPage.go
│   │   ├── menuPages.go
│   │   ├── nodePage.go
│   │   ├── page.go
│   │   ├── servicePage.go
│   │   └── stackPage.go
│   └── utils
│       └── utils.go
├── internal
│   └── docker
│       ├── alertmanager
│       │   ├── Dockerfile
│       │   └── conf
│       │       ├── alertmanager.yml
│       │       └── entrypoint.sh
│       ├── docker-compose.yml
│       ├── dockerd-exporter
│       │   ├── Dockerfile
│       │   └── conf
│       │       └── entrypoint.sh
│       ├── grafana
│       │   ├── Dockerfile
│       │   ├── dashboards
│       │   │   ├── nodes-dash.json
│       │   │   ├── prometheus-dash.json
│       │   │   └── services-dash.json
│       │   ├── dashboards.yml
│       │   └── datasources
│       │       └── prometheus.yaml
│       ├── node-exporter
│       │   ├── Dockerfile
│       │   └── conf
│       │       └── entrypoint.sh
│       ├── parsed.yml
│       └── prometheus
│           ├── Dockerfile
│           ├── conf
│           │   └── prometheus.yml
│           └── rules
│               ├── swarm_node.rules.yml
│               └── swarm_task.rules.yml
├── scripts
│   ├── build.sh
│   └── deploy.sh
```