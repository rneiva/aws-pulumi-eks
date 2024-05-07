package compute

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rneiva/aws-go-pulumi-eks/infra/networking"
)

func EksModule(ctx *pulumi.Context) error {
	cfg := config.New(ctx, "")
	clusterName := cfg.Require("clusterName")
	vpcId := networking.GetVpcId()

	cluster, err := eks.NewCluster(ctx, clusterName, &eks.ClusterArgs{
		VpcId:           vpcId,
		DesiredCapacity: pulumi.Int(2),
		MinSize:         pulumi.Int(1),
		MaxSize:         pulumi.Int(2),
		EnabledClusterLogTypes: pulumi.StringArray{
			pulumi.String("api"),
			pulumi.String("audit"),
			pulumi.String("authenticator"),
		},
	})
	if err != nil {
		log.Printf("Error creating EKS cluster: %s", err.Error())
		return err
	}

	ctx.Export("kubeConfig", cluster.Kubeconfig)

	return nil
}
