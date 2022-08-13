package pkg

import "gopkg.in/alecthomas/kingpin.v2"

var (
	ArgsApp       = kingpin.New("artifact-server", "A simple artifact server used with Act")
	ArgsAuthToken = ArgsApp.Flag("auth-token", "API token to be used for authenthication. (evironment variable AUTH_TOKEN)").Short('t').
			Default("ActIsAwesome").
			OverrideDefaultFromEnvar("AUTH_TOKEN").String()
	ArgsArtifactPath = ArgsApp.Flag("path", "The root directory where to store artifacts. (evironment variable ARTIFACT_PATH)").Short('d').
				Default("/tmp/artifacts").
				OverrideDefaultFromEnvar("ARTIFACT_PATH").String()
	ArgsServerBind = ArgsApp.Flag("bind", "HOST:PORT to bind the server. (evironment variable SERVER_BIND)").Short('b').
			Default("0.0.0.0:1234").
			OverrideDefaultFromEnvar("SERVER_BIND").String()
)
