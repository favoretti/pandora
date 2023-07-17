using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.EventGrid.v2022_06_15.TopicTypes;

internal class ListOperation : Pandora.Definitions.Operations.GetOperation
{
    public override Type? ResponseObject() => typeof(TopicTypesListResultModel);

    public override string? UriSuffix() => "/providers/Microsoft.EventGrid/topicTypes";


}
