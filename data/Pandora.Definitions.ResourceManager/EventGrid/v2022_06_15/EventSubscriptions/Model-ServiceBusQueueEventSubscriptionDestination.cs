using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.EventGrid.v2022_06_15.EventSubscriptions;

[ValueForType("ServiceBusQueue")]
internal class ServiceBusQueueEventSubscriptionDestinationModel : EventSubscriptionDestinationModel
{
    [JsonPropertyName("properties")]
    public ServiceBusQueueEventSubscriptionDestinationPropertiesModel? Properties { get; set; }
}
