package main

type input struct {
	Question string
	Answer   string
}

type param struct {
	Tag    string
	Domain string
	Schema string
	Node   struct {
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
}

var inputs = []input{
	{Question: "Docker stack name"},
	{Question: "Domain name"},
	{Question: "Admin username"},
	{Question: "Admin password"},
	{Question: "BasicAuth username"},
	{Question: "BasicAuth password"},
	{Question: "Slack Webhook URL"},
	{Question: "Username for Slack alerts"},
	{Question: "Traefik external port"},
	{Question: "HTTP schema"},
	{Question: "Docker Swarm metric port"},
	{Question: "Docker gwbridge IP"},
}
var p = param{Tag: "development", Node: struct{ ID string }{"{{.Node.ID}}"}}

func main() {
	configexist := fileExists(configPath)

	if configexist {
		loadConfig(configPath)
		setAnswers()
	}
	menuPage()
}
