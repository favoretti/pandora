using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Consumption.v2021_10_01.Lots;

internal class ListByBillingAccountOperation : Operations.ListOperation
{
    public override string? FieldContainingPaginationDetails() => "nextLink";

    public override ResourceID? ResourceId() => new BillingAccountId();

    public override Type NestedItemType() => typeof(LotSummaryModel);

    public override Type? OptionsObject() => typeof(ListByBillingAccountOperation.ListByBillingAccountOptions);

    public override string? UriSuffix() => "/providers/Microsoft.Consumption/lots";

    internal class ListByBillingAccountOptions
    {
        [QueryStringName("$filter")]
        [Optional]
        public string Filter { get; set; }
    }
}
