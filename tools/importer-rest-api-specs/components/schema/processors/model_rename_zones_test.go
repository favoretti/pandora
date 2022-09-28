package processors

import (
	"testing"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func TestProcessModel_RenameZones_Valid(t *testing.T) {
	testData := []struct {
		modelNameInput   string
		mappingsInput    resourcemanager.MappingDefinition
		modelsInput      map[string]resourcemanager.TerraformSchemaModelDefinition
		expectedMappings resourcemanager.MappingDefinition
		expectedModels   map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Disco",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"AvailabilityZones": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeList,
								NestedObject: &resourcemanager.TerraformSchemaFieldObjectDefinition{
									Type: resourcemanager.TerraformSchemaFieldTypeString,
								},
							},
						},
					},
				},
			},
			mappingsInput: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
			},
			expectedModels: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Zones": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeList,
								NestedObject: &resourcemanager.TerraformSchemaFieldObjectDefinition{
									Type: resourcemanager.TerraformSchemaFieldTypeString,
								},
							},
						},
					},
				},
			},
			expectedMappings: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actualModels, actualMappings, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput, v.mappingsInput)
		if err != nil {
			if v.expectedModels == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		modelDefinitionsMatch(t, actualModels, v.expectedModels)
		mappingDefinitionsMatch(t, actualMappings, v.expectedMappings)
	}
}

func TestProcessModel_RenameZone_Valid(t *testing.T) {
	testData := []struct {
		modelNameInput   string
		mappingsInput    resourcemanager.MappingDefinition
		modelsInput      map[string]resourcemanager.TerraformSchemaModelDefinition
		expectedMappings resourcemanager.MappingDefinition
		expectedModels   map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Disco",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"AvailabilityZone": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			mappingsInput: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
			},
			expectedModels: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Disco": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Zone": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			expectedMappings: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "Zone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actualModels, actualMappings, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput, v.mappingsInput)
		if err != nil {
			if v.expectedModels == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		modelDefinitionsMatch(t, actualModels, v.expectedModels)
		mappingDefinitionsMatch(t, actualMappings, v.expectedMappings)
	}
}

func TestProcessModel_RenameZones_Invalid(t *testing.T) {
	testData := []struct {
		modelNameInput   string
		mappingsInput    resourcemanager.MappingDefinition
		modelsInput      map[string]resourcemanager.TerraformSchemaModelDefinition
		expectedMappings resourcemanager.MappingDefinition
		expectedModels   map[string]resourcemanager.TerraformSchemaModelDefinition
	}{
		{
			modelNameInput: "Leopard",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Leopard": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			mappingsInput: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
			},
			// unchanged
			expectedModels: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Leopard": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
					},
				},
			},
			expectedMappings: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZones",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZones",
						},
					},
				},
			},
		},
		{
			modelNameInput: "Meerkat",
			modelsInput: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Meerkat": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
			},
			mappingsInput: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
			},
			// unchanged since the nested model doesn't match
			expectedModels: map[string]resourcemanager.TerraformSchemaModelDefinition{
				"Meerkat": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type:          resourcemanager.TerraformSchemaFieldTypeReference,
								ReferenceName: stringPointer("SubResource"),
							},
						},
					},
				},
				"SubResource": {
					Fields: map[string]resourcemanager.TerraformSchemaFieldDefinition{
						"Id": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeString,
							},
						},
						"Weight": {
							ObjectDefinition: resourcemanager.TerraformSchemaFieldObjectDefinition{
								Type: resourcemanager.TerraformSchemaFieldTypeInteger,
							},
						},
					},
				},
			},
			expectedMappings: resourcemanager.MappingDefinition{
				Create: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Update: &[]resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
				Read: []resourcemanager.FieldMappingDefinition{
					{
						Type: resourcemanager.DirectAssignmentMappingDefinitionType,
						DirectAssignment: &resourcemanager.FieldMappingDirectAssignmentDefinition{
							SchemaModelName: "Disco",
							SchemaFieldPath: "AvailabilityZone",
							SdkModelName:    "SomeModel",
							SdkFieldPath:    "AvailabilityZone",
						},
					},
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %s", v.modelNameInput)

		actualModels, actualMappings, err := modelRenameZones{}.ProcessModel(v.modelNameInput, v.modelsInput[v.modelNameInput], v.modelsInput, v.mappingsInput)
		if err != nil {
			if v.expectedModels == nil {
				continue
			}

			t.Fatalf("error: %+v", err)
		}

		modelDefinitionsMatch(t, actualModels, v.expectedModels)
		mappingDefinitionsMatch(t, actualMappings, v.expectedMappings)
	}
}
