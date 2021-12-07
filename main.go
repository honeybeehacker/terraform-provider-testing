  GNU nano 5.4                                     main.go                                               
package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
        "terraform-provider-testing/testing"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: testing.Provider})
}
