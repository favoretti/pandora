using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Resources.v2023_08_01.DeploymentScripts;


internal class ContainerConfigurationModel
{
    [JsonPropertyName("containerGroupName")]
    public string? ContainerGroupName { get; set; }

    [JsonPropertyName("subnetIds")]
    public List<ContainerGroupSubnetIdModel>? SubnetIds { get; set; }
}
