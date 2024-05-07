package main

import (
	"github.com/rneiva/aws-go-pulumi-eks/infra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Deploy networking resources like VPCs and subnets
		vpc, err := networking.DeployVpcModule(ctx)
		if err != nil {
			return err
		}

		// // Deploy compute resources such as EKS clusters
		// err = compute.DeployComputeModule(ctx, vpc)
		// if err != nil {
		// 	return err
		// }

		// // Deploy a load balancer
		// err = loadbalancing.DeployLoadBalancerModule(ctx, vpc)
		// if err != nil {
		// 	return err
		// }

		// // Deploy monitoring resources like CloudWatch
		// err = monitoring.DeployMonitoringModule(ctx, vpc)
		// if err != nil {
		// 	return err
		// }

		return nil
	})
}