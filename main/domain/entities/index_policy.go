package entities

type indexPolicyActionRollover struct {
	MinSize     string `json:"min_size,omitempty"`
	MinIndexAge string `json:"min_index_age,omitempty"`
}

type indexPolicyAction struct {
	Retry struct {
		Count   int    `json:"count"`
		Backoff string `json:"backoff"`
		Delay   string `json:"delay"`
	} `json:"retry"`
	Rollover *indexPolicyActionRollover `json:"rollover,omitempty"`
	ReadOnly *struct {
	} `json:"read_only,omitempty"`
	Delete *struct {
	} `json:"delete,omitempty"`
}

type indexPolicyTransitionConditions struct {
	MinIndexAge string `json:"min_index_age"`
}

type indexPolicyTransition struct {
	StateName  string                           `json:"state_name,omitempty"`
	Conditions *indexPolicyTransitionConditions `json:"conditions,omitempty"`
}

type indexPolicyState struct {
	Name        string                  `json:"name"`
	Actions     []indexPolicyAction     `json:"actions"`
	Transitions []indexPolicyTransition `json:"transitions"`
}

type indexPolicyIsmTemplate struct {
	Priority      int      `json:"priority"`
	IndexPatterns []string `json:"index_patterns"`
}

type IndexPolicy struct {
	// fields present in GET
	Id          string `json:"_id,omitempty"`
	Version     int    `json:"_version,omitempty"`
	SeqNo       int    `json:"_seq_no,omitempty"`
	PrimaryTerm int    `json:"_primary_term,omitempty"`

	// fields present in GET and PUT
	Policy struct {
		// fields present in GET
		PolicyId        string `json:"policy_id,omitempty"`
		LastUpdatedTime int64  `json:"last_updated_time,omitempty"`
		SchemaVersion   int    `json:"schema_version,omitempty"`

		// fields present in GET and PUT
		IsmTemplate       []indexPolicyIsmTemplate `json:"ism_template"`
		Description       string                   `json:"description"`
		ErrorNotification struct {
			Destination struct {
				CustomWebhook struct {
					HeaderParams struct {
						XApiKey string `json:"x-api-key"`
					} `json:"header_params"`
					Url string `json:"url"`
				} `json:"custom_webhook"`
			} `json:"destination"`
			MessageTemplate struct {
				Source string `json:"source"`
			} `json:"message_template"`
		} `json:"error_notification"`
		DefaultState string             `json:"default_state"`
		States       []indexPolicyState `json:"states"`
	} `json:"policy"`
}

func NewIndexPolicy() *IndexPolicy {
	policy := IndexPolicy{}

	policy.Policy.IsmTemplate = []indexPolicyIsmTemplate{
		{
			Priority:      100,
			IndexPatterns: []string{"exp_*", "fin_*"},
		},
	}
	policy.Policy.Description = "Policy to be used when the rollover is made with these steps: hot, open, close and delete"

	policy.Policy.ErrorNotification.Destination.CustomWebhook.HeaderParams.XApiKey = ""
	policy.Policy.ErrorNotification.Destination.CustomWebhook.Url = ""
	policy.Policy.ErrorNotification.MessageTemplate.Source = "The index {{ctx.index}} failed during policy execution."

	policy.Policy.DefaultState = "hot"

	policy.Policy.States = []indexPolicyState{
		generateState(
			"hot", "open",
			"", "6h",
			false, false,
			""),
		generateState(
			"open", "close",
			"1gb", "1d",
			false, false,
			""),
		generateState(
			"close", "delete",
			"", "",
			true, false,
			"30d"),
		generateState(
			"delete", "",
			"", "",
			false, true,
			""),
	}

	return &policy
}

func generateState(
	current, next,
	actionsMinSize, actionsMinAge string,
	actionsReadOnly, actionsDelete bool,
	transitionsMinAge string,
) indexPolicyState {
	action := indexPolicyAction{}
	if actionsMinSize != "" || actionsMinAge != "" {
		if action.Rollover == nil {
			action.Rollover = &indexPolicyActionRollover{}
		}
	}
	if actionsMinSize != "" {
		action.Rollover.MinSize = actionsMinSize
	}
	if actionsMinAge != "" {
		action.Rollover.MinIndexAge = actionsMinAge
	}
	if actionsReadOnly {
		action.ReadOnly = &struct{}{}
	}
	if actionsDelete {
		action.Delete = &struct{}{}
	}

	action.Retry.Count = 5
	action.Retry.Backoff = "exponential"
	action.Retry.Delay = "30m"

	var transition *indexPolicyTransition
	if next != "" {
		transition = &indexPolicyTransition{}
		transition.StateName = next
		if transitionsMinAge != "" {
			transition.Conditions = &indexPolicyTransitionConditions{}
			transition.Conditions.MinIndexAge = transitionsMinAge
		}
	}

	state := indexPolicyState{}
	state.Name = current
	state.Actions = []indexPolicyAction{action}
	if transition != nil {
		state.Transitions = []indexPolicyTransition{*transition}
	} else {
		state.Transitions = []indexPolicyTransition{}
	}

	return state
}
