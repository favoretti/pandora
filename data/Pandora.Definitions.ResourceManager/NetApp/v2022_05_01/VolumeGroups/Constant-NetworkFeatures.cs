using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.NetApp.v2022_05_01.VolumeGroups;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum NetworkFeaturesConstant
{
    [Description("Basic")]
    Basic,

    [Description("Standard")]
    Standard,
}
