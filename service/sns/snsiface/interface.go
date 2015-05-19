// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package snsiface provides an interface that satisfies the sns service.
package snsiface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/sns"
)

type SNSAPI interface {
	AddPermission(*sns.AddPermissionInput) (*sns.AddPermissionOutput, awserr.Error)

	ConfirmSubscription(*sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, awserr.Error)

	CreatePlatformApplication(*sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, awserr.Error)

	CreatePlatformEndpoint(*sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, awserr.Error)

	CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, awserr.Error)

	DeleteEndpoint(*sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, awserr.Error)

	DeletePlatformApplication(*sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, awserr.Error)

	DeleteTopic(*sns.DeleteTopicInput) (*sns.DeleteTopicOutput, awserr.Error)

	GetEndpointAttributes(*sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, awserr.Error)

	GetPlatformApplicationAttributes(*sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, awserr.Error)

	GetSubscriptionAttributes(*sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, awserr.Error)

	GetTopicAttributes(*sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, awserr.Error)

	ListEndpointsByPlatformApplication(*sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, awserr.Error)

	ListPlatformApplications(*sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, awserr.Error)

	ListSubscriptions(*sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, awserr.Error)

	ListSubscriptionsByTopic(*sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, awserr.Error)

	ListTopics(*sns.ListTopicsInput) (*sns.ListTopicsOutput, awserr.Error)

	Publish(*sns.PublishInput) (*sns.PublishOutput, awserr.Error)

	RemovePermission(*sns.RemovePermissionInput) (*sns.RemovePermissionOutput, awserr.Error)

	SetEndpointAttributes(*sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, awserr.Error)

	SetPlatformApplicationAttributes(*sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, awserr.Error)

	SetSubscriptionAttributes(*sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, awserr.Error)

	SetTopicAttributes(*sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, awserr.Error)

	Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, awserr.Error)

	Unsubscribe(*sns.UnsubscribeInput) (*sns.UnsubscribeOutput, awserr.Error)
}