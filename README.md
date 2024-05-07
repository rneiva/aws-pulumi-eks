# aws-go-pulumi-eks

## Prerequisites

1. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
2. [Configure AWS Credentials](https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/)
    * I suggest you use [Leapp](https://docs.leapp.cloud/latest/installation/install-leapp/)
3. [Install Go](https://www.pulumi.com/docs/intro/languages/go/)

## Go Modules

- Installing Go Modules
```
go get github.com/pulumi/pulumi/sdk/v3/go/pulumi
go get github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2
go get github.com/pulumi/pulumi/sdk/v3/go/pulumi
go get github.com/pulumi/pulumi/sdk/v3/go/pulumi/config
```

## Deploy your stack

```bash
$ pulumi preview
```

```bash
$ pulumi up
```

## Clean Up

Once you're finished experimenting, you can destroy your stack and remove it to avoid incurring any additional cost:

```bash
pulumi destroy
pulumi stack rm
```