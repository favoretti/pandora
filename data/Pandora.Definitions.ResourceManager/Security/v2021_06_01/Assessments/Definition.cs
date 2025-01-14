using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Security.v2021_06_01.Assessments;

internal class Definition : ResourceDefinition
{
    public string Name => "Assessments";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateOrUpdateOperation(),
        new DeleteOperation(),
        new GetOperation(),
        new ListOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(AssessmentStatusCodeConstant),
        typeof(AssessmentTypeConstant),
        typeof(CategoriesConstant),
        typeof(ExpandEnumConstant),
        typeof(ImplementationEffortConstant),
        typeof(SeverityConstant),
        typeof(SourceConstant),
        typeof(ThreatsConstant),
        typeof(UserImpactConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(AssessmentLinksModel),
        typeof(AssessmentStatusModel),
        typeof(AssessmentStatusResponseModel),
        typeof(ResourceDetailsModel),
        typeof(SecurityAssessmentModel),
        typeof(SecurityAssessmentMetadataPartnerDataModel),
        typeof(SecurityAssessmentMetadataPropertiesModel),
        typeof(SecurityAssessmentPartnerDataModel),
        typeof(SecurityAssessmentPropertiesModel),
        typeof(SecurityAssessmentPropertiesResponseModel),
        typeof(SecurityAssessmentResponseModel),
    };
}
