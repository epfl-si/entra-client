package main

import (
	"epfl-entra/cmd"
	_ "epfl-entra/cmd/certificate"
	_ "epfl-entra/cmd/claim"
	_ "epfl-entra/cmd/group"
	_ "epfl-entra/cmd/manifest"
	_ "epfl-entra/cmd/serviceprincipal"
	_ "epfl-entra/cmd/user"
)

func main() {
	cmd.Execute()
}
