using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Sql.v2023_02_01_preview.RecoverableDatabases;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum DatabaseKeyTypeConstant
{
    [Description("AzureKeyVault")]
    AzureKeyVault,
}
