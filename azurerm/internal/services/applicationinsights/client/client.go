package client

import (
	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2020-02-02/insights"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	AnalyticsItemsClient     *insights.AnalyticsItemsClient
	APIKeysClient            *insights.APIKeysClient
	ComponentsClient         *insights.ComponentsClient
	WebTestsClient           *insights.WebTestsClient
	BillingClient            *insights.ComponentCurrentBillingFeaturesClient
	SmartDetectionRuleClient *insights.ProactiveDetectionConfigurationsClient
}

func NewClient(o *common.ClientOptions) *Client {
	analyticsItemsClient := insights.NewAnalyticsItemsClient(o.SubscriptionId)
	o.ConfigureClient(&analyticsItemsClient.Client, o.ResourceManagerAuthorizer)

	apiKeysClient := insights.NewAPIKeysClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&apiKeysClient.Client, o.ResourceManagerAuthorizer)

	componentsClient := insights.NewComponentsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&componentsClient.Client, o.ResourceManagerAuthorizer)

	webTestsClient := insights.NewWebTestsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&webTestsClient.Client, o.ResourceManagerAuthorizer)

	billingClient := insights.NewComponentCurrentBillingFeaturesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&billingClient.Client, o.ResourceManagerAuthorizer)

	smartDetectionRuleClient := insights.NewProactiveDetectionConfigurationsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&smartDetectionRuleClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AnalyticsItemsClient:     &analyticsItemsClient,
		APIKeysClient:            &apiKeysClient,
		ComponentsClient:         &componentsClient,
		WebTestsClient:           &webTestsClient,
		BillingClient:            &billingClient,
		SmartDetectionRuleClient: &smartDetectionRuleClient,
	}
}
