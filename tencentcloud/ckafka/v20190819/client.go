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

package v20190819

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-08-19"

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


func NewBatchCreateAclRequest() (request *BatchCreateAclRequest) {
    request = &BatchCreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchCreateAcl")
    
    
    return
}

func NewBatchCreateAclResponse() (response *BatchCreateAclResponse) {
    response = &BatchCreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateAcl
// This API is used to create ACL policies in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) BatchCreateAcl(request *BatchCreateAclRequest) (response *BatchCreateAclResponse, err error) {
    return c.BatchCreateAclWithContext(context.Background(), request)
}

// BatchCreateAcl
// This API is used to create ACL policies in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) BatchCreateAclWithContext(ctx context.Context, request *BatchCreateAclRequest) (response *BatchCreateAclResponse, err error) {
    if request == nil {
        request = NewBatchCreateAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "BatchCreateAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyGroupOffsetsRequest() (request *BatchModifyGroupOffsetsRequest) {
    request = &BatchModifyGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchModifyGroupOffsets")
    
    
    return
}

func NewBatchModifyGroupOffsetsResponse() (response *BatchModifyGroupOffsetsResponse) {
    response = &BatchModifyGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyGroupOffsets
// This API is used to batch modify consumer group offsets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
func (c *Client) BatchModifyGroupOffsets(request *BatchModifyGroupOffsetsRequest) (response *BatchModifyGroupOffsetsResponse, err error) {
    return c.BatchModifyGroupOffsetsWithContext(context.Background(), request)
}

// BatchModifyGroupOffsets
// This API is used to batch modify consumer group offsets.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
func (c *Client) BatchModifyGroupOffsetsWithContext(ctx context.Context, request *BatchModifyGroupOffsetsRequest) (response *BatchModifyGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewBatchModifyGroupOffsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "BatchModifyGroupOffsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyTopicAttributesRequest() (request *BatchModifyTopicAttributesRequest) {
    request = &BatchModifyTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchModifyTopicAttributes")
    
    
    return
}

func NewBatchModifyTopicAttributesResponse() (response *BatchModifyTopicAttributesResponse) {
    response = &BatchModifyTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchModifyTopicAttributes
// This API is used to batch set topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchModifyTopicAttributes(request *BatchModifyTopicAttributesRequest) (response *BatchModifyTopicAttributesResponse, err error) {
    return c.BatchModifyTopicAttributesWithContext(context.Background(), request)
}

// BatchModifyTopicAttributes
// This API is used to batch set topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchModifyTopicAttributesWithContext(ctx context.Context, request *BatchModifyTopicAttributesRequest) (response *BatchModifyTopicAttributesResponse, err error) {
    if request == nil {
        request = NewBatchModifyTopicAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "BatchModifyTopicAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
    request = &CreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateAcl")
    
    
    return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
    response = &CreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAcl
// This API is used to add an ACL policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
    return c.CreateAclWithContext(context.Background(), request)
}

// CreateAcl
// This API is used to add an ACL policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAclWithContext(ctx context.Context, request *CreateAclRequest) (response *CreateAclResponse, err error) {
    if request == nil {
        request = NewCreateAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAclRuleRequest() (request *CreateAclRuleRequest) {
    request = &CreateAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateAclRule")
    
    
    return
}

func NewCreateAclRuleResponse() (response *CreateAclRuleResponse) {
    response = &CreateAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAclRule
// This API shows you how to create an ACL rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAclRule(request *CreateAclRuleRequest) (response *CreateAclRuleResponse, err error) {
    return c.CreateAclRuleWithContext(context.Background(), request)
}

// CreateAclRule
// This API shows you how to create an ACL rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAclRuleWithContext(ctx context.Context, request *CreateAclRuleRequest) (response *CreateAclRuleResponse, err error) {
    if request == nil {
        request = NewCreateAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
    request = &CreateConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateConsumer")
    
    
    return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
    response = &CreateConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumer
// This API is used to create a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    return c.CreateConsumerWithContext(context.Background(), request)
}

// CreateConsumer
// This API is used to create a consumer group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateConsumerWithContext(ctx context.Context, request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    if request == nil {
        request = NewCreateConsumerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateConsumer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatahubTopicRequest() (request *CreateDatahubTopicRequest) {
    request = &CreateDatahubTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateDatahubTopic")
    
    
    return
}

func NewCreateDatahubTopicResponse() (response *CreateDatahubTopicResponse) {
    response = &CreateDatahubTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatahubTopic
// This API is used to create a DIP topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICEXIST = "InvalidParameter.TopicExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateDatahubTopic(request *CreateDatahubTopicRequest) (response *CreateDatahubTopicResponse, err error) {
    return c.CreateDatahubTopicWithContext(context.Background(), request)
}

// CreateDatahubTopic
// This API is used to create a DIP topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICEXIST = "InvalidParameter.TopicExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateDatahubTopicWithContext(ctx context.Context, request *CreateDatahubTopicRequest) (response *CreateDatahubTopicResponse, err error) {
    if request == nil {
        request = NewCreateDatahubTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateDatahubTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatahubTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatahubTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancePreRequest() (request *CreateInstancePreRequest) {
    request = &CreateInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstancePre")
    
    
    return
}

func NewCreateInstancePreResponse() (response *CreateInstancePreResponse) {
    response = &CreateInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstancePre
// This API is used to create prepaid annual and monthly instances. It only supports creating Pro Edition instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePre(request *CreateInstancePreRequest) (response *CreateInstancePreResponse, err error) {
    return c.CreateInstancePreWithContext(context.Background(), request)
}

// CreateInstancePre
// This API is used to create prepaid annual and monthly instances. It only supports creating Pro Edition instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePreWithContext(ctx context.Context, request *CreateInstancePreRequest) (response *CreateInstancePreResponse, err error) {
    if request == nil {
        request = NewCreateInstancePreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateInstancePre")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePartitionRequest() (request *CreatePartitionRequest) {
    request = &CreatePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreatePartition")
    
    
    return
}

func NewCreatePartitionResponse() (response *CreatePartitionResponse) {
    response = &CreatePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePartition
// This API is used to add a partition in a topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePartition(request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
    return c.CreatePartitionWithContext(context.Background(), request)
}

// CreatePartition
// This API is used to add a partition in a topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePartitionWithContext(ctx context.Context, request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
    if request == nil {
        request = NewCreatePartitionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreatePartition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePostPaidInstanceRequest() (request *CreatePostPaidInstanceRequest) {
    request = &CreatePostPaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreatePostPaidInstance")
    
    
    return
}

func NewCreatePostPaidInstanceResponse() (response *CreatePostPaidInstanceResponse) {
    response = &CreatePostPaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePostPaidInstance
// This API is used to replace `CreateInstancePost`  to create a pay-as-you-go instance.  You can call this API via SDK or the TencentCloud API console to create a pay-as-you-go CKafka instance,  which is an alternate option for making a purchase in the console.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePostPaidInstance(request *CreatePostPaidInstanceRequest) (response *CreatePostPaidInstanceResponse, err error) {
    return c.CreatePostPaidInstanceWithContext(context.Background(), request)
}

// CreatePostPaidInstance
// This API is used to replace `CreateInstancePost`  to create a pay-as-you-go instance.  You can call this API via SDK or the TencentCloud API console to create a pay-as-you-go CKafka instance,  which is an alternate option for making a purchase in the console.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePostPaidInstanceWithContext(ctx context.Context, request *CreatePostPaidInstanceRequest) (response *CreatePostPaidInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePostPaidInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreatePostPaidInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePostPaidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePostPaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
    request = &CreateRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateRoute")
    
    
    return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
    response = &CreateRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoute
// This API is used to add instance routes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
    return c.CreateRouteWithContext(context.Background(), request)
}

// CreateRoute
// This API is used to add instance routes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateRouteWithContext(ctx context.Context, request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
    if request == nil {
        request = NewCreateRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopic
// This API is used to create a CKafka topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_TOPICNAMEALREADYEXIST = "InvalidParameterValue.TopicNameAlreadyExist"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// This API is used to create a CKafka topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_TOPICNAMEALREADYEXIST = "InvalidParameterValue.TopicNameAlreadyExist"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicIpWhiteListRequest() (request *CreateTopicIpWhiteListRequest) {
    request = &CreateTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopicIpWhiteList")
    
    
    return
}

func NewCreateTopicIpWhiteListResponse() (response *CreateTopicIpWhiteListResponse) {
    response = &CreateTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopicIpWhiteList
// This API is used to create a topic IP allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicIpWhiteList(request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
    return c.CreateTopicIpWhiteListWithContext(context.Background(), request)
}

// CreateTopicIpWhiteList
// This API is used to create a topic IP allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicIpWhiteListWithContext(ctx context.Context, request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateTopicIpWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateTopicIpWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopicIpWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// This API is used to add a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// This API is used to add a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRequest() (request *DeleteAclRequest) {
    request = &DeleteAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAcl")
    
    
    return
}

func NewDeleteAclResponse() (response *DeleteAclResponse) {
    response = &DeleteAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAcl
// This API is used to delete an ACL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAcl(request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
    return c.DeleteAclWithContext(context.Background(), request)
}

// DeleteAcl
// This API is used to delete an ACL.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclWithContext(ctx context.Context, request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
    if request == nil {
        request = NewDeleteAclRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteAcl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRuleRequest() (request *DeleteAclRuleRequest) {
    request = &DeleteAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAclRule")
    
    
    return
}

func NewDeleteAclRuleResponse() (response *DeleteAclRuleResponse) {
    response = &DeleteAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAclRule
// This API is used to delete an ACL rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclRule(request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
    return c.DeleteAclRuleWithContext(context.Background(), request)
}

// DeleteAclRule
// This API is used to delete an ACL rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclRuleWithContext(ctx context.Context, request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// Delete consumer groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// Delete consumer groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupSubscribeTopicRequest() (request *DeleteGroupSubscribeTopicRequest) {
    request = &DeleteGroupSubscribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteGroupSubscribeTopic")
    
    
    return
}

func NewDeleteGroupSubscribeTopicResponse() (response *DeleteGroupSubscribeTopicResponse) {
    response = &DeleteGroupSubscribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroupSubscribeTopic
// This API is used to delete topics subscribed by a consumption group. The consumption group status must be Empty.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroupSubscribeTopic(request *DeleteGroupSubscribeTopicRequest) (response *DeleteGroupSubscribeTopicResponse, err error) {
    return c.DeleteGroupSubscribeTopicWithContext(context.Background(), request)
}

// DeleteGroupSubscribeTopic
// This API is used to delete topics subscribed by a consumption group. The consumption group status must be Empty.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroupSubscribeTopicWithContext(ctx context.Context, request *DeleteGroupSubscribeTopicRequest) (response *DeleteGroupSubscribeTopicResponse, err error) {
    if request == nil {
        request = NewDeleteGroupSubscribeTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteGroupSubscribeTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroupSubscribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupSubscribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstancePostRequest() (request *DeleteInstancePostRequest) {
    request = &DeleteInstancePostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteInstancePost")
    
    
    return
}

func NewDeleteInstancePostResponse() (response *DeleteInstancePostResponse) {
    response = &DeleteInstancePostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstancePost
// This API is used to delete post-payment instances. It directly performs instance termination by calling API deletion without associating connectors and tasks in pre-check.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteInstancePost(request *DeleteInstancePostRequest) (response *DeleteInstancePostResponse, err error) {
    return c.DeleteInstancePostWithContext(context.Background(), request)
}

// DeleteInstancePost
// This API is used to delete post-payment instances. It directly performs instance termination by calling API deletion without associating connectors and tasks in pre-check.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteInstancePostWithContext(ctx context.Context, request *DeleteInstancePostRequest) (response *DeleteInstancePostResponse, err error) {
    if request == nil {
        request = NewDeleteInstancePostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteInstancePost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstancePost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstancePostResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstancePreRequest() (request *DeleteInstancePreRequest) {
    request = &DeleteInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteInstancePre")
    
    
    return
}

func NewDeleteInstancePreResponse() (response *DeleteInstancePreResponse) {
    response = &DeleteInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstancePre
// This API is used to delete prepaid instances. It performs isolation and deletion actions on the instance. After successful execution, the instance will be directly deleted and terminated. By calling API deletion, it directly performs instance termination without associating connectors and tasks in pre-check.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePre(request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    return c.DeleteInstancePreWithContext(context.Background(), request)
}

// DeleteInstancePre
// This API is used to delete prepaid instances. It performs isolation and deletion actions on the instance. After successful execution, the instance will be directly deleted and terminated. By calling API deletion, it directly performs instance termination without associating connectors and tasks in pre-check.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePreWithContext(ctx context.Context, request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    if request == nil {
        request = NewDeleteInstancePreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteInstancePre")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
    request = &DeleteRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteRoute")
    
    
    return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
    response = &DeleteRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoute
// This API is used to delete a route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
    return c.DeleteRouteWithContext(context.Background(), request)
}

// DeleteRoute
// This API is used to delete a route.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteRouteWithContext(ctx context.Context, request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
    if request == nil {
        request = NewDeleteRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteTriggerTimeRequest() (request *DeleteRouteTriggerTimeRequest) {
    request = &DeleteRouteTriggerTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteRouteTriggerTime")
    
    
    return
}

func NewDeleteRouteTriggerTimeResponse() (response *DeleteRouteTriggerTimeResponse) {
    response = &DeleteRouteTriggerTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRouteTriggerTime
// This API is used to modify the delayed trigger time of route deletion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRouteTriggerTime(request *DeleteRouteTriggerTimeRequest) (response *DeleteRouteTriggerTimeResponse, err error) {
    return c.DeleteRouteTriggerTimeWithContext(context.Background(), request)
}

// DeleteRouteTriggerTime
// This API is used to modify the delayed trigger time of route deletion.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRouteTriggerTimeWithContext(ctx context.Context, request *DeleteRouteTriggerTimeRequest) (response *DeleteRouteTriggerTimeResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTriggerTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteRouteTriggerTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRouteTriggerTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRouteTriggerTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// This API is used to delete a CKafka topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_FREQUENCYTOPICDELETEOPERATE = "UnsupportedOperation.FrequencyTopicDeleteOperate"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// This API is used to delete a CKafka topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_FREQUENCYTOPICDELETEOPERATE = "UnsupportedOperation.FrequencyTopicDeleteOperate"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicIpWhiteListRequest() (request *DeleteTopicIpWhiteListRequest) {
    request = &DeleteTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopicIpWhiteList")
    
    
    return
}

func NewDeleteTopicIpWhiteListResponse() (response *DeleteTopicIpWhiteListResponse) {
    response = &DeleteTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopicIpWhiteList
// This API is used to delete a topic IP allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicIpWhiteList(request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
    return c.DeleteTopicIpWhiteListWithContext(context.Background(), request)
}

// DeleteTopicIpWhiteList
// This API is used to delete a topic IP allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicIpWhiteListWithContext(ctx context.Context, request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteTopicIpWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteTopicIpWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopicIpWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// This API is used to delete a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// This API is used to delete a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DeleteUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeACLRequest() (request *DescribeACLRequest) {
    request = &DescribeACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeACL")
    
    
    return
}

func NewDescribeACLResponse() (response *DescribeACLResponse) {
    response = &DescribeACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeACL
// This API is used to enumerate ACLs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeACL(request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
    return c.DescribeACLWithContext(context.Background(), request)
}

// DescribeACL
// This API is used to enumerate ACLs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeACLWithContext(ctx context.Context, request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
    if request == nil {
        request = NewDescribeACLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeACL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAclRuleRequest() (request *DescribeAclRuleRequest) {
    request = &DescribeAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAclRule")
    
    
    return
}

func NewDescribeAclRuleResponse() (response *DescribeAclRuleResponse) {
    response = &DescribeAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAclRule
// This API is used to query the ACL rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeAclRule(request *DescribeAclRuleRequest) (response *DescribeAclRuleResponse, err error) {
    return c.DescribeAclRuleWithContext(context.Background(), request)
}

// DescribeAclRule
// This API is used to query the ACL rule list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeAclRuleWithContext(ctx context.Context, request *DescribeAclRuleRequest) (response *DescribeAclRuleResponse, err error) {
    if request == nil {
        request = NewDescribeAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCkafkaVersionRequest() (request *DescribeCkafkaVersionRequest) {
    request = &DescribeCkafkaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaVersion")
    
    
    return
}

func NewDescribeCkafkaVersionResponse() (response *DescribeCkafkaVersionResponse) {
    response = &DescribeCkafkaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCkafkaVersion
// This API is used to query instance version information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCkafkaVersion(request *DescribeCkafkaVersionRequest) (response *DescribeCkafkaVersionResponse, err error) {
    return c.DescribeCkafkaVersionWithContext(context.Background(), request)
}

// DescribeCkafkaVersion
// This API is used to query instance version information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCkafkaVersionWithContext(ctx context.Context, request *DescribeCkafkaVersionRequest) (response *DescribeCkafkaVersionResponse, err error) {
    if request == nil {
        request = NewDescribeCkafkaVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeCkafkaVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCkafkaVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCkafkaVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCkafkaZoneRequest() (request *DescribeCkafkaZoneRequest) {
    request = &DescribeCkafkaZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaZone")
    
    
    return
}

func NewDescribeCkafkaZoneResponse() (response *DescribeCkafkaZoneResponse) {
    response = &DescribeCkafkaZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCkafkaZone
// This API is used to view the AZ list of Ckafka.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCkafkaZone(request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
    return c.DescribeCkafkaZoneWithContext(context.Background(), request)
}

// DescribeCkafkaZone
// This API is used to view the AZ list of Ckafka.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCkafkaZoneWithContext(ctx context.Context, request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCkafkaZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeCkafkaZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCkafkaZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCkafkaZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
    request = &DescribeConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConsumerGroup")
    
    
    return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
    response = &DescribeConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroup
// This API is used to query consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    return c.DescribeConsumerGroupWithContext(context.Background(), request)
}

// DescribeConsumerGroup
// This API is used to query consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeConsumerGroupWithContext(ctx context.Context, request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmInfoRequest() (request *DescribeCvmInfoRequest) {
    request = &DescribeCvmInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCvmInfo")
    
    
    return
}

func NewDescribeCvmInfoResponse() (response *DescribeCvmInfoResponse) {
    response = &DescribeCvmInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCvmInfo
// This API is used to get instance information corresponding to backend CVM, including cvmId and ip. It is for Pro Edition, while Standard Edition returns empty data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCvmInfo(request *DescribeCvmInfoRequest) (response *DescribeCvmInfoResponse, err error) {
    return c.DescribeCvmInfoWithContext(context.Background(), request)
}

// DescribeCvmInfo
// This API is used to get instance information corresponding to backend CVM, including cvmId and ip. It is for Pro Edition, while Standard Edition returns empty data.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCvmInfoWithContext(ctx context.Context, request *DescribeCvmInfoRequest) (response *DescribeCvmInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCvmInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeCvmInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCvmInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCvmInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTopicRequest() (request *DescribeDatahubTopicRequest) {
    request = &DescribeDatahubTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTopic")
    
    
    return
}

func NewDescribeDatahubTopicResponse() (response *DescribeDatahubTopicResponse) {
    response = &DescribeDatahubTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatahubTopic
// This API is used to retrieve DIP topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopic(request *DescribeDatahubTopicRequest) (response *DescribeDatahubTopicResponse, err error) {
    return c.DescribeDatahubTopicWithContext(context.Background(), request)
}

// DescribeDatahubTopic
// This API is used to retrieve DIP topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopicWithContext(ctx context.Context, request *DescribeDatahubTopicRequest) (response *DescribeDatahubTopicResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeDatahubTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTopicsRequest() (request *DescribeDatahubTopicsRequest) {
    request = &DescribeDatahubTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTopics")
    
    
    return
}

func NewDescribeDatahubTopicsResponse() (response *DescribeDatahubTopicsResponse) {
    response = &DescribeDatahubTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatahubTopics
// This API is used to query the DataHub topic list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopics(request *DescribeDatahubTopicsRequest) (response *DescribeDatahubTopicsResponse, err error) {
    return c.DescribeDatahubTopicsWithContext(context.Background(), request)
}

// DescribeDatahubTopics
// This API is used to query the DataHub topic list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopicsWithContext(ctx context.Context, request *DescribeDatahubTopicsRequest) (response *DescribeDatahubTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTopicsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeDatahubTopics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroup
// This API is used to enumerate consumer groups (simplified).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    return c.DescribeGroupWithContext(context.Background(), request)
}

// DescribeGroup
// This API is used to enumerate consumer groups (simplified).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInfoRequest() (request *DescribeGroupInfoRequest) {
    request = &DescribeGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupInfo")
    
    
    return
}

func NewDescribeGroupInfoResponse() (response *DescribeGroupInfoResponse) {
    response = &DescribeGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupInfo
// This API is used to get consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupInfo(request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
    return c.DescribeGroupInfoWithContext(context.Background(), request)
}

// DescribeGroupInfo
// This API is used to get consumer group information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupInfoWithContext(ctx context.Context, request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupOffsetsRequest() (request *DescribeGroupOffsetsRequest) {
    request = &DescribeGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupOffsets")
    
    
    return
}

func NewDescribeGroupOffsetsResponse() (response *DescribeGroupOffsetsResponse) {
    response = &DescribeGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupOffsets
// This API is used to get the consumer group offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupOffsets(request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
    return c.DescribeGroupOffsetsWithContext(context.Background(), request)
}

// DescribeGroupOffsets
// This API is used to get the consumer group offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupOffsetsWithContext(ctx context.Context, request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupOffsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeGroupOffsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAttributesRequest() (request *DescribeInstanceAttributesRequest) {
    request = &DescribeInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstanceAttributes")
    
    
    return
}

func NewDescribeInstanceAttributesResponse() (response *DescribeInstanceAttributesResponse) {
    response = &DescribeInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceAttributes
// This API is used to obtain instance attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    return c.DescribeInstanceAttributesWithContext(context.Background(), request)
}

// DescribeInstanceAttributes
// This API is used to obtain instance attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstanceAttributesWithContext(ctx context.Context, request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeInstanceAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to search for a list of TDMQ CKafka instances under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to search for a list of TDMQ CKafka instances under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDetailRequest() (request *DescribeInstancesDetailRequest) {
    request = &DescribeInstancesDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstancesDetail")
    
    
    return
}

func NewDescribeInstancesDetailResponse() (response *DescribeInstancesDetailResponse) {
    response = &DescribeInstancesDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesDetail
// This API is used to get instance list details under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesDetail(request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
    return c.DescribeInstancesDetailWithContext(context.Background(), request)
}

// DescribeInstancesDetail
// This API is used to get instance list details under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesDetailWithContext(ctx context.Context, request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeInstancesDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyTypeRequest() (request *DescribeModifyTypeRequest) {
    request = &DescribeModifyTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeModifyType")
    
    
    return
}

func NewDescribeModifyTypeResponse() (response *DescribeModifyTypeResponse) {
    response = &DescribeModifyTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModifyType
// This API is used to query instance specification change types.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeModifyType(request *DescribeModifyTypeRequest) (response *DescribeModifyTypeResponse, err error) {
    return c.DescribeModifyTypeWithContext(context.Background(), request)
}

// DescribeModifyType
// This API is used to query instance specification change types.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeModifyTypeWithContext(ctx context.Context, request *DescribeModifyTypeRequest) (response *DescribeModifyTypeResponse, err error) {
    if request == nil {
        request = NewDescribeModifyTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeModifyType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModifyType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModifyTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
    request = &DescribeRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRegion")
    
    
    return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
    response = &DescribeRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegion
// This API is used to enumerate regions, and can be called only in Guangzhou.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    return c.DescribeRegionWithContext(context.Background(), request)
}

// DescribeRegion
// This API is used to enumerate regions, and can be called only in Guangzhou.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRegionWithContext(ctx context.Context, request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    if request == nil {
        request = NewDescribeRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
    request = &DescribeRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRoute")
    
    
    return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
    response = &DescribeRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoute
// This API is used to view route information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    return c.DescribeRouteWithContext(context.Background(), request)
}

// DescribeRoute
// This API is used to view route information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRouteWithContext(ctx context.Context, request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeRoute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRoutesRequest() (request *DescribeSecurityGroupRoutesRequest) {
    request = &DescribeSecurityGroupRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeSecurityGroupRoutes")
    
    
    return
}

func NewDescribeSecurityGroupRoutesResponse() (response *DescribeSecurityGroupRoutesResponse) {
    response = &DescribeSecurityGroupRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupRoutes
// This API is used to retrieve the security group route information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupRoutes(request *DescribeSecurityGroupRoutesRequest) (response *DescribeSecurityGroupRoutesResponse, err error) {
    return c.DescribeSecurityGroupRoutesWithContext(context.Background(), request)
}

// DescribeSecurityGroupRoutes
// This API is used to retrieve the security group route information list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityGroupRoutesWithContext(ctx context.Context, request *DescribeSecurityGroupRoutesRequest) (response *DescribeSecurityGroupRoutesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRoutesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeSecurityGroupRoutes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupRoutes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskStatus
// This API is used to query the task status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// This API is used to query the task status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopic
// API domain name: https://ckafka.tencentcloudapi.com
//
// This API is used to get the list of topics in a CKafka instance of a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// API domain name: https://ckafka.tencentcloudapi.com
//
// This API is used to get the list of topics in a CKafka instance of a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicWithContext(ctx context.Context, request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicAttributesRequest() (request *DescribeTopicAttributesRequest) {
    request = &DescribeTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicAttributes")
    
    
    return
}

func NewDescribeTopicAttributesResponse() (response *DescribeTopicAttributesResponse) {
    response = &DescribeTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicAttributes
// This API is used to retrieve topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicAttributes(request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
    return c.DescribeTopicAttributesWithContext(context.Background(), request)
}

// DescribeTopicAttributes
// This API is used to retrieve topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicAttributesWithContext(ctx context.Context, request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeTopicAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopicAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicDetailRequest() (request *DescribeTopicDetailRequest) {
    request = &DescribeTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicDetail")
    
    
    return
}

func NewDescribeTopicDetailResponse() (response *DescribeTopicDetailResponse) {
    response = &DescribeTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicDetail
// This API is used to get topic list details (only for call in the console).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    return c.DescribeTopicDetailWithContext(context.Background(), request)
}

// DescribeTopicDetail
// This API is used to get topic list details (only for call in the console).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicDetailWithContext(ctx context.Context, request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTopicDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopicDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicProduceConnectionRequest() (request *DescribeTopicProduceConnectionRequest) {
    request = &DescribeTopicProduceConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicProduceConnection")
    
    
    return
}

func NewDescribeTopicProduceConnectionResponse() (response *DescribeTopicProduceConnectionResponse) {
    response = &DescribeTopicProduceConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicProduceConnection
// This API is used to query the connection information of the topic producer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
func (c *Client) DescribeTopicProduceConnection(request *DescribeTopicProduceConnectionRequest) (response *DescribeTopicProduceConnectionResponse, err error) {
    return c.DescribeTopicProduceConnectionWithContext(context.Background(), request)
}

// DescribeTopicProduceConnection
// This API is used to query the connection information of the topic producer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
func (c *Client) DescribeTopicProduceConnectionWithContext(ctx context.Context, request *DescribeTopicProduceConnectionRequest) (response *DescribeTopicProduceConnectionResponse, err error) {
    if request == nil {
        request = NewDescribeTopicProduceConnectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopicProduceConnection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicProduceConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicProduceConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicSubscribeGroupRequest() (request *DescribeTopicSubscribeGroupRequest) {
    request = &DescribeTopicSubscribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicSubscribeGroup")
    
    
    return
}

func NewDescribeTopicSubscribeGroupResponse() (response *DescribeTopicSubscribeGroupResponse) {
    response = &DescribeTopicSubscribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicSubscribeGroup
// This API is used to search and subscribe the message group information of a topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicSubscribeGroup(request *DescribeTopicSubscribeGroupRequest) (response *DescribeTopicSubscribeGroupResponse, err error) {
    return c.DescribeTopicSubscribeGroupWithContext(context.Background(), request)
}

// DescribeTopicSubscribeGroup
// This API is used to search and subscribe the message group information of a topic.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicSubscribeGroupWithContext(ctx context.Context, request *DescribeTopicSubscribeGroupRequest) (response *DescribeTopicSubscribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeTopicSubscribeGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopicSubscribeGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicSubscribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicSubscribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicSyncReplicaRequest() (request *DescribeTopicSyncReplicaRequest) {
    request = &DescribeTopicSyncReplicaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicSyncReplica")
    
    
    return
}

func NewDescribeTopicSyncReplicaResponse() (response *DescribeTopicSyncReplicaResponse) {
    response = &DescribeTopicSyncReplicaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicSyncReplica
// This API is used to get the details of a synced topic replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicSyncReplica(request *DescribeTopicSyncReplicaRequest) (response *DescribeTopicSyncReplicaResponse, err error) {
    return c.DescribeTopicSyncReplicaWithContext(context.Background(), request)
}

// DescribeTopicSyncReplica
// This API is used to get the details of a synced topic replica.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicSyncReplicaWithContext(ctx context.Context, request *DescribeTopicSyncReplicaRequest) (response *DescribeTopicSyncReplicaResponse, err error) {
    if request == nil {
        request = NewDescribeTopicSyncReplicaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTopicSyncReplica")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicSyncReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicSyncReplicaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTypeInstancesRequest() (request *DescribeTypeInstancesRequest) {
    request = &DescribeTypeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTypeInstances")
    
    
    return
}

func NewDescribeTypeInstancesResponse() (response *DescribeTypeInstancesResponse) {
    response = &DescribeTypeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTypeInstances
// This API is used to search for a list of TDMQ CKafka instances of the specified type under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTypeInstances(request *DescribeTypeInstancesRequest) (response *DescribeTypeInstancesResponse, err error) {
    return c.DescribeTypeInstancesWithContext(context.Background(), request)
}

// DescribeTypeInstances
// This API is used to search for a list of TDMQ CKafka instances of the specified type under a user account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTypeInstancesWithContext(ctx context.Context, request *DescribeTypeInstancesRequest) (response *DescribeTypeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeTypeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeTypeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTypeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTypeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUser
// This API is used to query user information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// This API is used to query user information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "DescribeUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewFetchMessageByOffsetRequest() (request *FetchMessageByOffsetRequest) {
    request = &FetchMessageByOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageByOffset")
    
    
    return
}

func NewFetchMessageByOffsetResponse() (response *FetchMessageByOffsetResponse) {
    response = &FetchMessageByOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FetchMessageByOffset
// This API is used to query messages based on a specified offset position.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchMessageByOffset(request *FetchMessageByOffsetRequest) (response *FetchMessageByOffsetResponse, err error) {
    return c.FetchMessageByOffsetWithContext(context.Background(), request)
}

// FetchMessageByOffset
// This API is used to query messages based on a specified offset position.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchMessageByOffsetWithContext(ctx context.Context, request *FetchMessageByOffsetRequest) (response *FetchMessageByOffsetResponse, err error) {
    if request == nil {
        request = NewFetchMessageByOffsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "FetchMessageByOffset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageByOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewFetchMessageListByOffsetRequest() (request *FetchMessageListByOffsetRequest) {
    request = &FetchMessageListByOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageListByOffset")
    
    
    return
}

func NewFetchMessageListByOffsetResponse() (response *FetchMessageListByOffsetResponse) {
    response = &FetchMessageListByOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FetchMessageListByOffset
// This API is used to query the message list based on an offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByOffset(request *FetchMessageListByOffsetRequest) (response *FetchMessageListByOffsetResponse, err error) {
    return c.FetchMessageListByOffsetWithContext(context.Background(), request)
}

// FetchMessageListByOffset
// This API is used to query the message list based on an offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByOffsetWithContext(ctx context.Context, request *FetchMessageListByOffsetRequest) (response *FetchMessageListByOffsetResponse, err error) {
    if request == nil {
        request = NewFetchMessageListByOffsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "FetchMessageListByOffset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageListByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageListByOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewFetchMessageListByTimestampRequest() (request *FetchMessageListByTimestampRequest) {
    request = &FetchMessageListByTimestampRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageListByTimestamp")
    
    
    return
}

func NewFetchMessageListByTimestampResponse() (response *FetchMessageListByTimestampResponse) {
    response = &FetchMessageListByTimestampResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FetchMessageListByTimestamp
// This API is used to query a message list by timestamp.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByTimestamp(request *FetchMessageListByTimestampRequest) (response *FetchMessageListByTimestampResponse, err error) {
    return c.FetchMessageListByTimestampWithContext(context.Background(), request)
}

// FetchMessageListByTimestamp
// This API is used to query a message list by timestamp.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByTimestampWithContext(ctx context.Context, request *FetchMessageListByTimestampRequest) (response *FetchMessageListByTimestampResponse, err error) {
    if request == nil {
        request = NewFetchMessageListByTimestampRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "FetchMessageListByTimestamp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageListByTimestamp require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageListByTimestampResponse()
    err = c.Send(request, response)
    return
}

func NewInquireCkafkaPriceRequest() (request *InquireCkafkaPriceRequest) {
    request = &InquireCkafkaPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "InquireCkafkaPrice")
    
    
    return
}

func NewInquireCkafkaPriceResponse() (response *InquireCkafkaPriceResponse) {
    response = &InquireCkafkaPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquireCkafkaPrice
// This API is used to purchase a CKafka instance or query the instance renewal price.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquireCkafkaPrice(request *InquireCkafkaPriceRequest) (response *InquireCkafkaPriceResponse, err error) {
    return c.InquireCkafkaPriceWithContext(context.Background(), request)
}

// InquireCkafkaPrice
// This API is used to purchase a CKafka instance or query the instance renewal price.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquireCkafkaPriceWithContext(ctx context.Context, request *InquireCkafkaPriceRequest) (response *InquireCkafkaPriceResponse, err error) {
    if request == nil {
        request = NewInquireCkafkaPriceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "InquireCkafkaPrice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireCkafkaPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireCkafkaPriceResponse()
    err = c.Send(request, response)
    return
}

func NewInstanceScalingDownRequest() (request *InstanceScalingDownRequest) {
    request = &InstanceScalingDownRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "InstanceScalingDown")
    
    
    return
}

func NewInstanceScalingDownResponse() (response *InstanceScalingDownResponse) {
    response = &InstanceScalingDownResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstanceScalingDown
// This API is used to perform downsizing on a pay-as-you-go instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) InstanceScalingDown(request *InstanceScalingDownRequest) (response *InstanceScalingDownResponse, err error) {
    return c.InstanceScalingDownWithContext(context.Background(), request)
}

// InstanceScalingDown
// This API is used to perform downsizing on a pay-as-you-go instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) InstanceScalingDownWithContext(ctx context.Context, request *InstanceScalingDownRequest) (response *InstanceScalingDownResponse, err error) {
    if request == nil {
        request = NewInstanceScalingDownRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "InstanceScalingDown")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstanceScalingDown require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstanceScalingDownResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAclRuleRequest() (request *ModifyAclRuleRequest) {
    request = &ModifyAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyAclRule")
    
    
    return
}

func NewModifyAclRuleResponse() (response *ModifyAclRuleResponse) {
    response = &ModifyAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAclRule
// This API is used to modify ACL policy, currently only support whether to apply preset rules to newly-added topics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) ModifyAclRule(request *ModifyAclRuleRequest) (response *ModifyAclRuleResponse, err error) {
    return c.ModifyAclRuleWithContext(context.Background(), request)
}

// ModifyAclRule
// This API is used to modify ACL policy, currently only support whether to apply preset rules to newly-added topics.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) ModifyAclRuleWithContext(ctx context.Context, request *ModifyAclRuleRequest) (response *ModifyAclRuleResponse, err error) {
    if request == nil {
        request = NewModifyAclRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyAclRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatahubTopicRequest() (request *ModifyDatahubTopicRequest) {
    request = &ModifyDatahubTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyDatahubTopic")
    
    
    return
}

func NewModifyDatahubTopicResponse() (response *ModifyDatahubTopicResponse) {
    response = &ModifyDatahubTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDatahubTopic
// This API is used to modify DIP topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyDatahubTopic(request *ModifyDatahubTopicRequest) (response *ModifyDatahubTopicResponse, err error) {
    return c.ModifyDatahubTopicWithContext(context.Background(), request)
}

// ModifyDatahubTopic
// This API is used to modify DIP topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyDatahubTopicWithContext(ctx context.Context, request *ModifyDatahubTopicRequest) (response *ModifyDatahubTopicResponse, err error) {
    if request == nil {
        request = NewModifyDatahubTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyDatahubTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatahubTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatahubTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupOffsetsRequest() (request *ModifyGroupOffsetsRequest) {
    request = &ModifyGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyGroupOffsets")
    
    
    return
}

func NewModifyGroupOffsetsResponse() (response *ModifyGroupOffsetsResponse) {
    response = &ModifyGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGroupOffsets
// This API is used to set the consumer group (Groups) offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyGroupOffsets(request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
    return c.ModifyGroupOffsetsWithContext(context.Background(), request)
}

// ModifyGroupOffsets
// This API is used to set the consumer group (Groups) offset.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyGroupOffsetsWithContext(ctx context.Context, request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewModifyGroupOffsetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyGroupOffsets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAttributesRequest() (request *ModifyInstanceAttributesRequest) {
    request = &ModifyInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstanceAttributes")
    
    
    return
}

func NewModifyInstanceAttributesResponse() (response *ModifyInstanceAttributesResponse) {
    response = &ModifyInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAttributes
// This API is used to set instance attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstanceAttributes(request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
    return c.ModifyInstanceAttributesWithContext(context.Background(), request)
}

// ModifyInstanceAttributes
// This API is used to set instance attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstanceAttributesWithContext(ctx context.Context, request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyInstanceAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancePreRequest() (request *ModifyInstancePreRequest) {
    request = &ModifyInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstancePre")
    
    
    return
}

func NewModifyInstancePreResponse() (response *ModifyInstancePreResponse) {
    response = &ModifyInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstancePre
// This API is used to change the configuration of prepaid instances, adjust disks, modify bandwidth, and manage partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstancePre(request *ModifyInstancePreRequest) (response *ModifyInstancePreResponse, err error) {
    return c.ModifyInstancePreWithContext(context.Background(), request)
}

// ModifyInstancePre
// This API is used to change the configuration of prepaid instances, adjust disks, modify bandwidth, and manage partitions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstancePreWithContext(ctx context.Context, request *ModifyInstancePreRequest) (response *ModifyInstancePreResponse, err error) {
    if request == nil {
        request = NewModifyInstancePreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyInstancePre")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
    request = &ModifyPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyPassword")
    
    
    return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
    response = &ModifyPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPassword
// This API is used to change the password.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
    return c.ModifyPasswordWithContext(context.Background(), request)
}

// ModifyPassword
// This API is used to change the password.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyPasswordWithContext(ctx context.Context, request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
    if request == nil {
        request = NewModifyPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoutineMaintenanceTaskRequest() (request *ModifyRoutineMaintenanceTaskRequest) {
    request = &ModifyRoutineMaintenanceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyRoutineMaintenanceTask")
    
    
    return
}

func NewModifyRoutineMaintenanceTaskResponse() (response *ModifyRoutineMaintenanceTaskResponse) {
    response = &ModifyRoutineMaintenanceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRoutineMaintenanceTask
// This API is used to set automated ops attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyRoutineMaintenanceTask(request *ModifyRoutineMaintenanceTaskRequest) (response *ModifyRoutineMaintenanceTaskResponse, err error) {
    return c.ModifyRoutineMaintenanceTaskWithContext(context.Background(), request)
}

// ModifyRoutineMaintenanceTask
// This API is used to set automated ops attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyRoutineMaintenanceTaskWithContext(ctx context.Context, request *ModifyRoutineMaintenanceTaskRequest) (response *ModifyRoutineMaintenanceTaskResponse, err error) {
    if request == nil {
        request = NewModifyRoutineMaintenanceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyRoutineMaintenanceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoutineMaintenanceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoutineMaintenanceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicAttributesRequest() (request *ModifyTopicAttributesRequest) {
    request = &ModifyTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyTopicAttributes")
    
    
    return
}

func NewModifyTopicAttributesResponse() (response *ModifyTopicAttributesResponse) {
    response = &ModifyTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTopicAttributes
// This API is used to modify topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyTopicAttributes(request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
    return c.ModifyTopicAttributesWithContext(context.Background(), request)
}

// ModifyTopicAttributes
// This API is used to modify topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyTopicAttributesWithContext(ctx context.Context, request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
    if request == nil {
        request = NewModifyTopicAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "ModifyTopicAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewSendMessageRequest() (request *SendMessageRequest) {
    request = &SendMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "SendMessage")
    
    
    return
}

func NewSendMessageResponse() (response *SendMessageResponse) {
    response = &SendMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendMessage
// This API is used to send messages through the HTTP access layer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCETASKPAUSED = "OperationDenied.ResourceTaskPaused"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_KAFKASTORAGEERROR = "ResourceUnavailable.KafkaStorageError"
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
    return c.SendMessageWithContext(context.Background(), request)
}

// SendMessage
// This API is used to send messages through the HTTP access layer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCETASKPAUSED = "OperationDenied.ResourceTaskPaused"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_KAFKASTORAGEERROR = "ResourceUnavailable.KafkaStorageError"
func (c *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest) (response *SendMessageResponse, err error) {
    if request == nil {
        request = NewSendMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "SendMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessageResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeBrokerVersionRequest() (request *UpgradeBrokerVersionRequest) {
    request = &UpgradeBrokerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "UpgradeBrokerVersion")
    
    
    return
}

func NewUpgradeBrokerVersionResponse() (response *UpgradeBrokerVersionResponse) {
    response = &UpgradeBrokerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeBrokerVersion
// This API is used to upgrade the broker version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpgradeBrokerVersion(request *UpgradeBrokerVersionRequest) (response *UpgradeBrokerVersionResponse, err error) {
    return c.UpgradeBrokerVersionWithContext(context.Background(), request)
}

// UpgradeBrokerVersion
// This API is used to upgrade the broker version.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpgradeBrokerVersionWithContext(ctx context.Context, request *UpgradeBrokerVersionRequest) (response *UpgradeBrokerVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeBrokerVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ckafka", APIVersion, "UpgradeBrokerVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeBrokerVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeBrokerVersionResponse()
    err = c.Send(request, response)
    return
}
