package pkg

var CLI struct {
	Log struct {
		Level string `short:"l" enum:"trace,debug,info,warn,error" default:"debug"`
		Type  string `enum:"json,console" default:"json"`
	} `embed:"" prefix:"log-"`

	Server struct {
		Token        string `short:"t" help:"API token to be used for authenthication." default:"ActIsAwesome" env:"TOKEN"`
		ArtifactPath string `short:"d" help:"The root directory where to store artifacts." default:"/tmp/artifacts" env:"ARTIFACT_PATH"`
		Bind         string `short:"b" help:"Host to listen on." default:"0.0.0.0:1234" env:"BIND"`
	} `embed:"" prefix:"server-" envprefix:"SERVER_"`
}
