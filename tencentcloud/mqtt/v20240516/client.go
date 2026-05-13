// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20240516

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2024-05-16"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddClientSubscriptionRequest() (request *AddClientSubscriptionRequest) {
    request = &AddClientSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "AddClientSubscription")
    
    
    return
}

func NewAddClientSubscriptionResponse() (response *AddClientSubscriptionResponse) {
    response = &AddClientSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddClientSubscription
// This API is used to add a subscription for an MQTT client.
func (c *Client) AddClientSubscription(request *AddClientSubscriptionRequest) (response *AddClientSubscriptionResponse, err error) {
    return c.AddClientSubscriptionWithContext(context.Background(), request)
}

// AddClientSubscription
// This API is used to add a subscription for an MQTT client.
func (c *Client) AddClientSubscriptionWithContext(ctx context.Context, request *AddClientSubscriptionRequest) (response *AddClientSubscriptionResponse, err error) {
    if request == nil {
        request = NewAddClientSubscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "AddClientSubscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClientSubscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClientSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthorizationPolicyRequest() (request *CreateAuthorizationPolicyRequest) {
    request = &CreateAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateAuthorizationPolicy")
    
    
    return
}

func NewCreateAuthorizationPolicyResponse() (response *CreateAuthorizationPolicyResponse) {
    response = &CreateAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuthorizationPolicy
// This API is used to create a performance test task for an MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEPOLICY = "FailedOperation.DuplicatePolicy"
//  FAILEDOPERATION_DUPLICATEPRIORITY = "FailedOperation.DuplicatePriority"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) CreateAuthorizationPolicy(request *CreateAuthorizationPolicyRequest) (response *CreateAuthorizationPolicyResponse, err error) {
    return c.CreateAuthorizationPolicyWithContext(context.Background(), request)
}

// CreateAuthorizationPolicy
// This API is used to create a performance test task for an MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEPOLICY = "FailedOperation.DuplicatePolicy"
//  FAILEDOPERATION_DUPLICATEPRIORITY = "FailedOperation.DuplicatePriority"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) CreateAuthorizationPolicyWithContext(ctx context.Context, request *CreateAuthorizationPolicyRequest) (response *CreateAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAuthorizationPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "CreateAuthorizationPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceIdentityRequest() (request *CreateDeviceIdentityRequest) {
    request = &CreateDeviceIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateDeviceIdentity")
    
    
    return
}

func NewCreateDeviceIdentityResponse() (response *CreateDeviceIdentityResponse) {
    response = &CreateDeviceIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeviceIdentity
// Create a device signature for per-device secret
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateDeviceIdentity(request *CreateDeviceIdentityRequest) (response *CreateDeviceIdentityResponse, err error) {
    return c.CreateDeviceIdentityWithContext(context.Background(), request)
}

// CreateDeviceIdentity
// Create a device signature for per-device secret
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateDeviceIdentityWithContext(ctx context.Context, request *CreateDeviceIdentityRequest) (response *CreateDeviceIdentityResponse, err error) {
    if request == nil {
        request = NewCreateDeviceIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "CreateDeviceIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeviceIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// This API is used to purchase a new MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INSTANCETYPENOTMATCH = "InvalidParameter.InstanceTypeNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PUBLICNETWORKINVALIDPARAMETERVALUE = "InvalidParameterValue.PublicNetworkInvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// This API is used to purchase a new MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INSTANCETYPENOTMATCH = "InvalidParameter.InstanceTypeNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PUBLICNETWORKINVALIDPARAMETERVALUE = "InvalidParameterValue.PublicNetworkInvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "CreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMessageEnrichmentRuleRequest() (request *CreateMessageEnrichmentRuleRequest) {
    request = &CreateMessageEnrichmentRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateMessageEnrichmentRule")
    
    
    return
}

func NewCreateMessageEnrichmentRuleResponse() (response *CreateMessageEnrichmentRuleResponse) {
    response = &CreateMessageEnrichmentRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMessageEnrichmentRule
// This API is used to create a message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
func (c *Client) CreateMessageEnrichmentRule(request *CreateMessageEnrichmentRuleRequest) (response *CreateMessageEnrichmentRuleResponse, err error) {
    return c.CreateMessageEnrichmentRuleWithContext(context.Background(), request)
}

// CreateMessageEnrichmentRule
// This API is used to create a message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
func (c *Client) CreateMessageEnrichmentRuleWithContext(ctx context.Context, request *CreateMessageEnrichmentRuleRequest) (response *CreateMessageEnrichmentRuleResponse, err error) {
    if request == nil {
        request = NewCreateMessageEnrichmentRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "CreateMessageEnrichmentRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMessageEnrichmentRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMessageEnrichmentRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to add an MQTT role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to add an MQTT role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuthorizationPolicyRequest() (request *DeleteAuthorizationPolicyRequest) {
    request = &DeleteAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteAuthorizationPolicy")
    
    
    return
}

func NewDeleteAuthorizationPolicyResponse() (response *DeleteAuthorizationPolicyResponse) {
    response = &DeleteAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuthorizationPolicy
// This API is used to delete policy rules.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_AUTHORIZATIONPOLICY = "ResourceNotFound.AuthorizationPolicy"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteAuthorizationPolicy(request *DeleteAuthorizationPolicyRequest) (response *DeleteAuthorizationPolicyResponse, err error) {
    return c.DeleteAuthorizationPolicyWithContext(context.Background(), request)
}

// DeleteAuthorizationPolicy
// This API is used to delete policy rules.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_AUTHORIZATIONPOLICY = "ResourceNotFound.AuthorizationPolicy"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteAuthorizationPolicyWithContext(ctx context.Context, request *DeleteAuthorizationPolicyRequest) (response *DeleteAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAuthorizationPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteAuthorizationPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClientSubscriptionRequest() (request *DeleteClientSubscriptionRequest) {
    request = &DeleteClientSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteClientSubscription")
    
    
    return
}

func NewDeleteClientSubscriptionResponse() (response *DeleteClientSubscriptionResponse) {
    response = &DeleteClientSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClientSubscription
// This API is used to delete a subscription of an MQTT client.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_AUTHORIZATIONPOLICY = "ResourceNotFound.AuthorizationPolicy"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteClientSubscription(request *DeleteClientSubscriptionRequest) (response *DeleteClientSubscriptionResponse, err error) {
    return c.DeleteClientSubscriptionWithContext(context.Background(), request)
}

// DeleteClientSubscription
// This API is used to delete a subscription of an MQTT client.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_AUTHORIZATIONPOLICY = "ResourceNotFound.AuthorizationPolicy"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteClientSubscriptionWithContext(ctx context.Context, request *DeleteClientSubscriptionRequest) (response *DeleteClientSubscriptionResponse, err error) {
    if request == nil {
        request = NewDeleteClientSubscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteClientSubscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClientSubscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClientSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceIdentityRequest() (request *DeleteDeviceIdentityRequest) {
    request = &DeleteDeviceIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteDeviceIdentity")
    
    
    return
}

func NewDeleteDeviceIdentityResponse() (response *DeleteDeviceIdentityResponse) {
    response = &DeleteDeviceIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeviceIdentity
// Delete a device signature
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) DeleteDeviceIdentity(request *DeleteDeviceIdentityRequest) (response *DeleteDeviceIdentityResponse, err error) {
    return c.DeleteDeviceIdentityWithContext(context.Background(), request)
}

// DeleteDeviceIdentity
// Delete a device signature
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) DeleteDeviceIdentityWithContext(ctx context.Context, request *DeleteDeviceIdentityRequest) (response *DeleteDeviceIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteDeviceIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeviceIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstance
// This API is used to delete an MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// This API is used to delete an MQTT instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMessageEnrichmentRuleRequest() (request *DeleteMessageEnrichmentRuleRequest) {
    request = &DeleteMessageEnrichmentRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteMessageEnrichmentRule")
    
    
    return
}

func NewDeleteMessageEnrichmentRuleResponse() (response *DeleteMessageEnrichmentRuleResponse) {
    response = &DeleteMessageEnrichmentRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMessageEnrichmentRule
// This API is used to delete the message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMessageEnrichmentRule(request *DeleteMessageEnrichmentRuleRequest) (response *DeleteMessageEnrichmentRuleResponse, err error) {
    return c.DeleteMessageEnrichmentRuleWithContext(context.Background(), request)
}

// DeleteMessageEnrichmentRule
// This API is used to delete the message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteMessageEnrichmentRuleWithContext(ctx context.Context, request *DeleteMessageEnrichmentRuleRequest) (response *DeleteMessageEnrichmentRuleResponse, err error) {
    if request == nil {
        request = NewDeleteMessageEnrichmentRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteMessageEnrichmentRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMessageEnrichmentRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMessageEnrichmentRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// This API is used to delete an MQTT topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// This API is used to delete an MQTT topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete an MQTT user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete an MQTT user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DeleteUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthorizationPoliciesRequest() (request *DescribeAuthorizationPoliciesRequest) {
    request = &DescribeAuthorizationPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeAuthorizationPolicies")
    
    
    return
}

func NewDescribeAuthorizationPoliciesResponse() (response *DescribeAuthorizationPoliciesResponse) {
    response = &DescribeAuthorizationPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuthorizationPolicies
// This API is used to query authorization rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeAuthorizationPolicies(request *DescribeAuthorizationPoliciesRequest) (response *DescribeAuthorizationPoliciesResponse, err error) {
    return c.DescribeAuthorizationPoliciesWithContext(context.Background(), request)
}

// DescribeAuthorizationPolicies
// This API is used to query authorization rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeAuthorizationPoliciesWithContext(ctx context.Context, request *DescribeAuthorizationPoliciesRequest) (response *DescribeAuthorizationPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAuthorizationPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeAuthorizationPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthorizationPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthorizationPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientListRequest() (request *DescribeClientListRequest) {
    request = &DescribeClientListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeClientList")
    
    
    return
}

func NewDescribeClientListResponse() (response *DescribeClientListResponse) {
    response = &DescribeClientListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClientList
// This API is used to query the MQTT client details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeClientList(request *DescribeClientListRequest) (response *DescribeClientListResponse, err error) {
    return c.DescribeClientListWithContext(context.Background(), request)
}

// DescribeClientList
// This API is used to query the MQTT client details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeClientListWithContext(ctx context.Context, request *DescribeClientListRequest) (response *DescribeClientListResponse, err error) {
    if request == nil {
        request = NewDescribeClientListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeClientList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceCertificatesRequest() (request *DescribeDeviceCertificatesRequest) {
    request = &DescribeDeviceCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeDeviceCertificates")
    
    
    return
}

func NewDescribeDeviceCertificatesResponse() (response *DescribeDeviceCertificatesResponse) {
    response = &DescribeDeviceCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceCertificates
// Query device certificates with paging
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeDeviceCertificates(request *DescribeDeviceCertificatesRequest) (response *DescribeDeviceCertificatesResponse, err error) {
    return c.DescribeDeviceCertificatesWithContext(context.Background(), request)
}

// DescribeDeviceCertificates
// Query device certificates with paging
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"
func (c *Client) DescribeDeviceCertificatesWithContext(ctx context.Context, request *DescribeDeviceCertificatesRequest) (response *DescribeDeviceCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeDeviceCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceIdentitiesRequest() (request *DescribeDeviceIdentitiesRequest) {
    request = &DescribeDeviceIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeDeviceIdentities")
    
    
    return
}

func NewDescribeDeviceIdentitiesResponse() (response *DescribeDeviceIdentitiesResponse) {
    response = &DescribeDeviceIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceIdentities
// Query the list of device identifiers in a cluster
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
func (c *Client) DescribeDeviceIdentities(request *DescribeDeviceIdentitiesRequest) (response *DescribeDeviceIdentitiesResponse, err error) {
    return c.DescribeDeviceIdentitiesWithContext(context.Background(), request)
}

// DescribeDeviceIdentities
// Query the list of device identifiers in a cluster
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
func (c *Client) DescribeDeviceIdentitiesWithContext(ctx context.Context, request *DescribeDeviceIdentitiesRequest) (response *DescribeDeviceIdentitiesResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceIdentitiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeDeviceIdentities")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceIdentityRequest() (request *DescribeDeviceIdentityRequest) {
    request = &DescribeDeviceIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeDeviceIdentity")
    
    
    return
}

func NewDescribeDeviceIdentityResponse() (response *DescribeDeviceIdentityResponse) {
    response = &DescribeDeviceIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceIdentity
// Query device identification
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceIdentity(request *DescribeDeviceIdentityRequest) (response *DescribeDeviceIdentityResponse, err error) {
    return c.DescribeDeviceIdentityWithContext(context.Background(), request)
}

// DescribeDeviceIdentity
// Query device identification
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeviceIdentityWithContext(ctx context.Context, request *DescribeDeviceIdentityRequest) (response *DescribeDeviceIdentityResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeDeviceIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// This API is used to query instance information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// This API is used to query instance information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// Get instance list. Description of the Filters parameter use is as follows:
//
// 1. InstanceName, fuzzy search by name
//
// 2. InstanceId, query by instance ID
//
// 3. InstanceStatus, instance status query, supports multiple selections
//
// 
//
// When using the TagFilters parameter for search, the filters parameter is invalid.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// Get instance list. Description of the Filters parameter use is as follows:
//
// 1. InstanceName, fuzzy search by name
//
// 2. InstanceId, query by instance ID
//
// 3. InstanceStatus, instance status query, supports multiple selections
//
// 
//
// When using the TagFilters parameter for search, the filters parameter is invalid.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageByTopicRequest() (request *DescribeMessageByTopicRequest) {
    request = &DescribeMessageByTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeMessageByTopic")
    
    
    return
}

func NewDescribeMessageByTopicResponse() (response *DescribeMessageByTopicResponse) {
    response = &DescribeMessageByTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageByTopic
// This API is used to query messages based on subscription.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageByTopic(request *DescribeMessageByTopicRequest) (response *DescribeMessageByTopicResponse, err error) {
    return c.DescribeMessageByTopicWithContext(context.Background(), request)
}

// DescribeMessageByTopic
// This API is used to query messages based on subscription.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageByTopicWithContext(ctx context.Context, request *DescribeMessageByTopicRequest) (response *DescribeMessageByTopicResponse, err error) {
    if request == nil {
        request = NewDescribeMessageByTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeMessageByTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageByTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageByTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageDetailsRequest() (request *DescribeMessageDetailsRequest) {
    request = &DescribeMessageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeMessageDetails")
    
    
    return
}

func NewDescribeMessageDetailsResponse() (response *DescribeMessageDetailsResponse) {
    response = &DescribeMessageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageDetails
// This API is used to query the MQTT message details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessageDetails(request *DescribeMessageDetailsRequest) (response *DescribeMessageDetailsResponse, err error) {
    return c.DescribeMessageDetailsWithContext(context.Background(), request)
}

// DescribeMessageDetails
// This API is used to query the MQTT message details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessageDetailsWithContext(ctx context.Context, request *DescribeMessageDetailsRequest) (response *DescribeMessageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeMessageDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeMessageDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageEnrichmentRulesRequest() (request *DescribeMessageEnrichmentRulesRequest) {
    request = &DescribeMessageEnrichmentRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeMessageEnrichmentRules")
    
    
    return
}

func NewDescribeMessageEnrichmentRulesResponse() (response *DescribeMessageEnrichmentRulesResponse) {
    response = &DescribeMessageEnrichmentRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageEnrichmentRules
// This API is used to query message enrichment rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageEnrichmentRules(request *DescribeMessageEnrichmentRulesRequest) (response *DescribeMessageEnrichmentRulesResponse, err error) {
    return c.DescribeMessageEnrichmentRulesWithContext(context.Background(), request)
}

// DescribeMessageEnrichmentRules
// This API is used to query message enrichment rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageEnrichmentRulesWithContext(ctx context.Context, request *DescribeMessageEnrichmentRulesRequest) (response *DescribeMessageEnrichmentRulesResponse, err error) {
    if request == nil {
        request = NewDescribeMessageEnrichmentRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeMessageEnrichmentRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageEnrichmentRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageEnrichmentRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductSKUListRequest() (request *DescribeProductSKUListRequest) {
    request = &DescribeProductSKUListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeProductSKUList")
    
    
    return
}

func NewDescribeProductSKUListResponse() (response *DescribeProductSKUListResponse) {
    response = &DescribeProductSKUListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductSKUList
// This API is used to obtain product sales specifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeProductSKUList(request *DescribeProductSKUListRequest) (response *DescribeProductSKUListResponse, err error) {
    return c.DescribeProductSKUListWithContext(context.Background(), request)
}

// DescribeProductSKUList
// This API is used to obtain product sales specifications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeProductSKUListWithContext(ctx context.Context, request *DescribeProductSKUListRequest) (response *DescribeProductSKUListResponse, err error) {
    if request == nil {
        request = NewDescribeProductSKUListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeProductSKUList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductSKUList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductSKUListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedSubscriptionClientRequest() (request *DescribeSharedSubscriptionClientRequest) {
    request = &DescribeSharedSubscriptionClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeSharedSubscriptionClient")
    
    
    return
}

func NewDescribeSharedSubscriptionClientResponse() (response *DescribeSharedSubscriptionClientResponse) {
    response = &DescribeSharedSubscriptionClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSharedSubscriptionClient
// Query shared subscription group detailed information
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptionClient(request *DescribeSharedSubscriptionClientRequest) (response *DescribeSharedSubscriptionClientResponse, err error) {
    return c.DescribeSharedSubscriptionClientWithContext(context.Background(), request)
}

// DescribeSharedSubscriptionClient
// Query shared subscription group detailed information
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptionClientWithContext(ctx context.Context, request *DescribeSharedSubscriptionClientRequest) (response *DescribeSharedSubscriptionClientResponse, err error) {
    if request == nil {
        request = NewDescribeSharedSubscriptionClientRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeSharedSubscriptionClient")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedSubscriptionClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSharedSubscriptionClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedSubscriptionGroupsRequest() (request *DescribeSharedSubscriptionGroupsRequest) {
    request = &DescribeSharedSubscriptionGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeSharedSubscriptionGroups")
    
    
    return
}

func NewDescribeSharedSubscriptionGroupsResponse() (response *DescribeSharedSubscriptionGroupsResponse) {
    response = &DescribeSharedSubscriptionGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSharedSubscriptionGroups
// This API is used to query the subscription group list shared within the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptionGroups(request *DescribeSharedSubscriptionGroupsRequest) (response *DescribeSharedSubscriptionGroupsResponse, err error) {
    return c.DescribeSharedSubscriptionGroupsWithContext(context.Background(), request)
}

// DescribeSharedSubscriptionGroups
// This API is used to query the subscription group list shared within the cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptionGroupsWithContext(ctx context.Context, request *DescribeSharedSubscriptionGroupsRequest) (response *DescribeSharedSubscriptionGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSharedSubscriptionGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeSharedSubscriptionGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedSubscriptionGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSharedSubscriptionGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSharedSubscriptionsRequest() (request *DescribeSharedSubscriptionsRequest) {
    request = &DescribeSharedSubscriptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeSharedSubscriptions")
    
    
    return
}

func NewDescribeSharedSubscriptionsResponse() (response *DescribeSharedSubscriptionsResponse) {
    response = &DescribeSharedSubscriptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSharedSubscriptions
// Query the subscription list of a shared subscription group
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptions(request *DescribeSharedSubscriptionsRequest) (response *DescribeSharedSubscriptionsResponse, err error) {
    return c.DescribeSharedSubscriptionsWithContext(context.Background(), request)
}

// DescribeSharedSubscriptions
// Query the subscription list of a shared subscription group
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeSharedSubscriptionsWithContext(ctx context.Context, request *DescribeSharedSubscriptionsRequest) (response *DescribeSharedSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSharedSubscriptionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeSharedSubscriptions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSharedSubscriptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSharedSubscriptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopic
// This API is used to query the MQTT topic details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// This API is used to query the MQTT topic details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopicWithContext(ctx context.Context, request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// This API is used to query the user list. Filter parameter usage instructions are as follows:.
//
// 
//
// This API is used to perform Username fuzzy search.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// This API is used to query the user list. Filter parameter usage instructions are as follows:.
//
// 
//
// This API is used to perform Username fuzzy search.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "DescribeUserList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewKickOutClientRequest() (request *KickOutClientRequest) {
    request = &KickOutClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "KickOutClient")
    
    
    return
}

func NewKickOutClientResponse() (response *KickOutClientResponse) {
    response = &KickOutClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KickOutClient
// This API is used to kick out a client.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) KickOutClient(request *KickOutClientRequest) (response *KickOutClientResponse, err error) {
    return c.KickOutClientWithContext(context.Background(), request)
}

// KickOutClient
// This API is used to kick out a client.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) KickOutClientWithContext(ctx context.Context, request *KickOutClientRequest) (response *KickOutClientResponse, err error) {
    if request == nil {
        request = NewKickOutClientRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "KickOutClient")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KickOutClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewKickOutClientResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuthorizationPolicyRequest() (request *ModifyAuthorizationPolicyRequest) {
    request = &ModifyAuthorizationPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyAuthorizationPolicy")
    
    
    return
}

func NewModifyAuthorizationPolicyResponse() (response *ModifyAuthorizationPolicyResponse) {
    response = &ModifyAuthorizationPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuthorizationPolicy
// This API is used to modify policy rules. See the data plane authorization policy description (https://www.tencentcloud.com/document/product/1778/109715?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEPOLICY = "FailedOperation.DuplicatePolicy"
//  FAILEDOPERATION_DUPLICATEPRIORITY = "FailedOperation.DuplicatePriority"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyAuthorizationPolicy(request *ModifyAuthorizationPolicyRequest) (response *ModifyAuthorizationPolicyResponse, err error) {
    return c.ModifyAuthorizationPolicyWithContext(context.Background(), request)
}

// ModifyAuthorizationPolicy
// This API is used to modify policy rules. See the data plane authorization policy description (https://www.tencentcloud.com/document/product/1778/109715?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEPOLICY = "FailedOperation.DuplicatePolicy"
//  FAILEDOPERATION_DUPLICATEPRIORITY = "FailedOperation.DuplicatePriority"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyAuthorizationPolicyWithContext(ctx context.Context, request *ModifyAuthorizationPolicyRequest) (response *ModifyAuthorizationPolicyResponse, err error) {
    if request == nil {
        request = NewModifyAuthorizationPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "ModifyAuthorizationPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuthorizationPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuthorizationPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceIdentityRequest() (request *ModifyDeviceIdentityRequest) {
    request = &ModifyDeviceIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyDeviceIdentity")
    
    
    return
}

func NewModifyDeviceIdentityResponse() (response *ModifyDeviceIdentityResponse) {
    response = &ModifyDeviceIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDeviceIdentity
// Modify the device signature
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) ModifyDeviceIdentity(request *ModifyDeviceIdentityRequest) (response *ModifyDeviceIdentityResponse, err error) {
    return c.ModifyDeviceIdentityWithContext(context.Background(), request)
}

// ModifyDeviceIdentity
// Modify the device signature
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) ModifyDeviceIdentityWithContext(ctx context.Context, request *ModifyDeviceIdentityRequest) (response *ModifyDeviceIdentityResponse, err error) {
    if request == nil {
        request = NewModifyDeviceIdentityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "ModifyDeviceIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// This API is used to modify instance attributes. Only running clusters can call this API to perform configuration change.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  FAILEDOPERATION_NOTSUPPORTDISABLEAUTHORIZATIONPOLICY = "FailedOperation.NotSupportDisableAuthorizationPolicy"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to modify instance attributes. Only running clusters can call this API to perform configuration change.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  FAILEDOPERATION_NOTSUPPORTDISABLEAUTHORIZATIONPOLICY = "FailedOperation.NotSupportDisableAuthorizationPolicy"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "ModifyInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMessageEnrichmentRuleRequest() (request *ModifyMessageEnrichmentRuleRequest) {
    request = &ModifyMessageEnrichmentRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyMessageEnrichmentRule")
    
    
    return
}

func NewModifyMessageEnrichmentRuleResponse() (response *ModifyMessageEnrichmentRuleResponse) {
    response = &ModifyMessageEnrichmentRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMessageEnrichmentRule
// This API is used to modify message enrichment rules.
//
// Note: All attributes of the current rule must be submitted, even if specific fields are not modified.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyMessageEnrichmentRule(request *ModifyMessageEnrichmentRuleRequest) (response *ModifyMessageEnrichmentRuleResponse, err error) {
    return c.ModifyMessageEnrichmentRuleWithContext(context.Background(), request)
}

// ModifyMessageEnrichmentRule
// This API is used to modify message enrichment rules.
//
// Note: All attributes of the current rule must be submitted, even if specific fields are not modified.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyMessageEnrichmentRuleWithContext(ctx context.Context, request *ModifyMessageEnrichmentRuleRequest) (response *ModifyMessageEnrichmentRuleResponse, err error) {
    if request == nil {
        request = NewModifyMessageEnrichmentRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "ModifyMessageEnrichmentRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMessageEnrichmentRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMessageEnrichmentRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUser
// This API is used to modify an MQTT role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    return c.ModifyUserWithContext(context.Background(), request)
}

// ModifyUser
// This API is used to modify an MQTT role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "ModifyUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuthorizationPolicyPriorityRequest() (request *UpdateAuthorizationPolicyPriorityRequest) {
    request = &UpdateAuthorizationPolicyPriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "UpdateAuthorizationPolicyPriority")
    
    
    return
}

func NewUpdateAuthorizationPolicyPriorityResponse() (response *UpdateAuthorizationPolicyPriorityResponse) {
    response = &UpdateAuthorizationPolicyPriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAuthorizationPolicyPriority
// This API is used to modify policy rule priority.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEAUTHORIZATIONIDORPRIORITY = "FailedOperation.DuplicateAuthorizationIdOrPriority"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateAuthorizationPolicyPriority(request *UpdateAuthorizationPolicyPriorityRequest) (response *UpdateAuthorizationPolicyPriorityResponse, err error) {
    return c.UpdateAuthorizationPolicyPriorityWithContext(context.Background(), request)
}

// UpdateAuthorizationPolicyPriority
// This API is used to modify policy rule priority.
//
// error code that may be returned:
//  FAILEDOPERATION_DUPLICATEAUTHORIZATIONIDORPRIORITY = "FailedOperation.DuplicateAuthorizationIdOrPriority"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateAuthorizationPolicyPriorityWithContext(ctx context.Context, request *UpdateAuthorizationPolicyPriorityRequest) (response *UpdateAuthorizationPolicyPriorityResponse, err error) {
    if request == nil {
        request = NewUpdateAuthorizationPolicyPriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "UpdateAuthorizationPolicyPriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAuthorizationPolicyPriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAuthorizationPolicyPriorityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateMessageEnrichmentRulePriorityRequest() (request *UpdateMessageEnrichmentRulePriorityRequest) {
    request = &UpdateMessageEnrichmentRulePriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mqtt", APIVersion, "UpdateMessageEnrichmentRulePriority")
    
    
    return
}

func NewUpdateMessageEnrichmentRulePriorityResponse() (response *UpdateMessageEnrichmentRulePriorityResponse) {
    response = &UpdateMessageEnrichmentRulePriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateMessageEnrichmentRulePriority
// This API is used to modify the priority for message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateMessageEnrichmentRulePriority(request *UpdateMessageEnrichmentRulePriorityRequest) (response *UpdateMessageEnrichmentRulePriorityResponse, err error) {
    return c.UpdateMessageEnrichmentRulePriorityWithContext(context.Background(), request)
}

// UpdateMessageEnrichmentRulePriority
// This API is used to modify the priority for message enrichment rule.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) UpdateMessageEnrichmentRulePriorityWithContext(ctx context.Context, request *UpdateMessageEnrichmentRulePriorityRequest) (response *UpdateMessageEnrichmentRulePriorityResponse, err error) {
    if request == nil {
        request = NewUpdateMessageEnrichmentRulePriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mqtt", APIVersion, "UpdateMessageEnrichmentRulePriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateMessageEnrichmentRulePriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateMessageEnrichmentRulePriorityResponse()
    err = c.Send(request, response)
    return
}
