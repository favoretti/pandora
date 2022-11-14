using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Storage.v2021_04_01.FileShares;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ShareAccessTierConstant
{
    [Description("Cool")]
    Cool,

    [Description("Hot")]
    Hot,

    [Description("Premium")]
    Premium,

    [Description("TransactionOptimized")]
    TransactionOptimized,
}