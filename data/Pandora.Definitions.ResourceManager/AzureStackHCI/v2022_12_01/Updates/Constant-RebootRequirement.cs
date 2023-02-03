using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.AzureStackHCI.v2022_12_01.Updates;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum RebootRequirementConstant
{
    [Description("False")]
    False,

    [Description("True")]
    True,

    [Description("Unknown")]
    Unknown,
}
