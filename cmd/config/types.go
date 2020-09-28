package config

type input struct {
	Question string
	Answer   string
}

type params struct {
	Domain string
	Schema string
	Cgroup struct {
		Path    string
		Enabled string
	}
	Node struct {
		ID string
	}
	AdminUser struct {
		Name     string
		Password string
	}
	Slack struct {
		Webhook   string
		AlertUser string
		Channel   string
	}
	Traefik struct {
		Port                  string
		BAPassword            string
		BAUser                string
		PrometheusSubDomain   string
		GrafanaSubDomain      string
		AlertmanagerSubDomain string
	}
	Docker struct {
		Tag        string
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
