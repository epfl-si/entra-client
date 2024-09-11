package main

import (
	"entra-client/cmd"
	_ "entra-client/cmd/application"
	_ "entra-client/cmd/certificate"
	_ "entra-client/cmd/claim"
	_ "entra-client/cmd/group"
	_ "entra-client/cmd/manifest"
	_ "entra-client/cmd/serviceprincipal"
	_ "entra-client/cmd/token"
	_ "entra-client/cmd/user"
)

func main() {
	cmd.Execute()
}
