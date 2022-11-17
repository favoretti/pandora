using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Batch.v2022_10_01.Pool;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum StorageAccountTypeConstant
{
    [Description("Premium_LRS")]
    PremiumLRS,

    [Description("Standard_LRS")]
    StandardLRS,
}
