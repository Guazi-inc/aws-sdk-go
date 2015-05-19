// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sqsiface provides an interface that satisfies the sqs service.
package sqsiface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/sqs"
)

type SQSAPI interface {
	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, awserr.Error)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, awserr.Error)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, awserr.Error)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, awserr.Error)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, awserr.Error)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, awserr.Error)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, awserr.Error)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, awserr.Error)

	GetQueueURL(*sqs.GetQueueURLInput) (*sqs.GetQueueURLOutput, awserr.Error)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, awserr.Error)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, awserr.Error)

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, awserr.Error)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, awserr.Error)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, awserr.Error)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, awserr.Error)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, awserr.Error)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, awserr.Error)
}