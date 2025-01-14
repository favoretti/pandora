using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Orbital.v2022_11_01.ContactProfile;


internal class ContactProfileThirdPartyConfigurationModel
{
    [JsonPropertyName("missionConfiguration")]
    [Required]
    public string MissionConfiguration { get; set; }

    [JsonPropertyName("providerName")]
    [Required]
    public string ProviderName { get; set; }
}
