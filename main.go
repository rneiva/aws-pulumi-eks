package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rneiva/aws-go-pulumi-eks/infra/networking"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return vpc.vpcModule(ctx)
	})
}