package migration

import (
	"context"
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/notificationhub/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/pluginsdk"
)

var _ pluginsdk.StateUpgrade = NotificationHubAuthorizationRuleResourceV0ToV1{}

type NotificationHubAuthorizationRuleResourceV0ToV1 struct{}

func (NotificationHubAuthorizationRuleResourceV0ToV1) Schema() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"notification_hub_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"namespace_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"resource_group_name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
		},

		"manage": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"send": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"listen": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"primary_access_key": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"secondary_access_key": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},
	}
}

func (NotificationHubAuthorizationRuleResourceV0ToV1) UpgradeFunc() pluginsdk.StateUpgraderFunc {
	return func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
		oldIdRaw := rawState["id"].(string)
		oldId, err := parse.NotificationHubAuthorizationRuleIDInsensitively(oldIdRaw)
		if err != nil {
			return rawState, fmt.Errorf("parsing ID %q to upgrade: %+v", oldIdRaw, err)
		}

		rawState["id"] = oldId.ID()
		return rawState, nil
	}
}
