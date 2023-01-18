using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Authorization.v2020_10_01.RoleEligibilityScheduleRequests;

internal class ScopedRoleEligibilityScheduleRequestId : ResourceID
{
    public string? CommonAlias => null;

    public string ID => "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}";

    public List<ResourceIDSegment> Segments => new List<ResourceIDSegment>
    {
        ResourceIDSegment.Scope("scope"),
        ResourceIDSegment.Static("staticProviders", "providers"),
        ResourceIDSegment.ResourceProvider("staticMicrosoftAuthorization", "Microsoft.Authorization"),
        ResourceIDSegment.Static("staticRoleEligibilityScheduleRequests", "roleEligibilityScheduleRequests"),
        ResourceIDSegment.UserSpecified("roleEligibilityScheduleRequestName"),
    };
}
