using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.DataMigration.v2021_06_30.PUT;

internal class ServiceTaskId : ResourceID
{
    public string? CommonAlias => null;

    public string ID => "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{serviceTaskName}";

    public List<ResourceIDSegment> Segments => new List<ResourceIDSegment>
    {
        ResourceIDSegment.Static("staticSubscriptions", "subscriptions"),
        ResourceIDSegment.SubscriptionId("subscriptionId"),
        ResourceIDSegment.Static("staticResourceGroups", "resourceGroups"),
        ResourceIDSegment.UserSpecified("resourceGroupName"),
        ResourceIDSegment.Static("staticProviders", "providers"),
        ResourceIDSegment.ResourceProvider("staticMicrosoftDataMigration", "Microsoft.DataMigration"),
        ResourceIDSegment.Static("staticServices", "services"),
        ResourceIDSegment.UserSpecified("serviceName"),
        ResourceIDSegment.Static("staticServiceTasks", "serviceTasks"),
        ResourceIDSegment.UserSpecified("serviceTaskName"),
    };
}
