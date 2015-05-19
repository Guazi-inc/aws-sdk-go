// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudsearchdomainiface provides an interface that satisfies the cloudsearchdomain service.
package cloudsearchdomainiface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/cloudsearchdomain"
)

type CloudSearchDomainAPI interface {
	Search(*cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, awserr.Error)

	Suggest(*cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, awserr.Error)

	UploadDocuments(*cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, awserr.Error)
}