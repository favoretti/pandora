using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.StreamAnalytics.v2021_10_01_preview.StreamingJobs;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum EventsOutOfOrderPolicyConstant
{
    [Description("Adjust")]
    Adjust,

    [Description("Drop")]
    Drop,
}
