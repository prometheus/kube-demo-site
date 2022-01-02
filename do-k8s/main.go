// Copyright 2022 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a DigitalOcean resource (Domain)
		domain, err := digitalocean.NewDomain(ctx, "kube-prometheus-io", &digitalocean.DomainArgs{
			Name: pulumi.String("kube.prometheus.io"),
		})
		if err != nil {
			return err
		}

		// Export the name of the domain
		ctx.Export("domainName", domain.Name)
		return nil
	})
}
