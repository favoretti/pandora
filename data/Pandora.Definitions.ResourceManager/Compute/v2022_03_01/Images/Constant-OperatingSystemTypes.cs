using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Compute.v2022_03_01.Images;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum OperatingSystemTypesConstant
{
    [Description("Linux")]
    Linux,

    [Description("Windows")]
    Windows,
}
