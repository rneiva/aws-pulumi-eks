package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rneiva/aws-go-pulumi-eks/infra/networking"
	"github.com/rneiva/aws-go-pulumi-eks/infra/compute"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		if err := networking.VpcModule(ctx); err != nil {
			return err
		}
		if err := compute.EksModule(ctx); err != nil {
			return err
		}
		return nil
	})
}