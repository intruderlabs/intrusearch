package intrusearch

import (
	"fmt"
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	logger "github.com/sirupsen/logrus"
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
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &policy)
	}

	return policy, mapped
}
