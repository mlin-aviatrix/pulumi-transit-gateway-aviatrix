package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/mlin-aviatrix/pulumi-transit-gateway-aviatrix/sdk/go/aviatrix"
)

func main() {
	pulumi.Providers()

	pulumi.Run(func(ctx *pulumi.Context) error {
		//aviatrix.TransitModule
		//aviatrix.NewTransitModule()
	})
}