package main

import (
	"fmt"
	"github.com/xiam/nmapr"
)

func main() {

	report, err := nmapr.Open("input.xml")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Report for: %s (%s)\n\n", report.Args, report.StartStr)

	for _, host := range report.Host {
		fmt.Printf("%s (%s)\n", host.Address.Addr, host.Hostnames[0].Name)

		oports := []nmapr.Port{}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				oports = append(oports, port)
			}
		}

		if len(oports) > 0 {
			fmt.Printf("Open ports:\n")
			for _, port := range oports {
				fmt.Printf("- %d\t%s\t%s\t%s\n", port.PortID, port.Service.Name, port.Service.Product, port.Service.Version)
			}
		} else {
			fmt.Printf("No open ports.\n")
		}

		if len(host.OS) > 0 {
			fmt.Printf("Matching OSes:\n")
			for _, osm := range host.OS {
				fmt.Printf("- %s / %s / %s (%d%%)\n", osm.Class.Family, osm.Class.Vendor, osm.Class.Type, osm.Class.Accuracy)
			}
		}

		fmt.Printf("\n")
	}

}
