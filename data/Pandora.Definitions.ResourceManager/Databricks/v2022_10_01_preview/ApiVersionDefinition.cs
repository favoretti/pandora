using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.Databricks.v2022_10_01_preview;

public partial class Definition : ApiVersionDefinition
{
    public string ApiVersion => "2022-10-01-preview";
    public bool Preview => true;
    public Source Source => Source.ResourceManagerRestApiSpecs;

    public IEnumerable<ResourceDefinition> Resources => new List<ResourceDefinition>
    {
        new AccessConnector.Definition(),
    };
}
