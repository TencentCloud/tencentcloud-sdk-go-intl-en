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

package v20200217

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-02-17"

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


func NewAcknowledgeMessageRequest() (request *AcknowledgeMessageRequest) {
    request = &AcknowledgeMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "AcknowledgeMessage")
    
    
    return
}

func NewAcknowledgeMessageResponse() (response *AcknowledgeMessageResponse) {
    response = &AcknowledgeMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcknowledgeMessage
// This API is used to acknowledge the message in the specified topic by the provided `MessageID`.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) AcknowledgeMessage(request *AcknowledgeMessageRequest) (response *AcknowledgeMessageResponse, err error) {
    return c.AcknowledgeMessageWithContext(context.Background(), request)
}

// AcknowledgeMessage
// This API is used to acknowledge the message in the specified topic by the provided `MessageID`.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) AcknowledgeMessageWithContext(ctx context.Context, request *AcknowledgeMessageRequest) (response *AcknowledgeMessageResponse, err error) {
    if request == nil {
        request = NewAcknowledgeMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "AcknowledgeMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcknowledgeMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcknowledgeMessageResponse()
    err = c.Send(request, response)
    return
}

func NewClearCmqQueueRequest() (request *ClearCmqQueueRequest) {
    request = &ClearCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqQueue")
    
    
    return
}

func NewClearCmqQueueResponse() (response *ClearCmqQueueResponse) {
    response = &ClearCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearCmqQueue
// This API is used to clear the messages in the CMQ message queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ClearCmqQueue(request *ClearCmqQueueRequest) (response *ClearCmqQueueResponse, err error) {
    return c.ClearCmqQueueWithContext(context.Background(), request)
}

// ClearCmqQueue
// This API is used to clear the messages in the CMQ message queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ClearCmqQueueWithContext(ctx context.Context, request *ClearCmqQueueRequest) (response *ClearCmqQueueResponse, err error) {
    if request == nil {
        request = NewClearCmqQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ClearCmqQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewClearCmqSubscriptionFilterTagsRequest() (request *ClearCmqSubscriptionFilterTagsRequest) {
    request = &ClearCmqSubscriptionFilterTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqSubscriptionFilterTags")
    
    
    return
}

func NewClearCmqSubscriptionFilterTagsResponse() (response *ClearCmqSubscriptionFilterTagsResponse) {
    response = &ClearCmqSubscriptionFilterTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearCmqSubscriptionFilterTags
// This API is used to clear the message tags of a subscriber.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ClearCmqSubscriptionFilterTags(request *ClearCmqSubscriptionFilterTagsRequest) (response *ClearCmqSubscriptionFilterTagsResponse, err error) {
    return c.ClearCmqSubscriptionFilterTagsWithContext(context.Background(), request)
}

// ClearCmqSubscriptionFilterTags
// This API is used to clear the message tags of a subscriber.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ClearCmqSubscriptionFilterTagsWithContext(ctx context.Context, request *ClearCmqSubscriptionFilterTagsRequest) (response *ClearCmqSubscriptionFilterTagsResponse, err error) {
    if request == nil {
        request = NewClearCmqSubscriptionFilterTagsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ClearCmqSubscriptionFilterTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearCmqSubscriptionFilterTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearCmqSubscriptionFilterTagsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCluster
// This API is used to create a cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_CLUSTER = "ResourceInUse.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    return c.CreateClusterWithContext(context.Background(), request)
}

// CreateCluster
// This API is used to create a cluster.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_CLUSTER = "ResourceInUse.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqQueueRequest() (request *CreateCmqQueueRequest) {
    request = &CreateCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqQueue")
    
    
    return
}

func NewCreateCmqQueueResponse() (response *CreateCmqQueueResponse) {
    response = &CreateCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCmqQueue
// This API is used to create a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_QUEUE = "ResourceInUse.Queue"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmqQueue(request *CreateCmqQueueRequest) (response *CreateCmqQueueResponse, err error) {
    return c.CreateCmqQueueWithContext(context.Background(), request)
}

// CreateCmqQueue
// This API is used to create a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_QUEUE = "ResourceInUse.Queue"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCmqQueueWithContext(ctx context.Context, request *CreateCmqQueueRequest) (response *CreateCmqQueueResponse, err error) {
    if request == nil {
        request = NewCreateCmqQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateCmqQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqSubscribeRequest() (request *CreateCmqSubscribeRequest) {
    request = &CreateCmqSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqSubscribe")
    
    
    return
}

func NewCreateCmqSubscribeResponse() (response *CreateCmqSubscribeResponse) {
    response = &CreateCmqSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCmqSubscribe
// This API is used to create a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqSubscribe(request *CreateCmqSubscribeRequest) (response *CreateCmqSubscribeResponse, err error) {
    return c.CreateCmqSubscribeWithContext(context.Background(), request)
}

// CreateCmqSubscribe
// This API is used to create a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqSubscribeWithContext(ctx context.Context, request *CreateCmqSubscribeRequest) (response *CreateCmqSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateCmqSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateCmqSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCmqTopicRequest() (request *CreateCmqTopicRequest) {
    request = &CreateCmqTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqTopic")
    
    
    return
}

func NewCreateCmqTopicResponse() (response *CreateCmqTopicResponse) {
    response = &CreateCmqTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCmqTopic
// This API is used to create a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqTopic(request *CreateCmqTopicRequest) (response *CreateCmqTopicResponse, err error) {
    return c.CreateCmqTopicWithContext(context.Background(), request)
}

// CreateCmqTopic
// This API is used to create a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateCmqTopicWithContext(ctx context.Context, request *CreateCmqTopicRequest) (response *CreateCmqTopicResponse, err error) {
    if request == nil {
        request = NewCreateCmqTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateCmqTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCmqTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCmqTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironment")
    
    
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvironment
// This API is used to create a TDMQ namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"
//  LIMITEXCEEDED_NAMESPACES = "LimitExceeded.Namespaces"
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// This API is used to create a TDMQ namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"
//  LIMITEXCEEDED_NAMESPACES = "LimitExceeded.Namespaces"
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateEnvironment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRoleRequest() (request *CreateEnvironmentRoleRequest) {
    request = &CreateEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironmentRole")
    
    
    return
}

func NewCreateEnvironmentRoleResponse() (response *CreateEnvironmentRoleResponse) {
    response = &CreateEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvironmentRole
// This API is used to create an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) CreateEnvironmentRole(request *CreateEnvironmentRoleRequest) (response *CreateEnvironmentRoleResponse, err error) {
    return c.CreateEnvironmentRoleWithContext(context.Background(), request)
}

// CreateEnvironmentRole
// This API is used to create an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) CreateEnvironmentRoleWithContext(ctx context.Context, request *CreateEnvironmentRoleRequest) (response *CreateEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateEnvironmentRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQUserRequest() (request *CreateRabbitMQUserRequest) {
    request = &CreateRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQUser")
    
    
    return
}

func NewCreateRabbitMQUserResponse() (response *CreateRabbitMQUserResponse) {
    response = &CreateRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQUser
// This API is used to create a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQUser(request *CreateRabbitMQUserRequest) (response *CreateRabbitMQUserResponse, err error) {
    return c.CreateRabbitMQUserWithContext(context.Background(), request)
}

// CreateRabbitMQUser
// This API is used to create a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQUserWithContext(ctx context.Context, request *CreateRabbitMQUserRequest) (response *CreateRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRabbitMQUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQVipInstanceRequest() (request *CreateRabbitMQVipInstanceRequest) {
    request = &CreateRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVipInstance")
    
    
    return
}

func NewCreateRabbitMQVipInstanceResponse() (response *CreateRabbitMQVipInstanceResponse) {
    response = &CreateRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQVipInstance
// This API is used to create a TDMQ for RabbitMQ exclusive instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVipInstance(request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
    return c.CreateRabbitMQVipInstanceWithContext(context.Background(), request)
}

// CreateRabbitMQVipInstance
// This API is used to create a TDMQ for RabbitMQ exclusive instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVipInstanceWithContext(ctx context.Context, request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQVipInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRabbitMQVipInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQVirtualHostRequest() (request *CreateRabbitMQVirtualHostRequest) {
    request = &CreateRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVirtualHost")
    
    
    return
}

func NewCreateRabbitMQVirtualHostResponse() (response *CreateRabbitMQVirtualHostResponse) {
    response = &CreateRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQVirtualHost
// This API is used to create a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVirtualHost(request *CreateRabbitMQVirtualHostRequest) (response *CreateRabbitMQVirtualHostResponse, err error) {
    return c.CreateRabbitMQVirtualHostWithContext(context.Background(), request)
}

// CreateRabbitMQVirtualHost
// This API is used to create a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVirtualHostWithContext(ctx context.Context, request *CreateRabbitMQVirtualHostRequest) (response *CreateRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQVirtualHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRabbitMQVirtualHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQClusterRequest() (request *CreateRocketMQClusterRequest) {
    request = &CreateRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQCluster")
    
    
    return
}

func NewCreateRocketMQClusterResponse() (response *CreateRocketMQClusterResponse) {
    response = &CreateRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQCluster
// This API is used to create a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQCluster(request *CreateRocketMQClusterRequest) (response *CreateRocketMQClusterResponse, err error) {
    return c.CreateRocketMQClusterWithContext(context.Background(), request)
}

// CreateRocketMQCluster
// This API is used to create a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQClusterWithContext(ctx context.Context, request *CreateRocketMQClusterRequest) (response *CreateRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQGroupRequest() (request *CreateRocketMQGroupRequest) {
    request = &CreateRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQGroup")
    
    
    return
}

func NewCreateRocketMQGroupResponse() (response *CreateRocketMQGroupResponse) {
    response = &CreateRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQGroup
// This API is used to create a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateRocketMQGroup(request *CreateRocketMQGroupRequest) (response *CreateRocketMQGroupResponse, err error) {
    return c.CreateRocketMQGroupWithContext(context.Background(), request)
}

// CreateRocketMQGroup
// This API is used to create a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) CreateRocketMQGroupWithContext(ctx context.Context, request *CreateRocketMQGroupRequest) (response *CreateRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQNamespaceRequest() (request *CreateRocketMQNamespaceRequest) {
    request = &CreateRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQNamespace")
    
    
    return
}

func NewCreateRocketMQNamespaceResponse() (response *CreateRocketMQNamespaceResponse) {
    response = &CreateRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQNamespace
// This API is used to create a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQNamespace(request *CreateRocketMQNamespaceRequest) (response *CreateRocketMQNamespaceResponse, err error) {
    return c.CreateRocketMQNamespaceWithContext(context.Background(), request)
}

// CreateRocketMQNamespace
// This API is used to create a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQNamespaceWithContext(ctx context.Context, request *CreateRocketMQNamespaceRequest) (response *CreateRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRocketMQTopicRequest() (request *CreateRocketMQTopicRequest) {
    request = &CreateRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQTopic")
    
    
    return
}

func NewCreateRocketMQTopicResponse() (response *CreateRocketMQTopicResponse) {
    response = &CreateRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQTopic
// This API is used to create a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQTopic(request *CreateRocketMQTopicRequest) (response *CreateRocketMQTopicResponse, err error) {
    return c.CreateRocketMQTopicWithContext(context.Background(), request)
}

// CreateRocketMQTopic
// This API is used to create a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRocketMQTopicWithContext(ctx context.Context, request *CreateRocketMQTopicRequest) (response *CreateRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRole")
    
    
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRole
// This API is used to create a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_CREATESECRETKEY = "FailedOperation.CreateSecretKey"
//  FAILEDOPERATION_SAVESECRETKEY = "FailedOperation.SaveSecretKey"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ROLE = "ResourceInUse.Role"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

// CreateRole
// This API is used to create a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_CREATESECRETKEY = "FailedOperation.CreateSecretKey"
//  FAILEDOPERATION_SAVESECRETKEY = "FailedOperation.SaveSecretKey"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ROLE = "ResourceInUse.Role"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscriptionRequest() (request *CreateSubscriptionRequest) {
    request = &CreateSubscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateSubscription")
    
    
    return
}

func NewCreateSubscriptionResponse() (response *CreateSubscriptionResponse) {
    response = &CreateSubscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubscription
// This API is used to create a subscription to a topic.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_RETRY = "InternalError.Retry"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateSubscription(request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
    return c.CreateSubscriptionWithContext(context.Background(), request)
}

// CreateSubscription
// This API is used to create a subscription to a topic.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_RETRY = "InternalError.Retry"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateSubscriptionWithContext(ctx context.Context, request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
    if request == nil {
        request = NewCreateSubscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateSubscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopic
// This API is used to add a message topic in the specified partition and type.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// This API is used to add a message topic in the specified partition and type.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCluster
// This API is used to delete a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_VPCINUSE = "FailedOperation.VpcInUse"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    return c.DeleteClusterWithContext(context.Background(), request)
}

// DeleteCluster
// This API is used to delete a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_VPCINUSE = "FailedOperation.VpcInUse"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqQueueRequest() (request *DeleteCmqQueueRequest) {
    request = &DeleteCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqQueue")
    
    
    return
}

func NewDeleteCmqQueueResponse() (response *DeleteCmqQueueResponse) {
    response = &DeleteCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCmqQueue
// This API is used to delete a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqQueue(request *DeleteCmqQueueRequest) (response *DeleteCmqQueueResponse, err error) {
    return c.DeleteCmqQueueWithContext(context.Background(), request)
}

// DeleteCmqQueue
// This API is used to delete a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqQueueWithContext(ctx context.Context, request *DeleteCmqQueueRequest) (response *DeleteCmqQueueResponse, err error) {
    if request == nil {
        request = NewDeleteCmqQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteCmqQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqSubscribeRequest() (request *DeleteCmqSubscribeRequest) {
    request = &DeleteCmqSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqSubscribe")
    
    
    return
}

func NewDeleteCmqSubscribeResponse() (response *DeleteCmqSubscribeResponse) {
    response = &DeleteCmqSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCmqSubscribe
// This API is used to delete a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqSubscribe(request *DeleteCmqSubscribeRequest) (response *DeleteCmqSubscribeResponse, err error) {
    return c.DeleteCmqSubscribeWithContext(context.Background(), request)
}

// DeleteCmqSubscribe
// This API is used to delete a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqSubscribeWithContext(ctx context.Context, request *DeleteCmqSubscribeRequest) (response *DeleteCmqSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteCmqSubscribeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteCmqSubscribe")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCmqTopicRequest() (request *DeleteCmqTopicRequest) {
    request = &DeleteCmqTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqTopic")
    
    
    return
}

func NewDeleteCmqTopicResponse() (response *DeleteCmqTopicResponse) {
    response = &DeleteCmqTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCmqTopic
// This API is used to delete a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqTopic(request *DeleteCmqTopicRequest) (response *DeleteCmqTopicResponse, err error) {
    return c.DeleteCmqTopicWithContext(context.Background(), request)
}

// DeleteCmqTopic
// This API is used to delete a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCmqTopicWithContext(ctx context.Context, request *DeleteCmqTopicRequest) (response *DeleteCmqTopicResponse, err error) {
    if request == nil {
        request = NewDeleteCmqTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteCmqTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCmqTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCmqTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentRolesRequest() (request *DeleteEnvironmentRolesRequest) {
    request = &DeleteEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironmentRoles")
    
    
    return
}

func NewDeleteEnvironmentRolesResponse() (response *DeleteEnvironmentRolesResponse) {
    response = &DeleteEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnvironmentRoles
// This API is used to delete an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteEnvironmentRoles(request *DeleteEnvironmentRolesRequest) (response *DeleteEnvironmentRolesResponse, err error) {
    return c.DeleteEnvironmentRolesWithContext(context.Background(), request)
}

// DeleteEnvironmentRoles
// This API is used to delete an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteEnvironmentRolesWithContext(ctx context.Context, request *DeleteEnvironmentRolesRequest) (response *DeleteEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteEnvironmentRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentsRequest() (request *DeleteEnvironmentsRequest) {
    request = &DeleteEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironments")
    
    
    return
}

func NewDeleteEnvironmentsResponse() (response *DeleteEnvironmentsResponse) {
    response = &DeleteEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnvironments
// This API is used to batch delete namespaces under a tenant.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_TOPICINUSE = "FailedOperation.TopicInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    return c.DeleteEnvironmentsWithContext(context.Background(), request)
}

// DeleteEnvironments
// This API is used to batch delete namespaces under a tenant.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  FAILEDOPERATION_TOPICINUSE = "FailedOperation.TopicInUse"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DeleteEnvironmentsWithContext(ctx context.Context, request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQUserRequest() (request *DeleteRabbitMQUserRequest) {
    request = &DeleteRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQUser")
    
    
    return
}

func NewDeleteRabbitMQUserResponse() (response *DeleteRabbitMQUserResponse) {
    response = &DeleteRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQUser
// This API is used to delete a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQUser(request *DeleteRabbitMQUserRequest) (response *DeleteRabbitMQUserResponse, err error) {
    return c.DeleteRabbitMQUserWithContext(context.Background(), request)
}

// DeleteRabbitMQUser
// This API is used to delete a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQUserWithContext(ctx context.Context, request *DeleteRabbitMQUserRequest) (response *DeleteRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRabbitMQUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQVipInstanceRequest() (request *DeleteRabbitMQVipInstanceRequest) {
    request = &DeleteRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVipInstance")
    
    
    return
}

func NewDeleteRabbitMQVipInstanceResponse() (response *DeleteRabbitMQVipInstanceResponse) {
    response = &DeleteRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQVipInstance
// This API is used to delete a TDMQ for RabbitMQ exclusive instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQVipInstance(request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
    return c.DeleteRabbitMQVipInstanceWithContext(context.Background(), request)
}

// DeleteRabbitMQVipInstance
// This API is used to delete a TDMQ for RabbitMQ exclusive instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQVipInstanceWithContext(ctx context.Context, request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQVipInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRabbitMQVipInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQVipInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQVirtualHostRequest() (request *DeleteRabbitMQVirtualHostRequest) {
    request = &DeleteRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVirtualHost")
    
    
    return
}

func NewDeleteRabbitMQVirtualHostResponse() (response *DeleteRabbitMQVirtualHostResponse) {
    response = &DeleteRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQVirtualHost
// This API is used to delete a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQVirtualHost(request *DeleteRabbitMQVirtualHostRequest) (response *DeleteRabbitMQVirtualHostResponse, err error) {
    return c.DeleteRabbitMQVirtualHostWithContext(context.Background(), request)
}

// DeleteRabbitMQVirtualHost
// This API is used to delete a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRabbitMQVirtualHostWithContext(ctx context.Context, request *DeleteRabbitMQVirtualHostRequest) (response *DeleteRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQVirtualHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRabbitMQVirtualHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQClusterRequest() (request *DeleteRocketMQClusterRequest) {
    request = &DeleteRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQCluster")
    
    
    return
}

func NewDeleteRocketMQClusterResponse() (response *DeleteRocketMQClusterResponse) {
    response = &DeleteRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQCluster
// This API is used to delete a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteRocketMQCluster(request *DeleteRocketMQClusterRequest) (response *DeleteRocketMQClusterResponse, err error) {
    return c.DeleteRocketMQClusterWithContext(context.Background(), request)
}

// DeleteRocketMQCluster
// This API is used to delete a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteRocketMQClusterWithContext(ctx context.Context, request *DeleteRocketMQClusterRequest) (response *DeleteRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQGroupRequest() (request *DeleteRocketMQGroupRequest) {
    request = &DeleteRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQGroup")
    
    
    return
}

func NewDeleteRocketMQGroupResponse() (response *DeleteRocketMQGroupResponse) {
    response = &DeleteRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQGroup
// This API is used to delete a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQGroup(request *DeleteRocketMQGroupRequest) (response *DeleteRocketMQGroupResponse, err error) {
    return c.DeleteRocketMQGroupWithContext(context.Background(), request)
}

// DeleteRocketMQGroup
// This API is used to delete a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQGroupWithContext(ctx context.Context, request *DeleteRocketMQGroupRequest) (response *DeleteRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQNamespaceRequest() (request *DeleteRocketMQNamespaceRequest) {
    request = &DeleteRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQNamespace")
    
    
    return
}

func NewDeleteRocketMQNamespaceResponse() (response *DeleteRocketMQNamespaceResponse) {
    response = &DeleteRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQNamespace
// This API is used to delete a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQNamespace(request *DeleteRocketMQNamespaceRequest) (response *DeleteRocketMQNamespaceResponse, err error) {
    return c.DeleteRocketMQNamespaceWithContext(context.Background(), request)
}

// DeleteRocketMQNamespace
// This API is used to delete a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQNamespaceWithContext(ctx context.Context, request *DeleteRocketMQNamespaceRequest) (response *DeleteRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRocketMQTopicRequest() (request *DeleteRocketMQTopicRequest) {
    request = &DeleteRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQTopic")
    
    
    return
}

func NewDeleteRocketMQTopicResponse() (response *DeleteRocketMQTopicResponse) {
    response = &DeleteRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQTopic
// This API is used to delete a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQTopic(request *DeleteRocketMQTopicRequest) (response *DeleteRocketMQTopicResponse, err error) {
    return c.DeleteRocketMQTopicWithContext(context.Background(), request)
}

// DeleteRocketMQTopic
// This API is used to delete a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRocketMQTopicWithContext(ctx context.Context, request *DeleteRocketMQTopicRequest) (response *DeleteRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRolesRequest() (request *DeleteRolesRequest) {
    request = &DeleteRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRoles")
    
    
    return
}

func NewDeleteRolesResponse() (response *DeleteRolesResponse) {
    response = &DeleteRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoles
// This API is used to delete one or multiple roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEROLES = "FailedOperation.DeleteRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRoles(request *DeleteRolesRequest) (response *DeleteRolesResponse, err error) {
    return c.DeleteRolesWithContext(context.Background(), request)
}

// DeleteRoles
// This API is used to delete one or multiple roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEROLES = "FailedOperation.DeleteRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRolesWithContext(ctx context.Context, request *DeleteRolesRequest) (response *DeleteRolesResponse, err error) {
    if request == nil {
        request = NewDeleteRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubscriptionsRequest() (request *DeleteSubscriptionsRequest) {
    request = &DeleteSubscriptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteSubscriptions")
    
    
    return
}

func NewDeleteSubscriptionsResponse() (response *DeleteSubscriptionsResponse) {
    response = &DeleteSubscriptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSubscriptions
// This API is used to delete a subscription.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_CONSUMERRUNNING = "OperationDenied.ConsumerRunning"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteSubscriptions(request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
    return c.DeleteSubscriptionsWithContext(context.Background(), request)
}

// DeleteSubscriptions
// This API is used to delete a subscription.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_CONSUMERRUNNING = "OperationDenied.ConsumerRunning"
//  RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteSubscriptionsWithContext(ctx context.Context, request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDeleteSubscriptionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteSubscriptions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSubscriptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSubscriptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicsRequest() (request *DeleteTopicsRequest) {
    request = &DeleteTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteTopics")
    
    
    return
}

func NewDeleteTopicsResponse() (response *DeleteTopicsResponse) {
    response = &DeleteTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopics
// This API is used to batch delete topics.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopics(request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    return c.DeleteTopicsWithContext(context.Background(), request)
}

// DeleteTopics
// This API is used to batch delete topics.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopicsWithContext(ctx context.Context, request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
    if request == nil {
        request = NewDeleteTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindClustersRequest() (request *DescribeBindClustersRequest) {
    request = &DescribeBindClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindClusters")
    
    
    return
}

func NewDescribeBindClustersResponse() (response *DescribeBindClustersResponse) {
    response = &DescribeBindClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBindClusters
// This API is used to get the list of dedicated clusters bound to a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SYSTEMUPGRADE = "ResourceUnavailable.SystemUpgrade"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindClusters(request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
    return c.DescribeBindClustersWithContext(context.Background(), request)
}

// DescribeBindClusters
// This API is used to get the list of dedicated clusters bound to a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SYSTEMUPGRADE = "ResourceUnavailable.SystemUpgrade"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindClustersWithContext(ctx context.Context, request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
    if request == nil {
        request = NewDescribeBindClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeBindClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindVpcsRequest() (request *DescribeBindVpcsRequest) {
    request = &DescribeBindVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindVpcs")
    
    
    return
}

func NewDescribeBindVpcsResponse() (response *DescribeBindVpcsResponse) {
    response = &DescribeBindVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBindVpcs
// This API is used to get the tenant-VPC binding relationship.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeBindVpcs(request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
    return c.DescribeBindVpcsWithContext(context.Background(), request)
}

// DescribeBindVpcs
// This API is used to get the tenant-VPC binding relationship.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeBindVpcsWithContext(ctx context.Context, request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeBindVpcsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeBindVpcs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindVpcs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterDetail")
    
    
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetail
// This API is used to get the details of a cluster.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// This API is used to get the details of a cluster.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// This API is used to get the list of clusters.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// This API is used to get the list of clusters.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqDeadLetterSourceQueuesRequest() (request *DescribeCmqDeadLetterSourceQueuesRequest) {
    request = &DescribeCmqDeadLetterSourceQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqDeadLetterSourceQueues")
    
    
    return
}

func NewDescribeCmqDeadLetterSourceQueuesResponse() (response *DescribeCmqDeadLetterSourceQueuesResponse) {
    response = &DescribeCmqDeadLetterSourceQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqDeadLetterSourceQueues
// This API is used to enumerate the source queues of a CMQ dead letter queue.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeCmqDeadLetterSourceQueues(request *DescribeCmqDeadLetterSourceQueuesRequest) (response *DescribeCmqDeadLetterSourceQueuesResponse, err error) {
    return c.DescribeCmqDeadLetterSourceQueuesWithContext(context.Background(), request)
}

// DescribeCmqDeadLetterSourceQueues
// This API is used to enumerate the source queues of a CMQ dead letter queue.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeCmqDeadLetterSourceQueuesWithContext(ctx context.Context, request *DescribeCmqDeadLetterSourceQueuesRequest) (response *DescribeCmqDeadLetterSourceQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeCmqDeadLetterSourceQueuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqDeadLetterSourceQueues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqDeadLetterSourceQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqDeadLetterSourceQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqQueueDetailRequest() (request *DescribeCmqQueueDetailRequest) {
    request = &DescribeCmqQueueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueueDetail")
    
    
    return
}

func NewDescribeCmqQueueDetailResponse() (response *DescribeCmqQueueDetailResponse) {
    response = &DescribeCmqQueueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqQueueDetail
// This API is used to query the details of a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmqQueueDetail(request *DescribeCmqQueueDetailRequest) (response *DescribeCmqQueueDetailResponse, err error) {
    return c.DescribeCmqQueueDetailWithContext(context.Background(), request)
}

// DescribeCmqQueueDetail
// This API is used to query the details of a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCmqQueueDetailWithContext(ctx context.Context, request *DescribeCmqQueueDetailRequest) (response *DescribeCmqQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqQueueDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqQueueDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqQueueDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqQueueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqQueuesRequest() (request *DescribeCmqQueuesRequest) {
    request = &DescribeCmqQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueues")
    
    
    return
}

func NewDescribeCmqQueuesResponse() (response *DescribeCmqQueuesResponse) {
    response = &DescribeCmqQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqQueues
// This API is used to query all CMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqQueues(request *DescribeCmqQueuesRequest) (response *DescribeCmqQueuesResponse, err error) {
    return c.DescribeCmqQueuesWithContext(context.Background(), request)
}

// DescribeCmqQueues
// This API is used to query all CMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqQueuesWithContext(ctx context.Context, request *DescribeCmqQueuesRequest) (response *DescribeCmqQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeCmqQueuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqQueues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqSubscriptionDetailRequest() (request *DescribeCmqSubscriptionDetailRequest) {
    request = &DescribeCmqSubscriptionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqSubscriptionDetail")
    
    
    return
}

func NewDescribeCmqSubscriptionDetailResponse() (response *DescribeCmqSubscriptionDetailResponse) {
    response = &DescribeCmqSubscriptionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqSubscriptionDetail
// This API is used to query the CMQ subscription details.
//
// error code that may be returned:
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeCmqSubscriptionDetail(request *DescribeCmqSubscriptionDetailRequest) (response *DescribeCmqSubscriptionDetailResponse, err error) {
    return c.DescribeCmqSubscriptionDetailWithContext(context.Background(), request)
}

// DescribeCmqSubscriptionDetail
// This API is used to query the CMQ subscription details.
//
// error code that may be returned:
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeCmqSubscriptionDetailWithContext(ctx context.Context, request *DescribeCmqSubscriptionDetailRequest) (response *DescribeCmqSubscriptionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqSubscriptionDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqSubscriptionDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqSubscriptionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqSubscriptionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqTopicDetailRequest() (request *DescribeCmqTopicDetailRequest) {
    request = &DescribeCmqTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopicDetail")
    
    
    return
}

func NewDescribeCmqTopicDetailResponse() (response *DescribeCmqTopicDetailResponse) {
    response = &DescribeCmqTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqTopicDetail
// This API is used to query the details of a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqTopicDetail(request *DescribeCmqTopicDetailRequest) (response *DescribeCmqTopicDetailResponse, err error) {
    return c.DescribeCmqTopicDetailWithContext(context.Background(), request)
}

// DescribeCmqTopicDetail
// This API is used to query the details of a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCmqTopicDetailWithContext(ctx context.Context, request *DescribeCmqTopicDetailRequest) (response *DescribeCmqTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCmqTopicDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqTopicDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqTopicDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCmqTopicsRequest() (request *DescribeCmqTopicsRequest) {
    request = &DescribeCmqTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopics")
    
    
    return
}

func NewDescribeCmqTopicsResponse() (response *DescribeCmqTopicsResponse) {
    response = &DescribeCmqTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCmqTopics
// This API is used to enumerate all CMQ topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
func (c *Client) DescribeCmqTopics(request *DescribeCmqTopicsRequest) (response *DescribeCmqTopicsResponse, err error) {
    return c.DescribeCmqTopicsWithContext(context.Background(), request)
}

// DescribeCmqTopics
// This API is used to enumerate all CMQ topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"
func (c *Client) DescribeCmqTopicsWithContext(ctx context.Context, request *DescribeCmqTopicsRequest) (response *DescribeCmqTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeCmqTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeCmqTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCmqTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCmqTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentAttributesRequest() (request *DescribeEnvironmentAttributesRequest) {
    request = &DescribeEnvironmentAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentAttributes")
    
    
    return
}

func NewDescribeEnvironmentAttributesResponse() (response *DescribeEnvironmentAttributesResponse) {
    response = &DescribeEnvironmentAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironmentAttributes
// This API is used to get the attributes of the specified namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeEnvironmentAttributes(request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
    return c.DescribeEnvironmentAttributesWithContext(context.Background(), request)
}

// DescribeEnvironmentAttributes
// This API is used to get the attributes of the specified namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeEnvironmentAttributesWithContext(ctx context.Context, request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeEnvironmentAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironmentAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentRolesRequest() (request *DescribeEnvironmentRolesRequest) {
    request = &DescribeEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentRoles")
    
    
    return
}

func NewDescribeEnvironmentRolesResponse() (response *DescribeEnvironmentRolesResponse) {
    response = &DescribeEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironmentRoles
// This API is used to get the list of namespace roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeEnvironmentRoles(request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
    return c.DescribeEnvironmentRolesWithContext(context.Background(), request)
}

// DescribeEnvironmentRoles
// This API is used to get the list of namespace roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeEnvironmentRolesWithContext(ctx context.Context, request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeEnvironmentRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// This API is used to get the list of namespaces under a tenant.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// This API is used to get the list of namespaces under a tenant.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublisherSummaryRequest() (request *DescribePublisherSummaryRequest) {
    request = &DescribePublisherSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePublisherSummary")
    
    
    return
}

func NewDescribePublisherSummaryResponse() (response *DescribePublisherSummaryResponse) {
    response = &DescribePublisherSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublisherSummary
// This API is used to obtain message production overview information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublisherSummary(request *DescribePublisherSummaryRequest) (response *DescribePublisherSummaryResponse, err error) {
    return c.DescribePublisherSummaryWithContext(context.Background(), request)
}

// DescribePublisherSummary
// This API is used to obtain message production overview information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublisherSummaryWithContext(ctx context.Context, request *DescribePublisherSummaryRequest) (response *DescribePublisherSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublisherSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribePublisherSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublisherSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublisherSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublishersRequest() (request *DescribePublishersRequest) {
    request = &DescribePublishersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePublishers")
    
    
    return
}

func NewDescribePublishersResponse() (response *DescribePublishersResponse) {
    response = &DescribePublishersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePublishers
// This API is used to obtain the list of producer information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublishers(request *DescribePublishersRequest) (response *DescribePublishersResponse, err error) {
    return c.DescribePublishersWithContext(context.Background(), request)
}

// DescribePublishers
// This API is used to obtain the list of producer information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublishersWithContext(ctx context.Context, request *DescribePublishersRequest) (response *DescribePublishersResponse, err error) {
    if request == nil {
        request = NewDescribePublishersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribePublishers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublishers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublishersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePulsarProInstanceDetailRequest() (request *DescribePulsarProInstanceDetailRequest) {
    request = &DescribePulsarProInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePulsarProInstanceDetail")
    
    
    return
}

func NewDescribePulsarProInstanceDetailResponse() (response *DescribePulsarProInstanceDetailResponse) {
    response = &DescribePulsarProInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePulsarProInstanceDetail
// This API is used to obtain the information of a TDMQ for Pulsar pro cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstanceDetail(request *DescribePulsarProInstanceDetailRequest) (response *DescribePulsarProInstanceDetailResponse, err error) {
    return c.DescribePulsarProInstanceDetailWithContext(context.Background(), request)
}

// DescribePulsarProInstanceDetail
// This API is used to obtain the information of a TDMQ for Pulsar pro cluster instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstanceDetailWithContext(ctx context.Context, request *DescribePulsarProInstanceDetailRequest) (response *DescribePulsarProInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribePulsarProInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribePulsarProInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePulsarProInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePulsarProInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePulsarProInstancesRequest() (request *DescribePulsarProInstancesRequest) {
    request = &DescribePulsarProInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribePulsarProInstances")
    
    
    return
}

func NewDescribePulsarProInstancesResponse() (response *DescribePulsarProInstancesResponse) {
    response = &DescribePulsarProInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePulsarProInstances
// This API is used to query the list of the purchased TDMQ for Pulsar pro instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstances(request *DescribePulsarProInstancesRequest) (response *DescribePulsarProInstancesResponse, err error) {
    return c.DescribePulsarProInstancesWithContext(context.Background(), request)
}

// DescribePulsarProInstances
// This API is used to query the list of the purchased TDMQ for Pulsar pro instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribePulsarProInstancesWithContext(ctx context.Context, request *DescribePulsarProInstancesRequest) (response *DescribePulsarProInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePulsarProInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribePulsarProInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePulsarProInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePulsarProInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQNodeListRequest() (request *DescribeRabbitMQNodeListRequest) {
    request = &DescribeRabbitMQNodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQNodeList")
    
    
    return
}

func NewDescribeRabbitMQNodeListResponse() (response *DescribeRabbitMQNodeListResponse) {
    response = &DescribeRabbitMQNodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQNodeList
// This API is used to query the list of TDMQ for RabbitMQ exclusive cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQNodeList(request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
    return c.DescribeRabbitMQNodeListWithContext(context.Background(), request)
}

// DescribeRabbitMQNodeList
// This API is used to query the list of TDMQ for RabbitMQ exclusive cluster nodes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQNodeListWithContext(ctx context.Context, request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQNodeListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQNodeList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQNodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQNodeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQUserRequest() (request *DescribeRabbitMQUserRequest) {
    request = &DescribeRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQUser")
    
    
    return
}

func NewDescribeRabbitMQUserResponse() (response *DescribeRabbitMQUserResponse) {
    response = &DescribeRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQUser
// This API is used to query the list of TDMQ for RabbitMQ users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQUser(request *DescribeRabbitMQUserRequest) (response *DescribeRabbitMQUserResponse, err error) {
    return c.DescribeRabbitMQUserWithContext(context.Background(), request)
}

// DescribeRabbitMQUser
// This API is used to query the list of TDMQ for RabbitMQ users.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQUserWithContext(ctx context.Context, request *DescribeRabbitMQUserRequest) (response *DescribeRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVipInstancesRequest() (request *DescribeRabbitMQVipInstancesRequest) {
    request = &DescribeRabbitMQVipInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstances")
    
    
    return
}

func NewDescribeRabbitMQVipInstancesResponse() (response *DescribeRabbitMQVipInstancesResponse) {
    response = &DescribeRabbitMQVipInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQVipInstances
// This API is used to query the list of the purchased TDMQ for RabbitMQ exclusive instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVipInstances(request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
    return c.DescribeRabbitMQVipInstancesWithContext(context.Background(), request)
}

// DescribeRabbitMQVipInstances
// This API is used to query the list of the purchased TDMQ for RabbitMQ exclusive instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVipInstancesWithContext(ctx context.Context, request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVipInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQVipInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVipInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVipInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVirtualHostRequest() (request *DescribeRabbitMQVirtualHostRequest) {
    request = &DescribeRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVirtualHost")
    
    
    return
}

func NewDescribeRabbitMQVirtualHostResponse() (response *DescribeRabbitMQVirtualHostResponse) {
    response = &DescribeRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQVirtualHost
// This API is used to query the list of TDMQ for RabbitMQ vhosts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVirtualHost(request *DescribeRabbitMQVirtualHostRequest) (response *DescribeRabbitMQVirtualHostResponse, err error) {
    return c.DescribeRabbitMQVirtualHostWithContext(context.Background(), request)
}

// DescribeRabbitMQVirtualHost
// This API is used to query the list of TDMQ for RabbitMQ vhosts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQVirtualHostWithContext(ctx context.Context, request *DescribeRabbitMQVirtualHostRequest) (response *DescribeRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVirtualHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQVirtualHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQVirtualHostListRequest() (request *DescribeRabbitMQVirtualHostListRequest) {
    request = &DescribeRabbitMQVirtualHostListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVirtualHostList")
    
    
    return
}

func NewDescribeRabbitMQVirtualHostListResponse() (response *DescribeRabbitMQVirtualHostListResponse) {
    response = &DescribeRabbitMQVirtualHostListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQVirtualHostList
// This API is used to query the list of TDMQ for RabbitMQ exclusive vhosts.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVirtualHostList(request *DescribeRabbitMQVirtualHostListRequest) (response *DescribeRabbitMQVirtualHostListResponse, err error) {
    return c.DescribeRabbitMQVirtualHostListWithContext(context.Background(), request)
}

// DescribeRabbitMQVirtualHostList
// This API is used to query the list of TDMQ for RabbitMQ exclusive vhosts.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVirtualHostListWithContext(ctx context.Context, request *DescribeRabbitMQVirtualHostListRequest) (response *DescribeRabbitMQVirtualHostListResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVirtualHostListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQVirtualHostList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVirtualHostList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVirtualHostListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQClusterRequest() (request *DescribeRocketMQClusterRequest) {
    request = &DescribeRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQCluster")
    
    
    return
}

func NewDescribeRocketMQClusterResponse() (response *DescribeRocketMQClusterResponse) {
    response = &DescribeRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQCluster
// This API is used to get the information of a specific RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQCluster(request *DescribeRocketMQClusterRequest) (response *DescribeRocketMQClusterResponse, err error) {
    return c.DescribeRocketMQClusterWithContext(context.Background(), request)
}

// DescribeRocketMQCluster
// This API is used to get the information of a specific RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQClusterWithContext(ctx context.Context, request *DescribeRocketMQClusterRequest) (response *DescribeRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQClustersRequest() (request *DescribeRocketMQClustersRequest) {
    request = &DescribeRocketMQClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQClusters")
    
    
    return
}

func NewDescribeRocketMQClustersResponse() (response *DescribeRocketMQClustersResponse) {
    response = &DescribeRocketMQClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQClusters
// This API is used to get the list of RocketMQ clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRocketMQClusters(request *DescribeRocketMQClustersRequest) (response *DescribeRocketMQClustersResponse, err error) {
    return c.DescribeRocketMQClustersWithContext(context.Background(), request)
}

// DescribeRocketMQClusters
// This API is used to get the list of RocketMQ clusters.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRocketMQClustersWithContext(ctx context.Context, request *DescribeRocketMQClustersRequest) (response *DescribeRocketMQClustersResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQGroupsRequest() (request *DescribeRocketMQGroupsRequest) {
    request = &DescribeRocketMQGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQGroups")
    
    
    return
}

func NewDescribeRocketMQGroupsResponse() (response *DescribeRocketMQGroupsResponse) {
    response = &DescribeRocketMQGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQGroups
// This API is used to get the list of RocketMQ consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQGroups(request *DescribeRocketMQGroupsRequest) (response *DescribeRocketMQGroupsResponse, err error) {
    return c.DescribeRocketMQGroupsWithContext(context.Background(), request)
}

// DescribeRocketMQGroups
// This API is used to get the list of RocketMQ consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQGroupsWithContext(ctx context.Context, request *DescribeRocketMQGroupsRequest) (response *DescribeRocketMQGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQMsgRequest() (request *DescribeRocketMQMsgRequest) {
    request = &DescribeRocketMQMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMsg")
    
    
    return
}

func NewDescribeRocketMQMsgResponse() (response *DescribeRocketMQMsgResponse) {
    response = &DescribeRocketMQMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQMsg
// This API is used to query the TDMQ for RocketMQ message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsg(request *DescribeRocketMQMsgRequest) (response *DescribeRocketMQMsgResponse, err error) {
    return c.DescribeRocketMQMsgWithContext(context.Background(), request)
}

// DescribeRocketMQMsg
// This API is used to query the TDMQ for RocketMQ message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsgWithContext(ctx context.Context, request *DescribeRocketMQMsgRequest) (response *DescribeRocketMQMsgResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQNamespacesRequest() (request *DescribeRocketMQNamespacesRequest) {
    request = &DescribeRocketMQNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQNamespaces")
    
    
    return
}

func NewDescribeRocketMQNamespacesResponse() (response *DescribeRocketMQNamespacesResponse) {
    response = &DescribeRocketMQNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQNamespaces
// This API is used to get the list of RocketMQ namespaces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQNamespaces(request *DescribeRocketMQNamespacesRequest) (response *DescribeRocketMQNamespacesResponse, err error) {
    return c.DescribeRocketMQNamespacesWithContext(context.Background(), request)
}

// DescribeRocketMQNamespaces
// This API is used to get the list of RocketMQ namespaces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQNamespacesWithContext(ctx context.Context, request *DescribeRocketMQNamespacesRequest) (response *DescribeRocketMQNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopicsRequest() (request *DescribeRocketMQTopicsRequest) {
    request = &DescribeRocketMQTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopics")
    
    
    return
}

func NewDescribeRocketMQTopicsResponse() (response *DescribeRocketMQTopicsResponse) {
    response = &DescribeRocketMQTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopics
// This API is used to get the list of RocketMQ topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQTopics(request *DescribeRocketMQTopicsRequest) (response *DescribeRocketMQTopicsResponse, err error) {
    return c.DescribeRocketMQTopicsWithContext(context.Background(), request)
}

// DescribeRocketMQTopics
// This API is used to get the list of RocketMQ topics.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQTopicsWithContext(ctx context.Context, request *DescribeRocketMQTopicsRequest) (response *DescribeRocketMQTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQVipInstanceDetailRequest() (request *DescribeRocketMQVipInstanceDetailRequest) {
    request = &DescribeRocketMQVipInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstanceDetail")
    
    
    return
}

func NewDescribeRocketMQVipInstanceDetailResponse() (response *DescribeRocketMQVipInstanceDetailResponse) {
    response = &DescribeRocketMQVipInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQVipInstanceDetail
// This API is used to get the information of a specific TDMQ for RocketMQ exclusive cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQVipInstanceDetail(request *DescribeRocketMQVipInstanceDetailRequest) (response *DescribeRocketMQVipInstanceDetailResponse, err error) {
    return c.DescribeRocketMQVipInstanceDetailWithContext(context.Background(), request)
}

// DescribeRocketMQVipInstanceDetail
// This API is used to get the information of a specific TDMQ for RocketMQ exclusive cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQVipInstanceDetailWithContext(ctx context.Context, request *DescribeRocketMQVipInstanceDetailRequest) (response *DescribeRocketMQVipInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQVipInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQVipInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQVipInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQVipInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQVipInstancesRequest() (request *DescribeRocketMQVipInstancesRequest) {
    request = &DescribeRocketMQVipInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstances")
    
    
    return
}

func NewDescribeRocketMQVipInstancesResponse() (response *DescribeRocketMQVipInstancesResponse) {
    response = &DescribeRocketMQVipInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQVipInstances
// This API is used to query the list of the purchased TDMQ for RocketMQ exclusive instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQVipInstances(request *DescribeRocketMQVipInstancesRequest) (response *DescribeRocketMQVipInstancesResponse, err error) {
    return c.DescribeRocketMQVipInstancesWithContext(context.Background(), request)
}

// DescribeRocketMQVipInstances
// This API is used to query the list of the purchased TDMQ for RocketMQ exclusive instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRocketMQVipInstancesWithContext(ctx context.Context, request *DescribeRocketMQVipInstancesRequest) (response *DescribeRocketMQVipInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQVipInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQVipInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQVipInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQVipInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
    request = &DescribeRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRoles")
    
    
    return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
    response = &DescribeRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoles
// This API is used to get the list of roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
    return c.DescribeRolesWithContext(context.Background(), request)
}

// DescribeRoles
// This API is used to get the list of roles.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRolesWithContext(ctx context.Context, request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
    if request == nil {
        request = NewDescribeRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscriptionsRequest() (request *DescribeSubscriptionsRequest) {
    request = &DescribeSubscriptionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeSubscriptions")
    
    
    return
}

func NewDescribeSubscriptionsResponse() (response *DescribeSubscriptionsResponse) {
    response = &DescribeSubscriptionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscriptions
// This API is used to query the list of subscribers under the specified environment and topic.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBESUBSCRIPTION = "FailedOperation.DescribeSubscription"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
    return c.DescribeSubscriptionsWithContext(context.Background(), request)
}

// DescribeSubscriptions
// This API is used to query the list of subscribers under the specified environment and topic.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBESUBSCRIPTION = "FailedOperation.DescribeSubscription"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeSubscriptionsWithContext(ctx context.Context, request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
    if request == nil {
        request = NewDescribeSubscriptionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeSubscriptions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscriptions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscriptionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
    request = &DescribeTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopics")
    
    
    return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
    response = &DescribeTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopics
// This API is used to get the list of topics under an environment.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDADMINURL = "InvalidParameter.InvalidAdminUrl"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    return c.DescribeTopicsWithContext(context.Background(), request)
}

// DescribeTopics
// This API is used to get the list of topics under an environment.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDADMINURL = "InvalidParameter.InvalidAdminUrl"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
func (c *Client) DescribeTopicsWithContext(ctx context.Context, request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
    request = &ModifyClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCluster")
    
    
    return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
    response = &ModifyClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCluster
// This API is used to update a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEDUPLICATION = "InvalidParameterValue.ClusterNameDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    return c.ModifyClusterWithContext(context.Background(), request)
}

// ModifyCluster
// This API is used to update a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEDUPLICATION = "InvalidParameterValue.ClusterNameDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyClusterWithContext(ctx context.Context, request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
    if request == nil {
        request = NewModifyClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqQueueAttributeRequest() (request *ModifyCmqQueueAttributeRequest) {
    request = &ModifyCmqQueueAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqQueueAttribute")
    
    
    return
}

func NewModifyCmqQueueAttributeResponse() (response *ModifyCmqQueueAttributeResponse) {
    response = &ModifyCmqQueueAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCmqQueueAttribute
// This API is used to modify the attributes of a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCmqQueueAttribute(request *ModifyCmqQueueAttributeRequest) (response *ModifyCmqQueueAttributeResponse, err error) {
    return c.ModifyCmqQueueAttributeWithContext(context.Background(), request)
}

// ModifyCmqQueueAttribute
// This API is used to modify the attributes of a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyCmqQueueAttributeWithContext(ctx context.Context, request *ModifyCmqQueueAttributeRequest) (response *ModifyCmqQueueAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqQueueAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyCmqQueueAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqQueueAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqQueueAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqSubscriptionAttributeRequest() (request *ModifyCmqSubscriptionAttributeRequest) {
    request = &ModifyCmqSubscriptionAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqSubscriptionAttribute")
    
    
    return
}

func NewModifyCmqSubscriptionAttributeResponse() (response *ModifyCmqSubscriptionAttributeResponse) {
    response = &ModifyCmqSubscriptionAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCmqSubscriptionAttribute
// This API is used to modify the attributes of a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyCmqSubscriptionAttribute(request *ModifyCmqSubscriptionAttributeRequest) (response *ModifyCmqSubscriptionAttributeResponse, err error) {
    return c.ModifyCmqSubscriptionAttributeWithContext(context.Background(), request)
}

// ModifyCmqSubscriptionAttribute
// This API is used to modify the attributes of a CMQ subscription.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyCmqSubscriptionAttributeWithContext(ctx context.Context, request *ModifyCmqSubscriptionAttributeRequest) (response *ModifyCmqSubscriptionAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqSubscriptionAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyCmqSubscriptionAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqSubscriptionAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqSubscriptionAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCmqTopicAttributeRequest() (request *ModifyCmqTopicAttributeRequest) {
    request = &ModifyCmqTopicAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqTopicAttribute")
    
    
    return
}

func NewModifyCmqTopicAttributeResponse() (response *ModifyCmqTopicAttributeResponse) {
    response = &ModifyCmqTopicAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCmqTopicAttribute
// This API is used to modify the attributes of a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyCmqTopicAttribute(request *ModifyCmqTopicAttributeRequest) (response *ModifyCmqTopicAttributeResponse, err error) {
    return c.ModifyCmqTopicAttributeWithContext(context.Background(), request)
}

// ModifyCmqTopicAttribute
// This API is used to modify the attributes of a CMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyCmqTopicAttributeWithContext(ctx context.Context, request *ModifyCmqTopicAttributeRequest) (response *ModifyCmqTopicAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCmqTopicAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyCmqTopicAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCmqTopicAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCmqTopicAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentAttributesRequest() (request *ModifyEnvironmentAttributesRequest) {
    request = &ModifyEnvironmentAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentAttributes")
    
    
    return
}

func NewModifyEnvironmentAttributesResponse() (response *ModifyEnvironmentAttributesResponse) {
    response = &ModifyEnvironmentAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnvironmentAttributes
// This API is used to modify the attributes of the specified namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentAttributes(request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
    return c.ModifyEnvironmentAttributesWithContext(context.Background(), request)
}

// ModifyEnvironmentAttributes
// This API is used to modify the attributes of the specified namespace.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"
//  INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"
//  LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentAttributesWithContext(ctx context.Context, request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyEnvironmentAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironmentAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvironmentAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentRoleRequest() (request *ModifyEnvironmentRoleRequest) {
    request = &ModifyEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentRole")
    
    
    return
}

func NewModifyEnvironmentRoleResponse() (response *ModifyEnvironmentRoleResponse) {
    response = &ModifyEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyEnvironmentRole
// This API is used to modify an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentRole(request *ModifyEnvironmentRoleRequest) (response *ModifyEnvironmentRoleResponse, err error) {
    return c.ModifyEnvironmentRoleWithContext(context.Background(), request)
}

// ModifyEnvironmentRole
// This API is used to modify an environment role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyEnvironmentRoleWithContext(ctx context.Context, request *ModifyEnvironmentRoleRequest) (response *ModifyEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyEnvironmentRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvironmentRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQUserRequest() (request *ModifyRabbitMQUserRequest) {
    request = &ModifyRabbitMQUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQUser")
    
    
    return
}

func NewModifyRabbitMQUserResponse() (response *ModifyRabbitMQUserResponse) {
    response = &ModifyRabbitMQUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQUser
// This API is used to modify a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQUser(request *ModifyRabbitMQUserRequest) (response *ModifyRabbitMQUserResponse, err error) {
    return c.ModifyRabbitMQUserWithContext(context.Background(), request)
}

// ModifyRabbitMQUser
// This API is used to modify a TDMQ for RabbitMQ user.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQUserWithContext(ctx context.Context, request *ModifyRabbitMQUserRequest) (response *ModifyRabbitMQUserResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRabbitMQUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQVirtualHostRequest() (request *ModifyRabbitMQVirtualHostRequest) {
    request = &ModifyRabbitMQVirtualHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQVirtualHost")
    
    
    return
}

func NewModifyRabbitMQVirtualHostResponse() (response *ModifyRabbitMQVirtualHostResponse) {
    response = &ModifyRabbitMQVirtualHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQVirtualHost
// This API is used to modify a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQVirtualHost(request *ModifyRabbitMQVirtualHostRequest) (response *ModifyRabbitMQVirtualHostResponse, err error) {
    return c.ModifyRabbitMQVirtualHostWithContext(context.Background(), request)
}

// ModifyRabbitMQVirtualHost
// This API is used to modify a TDMQ for RabbitMQ vhost.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQVirtualHostWithContext(ctx context.Context, request *ModifyRabbitMQVirtualHostRequest) (response *ModifyRabbitMQVirtualHostResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQVirtualHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRabbitMQVirtualHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQVirtualHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQVirtualHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQClusterRequest() (request *ModifyRocketMQClusterRequest) {
    request = &ModifyRocketMQClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQCluster")
    
    
    return
}

func NewModifyRocketMQClusterResponse() (response *ModifyRocketMQClusterResponse) {
    response = &ModifyRocketMQClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQCluster
// This API is used to update a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQCluster(request *ModifyRocketMQClusterRequest) (response *ModifyRocketMQClusterResponse, err error) {
    return c.ModifyRocketMQClusterWithContext(context.Background(), request)
}

// ModifyRocketMQCluster
// This API is used to update a RocketMQ cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQClusterWithContext(ctx context.Context, request *ModifyRocketMQClusterRequest) (response *ModifyRocketMQClusterResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQClusterResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQGroupRequest() (request *ModifyRocketMQGroupRequest) {
    request = &ModifyRocketMQGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQGroup")
    
    
    return
}

func NewModifyRocketMQGroupResponse() (response *ModifyRocketMQGroupResponse) {
    response = &ModifyRocketMQGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQGroup
// This API is used to update a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQGroup(request *ModifyRocketMQGroupRequest) (response *ModifyRocketMQGroupResponse, err error) {
    return c.ModifyRocketMQGroupWithContext(context.Background(), request)
}

// ModifyRocketMQGroup
// This API is used to update a RocketMQ consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQGroupWithContext(ctx context.Context, request *ModifyRocketMQGroupRequest) (response *ModifyRocketMQGroupResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQInstanceSpecRequest() (request *ModifyRocketMQInstanceSpecRequest) {
    request = &ModifyRocketMQInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQInstanceSpec")
    
    
    return
}

func NewModifyRocketMQInstanceSpecResponse() (response *ModifyRocketMQInstanceSpecResponse) {
    response = &ModifyRocketMQInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQInstanceSpec
// This API is used to modify the configurations of a TDMQ for RocketMQ exclusive instance, including the upgrade of the instance specification, node count, and storage, and the downgrade of the instance specification. After you call this API to place the order and make payments, the configuration modification will be in progress. You can query whether the modification has been completed through the `DescribeRocketMQVipInstances` API`.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  INVALIDPARAMETERVALUE_ATLEASTONE = "InvalidParameterValue.AtLeastOne"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INSTANCEDOWNGRADE = "UnsupportedOperation.InstanceDowngrade"
func (c *Client) ModifyRocketMQInstanceSpec(request *ModifyRocketMQInstanceSpecRequest) (response *ModifyRocketMQInstanceSpecResponse, err error) {
    return c.ModifyRocketMQInstanceSpecWithContext(context.Background(), request)
}

// ModifyRocketMQInstanceSpec
// This API is used to modify the configurations of a TDMQ for RocketMQ exclusive instance, including the upgrade of the instance specification, node count, and storage, and the downgrade of the instance specification. After you call this API to place the order and make payments, the configuration modification will be in progress. You can query whether the modification has been completed through the `DescribeRocketMQVipInstances` API`.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  INVALIDPARAMETERVALUE_ATLEASTONE = "InvalidParameterValue.AtLeastOne"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_INSTANCEDOWNGRADE = "UnsupportedOperation.InstanceDowngrade"
func (c *Client) ModifyRocketMQInstanceSpecWithContext(ctx context.Context, request *ModifyRocketMQInstanceSpecRequest) (response *ModifyRocketMQInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQInstanceSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQInstanceSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQNamespaceRequest() (request *ModifyRocketMQNamespaceRequest) {
    request = &ModifyRocketMQNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQNamespace")
    
    
    return
}

func NewModifyRocketMQNamespaceResponse() (response *ModifyRocketMQNamespaceResponse) {
    response = &ModifyRocketMQNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQNamespace
// This API is used to update a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQNamespace(request *ModifyRocketMQNamespaceRequest) (response *ModifyRocketMQNamespaceResponse, err error) {
    return c.ModifyRocketMQNamespaceWithContext(context.Background(), request)
}

// ModifyRocketMQNamespace
// This API is used to update a RocketMQ namespace.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRocketMQNamespaceWithContext(ctx context.Context, request *ModifyRocketMQNamespaceRequest) (response *ModifyRocketMQNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRocketMQTopicRequest() (request *ModifyRocketMQTopicRequest) {
    request = &ModifyRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQTopic")
    
    
    return
}

func NewModifyRocketMQTopicResponse() (response *ModifyRocketMQTopicResponse) {
    response = &ModifyRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQTopic
// This API is used to update a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQTopic(request *ModifyRocketMQTopicRequest) (response *ModifyRocketMQTopicResponse, err error) {
    return c.ModifyRocketMQTopicWithContext(context.Background(), request)
}

// ModifyRocketMQTopic
// This API is used to update a RocketMQ topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRocketMQTopicWithContext(ctx context.Context, request *ModifyRocketMQTopicRequest) (response *ModifyRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
    request = &ModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRole")
    
    
    return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
    response = &ModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRole
// This API is used to modify a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UPDATEROLE = "FailedOperation.UpdateRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    return c.ModifyRoleWithContext(context.Background(), request)
}

// ModifyRole
// This API is used to modify a role.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_UPDATEROLE = "FailedOperation.UpdateRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ModifyRoleWithContext(ctx context.Context, request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    if request == nil {
        request = NewModifyRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTopic")
    
    
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTopic
// This API is used to modify the topic remarks and number of partitions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// This API is used to modify the topic remarks and number of partitions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}

func NewPublishCmqMsgRequest() (request *PublishCmqMsgRequest) {
    request = &PublishCmqMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "PublishCmqMsg")
    
    
    return
}

func NewPublishCmqMsgResponse() (response *PublishCmqMsgResponse) {
    response = &PublishCmqMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PublishCmqMsg
// This API is used to send a CMQ topic message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PublishCmqMsg(request *PublishCmqMsgRequest) (response *PublishCmqMsgResponse, err error) {
    return c.PublishCmqMsgWithContext(context.Background(), request)
}

// PublishCmqMsg
// This API is used to send a CMQ topic message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PublishCmqMsgWithContext(ctx context.Context, request *PublishCmqMsgRequest) (response *PublishCmqMsgResponse, err error) {
    if request == nil {
        request = NewPublishCmqMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "PublishCmqMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishCmqMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishCmqMsgResponse()
    err = c.Send(request, response)
    return
}

func NewReceiveMessageRequest() (request *ReceiveMessageRequest) {
    request = &ReceiveMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ReceiveMessage")
    
    
    return
}

func NewReceiveMessageResponse() (response *ReceiveMessageResponse) {
    response = &ReceiveMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReceiveMessage
// Currently, the `ReceiveMessage` API only supports partitioned topics. It is used to receive messages sent to a specified partitioned topic. If it is called when there are no messages in the partitioned topic, the `ReceiveTimeout` exception will be reported.
//
// 
//
// Instructions on how to use `BatchReceivePolicy`:
//
// 
//
// `BatchReceive` has three parameters:
//
// 
//
// ● `MaxNumMessages`: The maximum number of messages returned by `Receive` when `BatchReceive` is used.
//
// ● `MaxNumBytes`: The maximum size (in bytes) of the message returned by `Receive` when `BatchReceive` is used.
//
// ● `Timeout`: The maximum timeout period (in milliseconds) of calling `Receive` when `BatchReceive` is used.
//
// 
//
// By default, if you don’t specify any of the three parameters, the `BatchReceive` feature is disabled; if one of the three parameter values is above zero, `BatchReceive` is enabled. `BatchReceive` will be disabled when any of the three parameter values reaches the threshold you specify.
//
// 
//
// Note: The values of both `MaxNumMessages` and `MaxNumBytes` are subject to the value of `ReceiveQueueSize`. If the values of `ReceiveQueueSize` and `MaxNumMessages` are 5 and 10, respectively, you can receive up to five rather than 10 messages when `BatchReceive` is used.
//
// 
//
// 
//
// 
//
// The API configured with `BatchReceivePolicy` returns multiple messages at a time.
//
// 
//
// 1. These messages are separated by “###”. After receiving them, you can separate them with split tools in different languages.
//
// 2. MessageIDs are separated by “###”. After receiving the messages, you can separate the MessageIDs with split tools in different languages, so that you can obtain the `MessageID` field information required for calling the `AcknowledgeMessage` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) ReceiveMessage(request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
    return c.ReceiveMessageWithContext(context.Background(), request)
}

// ReceiveMessage
// Currently, the `ReceiveMessage` API only supports partitioned topics. It is used to receive messages sent to a specified partitioned topic. If it is called when there are no messages in the partitioned topic, the `ReceiveTimeout` exception will be reported.
//
// 
//
// Instructions on how to use `BatchReceivePolicy`:
//
// 
//
// `BatchReceive` has three parameters:
//
// 
//
// ● `MaxNumMessages`: The maximum number of messages returned by `Receive` when `BatchReceive` is used.
//
// ● `MaxNumBytes`: The maximum size (in bytes) of the message returned by `Receive` when `BatchReceive` is used.
//
// ● `Timeout`: The maximum timeout period (in milliseconds) of calling `Receive` when `BatchReceive` is used.
//
// 
//
// By default, if you don’t specify any of the three parameters, the `BatchReceive` feature is disabled; if one of the three parameter values is above zero, `BatchReceive` is enabled. `BatchReceive` will be disabled when any of the three parameter values reaches the threshold you specify.
//
// 
//
// Note: The values of both `MaxNumMessages` and `MaxNumBytes` are subject to the value of `ReceiveQueueSize`. If the values of `ReceiveQueueSize` and `MaxNumMessages` are 5 and 10, respectively, you can receive up to five rather than 10 messages when `BatchReceive` is used.
//
// 
//
// 
//
// 
//
// The API configured with `BatchReceivePolicy` returns multiple messages at a time.
//
// 
//
// 1. These messages are separated by “###”. After receiving them, you can separate them with split tools in different languages.
//
// 2. MessageIDs are separated by “###”. After receiving the messages, you can separate the MessageIDs with split tools in different languages, so that you can obtain the `MessageID` field information required for calling the `AcknowledgeMessage` API.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) ReceiveMessageWithContext(ctx context.Context, request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
    if request == nil {
        request = NewReceiveMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ReceiveMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReceiveMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewReceiveMessageResponse()
    err = c.Send(request, response)
    return
}

func NewResetMsgSubOffsetByTimestampRequest() (request *ResetMsgSubOffsetByTimestampRequest) {
    request = &ResetMsgSubOffsetByTimestampRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ResetMsgSubOffsetByTimestamp")
    
    
    return
}

func NewResetMsgSubOffsetByTimestampResponse() (response *ResetMsgSubOffsetByTimestampResponse) {
    response = &ResetMsgSubOffsetByTimestampResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetMsgSubOffsetByTimestamp
// This API is used to rewind a message by timestamp, accurate down to the millisecond.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ResetMsgSubOffsetByTimestamp(request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
    return c.ResetMsgSubOffsetByTimestampWithContext(context.Background(), request)
}

// ResetMsgSubOffsetByTimestamp
// This API is used to rewind a message by timestamp, accurate down to the millisecond.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) ResetMsgSubOffsetByTimestampWithContext(ctx context.Context, request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
    if request == nil {
        request = NewResetMsgSubOffsetByTimestampRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ResetMsgSubOffsetByTimestamp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetMsgSubOffsetByTimestamp require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetMsgSubOffsetByTimestampResponse()
    err = c.Send(request, response)
    return
}

func NewResetRocketMQConsumerOffSetRequest() (request *ResetRocketMQConsumerOffSetRequest) {
    request = &ResetRocketMQConsumerOffSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ResetRocketMQConsumerOffSet")
    
    
    return
}

func NewResetRocketMQConsumerOffSetResponse() (response *ResetRocketMQConsumerOffSetResponse) {
    response = &ResetRocketMQConsumerOffSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetRocketMQConsumerOffSet
// This API is used to reset the consumption offset of a specified consumer group to a specified timestamp.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRocketMQConsumerOffSet(request *ResetRocketMQConsumerOffSetRequest) (response *ResetRocketMQConsumerOffSetResponse, err error) {
    return c.ResetRocketMQConsumerOffSetWithContext(context.Background(), request)
}

// ResetRocketMQConsumerOffSet
// This API is used to reset the consumption offset of a specified consumer group to a specified timestamp.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetRocketMQConsumerOffSetWithContext(ctx context.Context, request *ResetRocketMQConsumerOffSetRequest) (response *ResetRocketMQConsumerOffSetResponse, err error) {
    if request == nil {
        request = NewResetRocketMQConsumerOffSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ResetRocketMQConsumerOffSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRocketMQConsumerOffSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRocketMQConsumerOffSetResponse()
    err = c.Send(request, response)
    return
}

func NewRewindCmqQueueRequest() (request *RewindCmqQueueRequest) {
    request = &RewindCmqQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "RewindCmqQueue")
    
    
    return
}

func NewRewindCmqQueueResponse() (response *RewindCmqQueueResponse) {
    response = &RewindCmqQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RewindCmqQueue
// This API is used to rewind a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RewindCmqQueue(request *RewindCmqQueueRequest) (response *RewindCmqQueueResponse, err error) {
    return c.RewindCmqQueueWithContext(context.Background(), request)
}

// RewindCmqQueue
// This API is used to rewind a CMQ queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) RewindCmqQueueWithContext(ctx context.Context, request *RewindCmqQueueRequest) (response *RewindCmqQueueResponse, err error) {
    if request == nil {
        request = NewRewindCmqQueueRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "RewindCmqQueue")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RewindCmqQueue require credential")
    }

    request.SetContext(ctx)
    
    response = NewRewindCmqQueueResponse()
    err = c.Send(request, response)
    return
}

func NewSendBatchMessagesRequest() (request *SendBatchMessagesRequest) {
    request = &SendBatchMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendBatchMessages")
    
    
    return
}

func NewSendBatchMessagesResponse() (response *SendBatchMessagesResponse) {
    response = &SendBatchMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendBatchMessages
// This API is used to batch send messages.
//
// 
//
// Note: the batch message sending API in TDMQ is to package messages into a batch on the service side of TDMQ-HTTP and then send the batch as a TCP request inside the service. Therefore, when using this API, you should still follow the logic of sending a single message: each message is an independent HTTP request, and multiple HTTP requests are aggregated into one batch inside the TDMQ-HTTP service and then sent to the server; that is, batch sending messages is the same as sending a single message in terms of usage, and batch aggregation is completed inside the TDMQ-HTTP service.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendBatchMessages(request *SendBatchMessagesRequest) (response *SendBatchMessagesResponse, err error) {
    return c.SendBatchMessagesWithContext(context.Background(), request)
}

// SendBatchMessages
// This API is used to batch send messages.
//
// 
//
// Note: the batch message sending API in TDMQ is to package messages into a batch on the service side of TDMQ-HTTP and then send the batch as a TCP request inside the service. Therefore, when using this API, you should still follow the logic of sending a single message: each message is an independent HTTP request, and multiple HTTP requests are aggregated into one batch inside the TDMQ-HTTP service and then sent to the server; that is, batch sending messages is the same as sending a single message in terms of usage, and batch aggregation is completed inside the TDMQ-HTTP service.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendBatchMessagesWithContext(ctx context.Context, request *SendBatchMessagesRequest) (response *SendBatchMessagesResponse, err error) {
    if request == nil {
        request = NewSendBatchMessagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SendBatchMessages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendBatchMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendBatchMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewSendCmqMsgRequest() (request *SendCmqMsgRequest) {
    request = &SendCmqMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendCmqMsg")
    
    
    return
}

func NewSendCmqMsgResponse() (response *SendCmqMsgResponse) {
    response = &SendCmqMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendCmqMsg
// This API is used to send a CMQ message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendCmqMsg(request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
    return c.SendCmqMsgWithContext(context.Background(), request)
}

// SendCmqMsg
// This API is used to send a CMQ message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendCmqMsgWithContext(ctx context.Context, request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
    if request == nil {
        request = NewSendCmqMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SendCmqMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendCmqMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendCmqMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSendMessagesRequest() (request *SendMessagesRequest) {
    request = &SendMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendMessages")
    
    
    return
}

func NewSendMessagesResponse() (response *SendMessagesResponse) {
    response = &SendMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendMessages
// This API is used to send a single message.
//
// The message cannot be sent to a persistent topic.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendMessages(request *SendMessagesRequest) (response *SendMessagesResponse, err error) {
    return c.SendMessagesWithContext(context.Background(), request)
}

// SendMessages
// This API is used to send a single message.
//
// The message cannot be sent to a persistent topic.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"
//  FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"
//  FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"
//  FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"
//  FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"
//  FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"
//  FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"
//  FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"
//  INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"
//  INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
func (c *Client) SendMessagesWithContext(ctx context.Context, request *SendMessagesRequest) (response *SendMessagesResponse, err error) {
    if request == nil {
        request = NewSendMessagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SendMessages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewSendMsgRequest() (request *SendMsgRequest) {
    request = &SendMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendMsg")
    
    
    return
}

func NewSendMsgResponse() (response *SendMsgResponse) {
    response = &SendMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendMsg
// This API is used to test message sending. It cannot be used in the production environment.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) SendMsg(request *SendMsgRequest) (response *SendMsgResponse, err error) {
    return c.SendMsgWithContext(context.Background(), request)
}

// SendMsg
// This API is used to test message sending. It cannot be used in the production environment.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) SendMsgWithContext(ctx context.Context, request *SendMsgRequest) (response *SendMsgResponse, err error) {
    if request == nil {
        request = NewSendMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SendMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSendRocketMQMessageRequest() (request *SendRocketMQMessageRequest) {
    request = &SendRocketMQMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SendRocketMQMessage")
    
    
    return
}

func NewSendRocketMQMessageResponse() (response *SendRocketMQMessageResponse) {
    response = &SendRocketMQMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendRocketMQMessage
// This document is used to send a TDMQ for RocketMQ message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) SendRocketMQMessage(request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
    return c.SendRocketMQMessageWithContext(context.Background(), request)
}

// SendRocketMQMessage
// This document is used to send a TDMQ for RocketMQ message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) SendRocketMQMessageWithContext(ctx context.Context, request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
    if request == nil {
        request = NewSendRocketMQMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SendRocketMQMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendRocketMQMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendRocketMQMessageResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindCmqDeadLetterRequest() (request *UnbindCmqDeadLetterRequest) {
    request = &UnbindCmqDeadLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "UnbindCmqDeadLetter")
    
    
    return
}

func NewUnbindCmqDeadLetterResponse() (response *UnbindCmqDeadLetterResponse) {
    response = &UnbindCmqDeadLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindCmqDeadLetter
// This API is used to unbind a CMQ dead letter queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindCmqDeadLetter(request *UnbindCmqDeadLetterRequest) (response *UnbindCmqDeadLetterResponse, err error) {
    return c.UnbindCmqDeadLetterWithContext(context.Background(), request)
}

// UnbindCmqDeadLetter
// This API is used to unbind a CMQ dead letter queue.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UnbindCmqDeadLetterWithContext(ctx context.Context, request *UnbindCmqDeadLetterRequest) (response *UnbindCmqDeadLetterResponse, err error) {
    if request == nil {
        request = NewUnbindCmqDeadLetterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "UnbindCmqDeadLetter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCmqDeadLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindCmqDeadLetterResponse()
    err = c.Send(request, response)
    return
}
