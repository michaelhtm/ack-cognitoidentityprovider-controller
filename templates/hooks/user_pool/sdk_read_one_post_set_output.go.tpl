	output, err := rm.sdkapi.ListTagsForResourceWithContext(
		ctx, 
		&svcsdk.ListTagsForResourceInput{
			ResourceArn: (*string)(ko.Status.ACKResourceMetadata.ARN),
		},
	)
	rm.metrics.RecordAPICall("READ_MANY", "ListTagsForResource", err)
	if err != nil {
		return nil, err
	}

	ko.Spec.Tags = output.Tags
