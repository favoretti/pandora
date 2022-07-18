package generator

import (
	"fmt"
	"strings"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

type createFunctionComponents struct {
	readMethod            resourcemanager.ApiOperation
	readMethodName        string
	createMethod          resourcemanager.ApiOperation
	createMethodName      string
	resourceTypeName      string
	sdkResourceName       string
	newResourceIdFuncName string
	resourceId            resourcemanager.ResourceIdDefinition
}

func createFunctionForResource(input ResourceInput) string {
	if !input.Details.CreateMethod.Generate {
		return ""
	}

	createOperation, ok := input.Operations[input.Details.CreateMethod.MethodName]
	if !ok {
		// TODO: thread through errors
		panic(fmt.Sprintf("couldn't find create operation named %q", input.Details.CreateMethod.MethodName))
	}

	readOperation, ok := input.Operations[input.Details.ReadMethod.MethodName]
	if !ok {
		// TODO: thread through errors
		panic(fmt.Sprintf("couldn't find Read operation for create operation named %q", input.Details.ReadMethod.MethodName))
	}

	resourceId, ok := input.ResourceIds[input.Details.ResourceIdName]
	if !ok {
		// TODO: thread through errors
		panic(fmt.Sprintf("couldn't find Resource ID %q for Create Method", input.Details.ResourceIdName))
	}

	newResourceIdFuncName, err := input.newResourceIdFuncName()
	if err != nil {
		// TODO: thread through errors
		panic(fmt.Sprintf("obtaining New Resource ID Function for Create Method: %+v", err))
	}

	helper := createFunctionComponents{
		readMethod:            readOperation,
		readMethodName:        input.Details.ReadMethod.MethodName,
		createMethod:          createOperation,
		createMethodName:      input.Details.CreateMethod.MethodName,
		resourceTypeName:      input.ResourceTypeName,
		sdkResourceName:       input.SdkResourceName,
		newResourceIdFuncName: *newResourceIdFuncName,
		resourceId:            resourceId,
	}
	components := []string{
		helper.schemaDeserialization(),
		helper.idDefinitionAndMapping(),
		helper.requiresImport(),
		helper.payloadAndMappingsFromSchema(),
		helper.create(),
	}

	return fmt.Sprintf(`
func (r %[1]sResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: %[2]d * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.%[3]s.%[1]sClient

			%[4]s

			metadata.SetID(id.ID())
			return nil
		},
	}
}
`, input.ResourceTypeName, input.Details.ReadMethod.TimeoutInMinutes, input.ServiceName, strings.Join(components, "\n"))
}

func (h createFunctionComponents) schemaDeserialization() string {
	return fmt.Sprintf(`
			var config %[1]sResourceModel
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %%+v", err)
			}
`, h.resourceTypeName)
}

func (h createFunctionComponents) requiresImport() string {
	readMethodArguments := argumentsForApiOperationMethod(h.readMethod, h.sdkResourceName, h.readMethodName)
	return fmt.Sprintf(`
			existing, err := client.%[1]s(%[2]s)
			if err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %%s: %%+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}
`, h.readMethodName, readMethodArguments)
}

func (h createFunctionComponents) payloadAndMappingsFromSchema() string {
	// NOTE: whilst Payload is _technically_ optional in the API endpoint it's not, else it
	// wouldn't be a Create method
	createObjectName, err := h.createMethod.RequestObject.GolangTypeName(&h.sdkResourceName)
	if err != nil {
		// TODO: thread through errors
		panic(fmt.Sprintf("determining Golang Type name for Create Request Object: %+v", err))
	}

	return fmt.Sprintf(`
			payload := %[1]s{}
			// TODO: mapping from the Schema -> Payload
`, *createObjectName)
}

func (h createFunctionComponents) idDefinitionAndMapping() string {
	newIdFuncName := h.newResourceIdFuncName
	segments := make([]string, 0)

	subscriptionIdDefinition := ""
	for _, v := range h.resourceId.Segments {
		switch v.Type {
		case resourcemanager.SubscriptionIdSegment:
			{
				segments = append(segments, "subscriptionId")
				subscriptionIdDefinition = "subscriptionId := metadata.Client.Account.SubscriptionId"
				continue
			}

		default:
			{
				// TODO: once we have the mappings from the Schema -> Resource ID we should switch these out, but for now
				// let's just output the segments example value..
				segments = append(segments, fmt.Sprintf("%q", v.ExampleValue))
			}
		}
	}

	return fmt.Sprintf(`
%[3]s
id := %[1]s(%[2]s)
`, newIdFuncName, strings.Join(segments, ", "), subscriptionIdDefinition)
}

func (h createFunctionComponents) create() string {
	methodName := methodNameToCallForOperation(h.createMethod, h.createMethodName)
	methodArguments := argumentsForApiOperationMethod(h.createMethod, h.sdkResourceName, h.createMethodName)
	return fmt.Sprintf(`
			if err := client.%[1]s(%[2]s); err != nil {
				return fmt.Errorf("creating %%s: %%+v", id, err)
			}
`, methodName, methodArguments)
}
