using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.ContainerRegistry.v2023_07_01.CredentialSets;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum CredentialHealthStatusConstant
{
    [Description("Healthy")]
    Healthy,

    [Description("Unhealthy")]
    Unhealthy,
}
