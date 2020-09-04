package config

type input struct {
	Question string
	Answer   string
}

type params struct {
	Tag           string
	Domain        string
	Schema        string
	CgroupPath    string
	CgroupDisable string
	Node          struct {
		ID string
	}
	AdminUser struct {
		Name     string
		Password string
	}
	Slack struct {
		Webhook   string
		AlertUser string
	}
	Traefik struct {
		Port       string
		BAPassword string
		BAUser     string
	}
	Docker struct {
		StackName  string
		MetricPort string
		GwBridgeIP string
	}
	HostNamePath string // dev only
}

type paths struct {
	StackConfig string
	RawStack    string
	ParsedStack string
}
