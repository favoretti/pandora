using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Security.v2017_08_01_preview.IoTSecuritySolutionsAnalytics;

internal class Definition : ResourceDefinition
{
    public string Name => "IoTSecuritySolutionsAnalytics";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new AggregatedAlertDismissOperation(),
        new AggregatedAlertGetOperation(),
        new AggregatedAlertsListOperation(),
        new GetAllOperation(),
        new GetDefaultOperation(),
        new RecommendationGetOperation(),
        new RecommendationsListOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(ReportedSeverityConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(IoTSecurityAggregatedAlertModel),
        typeof(IoTSecurityAggregatedAlertPropertiesModel),
        typeof(IoTSecurityAggregatedRecommendationModel),
        typeof(IoTSecurityAggregatedRecommendationPropertiesModel),
        typeof(IoTSecurityAlertedDeviceModel),
        typeof(IoTSecurityAlertedDevicesListModel),
        typeof(IoTSecurityDeviceAlertModel),
        typeof(IoTSecurityDeviceAlertsListModel),
        typeof(IoTSecurityDeviceRecommendationModel),
        typeof(IoTSecurityDeviceRecommendationsListModel),
        typeof(IoTSecuritySolutionAnalyticsModelModel),
        typeof(IoTSecuritySolutionAnalyticsModelListModel),
        typeof(IoTSecuritySolutionAnalyticsModelPropertiesModel),
        typeof(IoTSecuritySolutionAnalyticsModelPropertiesDevicesMetricsInlinedModel),
        typeof(IoTSeverityMetricsModel),
    };
}
