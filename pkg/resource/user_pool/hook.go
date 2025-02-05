package user_pool

import (
	"context"

	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// syncTags examines the Tags in the supplied Resource and calls the
// TagResource and UntagResource APIs to ensure that the set of
// associated Tags stays in sync with the Resource.Spec.Tags
func (rm *resourceManager) SyncTags(
	ctx context.Context,
	resourceARN string,
	desiredTags map[string]*string,
	existingTags map[string]*string,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func() { exit(err) }()

	toAdd := map[string]string{}
	toDelete := []string{}

	for k, v := range desiredTags {
		if ev, found := existingTags[k]; !found || *ev != *v {
			toAdd[k] = *v
		}
	}

	for k, _ := range existingTags {
		if _, found := desiredTags[k]; !found {
			deleteKey := k
			toDelete = append(toDelete, deleteKey)
		}
	}

	if len(toAdd) > 0 {
		for k, v := range toAdd {
			rlog.Debug("adding tag to resource", "key", k, "value", v)
		}
		if err = rm.addTags(
			ctx,
			resourceARN,
			toAdd,
		); err != nil {
			return err
		}
	}
	if len(toDelete) > 0 {
		for _, k := range toDelete {
			rlog.Debug("removing tag from resource", "key", k)
		}
		if err = rm.removeTags(
			ctx,
			resourceARN,
			toDelete,
		); err != nil {
			return err
		}
	}

	return nil
}

// addTags adds the supplied Tags to the supplied resource
func (rm *resourceManager) addTags(
	ctx context.Context,
	resourceARN string,
	tags map[string]string,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.addTag")
	defer func() { exit(err) }()

	input := &svcsdk.TagResourceInput{
		ResourceArn: &resourceARN,
		Tags:        tags,
	}

	_, err = rm.sdkapi.TagResource(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "TagResource", err)
	return err
}

// removeTags removes the supplied Tags from the supplied resource
func (rm *resourceManager) removeTags(
	ctx context.Context,
	resourceARN string,
	tagKeys []string, // the set of tag keys to delete
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.removeTag")
	defer func() { exit(err) }()

	input := &svcsdk.UntagResourceInput{
		ResourceArn: &resourceARN,
		TagKeys:     tagKeys,
	}
	_, err = rm.sdkapi.UntagResource(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UntagResource", err)
	return err
}
