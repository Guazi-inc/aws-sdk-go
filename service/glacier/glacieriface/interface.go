// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package glacieriface provides an interface that satisfies the glacier service.
package glacieriface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/glacier"
)

type GlacierAPI interface {
	AbortMultipartUpload(*glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, awserr.Error)

	CompleteMultipartUpload(*glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, awserr.Error)

	CreateVault(*glacier.CreateVaultInput) (*glacier.CreateVaultOutput, awserr.Error)

	DeleteArchive(*glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, awserr.Error)

	DeleteVault(*glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, awserr.Error)

	DeleteVaultNotifications(*glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, awserr.Error)

	DescribeJob(*glacier.DescribeJobInput) (*glacier.GlacierJobDescription, awserr.Error)

	DescribeVault(*glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, awserr.Error)

	GetDataRetrievalPolicy(*glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, awserr.Error)

	GetJobOutput(*glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, awserr.Error)

	GetVaultNotifications(*glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, awserr.Error)

	InitiateJob(*glacier.InitiateJobInput) (*glacier.InitiateJobOutput, awserr.Error)

	InitiateMultipartUpload(*glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, awserr.Error)

	ListJobs(*glacier.ListJobsInput) (*glacier.ListJobsOutput, awserr.Error)

	ListMultipartUploads(*glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, awserr.Error)

	ListParts(*glacier.ListPartsInput) (*glacier.ListPartsOutput, awserr.Error)

	ListVaults(*glacier.ListVaultsInput) (*glacier.ListVaultsOutput, awserr.Error)

	SetDataRetrievalPolicy(*glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, awserr.Error)

	SetVaultNotifications(*glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, awserr.Error)

	UploadArchive(*glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, awserr.Error)

	UploadMultipartPart(*glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, awserr.Error)
}