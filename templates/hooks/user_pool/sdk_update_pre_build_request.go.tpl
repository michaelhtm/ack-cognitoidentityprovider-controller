	if delta.DifferentAt("Spec.Tags") {
		if err = rm.SyncTags(
			ctx, 
			string(*latest.ko.Status.ACKResourceMetadata.ARN), 
			desired.ko.Spec.Tags, 
			latest.ko.Spec.Tags,
		); err != nil {
			return nil, err
		}
	}
	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}
