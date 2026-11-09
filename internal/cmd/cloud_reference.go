// SPDX-FileCopyrightText: 2025 OVH SAS <opensource@ovh.net>
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/ovh/ovhcloud-cli/internal/services/cloud"
	"github.com/spf13/cobra"
)

// initCloudReferenceCommand holds the reference data that spans several products;
// product-specific reference data lives under its product command.
func initCloudReferenceCommand(cloudCmd *cobra.Command) {
	referenceCmd := &cobra.Command{
		Use:   "reference",
		Short: "Fetch reference data in the given cloud project",
	}
	referenceCmd.PersistentFlags().StringVar(&cloud.CloudProject, "cloud-project", "", "Cloud project ID")

	productAvailabilityCmd := &cobra.Command{
		Use:   "product-availability",
		Short: "Check the products available in the given cloud project",
	}
	referenceCmd.AddCommand(productAvailabilityCmd)

	productAvailabilityListCmd := withFilterFlag(&cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the available products, one row per product type, model and location",
		Run:     cloud.ListProductAvailability,
		Args:    cobra.NoArgs,
	})
	productAvailabilityListCmd.Flags().StringVar(&cloud.ProductAvailabilityFilters.Type, "type", "", "Product type (FLOATING_IP, GATEWAY, INSTANCE, INSTANCE_SNAPSHOT, LOADBALANCER, OBJECT_STORAGE, PUBLIC_IP, SHARE, SHARE_SNAPSHOT, SWIFT, VOLUME, VOLUME_BACKUP, VOLUME_SNAPSHOT)")
	productAvailabilityListCmd.Flags().StringVarP(&cloud.ProductAvailabilityFilters.Region, "region", "r", "", "Region, letter case ignored (e.g., GRA11, EU-WEST-PAR)")
	productAvailabilityListCmd.Flags().StringVar(&cloud.ProductAvailabilityFilters.Name, "name", "", "Exact model name (e.g., b3-8)")
	productAvailabilityListCmd.Flags().StringVar(&cloud.ProductAvailabilityFilters.NameContains, "name-contains", "", "Part of the model name, letter case ignored")
	productAvailabilityCmd.AddCommand(productAvailabilityListCmd)

	cloudCmd.AddCommand(referenceCmd)
}
