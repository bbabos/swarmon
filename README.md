# SwarMon
An out-of-the-box solution to completely monitor your Docker Swarm environments. The monitoring stack uses Prometheus, Grafana, Alertmanager, Traefik and some exporters like cAdvisor.

The integration and management process simplified with a Golang app, which can create the relevant config files based on your needs, deploy or update the services and some basic Docker option from CLI.

**The project is still in a development state, if you have any issues or bugs please contact me.**

## Prerequisites
- SSH access to a manager node
- Docker metrics exposed
- Golang installed on your local machine (to build the binary)

## Usage
You have to clone the repo, build the binary and copy the config folder and the built binary to the remote host.
- ```git clone https://github.com/bbabos/swarmon.git && cd swarmon/```
- ```GOOS=linux GOARCH=amd64 go build```
- copy the **configs** folder and the built **binary** from the root directory to the remote host
- start the binary with the command ```./swarmon```

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
├── configs
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