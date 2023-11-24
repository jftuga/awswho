/*
awswho.go
-John Taylor
2023-11-23

# Output the results of the sts get-caller-identity API

static compilation:
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const pgmName string = "awswho"
const pgmVersion string = "1.0.0"
const pgmUrl = "https://github.com/jftuga/awswho"

func main() {
	profile := ""
	flag.StringVar(&profile, "p", "", "an aws profile listed in ~/.aws/config (default \"default\")")
	version := false
	flag.BoolVar(&version, "v", false, "output version and then exit")
	region := ""
	flag.StringVar(&region, "r", "us-east-1", "aws region")
	all := false
	flag.BoolVar(&all, "a", false, "output Account as well as Arn and UserId")
	nl := false
	flag.BoolVar(&nl, "n", false, "output a newline character")
	flag.Parse()

	if version == true {
		fmt.Printf("%v v%v\n%v\n", pgmName, pgmVersion, pgmUrl)
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println(0)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(255)
	}
	fmt.Print(aws.ToString(identity.Account))
	if all == true {
		fmt.Print(" ")
		fmt.Print(aws.ToString(identity.Arn))
		fmt.Print(" ")
		fmt.Print(aws.ToString(identity.UserId))
	}
	if nl == true {
		fmt.Print("\n")
	}
}
