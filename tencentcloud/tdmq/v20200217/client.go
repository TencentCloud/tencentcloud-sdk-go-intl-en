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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER = "MissingParameter"
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"
//  MISSINGPARAMETER = "MissingParameter"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"
//  FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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

func NewCreateProClusterRequest() (request *CreateProClusterRequest) {
    request = &CreateProClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateProCluster")
    
    
    return
}

func NewCreateProClusterResponse() (response *CreateProClusterResponse) {
    response = &CreateProClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProCluster
// This api is used to create a professional cluster with prepayment via api calls.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEPROCLUSTERREGIONNOTEXIST = "FailedOperation.CreateProClusterRegionNotExist"
//  FAILEDOPERATION_GENERATEDEALSANDPAYERROR = "FailedOperation.GenerateDealsAndPayError"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_PRODUCTNOTEXIST = "FailedOperation.ProductNotExist"
//  INVALIDPARAMETER_VPC = "InvalidParameter.Vpc"
//  MISSINGPARAMETER_TAG = "MissingParameter.Tag"
func (c *Client) CreateProCluster(request *CreateProClusterRequest) (response *CreateProClusterResponse, err error) {
    return c.CreateProClusterWithContext(context.Background(), request)
}

// CreateProCluster
// This api is used to create a professional cluster with prepayment via api calls.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATEPROCLUSTERREGIONNOTEXIST = "FailedOperation.CreateProClusterRegionNotExist"
//  FAILEDOPERATION_GENERATEDEALSANDPAYERROR = "FailedOperation.GenerateDealsAndPayError"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_PRODUCTNOTEXIST = "FailedOperation.ProductNotExist"
//  INVALIDPARAMETER_VPC = "InvalidParameter.Vpc"
//  MISSINGPARAMETER_TAG = "MissingParameter.Tag"
func (c *Client) CreateProClusterWithContext(ctx context.Context, request *CreateProClusterRequest) (response *CreateProClusterResponse, err error) {
    if request == nil {
        request = NewCreateProClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateProCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRabbitMQBindingRequest() (request *CreateRabbitMQBindingRequest) {
    request = &CreateRabbitMQBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQBinding")
    
    
    return
}

func NewCreateRabbitMQBindingResponse() (response *CreateRabbitMQBindingResponse) {
    response = &CreateRabbitMQBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRabbitMQBinding
// This API is used to create a TDMQ for RabbitMQ routing relationship.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQBinding(request *CreateRabbitMQBindingRequest) (response *CreateRabbitMQBindingResponse, err error) {
    return c.CreateRabbitMQBindingWithContext(context.Background(), request)
}

// CreateRabbitMQBinding
// This API is used to create a TDMQ for RabbitMQ routing relationship.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQBindingWithContext(ctx context.Context, request *CreateRabbitMQBindingRequest) (response *CreateRabbitMQBindingResponse, err error) {
    if request == nil {
        request = NewCreateRabbitMQBindingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRabbitMQBinding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRabbitMQBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRabbitMQBindingResponse()
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
// This API is used to create a RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRabbitMQVipInstance(request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
    return c.CreateRabbitMQVipInstanceWithContext(context.Background(), request)
}

// CreateRabbitMQVipInstance
// This API is used to create a RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"
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

func NewCreateRocketMQEnvironmentRoleRequest() (request *CreateRocketMQEnvironmentRoleRequest) {
    request = &CreateRocketMQEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQEnvironmentRole")
    
    
    return
}

func NewCreateRocketMQEnvironmentRoleResponse() (response *CreateRocketMQEnvironmentRoleResponse) {
    response = &CreateRocketMQEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQEnvironmentRole
// Creates environment role authorization
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
func (c *Client) CreateRocketMQEnvironmentRole(request *CreateRocketMQEnvironmentRoleRequest) (response *CreateRocketMQEnvironmentRoleResponse, err error) {
    return c.CreateRocketMQEnvironmentRoleWithContext(context.Background(), request)
}

// CreateRocketMQEnvironmentRole
// Creates environment role authorization
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
func (c *Client) CreateRocketMQEnvironmentRoleWithContext(ctx context.Context, request *CreateRocketMQEnvironmentRoleRequest) (response *CreateRocketMQEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQEnvironmentRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQEnvironmentRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQEnvironmentRoleResponse()
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
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
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
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
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

func NewCreateRocketMQRoleRequest() (request *CreateRocketMQRoleRequest) {
    request = &CreateRocketMQRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQRole")
    
    
    return
}

func NewCreateRocketMQRoleResponse() (response *CreateRocketMQRoleResponse) {
    response = &CreateRocketMQRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQRole
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
func (c *Client) CreateRocketMQRole(request *CreateRocketMQRoleRequest) (response *CreateRocketMQRoleResponse, err error) {
    return c.CreateRocketMQRoleWithContext(context.Background(), request)
}

// CreateRocketMQRole
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
func (c *Client) CreateRocketMQRoleWithContext(ctx context.Context, request *CreateRocketMQRoleRequest) (response *CreateRocketMQRoleResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQRoleResponse()
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

func NewCreateRocketMQVipInstanceRequest() (request *CreateRocketMQVipInstanceRequest) {
    request = &CreateRocketMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQVipInstance")
    
    
    return
}

func NewCreateRocketMQVipInstanceResponse() (response *CreateRocketMQVipInstanceResponse) {
    response = &CreateRocketMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRocketMQVipInstance
// This API is used to create a RocketMQ Exclusive Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
func (c *Client) CreateRocketMQVipInstance(request *CreateRocketMQVipInstanceRequest) (response *CreateRocketMQVipInstanceResponse, err error) {
    return c.CreateRocketMQVipInstanceWithContext(context.Background(), request)
}

// CreateRocketMQVipInstance
// This API is used to create a RocketMQ Exclusive Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"
func (c *Client) CreateRocketMQVipInstanceWithContext(ctx context.Context, request *CreateRocketMQVipInstanceRequest) (response *CreateRocketMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewCreateRocketMQVipInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "CreateRocketMQVipInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRocketMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRocketMQVipInstanceResponse()
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
//  RESOURCEINUSE = "ResourceInUse"
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
//  RESOURCEINUSE = "ResourceInUse"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_ISOLATECONSUMERENABLE = "FailedOperation.IsolateConsumerEnable"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACKTIME = "InvalidParameterValue.AckTime"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNKNOWNPARAMETER_POLICY = "UnknownParameter.Policy"
//  UNSUPPORTEDOPERATION_TOPICUNACK = "UnsupportedOperation.TopicUnack"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// This API is used to add a message topic in the specified partition and type.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"
//  FAILEDOPERATION_ISOLATECONSUMERENABLE = "FailedOperation.IsolateConsumerEnable"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_REPLICATIONDESTCHECKFAILEDERROR = "FailedOperation.ReplicationDestCheckFailedError"
//  FAILEDOPERATION_REPLICATIONSOURCECHECKFAILEDERROR = "FailedOperation.ReplicationSourceCheckFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACKTIME = "InvalidParameterValue.AckTime"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
//  UNKNOWNPARAMETER_POLICY = "UnknownParameter.Policy"
//  UNSUPPORTEDOPERATION_TOPICUNACK = "UnsupportedOperation.TopicUnack"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"
//  FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"
//  FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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

func NewDeleteProClusterRequest() (request *DeleteProClusterRequest) {
    request = &DeleteProClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteProCluster")
    
    
    return
}

func NewDeleteProClusterResponse() (response *DeleteProClusterResponse) {
    response = &DeleteProClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProCluster
// This API is used to delete a professional cluster with prepayment via API calls.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETECLUSTERPROTECTION = "FailedOperation.DeleteClusterProtection"
//  FAILEDOPERATION_INSTANCECANNOTDELETE = "FailedOperation.InstanceCanNotDelete"
//  FAILEDOPERATION_ONLINEREFUNDRESOURCENOTEXIT = "FailedOperation.OnlineRefundResourceNotExit"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteProCluster(request *DeleteProClusterRequest) (response *DeleteProClusterResponse, err error) {
    return c.DeleteProClusterWithContext(context.Background(), request)
}

// DeleteProCluster
// This API is used to delete a professional cluster with prepayment via API calls.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETECLUSTERPROTECTION = "FailedOperation.DeleteClusterProtection"
//  FAILEDOPERATION_INSTANCECANNOTDELETE = "FailedOperation.InstanceCanNotDelete"
//  FAILEDOPERATION_ONLINEREFUNDRESOURCENOTEXIT = "FailedOperation.OnlineRefundResourceNotExit"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DeleteProClusterWithContext(ctx context.Context, request *DeleteProClusterRequest) (response *DeleteProClusterResponse, err error) {
    if request == nil {
        request = NewDeleteProClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteProCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQBindingRequest() (request *DeleteRabbitMQBindingRequest) {
    request = &DeleteRabbitMQBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQBinding")
    
    
    return
}

func NewDeleteRabbitMQBindingResponse() (response *DeleteRabbitMQBindingResponse) {
    response = &DeleteRabbitMQBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQBinding
// This API is used to unbind RabbitMQ routing relationships.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQBinding(request *DeleteRabbitMQBindingRequest) (response *DeleteRabbitMQBindingResponse, err error) {
    return c.DeleteRabbitMQBindingWithContext(context.Background(), request)
}

// DeleteRabbitMQBinding
// This API is used to unbind RabbitMQ routing relationships.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQBindingWithContext(ctx context.Context, request *DeleteRabbitMQBindingRequest) (response *DeleteRabbitMQBindingResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQBindingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRabbitMQBinding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQBinding require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQBindingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRabbitMQPermissionRequest() (request *DeleteRabbitMQPermissionRequest) {
    request = &DeleteRabbitMQPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQPermission")
    
    
    return
}

func NewDeleteRabbitMQPermissionResponse() (response *DeleteRabbitMQPermissionResponse) {
    response = &DeleteRabbitMQPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRabbitMQPermission
// This API is used to delete RabbitMQ permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQPermission(request *DeleteRabbitMQPermissionRequest) (response *DeleteRabbitMQPermissionResponse, err error) {
    return c.DeleteRabbitMQPermissionWithContext(context.Background(), request)
}

// DeleteRabbitMQPermission
// This API is used to delete RabbitMQ permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRabbitMQPermissionWithContext(ctx context.Context, request *DeleteRabbitMQPermissionRequest) (response *DeleteRabbitMQPermissionResponse, err error) {
    if request == nil {
        request = NewDeleteRabbitMQPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRabbitMQPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRabbitMQPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRabbitMQPermissionResponse()
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
// This API is used to delete a RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRabbitMQVipInstance(request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
    return c.DeleteRabbitMQVipInstanceWithContext(context.Background(), request)
}

// DeleteRabbitMQVipInstance
// This API is used to delete a RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDeleteRocketMQEnvironmentRolesRequest() (request *DeleteRocketMQEnvironmentRolesRequest) {
    request = &DeleteRocketMQEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQEnvironmentRoles")
    
    
    return
}

func NewDeleteRocketMQEnvironmentRolesResponse() (response *DeleteRocketMQEnvironmentRolesResponse) {
    response = &DeleteRocketMQEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQEnvironmentRoles
// Deletes environment role authorization
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteRocketMQEnvironmentRoles(request *DeleteRocketMQEnvironmentRolesRequest) (response *DeleteRocketMQEnvironmentRolesResponse, err error) {
    return c.DeleteRocketMQEnvironmentRolesWithContext(context.Background(), request)
}

// DeleteRocketMQEnvironmentRoles
// Deletes environment role authorization
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) DeleteRocketMQEnvironmentRolesWithContext(ctx context.Context, request *DeleteRocketMQEnvironmentRolesRequest) (response *DeleteRocketMQEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQEnvironmentRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQEnvironmentRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQEnvironmentRolesResponse()
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
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
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
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
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

func NewDeleteRocketMQRolesRequest() (request *DeleteRocketMQRolesRequest) {
    request = &DeleteRocketMQRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQRoles")
    
    
    return
}

func NewDeleteRocketMQRolesResponse() (response *DeleteRocketMQRolesResponse) {
    response = &DeleteRocketMQRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQRoles
// Deletes roles. Batch deletion is supported.
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
func (c *Client) DeleteRocketMQRoles(request *DeleteRocketMQRolesRequest) (response *DeleteRocketMQRolesResponse, err error) {
    return c.DeleteRocketMQRolesWithContext(context.Background(), request)
}

// DeleteRocketMQRoles
// Deletes roles. Batch deletion is supported.
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
func (c *Client) DeleteRocketMQRolesWithContext(ctx context.Context, request *DeleteRocketMQRolesRequest) (response *DeleteRocketMQRolesResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQRolesResponse()
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

func NewDeleteRocketMQVipInstanceRequest() (request *DeleteRocketMQVipInstanceRequest) {
    request = &DeleteRocketMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQVipInstance")
    
    
    return
}

func NewDeleteRocketMQVipInstanceResponse() (response *DeleteRocketMQVipInstanceResponse) {
    response = &DeleteRocketMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRocketMQVipInstance
// This API is used to delete a RocketMQ Exclusive Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteRocketMQVipInstance(request *DeleteRocketMQVipInstanceRequest) (response *DeleteRocketMQVipInstanceResponse, err error) {
    return c.DeleteRocketMQVipInstanceWithContext(context.Background(), request)
}

// DeleteRocketMQVipInstance
// This API is used to delete a RocketMQ Exclusive Edition instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteRocketMQVipInstanceWithContext(ctx context.Context, request *DeleteRocketMQVipInstanceRequest) (response *DeleteRocketMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteRocketMQVipInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DeleteRocketMQVipInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRocketMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRocketMQVipInstanceResponse()
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// This API is used to get the list of clusters.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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

func NewDescribeMqMsgTraceRequest() (request *DescribeMqMsgTraceRequest) {
    request = &DescribeMqMsgTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMqMsgTrace")
    
    
    return
}

func NewDescribeMqMsgTraceResponse() (response *DescribeMqMsgTraceResponse) {
    response = &DescribeMqMsgTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMqMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeMqMsgTrace(request *DescribeMqMsgTraceRequest) (response *DescribeMqMsgTraceResponse, err error) {
    return c.DescribeMqMsgTraceWithContext(context.Background(), request)
}

// DescribeMqMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeMqMsgTraceWithContext(ctx context.Context, request *DescribeMqMsgTraceRequest) (response *DescribeMqMsgTraceResponse, err error) {
    if request == nil {
        request = NewDescribeMqMsgTraceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeMqMsgTrace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMqMsgTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMqMsgTraceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgRequest() (request *DescribeMsgRequest) {
    request = &DescribeMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMsg")
    
    
    return
}

func NewDescribeMsgResponse() (response *DescribeMsgResponse) {
    response = &DescribeMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsg
// This API is used to get message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_MSG = "ResourceNotFound.Msg"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeMsg(request *DescribeMsgRequest) (response *DescribeMsgResponse, err error) {
    return c.DescribeMsgWithContext(context.Background(), request)
}

// DescribeMsg
// This API is used to get message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_MSG = "ResourceNotFound.Msg"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeMsgWithContext(ctx context.Context, request *DescribeMsgRequest) (response *DescribeMsgResponse, err error) {
    if request == nil {
        request = NewDescribeMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsgTraceRequest() (request *DescribeMsgTraceRequest) {
    request = &DescribeMsgTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMsgTrace")
    
    
    return
}

func NewDescribeMsgTraceResponse() (response *DescribeMsgTraceResponse) {
    response = &DescribeMsgTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_MSGPRODUCELOG = "ResourceNotFound.MsgProduceLog"
func (c *Client) DescribeMsgTrace(request *DescribeMsgTraceRequest) (response *DescribeMsgTraceResponse, err error) {
    return c.DescribeMsgTraceWithContext(context.Background(), request)
}

// DescribeMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_MSGPRODUCELOG = "ResourceNotFound.MsgProduceLog"
func (c *Client) DescribeMsgTraceWithContext(ctx context.Context, request *DescribeMsgTraceRequest) (response *DescribeMsgTraceResponse, err error) {
    if request == nil {
        request = NewDescribeMsgTraceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeMsgTrace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMsgTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMsgTraceResponse()
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublisherSummary(request *DescribePublisherSummaryRequest) (response *DescribePublisherSummaryResponse, err error) {
    return c.DescribePublisherSummaryWithContext(context.Background(), request)
}

// DescribePublisherSummary
// This API is used to obtain message production overview information.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribePublishers(request *DescribePublishersRequest) (response *DescribePublishersResponse, err error) {
    return c.DescribePublishersWithContext(context.Background(), request)
}

// DescribePublishers
// This API is used to obtain the list of producer information.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
func (c *Client) DescribePulsarProInstances(request *DescribePulsarProInstancesRequest) (response *DescribePulsarProInstancesResponse, err error) {
    return c.DescribePulsarProInstancesWithContext(context.Background(), request)
}

// DescribePulsarProInstances
// This API is used to query the list of the purchased TDMQ for Pulsar pro instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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

func NewDescribeRabbitMQBindingsRequest() (request *DescribeRabbitMQBindingsRequest) {
    request = &DescribeRabbitMQBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQBindings")
    
    
    return
}

func NewDescribeRabbitMQBindingsResponse() (response *DescribeRabbitMQBindingsResponse) {
    response = &DescribeRabbitMQBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQBindings
// This API is used to query the list of RabbitMQ route relations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQBindings(request *DescribeRabbitMQBindingsRequest) (response *DescribeRabbitMQBindingsResponse, err error) {
    return c.DescribeRabbitMQBindingsWithContext(context.Background(), request)
}

// DescribeRabbitMQBindings
// This API is used to query the list of RabbitMQ route relations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQBindingsWithContext(ctx context.Context, request *DescribeRabbitMQBindingsRequest) (response *DescribeRabbitMQBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQBindingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQBindings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQExchangesRequest() (request *DescribeRabbitMQExchangesRequest) {
    request = &DescribeRabbitMQExchangesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQExchanges")
    
    
    return
}

func NewDescribeRabbitMQExchangesResponse() (response *DescribeRabbitMQExchangesResponse) {
    response = &DescribeRabbitMQExchangesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQExchanges
// This API is used to query the list of TDMQ for RabbitMQ exchanges.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQExchanges(request *DescribeRabbitMQExchangesRequest) (response *DescribeRabbitMQExchangesResponse, err error) {
    return c.DescribeRabbitMQExchangesWithContext(context.Background(), request)
}

// DescribeRabbitMQExchanges
// This API is used to query the list of TDMQ for RabbitMQ exchanges.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQExchangesWithContext(ctx context.Context, request *DescribeRabbitMQExchangesRequest) (response *DescribeRabbitMQExchangesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQExchangesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQExchanges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQExchanges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQExchangesResponse()
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
// This API is used to query the RabbitMQ managed node list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQNodeList(request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
    return c.DescribeRabbitMQNodeListWithContext(context.Background(), request)
}

// DescribeRabbitMQNodeList
// This API is used to query the RabbitMQ managed node list.
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

func NewDescribeRabbitMQPermissionRequest() (request *DescribeRabbitMQPermissionRequest) {
    request = &DescribeRabbitMQPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQPermission")
    
    
    return
}

func NewDescribeRabbitMQPermissionResponse() (response *DescribeRabbitMQPermissionResponse) {
    response = &DescribeRabbitMQPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQPermission
// This API is used to query the list of TDMQ for RabbitMQ permissions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQPermission(request *DescribeRabbitMQPermissionRequest) (response *DescribeRabbitMQPermissionResponse, err error) {
    return c.DescribeRabbitMQPermissionWithContext(context.Background(), request)
}

// DescribeRabbitMQPermission
// This API is used to query the list of TDMQ for RabbitMQ permissions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQPermissionWithContext(ctx context.Context, request *DescribeRabbitMQPermissionRequest) (response *DescribeRabbitMQPermissionResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQQueueDetailRequest() (request *DescribeRabbitMQQueueDetailRequest) {
    request = &DescribeRabbitMQQueueDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQQueueDetail")
    
    
    return
}

func NewDescribeRabbitMQQueueDetailResponse() (response *DescribeRabbitMQQueueDetailResponse) {
    response = &DescribeRabbitMQQueueDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQQueueDetail
// This API is used to query the details of RabbitMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQQueueDetail(request *DescribeRabbitMQQueueDetailRequest) (response *DescribeRabbitMQQueueDetailResponse, err error) {
    return c.DescribeRabbitMQQueueDetailWithContext(context.Background(), request)
}

// DescribeRabbitMQQueueDetail
// This API is used to query the details of RabbitMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRabbitMQQueueDetailWithContext(ctx context.Context, request *DescribeRabbitMQQueueDetailRequest) (response *DescribeRabbitMQQueueDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQQueueDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQQueueDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQQueueDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQQueueDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRabbitMQQueuesRequest() (request *DescribeRabbitMQQueuesRequest) {
    request = &DescribeRabbitMQQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQQueues")
    
    
    return
}

func NewDescribeRabbitMQQueuesResponse() (response *DescribeRabbitMQQueuesResponse) {
    response = &DescribeRabbitMQQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQQueues
// This API is used to query the list of TDMQ for RabbitMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQQueues(request *DescribeRabbitMQQueuesRequest) (response *DescribeRabbitMQQueuesResponse, err error) {
    return c.DescribeRabbitMQQueuesWithContext(context.Background(), request)
}

// DescribeRabbitMQQueues
// This API is used to query the list of TDMQ for RabbitMQ queues.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRabbitMQQueuesWithContext(ctx context.Context, request *DescribeRabbitMQQueuesRequest) (response *DescribeRabbitMQQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQQueuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQQueues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQQueues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQQueuesResponse()
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

func NewDescribeRabbitMQVipInstanceRequest() (request *DescribeRabbitMQVipInstanceRequest) {
    request = &DescribeRabbitMQVipInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstance")
    
    
    return
}

func NewDescribeRabbitMQVipInstanceResponse() (response *DescribeRabbitMQVipInstanceResponse) {
    response = &DescribeRabbitMQVipInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRabbitMQVipInstance
// This API is used to obtain information about one RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstance(request *DescribeRabbitMQVipInstanceRequest) (response *DescribeRabbitMQVipInstanceResponse, err error) {
    return c.DescribeRabbitMQVipInstanceWithContext(context.Background(), request)
}

// DescribeRabbitMQVipInstance
// This API is used to obtain information about one RabbitMQ managed instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstanceWithContext(ctx context.Context, request *DescribeRabbitMQVipInstanceRequest) (response *DescribeRabbitMQVipInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeRabbitMQVipInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRabbitMQVipInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRabbitMQVipInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRabbitMQVipInstanceResponse()
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
// This API is used to query the RabbitMQ managed instance list of user purchases.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRabbitMQVipInstances(request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
    return c.DescribeRabbitMQVipInstancesWithContext(context.Background(), request)
}

// DescribeRabbitMQVipInstances
// This API is used to query the RabbitMQ managed instance list of user purchases.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
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

func NewDescribeRocketMQEnvironmentRolesRequest() (request *DescribeRocketMQEnvironmentRolesRequest) {
    request = &DescribeRocketMQEnvironmentRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQEnvironmentRoles")
    
    
    return
}

func NewDescribeRocketMQEnvironmentRolesResponse() (response *DescribeRocketMQEnvironmentRolesResponse) {
    response = &DescribeRocketMQEnvironmentRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQEnvironmentRoles
// Obtains the namespace role list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeRocketMQEnvironmentRoles(request *DescribeRocketMQEnvironmentRolesRequest) (response *DescribeRocketMQEnvironmentRolesResponse, err error) {
    return c.DescribeRocketMQEnvironmentRolesWithContext(context.Background(), request)
}

// DescribeRocketMQEnvironmentRoles
// Obtains the namespace role list
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DescribeRocketMQEnvironmentRolesWithContext(ctx context.Context, request *DescribeRocketMQEnvironmentRolesRequest) (response *DescribeRocketMQEnvironmentRolesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQEnvironmentRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQEnvironmentRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQEnvironmentRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQEnvironmentRolesResponse()
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

func NewDescribeRocketMQMsgTraceRequest() (request *DescribeRocketMQMsgTraceRequest) {
    request = &DescribeRocketMQMsgTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMsgTrace")
    
    
    return
}

func NewDescribeRocketMQMsgTraceResponse() (response *DescribeRocketMQMsgTraceResponse) {
    response = &DescribeRocketMQMsgTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsgTrace(request *DescribeRocketMQMsgTraceRequest) (response *DescribeRocketMQMsgTraceResponse, err error) {
    return c.DescribeRocketMQMsgTraceWithContext(context.Background(), request)
}

// DescribeRocketMQMsgTrace
// Queries message trajectory
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQMsgTraceWithContext(ctx context.Context, request *DescribeRocketMQMsgTraceRequest) (response *DescribeRocketMQMsgTraceResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQMsgTraceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQMsgTrace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQMsgTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQMsgTraceResponse()
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

func NewDescribeRocketMQProducersRequest() (request *DescribeRocketMQProducersRequest) {
    request = &DescribeRocketMQProducersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQProducers")
    
    
    return
}

func NewDescribeRocketMQProducersResponse() (response *DescribeRocketMQProducersResponse) {
    response = &DescribeRocketMQProducersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQProducers
// This API is used to query the producer client list under a specified topic in RocketMQ.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQProducers(request *DescribeRocketMQProducersRequest) (response *DescribeRocketMQProducersResponse, err error) {
    return c.DescribeRocketMQProducersWithContext(context.Background(), request)
}

// DescribeRocketMQProducers
// This API is used to query the producer client list under a specified topic in RocketMQ.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQProducersWithContext(ctx context.Context, request *DescribeRocketMQProducersRequest) (response *DescribeRocketMQProducersResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQProducersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQProducers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQProducers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQProducersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQPublicAccessMonitorDataRequest() (request *DescribeRocketMQPublicAccessMonitorDataRequest) {
    request = &DescribeRocketMQPublicAccessMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQPublicAccessMonitorData")
    
    
    return
}

func NewDescribeRocketMQPublicAccessMonitorDataResponse() (response *DescribeRocketMQPublicAccessMonitorDataResponse) {
    response = &DescribeRocketMQPublicAccessMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQPublicAccessMonitorData
// This API is used to pull public network metric monitoring data from TCOP. Currently, only inbound bandwidth and outbound bandwidth metrics from client to LB are supported.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQPublicAccessMonitorData(request *DescribeRocketMQPublicAccessMonitorDataRequest) (response *DescribeRocketMQPublicAccessMonitorDataResponse, err error) {
    return c.DescribeRocketMQPublicAccessMonitorDataWithContext(context.Background(), request)
}

// DescribeRocketMQPublicAccessMonitorData
// This API is used to pull public network metric monitoring data from TCOP. Currently, only inbound bandwidth and outbound bandwidth metrics from client to LB are supported.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQPublicAccessMonitorDataWithContext(ctx context.Context, request *DescribeRocketMQPublicAccessMonitorDataRequest) (response *DescribeRocketMQPublicAccessMonitorDataResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQPublicAccessMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQPublicAccessMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQPublicAccessMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQPublicAccessMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQPublicAccessPointRequest() (request *DescribeRocketMQPublicAccessPointRequest) {
    request = &DescribeRocketMQPublicAccessPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQPublicAccessPoint")
    
    
    return
}

func NewDescribeRocketMQPublicAccessPointResponse() (response *DescribeRocketMQPublicAccessPointResponse) {
    response = &DescribeRocketMQPublicAccessPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQPublicAccessPoint
// This API is used to query the public network access information of RocketMQ instances.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQPublicAccessPoint(request *DescribeRocketMQPublicAccessPointRequest) (response *DescribeRocketMQPublicAccessPointResponse, err error) {
    return c.DescribeRocketMQPublicAccessPointWithContext(context.Background(), request)
}

// DescribeRocketMQPublicAccessPoint
// This API is used to query the public network access information of RocketMQ instances.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_NAMSPACE = "ResourceNotFound.Namspace"
func (c *Client) DescribeRocketMQPublicAccessPointWithContext(ctx context.Context, request *DescribeRocketMQPublicAccessPointRequest) (response *DescribeRocketMQPublicAccessPointResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQPublicAccessPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQPublicAccessPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQPublicAccessPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQPublicAccessPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQRolesRequest() (request *DescribeRocketMQRolesRequest) {
    request = &DescribeRocketMQRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQRoles")
    
    
    return
}

func NewDescribeRocketMQRolesResponse() (response *DescribeRocketMQRolesResponse) {
    response = &DescribeRocketMQRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQRoles
// Obtains the list of roles
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQRoles(request *DescribeRocketMQRolesRequest) (response *DescribeRocketMQRolesResponse, err error) {
    return c.DescribeRocketMQRolesWithContext(context.Background(), request)
}

// DescribeRocketMQRoles
// Obtains the list of roles
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQRolesWithContext(ctx context.Context, request *DescribeRocketMQRolesRequest) (response *DescribeRocketMQRolesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopUsagesRequest() (request *DescribeRocketMQTopUsagesRequest) {
    request = &DescribeRocketMQTopUsagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopUsages")
    
    
    return
}

func NewDescribeRocketMQTopUsagesResponse() (response *DescribeRocketMQTopUsagesResponse) {
    response = &DescribeRocketMQTopUsagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopUsages
// Used to obtain the RocketMQ metric sorting list, such as sorting topics under a cluster instance by the most occupied storage space.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopUsages(request *DescribeRocketMQTopUsagesRequest) (response *DescribeRocketMQTopUsagesResponse, err error) {
    return c.DescribeRocketMQTopUsagesWithContext(context.Background(), request)
}

// DescribeRocketMQTopUsages
// Used to obtain the RocketMQ metric sorting list, such as sorting topics under a cluster instance by the most occupied storage space.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopUsagesWithContext(ctx context.Context, request *DescribeRocketMQTopUsagesRequest) (response *DescribeRocketMQTopUsagesResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopUsagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopUsages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopUsages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopUsagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopicRequest() (request *DescribeRocketMQTopicRequest) {
    request = &DescribeRocketMQTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopic")
    
    
    return
}

func NewDescribeRocketMQTopicResponse() (response *DescribeRocketMQTopicResponse) {
    response = &DescribeRocketMQTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopic
// This API is used to obtain RocketMQ topic details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeRocketMQTopic(request *DescribeRocketMQTopicRequest) (response *DescribeRocketMQTopicResponse, err error) {
    return c.DescribeRocketMQTopicWithContext(context.Background(), request)
}

// DescribeRocketMQTopic
// This API is used to obtain RocketMQ topic details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeRocketMQTopicWithContext(ctx context.Context, request *DescribeRocketMQTopicRequest) (response *DescribeRocketMQTopicResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopicMsgsRequest() (request *DescribeRocketMQTopicMsgsRequest) {
    request = &DescribeRocketMQTopicMsgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopicMsgs")
    
    
    return
}

func NewDescribeRocketMQTopicMsgsResponse() (response *DescribeRocketMQTopicMsgsResponse) {
    response = &DescribeRocketMQTopicMsgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopicMsgs
// Query RocketMQ messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicMsgs(request *DescribeRocketMQTopicMsgsRequest) (response *DescribeRocketMQTopicMsgsResponse, err error) {
    return c.DescribeRocketMQTopicMsgsWithContext(context.Background(), request)
}

// DescribeRocketMQTopicMsgs
// Query RocketMQ messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicMsgsWithContext(ctx context.Context, request *DescribeRocketMQTopicMsgsRequest) (response *DescribeRocketMQTopicMsgsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicMsgsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopicMsgs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopicMsgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicMsgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRocketMQTopicStatsRequest() (request *DescribeRocketMQTopicStatsRequest) {
    request = &DescribeRocketMQTopicStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopicStats")
    
    
    return
}

func NewDescribeRocketMQTopicStatsResponse() (response *DescribeRocketMQTopicStatsResponse) {
    response = &DescribeRocketMQTopicStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopicStats
// This API is used to obtain the topic production details list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicStats(request *DescribeRocketMQTopicStatsRequest) (response *DescribeRocketMQTopicStatsResponse, err error) {
    return c.DescribeRocketMQTopicStatsWithContext(context.Background(), request)
}

// DescribeRocketMQTopicStats
// This API is used to obtain the topic production details list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicStatsWithContext(ctx context.Context, request *DescribeRocketMQTopicStatsRequest) (response *DescribeRocketMQTopicStatsResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopicStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopicStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicStatsResponse()
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

func NewDescribeRocketMQTopicsByGroupRequest() (request *DescribeRocketMQTopicsByGroupRequest) {
    request = &DescribeRocketMQTopicsByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopicsByGroup")
    
    
    return
}

func NewDescribeRocketMQTopicsByGroupResponse() (response *DescribeRocketMQTopicsByGroupResponse) {
    response = &DescribeRocketMQTopicsByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRocketMQTopicsByGroup
// Obtains the list of topics subscribed under a specified consumer group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicsByGroup(request *DescribeRocketMQTopicsByGroupRequest) (response *DescribeRocketMQTopicsByGroupResponse, err error) {
    return c.DescribeRocketMQTopicsByGroupWithContext(context.Background(), request)
}

// DescribeRocketMQTopicsByGroup
// Obtains the list of topics subscribed under a specified consumer group
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) DescribeRocketMQTopicsByGroupWithContext(ctx context.Context, request *DescribeRocketMQTopicsByGroupRequest) (response *DescribeRocketMQTopicsByGroupResponse, err error) {
    if request == nil {
        request = NewDescribeRocketMQTopicsByGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeRocketMQTopicsByGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRocketMQTopicsByGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRocketMQTopicsByGroupResponse()
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

func NewDescribeTopicMsgsRequest() (request *DescribeTopicMsgsRequest) {
    request = &DescribeTopicMsgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopicMsgs")
    
    
    return
}

func NewDescribeTopicMsgsResponse() (response *DescribeTopicMsgsResponse) {
    response = &DescribeTopicMsgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicMsgs
// This API is used to query messages.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_MSGTIME = "LimitExceeded.MsgTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopicMsgs(request *DescribeTopicMsgsRequest) (response *DescribeTopicMsgsResponse, err error) {
    return c.DescribeTopicMsgsWithContext(context.Background(), request)
}

// DescribeTopicMsgs
// This API is used to query messages.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  LIMITEXCEEDED_MSGTIME = "LimitExceeded.MsgTime"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopicMsgsWithContext(ctx context.Context, request *DescribeTopicMsgsRequest) (response *DescribeTopicMsgsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicMsgsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "DescribeTopicMsgs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicMsgs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicMsgsResponse()
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
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

func NewExportRocketMQMessageDetailRequest() (request *ExportRocketMQMessageDetailRequest) {
    request = &ExportRocketMQMessageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ExportRocketMQMessageDetail")
    
    
    return
}

func NewExportRocketMQMessageDetailResponse() (response *ExportRocketMQMessageDetailResponse) {
    response = &ExportRocketMQMessageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportRocketMQMessageDetail
// Export the RocketMQ message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) ExportRocketMQMessageDetail(request *ExportRocketMQMessageDetailRequest) (response *ExportRocketMQMessageDetailResponse, err error) {
    return c.ExportRocketMQMessageDetailWithContext(context.Background(), request)
}

// ExportRocketMQMessageDetail
// Export the RocketMQ message details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_FAILEDOPERATION = "InternalError.FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) ExportRocketMQMessageDetailWithContext(ctx context.Context, request *ExportRocketMQMessageDetailRequest) (response *ExportRocketMQMessageDetailResponse, err error) {
    if request == nil {
        request = NewExportRocketMQMessageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ExportRocketMQMessageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportRocketMQMessageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportRocketMQMessageDetailResponse()
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_MODIFYCLUSTER = "FailedOperation.ModifyCluster"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_MODIFYCLUSTER = "FailedOperation.ModifyCluster"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
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
//  LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"
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

func NewModifyPublicNetworkSecurityPolicyRequest() (request *ModifyPublicNetworkSecurityPolicyRequest) {
    request = &ModifyPublicNetworkSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyPublicNetworkSecurityPolicy")
    
    
    return
}

func NewModifyPublicNetworkSecurityPolicyResponse() (response *ModifyPublicNetworkSecurityPolicyResponse) {
    response = &ModifyPublicNetworkSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPublicNetworkSecurityPolicy
// This API is used to modify the public network security policy for pulsar Pro Edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_INVALIDEXISTPUBLICACCESSPOINTERROR = "FailedOperation.InvalidExistPublicAccessPointError"
//  FAILEDOPERATION_INVALIDWHITELISTERROR = "FailedOperation.InvalidWhiteListError"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyPublicNetworkSecurityPolicy(request *ModifyPublicNetworkSecurityPolicyRequest) (response *ModifyPublicNetworkSecurityPolicyResponse, err error) {
    return c.ModifyPublicNetworkSecurityPolicyWithContext(context.Background(), request)
}

// ModifyPublicNetworkSecurityPolicy
// This API is used to modify the public network security policy for pulsar Pro Edition.
//
// error code that may be returned:
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_INVALIDEXISTPUBLICACCESSPOINTERROR = "FailedOperation.InvalidExistPublicAccessPointError"
//  FAILEDOPERATION_INVALIDWHITELISTERROR = "FailedOperation.InvalidWhiteListError"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyPublicNetworkSecurityPolicyWithContext(ctx context.Context, request *ModifyPublicNetworkSecurityPolicyRequest) (response *ModifyPublicNetworkSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPublicNetworkSecurityPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyPublicNetworkSecurityPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublicNetworkSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublicNetworkSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRabbitMQPermissionRequest() (request *ModifyRabbitMQPermissionRequest) {
    request = &ModifyRabbitMQPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQPermission")
    
    
    return
}

func NewModifyRabbitMQPermissionResponse() (response *ModifyRabbitMQPermissionResponse) {
    response = &ModifyRabbitMQPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRabbitMQPermission
// This API is used to modify RabbitMQ permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQPermission(request *ModifyRabbitMQPermissionRequest) (response *ModifyRabbitMQPermissionResponse, err error) {
    return c.ModifyRabbitMQPermissionWithContext(context.Background(), request)
}

// ModifyRabbitMQPermission
// This API is used to modify RabbitMQ permissions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyRabbitMQPermissionWithContext(ctx context.Context, request *ModifyRabbitMQPermissionRequest) (response *ModifyRabbitMQPermissionResponse, err error) {
    if request == nil {
        request = NewModifyRabbitMQPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRabbitMQPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRabbitMQPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRabbitMQPermissionResponse()
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

func NewModifyRocketMQEnvironmentRoleRequest() (request *ModifyRocketMQEnvironmentRoleRequest) {
    request = &ModifyRocketMQEnvironmentRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQEnvironmentRole")
    
    
    return
}

func NewModifyRocketMQEnvironmentRoleResponse() (response *ModifyRocketMQEnvironmentRoleResponse) {
    response = &ModifyRocketMQEnvironmentRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQEnvironmentRole
// Modifies environment role authorization
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) ModifyRocketMQEnvironmentRole(request *ModifyRocketMQEnvironmentRoleRequest) (response *ModifyRocketMQEnvironmentRoleResponse, err error) {
    return c.ModifyRocketMQEnvironmentRoleWithContext(context.Background(), request)
}

// ModifyRocketMQEnvironmentRole
// Modifies environment role authorization
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"
//  INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"
//  MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"
func (c *Client) ModifyRocketMQEnvironmentRoleWithContext(ctx context.Context, request *ModifyRocketMQEnvironmentRoleRequest) (response *ModifyRocketMQEnvironmentRoleResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQEnvironmentRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQEnvironmentRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQEnvironmentRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQEnvironmentRoleResponse()
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

func NewModifyRocketMQInstanceRequest() (request *ModifyRocketMQInstanceRequest) {
    request = &ModifyRocketMQInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQInstance")
    
    
    return
}

func NewModifyRocketMQInstanceResponse() (response *ModifyRocketMQInstanceResponse) {
    response = &ModifyRocketMQInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQInstance
// Modify the RocketMQ dedicated instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRocketMQInstance(request *ModifyRocketMQInstanceRequest) (response *ModifyRocketMQInstanceResponse, err error) {
    return c.ModifyRocketMQInstanceWithContext(context.Background(), request)
}

// ModifyRocketMQInstance
// Modify the RocketMQ dedicated instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRocketMQInstanceWithContext(ctx context.Context, request *ModifyRocketMQInstanceRequest) (response *ModifyRocketMQInstanceResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQInstanceResponse()
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

func NewModifyRocketMQRoleRequest() (request *ModifyRocketMQRoleRequest) {
    request = &ModifyRocketMQRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQRole")
    
    
    return
}

func NewModifyRocketMQRoleResponse() (response *ModifyRocketMQRoleResponse) {
    response = &ModifyRocketMQRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRocketMQRole
// Modifies roles
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
func (c *Client) ModifyRocketMQRole(request *ModifyRocketMQRoleRequest) (response *ModifyRocketMQRoleResponse, err error) {
    return c.ModifyRocketMQRoleWithContext(context.Background(), request)
}

// ModifyRocketMQRole
// Modifies roles
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
func (c *Client) ModifyRocketMQRoleWithContext(ctx context.Context, request *ModifyRocketMQRoleRequest) (response *ModifyRocketMQRoleResponse, err error) {
    if request == nil {
        request = NewModifyRocketMQRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "ModifyRocketMQRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRocketMQRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRocketMQRoleResponse()
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
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
// This API is used to receive messages sent to a specified Partitioned Topic. It supports only Partitioned Topic type. When there are no messages in the Partitioned Topic but the API is still called, it throws a ReceiveTimeout exception.
//
// 
//
// This API is used to batch receive policies.
//
// 
//
// This API is used to provide the following three parameters:.
//
// 
//
// MaxNumMessages: The maximum number of messages returned by the Receive API each time when using BatchReceive.
//
// MaxNumBytes: the maximum size of messages returned by the Receive API in a single BatchReceive operation, in bytes.
//
// Timeout: The maximum timeout period for each Receive API call when using BatchReceive is how long, in MS.
//
// 
//
// This API is used to disable the BatchReceive feature if none of the three parameters are specified. If any one of the three parameters has a value greater than 0, BatchReceive is enabled. BatchReceive completes when reaching the threshold specified in any one of the three parameters.
//
// 
//
// Note: MaxNumMessages and MaxNumBytes are subject to the size of ReceiveQueueSize for each receipt of messages. If ReceiveQueueSize is set to 5 and MaxNumMessages is set to 10, then BatchReceive will receive at most 5 messages at once rather than 10.
//
// 
//
// 
//
// 
//
// This API is used to return multiple messages at one time.
//
// 
//
// This API is used to Split multiple messages with the special character '###', allowing the business side to use Split tools in different languages to separate messages after receiving them.
//
// Multiple MessageIDs use the special character '###' to separate with each other. After receiving the message, the business side can leverage the Split tool provided by different languages to separate different messages. (Used for filling in the necessary MessageID field information when calling the AcknowledgeMessage API.).
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
// This API is used to receive messages sent to a specified Partitioned Topic. It supports only Partitioned Topic type. When there are no messages in the Partitioned Topic but the API is still called, it throws a ReceiveTimeout exception.
//
// 
//
// This API is used to batch receive policies.
//
// 
//
// This API is used to provide the following three parameters:.
//
// 
//
// MaxNumMessages: The maximum number of messages returned by the Receive API each time when using BatchReceive.
//
// MaxNumBytes: the maximum size of messages returned by the Receive API in a single BatchReceive operation, in bytes.
//
// Timeout: The maximum timeout period for each Receive API call when using BatchReceive is how long, in MS.
//
// 
//
// This API is used to disable the BatchReceive feature if none of the three parameters are specified. If any one of the three parameters has a value greater than 0, BatchReceive is enabled. BatchReceive completes when reaching the threshold specified in any one of the three parameters.
//
// 
//
// Note: MaxNumMessages and MaxNumBytes are subject to the size of ReceiveQueueSize for each receipt of messages. If ReceiveQueueSize is set to 5 and MaxNumMessages is set to 10, then BatchReceive will receive at most 5 messages at once rather than 10.
//
// 
//
// 
//
// 
//
// This API is used to return multiple messages at one time.
//
// 
//
// This API is used to Split multiple messages with the special character '###', allowing the business side to use Split tools in different languages to separate messages after receiving them.
//
// Multiple MessageIDs use the special character '###' to separate with each other. After receiving the message, the business side can leverage the Split tool provided by different languages to separate different messages. (Used for filling in the necessary MessageID field information when calling the AcknowledgeMessage API.).
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

func NewRetryRocketMQDlqMessageRequest() (request *RetryRocketMQDlqMessageRequest) {
    request = &RetryRocketMQDlqMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "RetryRocketMQDlqMessage")
    
    
    return
}

func NewRetryRocketMQDlqMessageResponse() (response *RetryRocketMQDlqMessageResponse) {
    response = &RetryRocketMQDlqMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRocketMQDlqMessage
// Resend the RocketMQ dead letter messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) RetryRocketMQDlqMessage(request *RetryRocketMQDlqMessageRequest) (response *RetryRocketMQDlqMessageResponse, err error) {
    return c.RetryRocketMQDlqMessageWithContext(context.Background(), request)
}

// RetryRocketMQDlqMessage
// Resend the RocketMQ dead letter messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
func (c *Client) RetryRocketMQDlqMessageWithContext(ctx context.Context, request *RetryRocketMQDlqMessageRequest) (response *RetryRocketMQDlqMessageResponse, err error) {
    if request == nil {
        request = NewRetryRocketMQDlqMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "RetryRocketMQDlqMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRocketMQDlqMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryRocketMQDlqMessageResponse()
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
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SendCmqMsg(request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
    return c.SendCmqMsgWithContext(context.Background(), request)
}

// SendCmqMsg
// This API is used to send a CMQ message.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_SENDMSG = "FailedOperation.SendMsg"
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
//  FAILEDOPERATION_CLOUDSERVICE = "FailedOperation.CloudService"
//  FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"
//  FAILEDOPERATION_OPERATELATER = "FailedOperation.OperateLater"
//  FAILEDOPERATION_SENDMSG = "FailedOperation.SendMsg"
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
// This API is used to send messages through RocketMQ. It is only used for sending a small number of test messages from the console and does not provide SLA. Cloud API is subject to traffic throttling. In real business scenarios, use the RocketMQ SDK to send messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
func (c *Client) SendRocketMQMessage(request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
    return c.SendRocketMQMessageWithContext(context.Background(), request)
}

// SendRocketMQMessage
// This API is used to send messages through RocketMQ. It is only used for sending a small number of test messages from the console and does not provide SLA. Cloud API is subject to traffic throttling. In real business scenarios, use the RocketMQ SDK to send messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"
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

func NewSetRocketMQPublicAccessPointRequest() (request *SetRocketMQPublicAccessPointRequest) {
    request = &SetRocketMQPublicAccessPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmq", APIVersion, "SetRocketMQPublicAccessPoint")
    
    
    return
}

func NewSetRocketMQPublicAccessPointResponse() (response *SetRocketMQPublicAccessPointResponse) {
    response = &SetRocketMQPublicAccessPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetRocketMQPublicAccessPoint
// This API is used to enable/disable public network access, and set the security access policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetRocketMQPublicAccessPoint(request *SetRocketMQPublicAccessPointRequest) (response *SetRocketMQPublicAccessPointResponse, err error) {
    return c.SetRocketMQPublicAccessPointWithContext(context.Background(), request)
}

// SetRocketMQPublicAccessPoint
// This API is used to enable/disable public network access, and set the security access policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SetRocketMQPublicAccessPointWithContext(ctx context.Context, request *SetRocketMQPublicAccessPointRequest) (response *SetRocketMQPublicAccessPointResponse, err error) {
    if request == nil {
        request = NewSetRocketMQPublicAccessPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmq", APIVersion, "SetRocketMQPublicAccessPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetRocketMQPublicAccessPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetRocketMQPublicAccessPointResponse()
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
