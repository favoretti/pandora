using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.RecoveryServices.v2022_09_10.Vaults;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum VaultPrivateEndpointStateConstant
{
    [Description("Enabled")]
    Enabled,

    [Description("None")]
    None,
}
