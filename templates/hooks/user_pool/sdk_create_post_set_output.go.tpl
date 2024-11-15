	err = rm.SyncTags(ctx, string(*ko.Status.ACKResourceMetadata.ARN), desired.ko.Spec.Tags, ko.Spec.Tags)
	if err != nil {
		return nil, err
	}
