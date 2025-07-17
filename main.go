package main

import (
	"github.com/epfl-si/entra-client/cmd"
	_ "github.com/epfl-si/entra-client/cmd/application"
	_ "github.com/epfl-si/entra-client/cmd/certificate"
	_ "github.com/epfl-si/entra-client/cmd/claimsmappingpolicy"
	_ "github.com/epfl-si/entra-client/cmd/group"
	_ "github.com/epfl-si/entra-client/cmd/manifest"
	_ "github.com/epfl-si/entra-client/cmd/secret"
	_ "github.com/epfl-si/entra-client/cmd/serviceprincipal"
	_ "github.com/epfl-si/entra-client/cmd/token"
	_ "github.com/epfl-si/entra-client/cmd/user"
)

func main() {
	cmd.Execute()
}
