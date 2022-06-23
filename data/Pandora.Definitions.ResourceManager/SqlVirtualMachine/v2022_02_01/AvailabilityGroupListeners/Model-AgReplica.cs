using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.SqlVirtualMachine.v2022_02_01.AvailabilityGroupListeners;


internal class AgReplicaModel
{
    [JsonPropertyName("commit")]
    public CommitConstant? Commit { get; set; }

    [JsonPropertyName("failover")]
    public FailoverConstant? Failover { get; set; }

    [JsonPropertyName("readableSecondary")]
    public ReadableSecondaryConstant? ReadableSecondary { get; set; }

    [JsonPropertyName("role")]
    public RoleConstant? Role { get; set; }

    [JsonPropertyName("sqlVirtualMachineInstanceId")]
    public string? SqlVirtualMachineInstanceId { get; set; }
}