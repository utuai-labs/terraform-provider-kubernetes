package kubernetes

import "github.com/hashicorp/terraform/helper/schema"

func podAffinityTermFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"topology_key": {
			Type:        schema.TypeString,
			Description: "idk",
			Optional:    true,
			Default:     1,
		},
		"label_selector": {
			Type:        schema.TypeList,
			Description: "No clue what this does",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: labelSelectorFields(true),
			},
		},
	}
	return s
}

func preferredDuringSchedulingIgnoredDuringExecutionFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"weight": {
			Type:        schema.TypeInt,
			Description: "Wieght of the anti affinity.",
			Optional:    true,
			Default:     1,
		},
		"pod_affinity_term": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: podAffinityTermFields(),
			},
		},
	}
	return s
}

func podAntiAffinityFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"preferred_during_scheduling_ignored_during_execution": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: preferredDuringSchedulingIgnoredDuringExecutionFields(),
			},
		},
	}
	return s
}

func affinityFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"pod_anti_affinity": {
			Type:        schema.TypeList,
			Description: "Anti affinity specifies pods that shouldn't exist on a machine together.",
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: podAntiAffinityFields(),
			},
		},
	}
	return s
}
