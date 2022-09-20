using System.Collections.Generic;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CommonSchema;

namespace Pandora.Definitions.ResourceManager.Compute.Terraform;

public class VirtualMachineResourceNetworkInterfaceReferencePropertiesSchema
{

    [HclName("delete_option")]
    [Optional]
    [PossibleValuesFromConstant(typeof(v2021_11_01.VirtualMachines.DeleteOptionsConstant))]
    public string DeleteOption { get; set; }


    [HclName("primary")]
    [Optional]
    public bool Primary { get; set; }

}