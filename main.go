package main

import (
	"flag"

	"github.com/dmacvicar/terraform-provider-libvirt/libvirt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	// Configure logging to prevent interference with Terraform's JSON output
	libvirt.ConfigureLogging()
	
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	
	defer libvirt.CleanupLibvirtConnections()

	opts := &plugin.ServeOpts{
		ProviderFunc: libvirt.Provider,
	}

	if debugMode {
		opts.Debug = true
		opts.ProviderAddr = "registry.terraform.io/dmacvicar/libvirt"
	}

	plugin.Serve(opts)
}