using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.HealthcareApis.v2022_06_01.DicomServices;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum PublicNetworkAccessConstant
{
    [Description("Disabled")]
    Disabled,

    [Description("Enabled")]
    Enabled,
}
