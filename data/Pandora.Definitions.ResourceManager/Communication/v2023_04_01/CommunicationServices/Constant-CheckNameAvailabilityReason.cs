using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Communication.v2023_04_01.CommunicationServices;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum CheckNameAvailabilityReasonConstant
{
    [Description("AlreadyExists")]
    AlreadyExists,

    [Description("Invalid")]
    Invalid,
}
