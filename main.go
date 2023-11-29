package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"sivchari.github.io/terraform-provider-exp/provider"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{}); err != nil {
		log.Fatal(err)
	}
}
