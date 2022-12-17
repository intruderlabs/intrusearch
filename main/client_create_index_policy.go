package intrusearch

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/entities"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/helpers"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/requests"
)

func (itself Client) CreateIndexPolicy(indexName string) (bool, []errors.GenericError) {
	// <size for rollover close>_<period for rollover delete>_<first_state>_rollover
	policyName := "1gb_30d_hot_rollover"

	logger.Infof("Creating/updating the index policy '%s' for the index '%s'...", policyName, indexName)

	if policy, errors := itself.getIndexPolicy(policyName); len(errors) <= 0 {
		logger.Infof("Policy already exists! Updating...")

		policyName = fmt.Sprintf(
			"%s?if_seq_no=%d&if_primary_term=%d",
			policyName, policy.SeqNo, policy.PrimaryTerm,
		)
	}

	body := entities.NewIndexPolicy()
	wrapper, mapped := requests.DoRequest(itself.client, requests.IsmPutIndexPolicyRequest{ // TODO: this policy is always running, why?
		Name: policyName,
		Body: helpers.NewSerializationHelper().ToReader(*body),
	})

	return wrapper.Success, mapped
}

func (itself Client) getIndexPolicy(name string) (entities.IndexPolicy, []errors.GenericError) {
	wrapper, mapped := requests.DoRequest(itself.client, requests.IsmGetIndexPolicyRequest{
		Name: name,
	})

	policy := entities.IndexPolicy{}
	if wrapper.Success { // if there's no HTTP error at all
		helpers.NewSerializationHelper().FromReader(wrapper.Body, &policy)
	}

	return policy, mapped
}
