	output, err := rm.sdkapi.ListTagsForResource(
		ctx,
		&svcsdk.ListTagsForResourceInput{
			ResourceArn: (*string)(ko.Status.ACKResourceMetadata.ARN),
		},
	)
	rm.metrics.RecordAPICall("READ_MANY", "ListTagsForResource", err)
	if err != nil {
		return nil, err
	}

	ko.Spec.Tags = aws.StringMap(output.Tags)