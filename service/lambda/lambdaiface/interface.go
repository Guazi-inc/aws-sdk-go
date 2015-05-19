// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambdaiface provides an interface that satisfies the lambda service.
package lambdaiface

import (
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/lambda"
)

type LambdaAPI interface {
	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, awserr.Error)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, awserr.Error)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, awserr.Error)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, awserr.Error)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, awserr.Error)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, awserr.Error)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, awserr.Error)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, awserr.Error)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, awserr.Error)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, awserr.Error)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, awserr.Error)

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, awserr.Error)

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, awserr.Error)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, awserr.Error)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, awserr.Error)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, awserr.Error)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, awserr.Error)
}