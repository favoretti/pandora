using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.MachineLearningServices.v2022_10_01.Job;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ListViewTypeConstant
{
    [Description("ActiveOnly")]
    ActiveOnly,

    [Description("All")]
    All,

    [Description("ArchivedOnly")]
    ArchivedOnly,
}
