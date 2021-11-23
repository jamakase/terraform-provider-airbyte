package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/jamakase/terraform-prodiver-airbyte/airbyte"
)

func main() {
	tfsdk.Serve(context.Background(), airbyte.New, tfsdk.ServeOpts{
		Name: "airbyte",
	})
}
