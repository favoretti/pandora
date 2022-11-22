using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.SecurityInsights.v2022_11_01.AutomationRules;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum IncidentLabelTypeConstant
{
    [Description("AutoAssigned")]
    AutoAssigned,

    [Description("User")]
    User,
}
