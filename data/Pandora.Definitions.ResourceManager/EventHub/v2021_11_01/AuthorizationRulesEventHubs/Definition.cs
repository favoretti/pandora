using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.EventHub.v2021_11_01.AuthorizationRulesEventHubs;

internal class Definition : ResourceDefinition
{
    public string Name => "AuthorizationRulesEventHubs";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new EventHubsCreateOrUpdateAuthorizationRuleOperation(),
        new EventHubsListAuthorizationRulesOperation(),
        new EventHubsListKeysOperation(),
        new EventHubsRegenerateKeysOperation(),
    };
}