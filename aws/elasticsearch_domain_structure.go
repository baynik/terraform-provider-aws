package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	elasticsearch "github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

func expandAdvancedSecurityOptions(m []interface{}) *elasticsearch.AdvancedSecurityOptionsInput {
	config := elasticsearch.AdvancedSecurityOptionsInput{}
	group := m[0].(map[string]interface{})

	if v, ok := group["enabled"].(bool); ok {
		config.Enabled = aws.Bool(v)
	}

	if v, ok := group["internal_user_database_enabled"].(bool); ok {
		config.InternalUserDatabaseEnabled = aws.Bool(v)
	}

	if v, ok := group["master_user_options"].([]interface{}); ok {
		if len(v) > 0 {
			muo := elasticsearch.MasterUserOptions{}
			masterUserOptions := v[0].(map[string]interface{})

			if v, ok := masterUserOptions["master_user_arn"].(string); ok {
				if v != "" {
					muo.MasterUserARN = aws.String(v)
				}
			}

			if v, ok := masterUserOptions["master_user_name"].(string); ok {
				if v != "" {
					muo.MasterUserName = aws.String(v)
				}
			}

			if v, ok := masterUserOptions["master_user_password"].(string); ok {
				if v != "" {
					muo.MasterUserPassword = aws.String(v)
				}
			}

			config.SetMasterUserOptions(&muo)
		}
	}

	return &config
}

func flattenAdvancedSecurityOptions(advancedSecurityOptions *elasticsearch.AdvancedSecurityOptions) []map[string]interface{} {
	if advancedSecurityOptions == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{
		"enabled":                        aws.BoolValue(advancedSecurityOptions.Enabled),
		"internal_user_database_enabled": aws.BoolValue(advancedSecurityOptions.InternalUserDatabaseEnabled),
	}

	return []map[string]interface{}{m}
}
