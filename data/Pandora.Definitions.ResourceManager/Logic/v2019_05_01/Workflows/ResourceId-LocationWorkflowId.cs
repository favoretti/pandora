using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Logic.v2019_05_01.Workflows;

internal class LocationWorkflowId : ResourceID
{
    public string? CommonAlias => null;

    public string ID => "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/locations/{location}/workflows/{workflowName}";

    public List<ResourceIDSegment> Segments => new List<ResourceIDSegment>
    {
        ResourceIDSegment.Static("staticSubscriptions", "subscriptions"),
        ResourceIDSegment.SubscriptionId("subscriptionId"),
        ResourceIDSegment.Static("staticResourceGroups", "resourceGroups"),
        ResourceIDSegment.ResourceGroup("resourceGroupName"),
        ResourceIDSegment.Static("staticProviders", "providers"),
        ResourceIDSegment.ResourceProvider("staticMicrosoftLogic", "Microsoft.Logic"),
        ResourceIDSegment.Static("staticLocations", "locations"),
        ResourceIDSegment.UserSpecified("location"),
        ResourceIDSegment.Static("staticWorkflows", "workflows"),
        ResourceIDSegment.UserSpecified("workflowName"),
    };
}
