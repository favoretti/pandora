using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.RecoveryServicesSiteRecovery.v2022_09_10.ReplicationEligibilityResults;

internal class ListOperation : Operations.GetOperation
{
    public override ResourceID? ResourceId() => new VirtualMachineId();

    public override Type? ResponseObject() => typeof(ReplicationEligibilityResultsCollectionModel);

    public override string? UriSuffix() => "/providers/Microsoft.RecoveryServices/replicationEligibilityResults";


}