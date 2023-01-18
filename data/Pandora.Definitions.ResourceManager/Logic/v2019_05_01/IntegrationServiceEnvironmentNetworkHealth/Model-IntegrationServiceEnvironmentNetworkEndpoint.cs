using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Logic.v2019_05_01.IntegrationServiceEnvironmentNetworkHealth;


internal class IntegrationServiceEnvironmentNetworkEndpointModel
{
    [JsonPropertyName("accessibility")]
    public IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateConstant? Accessibility { get; set; }

    [JsonPropertyName("domainName")]
    public string? DomainName { get; set; }

    [JsonPropertyName("ports")]
    public List<string>? Ports { get; set; }
}
