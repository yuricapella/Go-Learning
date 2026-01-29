package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/14_interface/interfaces"
)

func main() {
	fmt.Println("=== INTERFACES EM GO ===\n")

	interfaces.DemonstrarInterfacesBasicas()
	interfaces.DemonstrarInterfacesVazias()
	interfaces.DemonstrarTypeAssertions()
	interfaces.DemonstrarInterfacesMultiplosMetodos()
	interfaces.DemonstrarPolimorfismo()
}
