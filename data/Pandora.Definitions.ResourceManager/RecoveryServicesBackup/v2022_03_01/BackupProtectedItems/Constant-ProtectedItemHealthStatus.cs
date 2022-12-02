using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.RecoveryServicesBackup.v2022_03_01.BackupProtectedItems;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ProtectedItemHealthStatusConstant
{
    [Description("Healthy")]
    Healthy,

    [Description("IRPending")]
    IRPending,

    [Description("Invalid")]
    Invalid,

    [Description("NotReachable")]
    NotReachable,

    [Description("Unhealthy")]
    Unhealthy,
}
