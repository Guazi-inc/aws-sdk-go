// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatchlogsiface provides an interface that satisfies the cloudwatchlogs service.
package cloudwatchlogsiface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/cloudwatchlogs"
)

type CloudWatchLogsAPI interface {
	CreateLogGroup(*cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, awserr.Error)

	CreateLogStream(*cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, awserr.Error)

	DeleteLogGroup(*cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, awserr.Error)

	DeleteLogStream(*cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, awserr.Error)

	DeleteMetricFilter(*cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, awserr.Error)

	DeleteRetentionPolicy(*cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, awserr.Error)

	DescribeLogGroups(*cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, awserr.Error)

	DescribeLogStreams(*cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, awserr.Error)

	DescribeMetricFilters(*cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, awserr.Error)

	GetLogEvents(*cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, awserr.Error)

	PutLogEvents(*cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, awserr.Error)

	PutMetricFilter(*cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, awserr.Error)

	PutRetentionPolicy(*cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, awserr.Error)

	TestMetricFilter(*cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, awserr.Error)
}