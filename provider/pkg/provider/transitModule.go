package provider

import (
	"github.com/mlin-aviatrix/pulumi-aviatrix/sdk/go/aviatrix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TransitModuleArgs struct {
	AccountName pulumi.StringInput `pulumi:"accountName"`
	CloudType pulumi.IntInput `pulumi:"cloudType"`
	Region pulumi.StringInput `pulumi:"region"`
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	GatewaySize pulumi.StringInput `pulumi:"gatewaySize"`
}

type TransitModule struct {
	pulumi.ResourceState

	Vpc *aviatrix.Vpc `pulumi:"vpc"`
	TransitGateway *aviatrix.TransitGateway `pulumi:"transitGateway"`
}

func NewTransitGateway(ctx *pulumi.Context,
	name string, args *TransitModuleArgs, opts ...pulumi.ResourceOption) (*TransitModule, error) {
	if args == nil {
		args = &TransitModuleArgs{}
	}

	component := &TransitModule{}
	err := ctx.RegisterComponentResource("aviatrix:index:TransitModule", name, component, opts...)
	if err != nil {
		return nil, err
	}

	vpc, err := aviatrix.NewVpc(ctx, "mlin-vpc", &aviatrix.VpcArgs{
		AccountName: args.AccountName,
		CloudType: args.CloudType,
		Region: args.Region,
		Name: pulumi.String("mlin-vpc"),
		Cidr: pulumi.String("10.1.0.0/16"),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	gateway, err := aviatrix.NewTransitGateway(ctx, name, &aviatrix.TransitGatewayArgs{
		CloudType: vpc.CloudType,
		AccountName: vpc.AccountName,
		GwName: args.GatewayName,
		VpcId: vpc.VpcId,
		VpcReg: vpc.Region.Elem(),
		GwSize: args.GatewaySize,
		Subnet: vpc.PublicSubnets.Index(pulumi.Int(0)).Cidr().Elem(),
		Tags: pulumi.StringMap{
			"k1": pulumi.String("v1"),
		},
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	component.Vpc = vpc
	component.TransitGateway = gateway

	//if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
	//	"vpc": vpc,
	//	"transitGateway": gateway,
	//}); err != nil {
	//	return nil, err
	//}
	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"gatewayName": args.GatewayName,
	}); err != nil {
		return nil, err
	}

	return component, nil
}