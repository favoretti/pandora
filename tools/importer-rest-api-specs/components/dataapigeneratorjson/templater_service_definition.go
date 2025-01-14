package dataapigeneratorjson

import (
	"sort"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	dataApiModels "github.com/hashicorp/pandora/tools/importer-rest-api-specs/components/dataapigeneratorjson/models"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

func buildServiceDefinition(serviceName string, resourceProvider, terraformPackage *string, apiVersions []models.AzureApiDefinition) (*dataApiModels.ServiceDefinition, error) {
	output := dataApiModels.ServiceDefinition{
		Name:             serviceName,
		ResourceProvider: resourceProvider,
		Generate:         true,
	}

	if terraformPackage != nil {
		terraformResources := make([]string, 0)
		for _, apiVersion := range apiVersions {
			for _, resource := range apiVersion.Resources {
				if resource.Terraform == nil {
					continue
				}

				// TODO: support for Data Sources in the future
				for _, v := range resource.Terraform.Resources {
					terraformResources = append(terraformResources, v.ResourceName)
				}
			}
		}
		sort.Strings(terraformResources)

		output.Terraform = &dataApiModels.TerraformServiceDefinition{
			ServicePackageName: pointer.From(terraformPackage),
			Resources:          terraformResources,
		}
	}

	return &output, nil
}
