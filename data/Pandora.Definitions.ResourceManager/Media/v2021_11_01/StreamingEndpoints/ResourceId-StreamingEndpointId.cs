using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Media.v2021_11_01.StreamingEndpoints;

internal class StreamingEndpointId : ResourceID
{
    public string? CommonAlias => null;

    public string ID => "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{mediaServiceName}/streamingEndpoints/{streamingEndpointName}";

    public List<ResourceIDSegment> Segments => new List<ResourceIDSegment>
    {
        ResourceIDSegment.Static("staticSubscriptions", "subscriptions"),
        ResourceIDSegment.SubscriptionId("subscriptionId"),
        ResourceIDSegment.Static("staticResourceGroups", "resourceGroups"),
        ResourceIDSegment.ResourceGroup("resourceGroupName"),
        ResourceIDSegment.Static("staticProviders", "providers"),
        ResourceIDSegment.ResourceProvider("staticMicrosoftMedia", "Microsoft.Media"),
        ResourceIDSegment.Static("staticMediaServices", "mediaServices"),
        ResourceIDSegment.UserSpecified("mediaServiceName"),
        ResourceIDSegment.Static("staticStreamingEndpoints", "streamingEndpoints"),
        ResourceIDSegment.UserSpecified("streamingEndpointName"),
    };
}
