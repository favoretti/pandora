using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Compute.v2022_03_01.VirtualMachineScaleSets;


internal class UpgradePolicyModel
{
    [JsonPropertyName("automaticOSUpgradePolicy")]
    public AutomaticOSUpgradePolicyModel? AutomaticOSUpgradePolicy { get; set; }

    [JsonPropertyName("mode")]
    public UpgradeModeConstant? Mode { get; set; }

    [JsonPropertyName("rollingUpgradePolicy")]
    public RollingUpgradePolicyModel? RollingUpgradePolicy { get; set; }
}
