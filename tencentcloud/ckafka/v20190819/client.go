// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAclResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancePostRequest() (request *CreateInstancePostRequest) {
    request = &CreateInstancePostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstancePost")
    
    
    return
}

func NewCreateInstancePostResponse() (response *CreateInstancePostResponse) {
    response = &CreateInstancePostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstancePost
// This API is used to create a pay-as-you-go instance.
//
// error code that may be returned:
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
func (c *Client) CreateInstancePost(request *CreateInstancePostRequest) (response *CreateInstancePostResponse, err error) {
    return c.CreateInstancePostWithContext(context.Background(), request)
}

// CreateInstancePost
// This API is used to create a pay-as-you-go instance.
//
// error code that may be returned:
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
func (c *Client) CreateInstancePostWithContext(ctx context.Context, request *CreateInstancePostRequest) (response *CreateInstancePostResponse, err error) {
    if request == nil {
        request = NewCreateInstancePostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstancePost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancePostResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePartitionResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclResponse()
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
// This API is used to delete a monthly subscribed (prepaid) instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePre(request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    return c.DeleteInstancePreWithContext(context.Background(), request)
}

// DeleteInstancePre
// This API is used to delete a monthly subscribed (prepaid) instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePreWithContext(ctx context.Context, request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    if request == nil {
        request = NewDeleteInstancePreRequest()
    }
    
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
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppInfoRequest() (request *DescribeAppInfoRequest) {
    request = &DescribeAppInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAppInfo")
    
    
    return
}

func NewDescribeAppInfoResponse() (response *DescribeAppInfoResponse) {
    response = &DescribeAppInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppInfo
// This API is used to query the user list.
//
// error code that may be returned:
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
func (c *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
    return c.DescribeAppInfoWithContext(context.Background(), request)
}

// DescribeAppInfo
// This API is used to query the user list.
//
// error code that may be returned:
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
func (c *Client) DescribeAppInfoWithContext(ctx context.Context, request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAppInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupResponse()
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
// This API is used to get instance attributes. 
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
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    return c.DescribeInstanceAttributesWithContext(context.Background(), request)
}

// DescribeInstanceAttributes
// This API is used to get instance attributes. 
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
func (c *Client) DescribeInstanceAttributesWithContext(ctx context.Context, request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAttributesRequest()
    }
    
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
// This API is used to get the list of CKafka instances under a user account.
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
// This API is used to get the list of CKafka instances under a user account.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDetailResponse()
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
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRouteWithContext(ctx context.Context, request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteResponse()
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
// This API is used to get topic attributes.
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
// This API is used to get topic attributes.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicSyncReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicSyncReplicaResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageListByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageListByOffsetResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireCkafkaPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireCkafkaPriceResponse()
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
// This API is used to change the configurations of a prepaid instance, such as disk capacity and bandwidth.
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
// This API is used to change the configurations of a prepaid instance, such as disk capacity and bandwidth.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPasswordResponse()
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
func (c *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest) (response *SendMessageResponse, err error) {
    if request == nil {
        request = NewSendMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessageResponse()
    err = c.Send(request, response)
    return
}
