using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Logic.v2019_05_01.IntegrationServiceEnvironments;


internal class FlowEndpointsModel
{
    [JsonPropertyName("accessEndpointIpAddresses")]
    public List<IPAddressModel>? AccessEndpointIPAddresses { get; set; }

    [JsonPropertyName("outgoingIpAddresses")]
    public List<IPAddressModel>? OutgoingIPAddresses { get; set; }
}