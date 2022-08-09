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

package v20190319

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-19"

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


func NewCreateAuditRequest() (request *CreateAuditRequest) {
    request = &CreateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAudit")
    
    
    return
}

func NewCreateAuditResponse() (response *CreateAuditResponse) {
    response = &CreateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudit
// Parameter requirements:
//
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `osBucketName` are required.
//
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
//
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
//
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEBUCKETFAIL = "FailedOperation.CreateBucketFail"
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_CREATEAUDITERROR = "InternalError.CreateAuditError"
//  INVALIDPARAMETERVALUE_AUDITNAMEERROR = "InvalidParameterValue.AuditNameError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_ISCREATENEWBUCKETERROR = "InvalidParameterValue.IsCreateNewBucketError"
//  INVALIDPARAMETERVALUE_ISCREATENEWQUEUEERROR = "InvalidParameterValue.IsCreateNewQueueError"
//  INVALIDPARAMETERVALUE_ISENABLECMQNOTIFYERROR = "InvalidParameterValue.IsEnableCmqNotifyError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  MISSINGPARAMETER_MISSAUDITNAME = "MissingParameter.MissAuditName"
//  MISSINGPARAMETER_MISSCOSBUCKETNAME = "MissingParameter.MissCosBucketName"
//  MISSINGPARAMETER_MISSCOSREGION = "MissingParameter.MissCosRegion"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDIT = "ResourceInUse.AlreadyExistsSameAudit"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCMQCONFIG = "ResourceInUse.AlreadyExistsSameAuditCmqConfig"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCOSCONFIG = "ResourceInUse.AlreadyExistsSameAuditCosConfig"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
func (c *Client) CreateAudit(request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    return c.CreateAuditWithContext(context.Background(), request)
}

// CreateAudit
// Parameter requirements:
//
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `osBucketName` are required.
//
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
//
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
//
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEBUCKETFAIL = "FailedOperation.CreateBucketFail"
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_CREATEAUDITERROR = "InternalError.CreateAuditError"
//  INVALIDPARAMETERVALUE_AUDITNAMEERROR = "InvalidParameterValue.AuditNameError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_ISCREATENEWBUCKETERROR = "InvalidParameterValue.IsCreateNewBucketError"
//  INVALIDPARAMETERVALUE_ISCREATENEWQUEUEERROR = "InvalidParameterValue.IsCreateNewQueueError"
//  INVALIDPARAMETERVALUE_ISENABLECMQNOTIFYERROR = "InvalidParameterValue.IsEnableCmqNotifyError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  MISSINGPARAMETER_MISSAUDITNAME = "MissingParameter.MissAuditName"
//  MISSINGPARAMETER_MISSCOSBUCKETNAME = "MissingParameter.MissCosBucketName"
//  MISSINGPARAMETER_MISSCOSREGION = "MissingParameter.MissCosRegion"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDIT = "ResourceInUse.AlreadyExistsSameAudit"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCMQCONFIG = "ResourceInUse.AlreadyExistsSameAuditCmqConfig"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCOSCONFIG = "ResourceInUse.AlreadyExistsSameAuditCosConfig"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
func (c *Client) CreateAuditWithContext(ctx context.Context, request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    if request == nil {
        request = NewCreateAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditTrackRequest() (request *CreateAuditTrackRequest) {
    request = &CreateAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAuditTrack")
    
    
    return
}

func NewCreateAuditTrackResponse() (response *CreateAuditTrackResponse) {
    response = &CreateAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuditTrack
// This API is used to create a tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) CreateAuditTrack(request *CreateAuditTrackRequest) (response *CreateAuditTrackResponse, err error) {
    return c.CreateAuditTrackWithContext(context.Background(), request)
}

// CreateAuditTrack
// This API is used to create a tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) CreateAuditTrackWithContext(ctx context.Context, request *CreateAuditTrackRequest) (response *CreateAuditTrackResponse, err error) {
    if request == nil {
        request = NewCreateAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRequest() (request *DeleteAuditRequest) {
    request = &DeleteAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAudit")
    
    
    return
}

func NewDeleteAuditResponse() (response *DeleteAuditResponse) {
    response = &DeleteAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAudit
// This API is used to delete a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_DELETEAUDITERROR = "InternalError.DeleteAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAudit(request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    return c.DeleteAuditWithContext(context.Background(), request)
}

// DeleteAudit
// This API is used to delete a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_DELETEAUDITERROR = "InternalError.DeleteAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAuditWithContext(ctx context.Context, request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditTrackRequest() (request *DeleteAuditTrackRequest) {
    request = &DeleteAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAuditTrack")
    
    
    return
}

func NewDeleteAuditTrackResponse() (response *DeleteAuditTrackResponse) {
    response = &DeleteAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAuditTrack
// This API is used to delete a CloudAudit tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DeleteAuditTrack(request *DeleteAuditTrackRequest) (response *DeleteAuditTrackResponse, err error) {
    return c.DeleteAuditTrackWithContext(context.Background(), request)
}

// DeleteAuditTrack
// This API is used to delete a CloudAudit tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DeleteAuditTrackWithContext(ctx context.Context, request *DeleteAuditTrackRequest) (response *DeleteAuditTrackResponse, err error) {
    if request == nil {
        request = NewDeleteAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRequest() (request *DescribeAuditRequest) {
    request = &DescribeAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAudit")
    
    
    return
}

func NewDescribeAuditResponse() (response *DescribeAuditResponse) {
    response = &DescribeAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAudit
// This API is used to query the details of a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEAUDITERROR = "InternalError.DescribeAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAudit(request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    return c.DescribeAuditWithContext(context.Background(), request)
}

// DescribeAudit
// This API is used to query the details of a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEAUDITERROR = "InternalError.DescribeAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAuditWithContext(ctx context.Context, request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditTracksRequest() (request *DescribeAuditTracksRequest) {
    request = &DescribeAuditTracksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTracks")
    
    
    return
}

func NewDescribeAuditTracksResponse() (response *DescribeAuditTracksResponse) {
    response = &DescribeAuditTracksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditTracks
// This API is used to query the CloudAudit tracking set list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracks(request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    return c.DescribeAuditTracksWithContext(context.Background(), request)
}

// DescribeAuditTracks
// This API is used to query the CloudAudit tracking set list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracksWithContext(ctx context.Context, request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    if request == nil {
        request = NewDescribeAuditTracksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditTracks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditTracksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEvents
// This API is used to query CloudAudit logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// This API is used to query CloudAudit logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttributeKeyRequest() (request *GetAttributeKeyRequest) {
    request = &GetAttributeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "GetAttributeKey")
    
    
    return
}

func NewGetAttributeKeyResponse() (response *GetAttributeKeyResponse) {
    response = &GetAttributeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAttributeKey
// This API is used to query the valid values of `AttributeKey`.
//
// error code that may be returned:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
func (c *Client) GetAttributeKey(request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    return c.GetAttributeKeyWithContext(context.Background(), request)
}

// GetAttributeKey
// This API is used to query the valid values of `AttributeKey`.
//
// error code that may be returned:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
func (c *Client) GetAttributeKeyWithContext(ctx context.Context, request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    if request == nil {
        request = NewGetAttributeKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttributeKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttributeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewInquireAuditCreditRequest() (request *InquireAuditCreditRequest) {
    request = &InquireAuditCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "InquireAuditCredit")
    
    
    return
}

func NewInquireAuditCreditResponse() (response *InquireAuditCreditResponse) {
    response = &InquireAuditCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquireAuditCredit
// This API is used to query the number of tracking sets that can be created.
//
// error code that may be returned:
//  INTERNALERROR_INQUIREAUDITCREDITERROR = "InternalError.InquireAuditCreditError"
func (c *Client) InquireAuditCredit(request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    return c.InquireAuditCreditWithContext(context.Background(), request)
}

// InquireAuditCredit
// This API is used to query the number of tracking sets that can be created.
//
// error code that may be returned:
//  INTERNALERROR_INQUIREAUDITCREDITERROR = "InternalError.InquireAuditCreditError"
func (c *Client) InquireAuditCreditWithContext(ctx context.Context, request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    if request == nil {
        request = NewInquireAuditCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireAuditCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireAuditCreditResponse()
    err = c.Send(request, response)
    return
}

func NewListAuditsRequest() (request *ListAuditsRequest) {
    request = &ListAuditsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListAudits")
    
    
    return
}

func NewListAuditsResponse() (response *ListAuditsResponse) {
    response = &ListAuditsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAudits
// This API is used to query the summary of tracking sets.
//
// error code that may be returned:
//  INTERNALERROR_LISTAUDITSERROR = "InternalError.ListAuditsError"
func (c *Client) ListAudits(request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    return c.ListAuditsWithContext(context.Background(), request)
}

// ListAudits
// This API is used to query the summary of tracking sets.
//
// error code that may be returned:
//  INTERNALERROR_LISTAUDITSERROR = "InternalError.ListAuditsError"
func (c *Client) ListAuditsWithContext(ctx context.Context, request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    if request == nil {
        request = NewListAuditsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAudits require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuditsResponse()
    err = c.Send(request, response)
    return
}

func NewListCmqEnableRegionRequest() (request *ListCmqEnableRegionRequest) {
    request = &ListCmqEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCmqEnableRegion")
    
    
    return
}

func NewListCmqEnableRegionResponse() (response *ListCmqEnableRegionResponse) {
    response = &ListCmqEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListCmqEnableRegion
// This API is used to query CloudAudit-enabled CMQ AZs.
//
// error code that may be returned:
//  INTERNALERROR_LISTCMQENABLEREGIONERROR = "InternalError.ListCmqEnableRegionError"
func (c *Client) ListCmqEnableRegion(request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    return c.ListCmqEnableRegionWithContext(context.Background(), request)
}

// ListCmqEnableRegion
// This API is used to query CloudAudit-enabled CMQ AZs.
//
// error code that may be returned:
//  INTERNALERROR_LISTCMQENABLEREGIONERROR = "InternalError.ListCmqEnableRegionError"
func (c *Client) ListCmqEnableRegionWithContext(ctx context.Context, request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCmqEnableRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCmqEnableRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCmqEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListCosEnableRegionRequest() (request *ListCosEnableRegionRequest) {
    request = &ListCosEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCosEnableRegion")
    
    
    return
}

func NewListCosEnableRegionResponse() (response *ListCosEnableRegionResponse) {
    response = &ListCosEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListCosEnableRegion
// This API is used to query CloudAudit-enabled COS AZs.
//
// error code that may be returned:
//  INTERNALERROR_LISTCOSENABLEREGIONERROR = "InternalError.ListCosEnableRegionError"
func (c *Client) ListCosEnableRegion(request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    return c.ListCosEnableRegionWithContext(context.Background(), request)
}

// ListCosEnableRegion
// This API is used to query CloudAudit-enabled COS AZs.
//
// error code that may be returned:
//  INTERNALERROR_LISTCOSENABLEREGIONERROR = "InternalError.ListCosEnableRegionError"
func (c *Client) ListCosEnableRegionWithContext(ctx context.Context, request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCosEnableRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCosEnableRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCosEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewLookUpEventsRequest() (request *LookUpEventsRequest) {
    request = &LookUpEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "LookUpEvents")
    
    
    return
}

func NewLookUpEventsResponse() (response *LookUpEventsResponse) {
    response = &LookUpEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LookUpEvents
// This API is used to search for operation logs to help query relevant operation information.
//
// error code that may be returned:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER_TIME = "InvalidParameter.Time"
//  INVALIDPARAMETERVALUE_MAXRESULT = "InvalidParameterValue.MaxResult"
//  INVALIDPARAMETERVALUE_TIME = "InvalidParameterValue.Time"
//  INVALIDPARAMETERVALUE_ATTRIBUTEKEY = "InvalidParameterValue.attributeKey"
//  LIMITEXCEEDED_OVERTIME = "LimitExceeded.OverTime"
func (c *Client) LookUpEvents(request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    return c.LookUpEventsWithContext(context.Background(), request)
}

// LookUpEvents
// This API is used to search for operation logs to help query relevant operation information.
//
// error code that may be returned:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER_TIME = "InvalidParameter.Time"
//  INVALIDPARAMETERVALUE_MAXRESULT = "InvalidParameterValue.MaxResult"
//  INVALIDPARAMETERVALUE_TIME = "InvalidParameterValue.Time"
//  INVALIDPARAMETERVALUE_ATTRIBUTEKEY = "InvalidParameterValue.attributeKey"
//  LIMITEXCEEDED_OVERTIME = "LimitExceeded.OverTime"
func (c *Client) LookUpEventsWithContext(ctx context.Context, request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    if request == nil {
        request = NewLookUpEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LookUpEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewLookUpEventsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditTrackRequest() (request *ModifyAuditTrackRequest) {
    request = &ModifyAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ModifyAuditTrack")
    
    
    return
}

func NewModifyAuditTrackResponse() (response *ModifyAuditTrackResponse) {
    response = &ModifyAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAuditTrack
// This API is used to modify a CloudAudit tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) ModifyAuditTrack(request *ModifyAuditTrackRequest) (response *ModifyAuditTrackResponse, err error) {
    return c.ModifyAuditTrackWithContext(context.Background(), request)
}

// ModifyAuditTrack
// This API is used to modify a CloudAudit tracking set.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) ModifyAuditTrackWithContext(ctx context.Context, request *ModifyAuditTrackRequest) (response *ModifyAuditTrackResponse, err error) {
    if request == nil {
        request = NewModifyAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewStartLoggingRequest() (request *StartLoggingRequest) {
    request = &StartLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StartLogging")
    
    
    return
}

func NewStartLoggingResponse() (response *StartLoggingResponse) {
    response = &StartLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartLogging
// This API is used to enable a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_STARTLOGGINGERROR = "InternalError.StartLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StartLogging(request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    return c.StartLoggingWithContext(context.Background(), request)
}

// StartLogging
// This API is used to enable a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_STARTLOGGINGERROR = "InternalError.StartLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StartLoggingWithContext(ctx context.Context, request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    if request == nil {
        request = NewStartLoggingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLogging require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewStopLoggingRequest() (request *StopLoggingRequest) {
    request = &StopLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StopLogging")
    
    
    return
}

func NewStopLoggingResponse() (response *StopLoggingResponse) {
    response = &StopLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopLogging
// This API is used to disable a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_STOPLOGGINGERROR = "InternalError.StopLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StopLogging(request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    return c.StopLoggingWithContext(context.Background(), request)
}

// StopLogging
// This API is used to disable a tracking set.
//
// error code that may be returned:
//  INTERNALERROR_STOPLOGGINGERROR = "InternalError.StopLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StopLoggingWithContext(ctx context.Context, request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    if request == nil {
        request = NewStopLoggingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLogging require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuditRequest() (request *UpdateAuditRequest) {
    request = &UpdateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "UpdateAudit")
    
    
    return
}

func NewUpdateAuditResponse() (response *UpdateAuditResponse) {
    response = &UpdateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAudit
// Parameter requirements:
//
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `osBucketName` are required.
//
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
//
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
//
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
//
// error code that may be returned:
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_UPDATEAUDITERROR = "InternalError.UpdateAuditError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) UpdateAudit(request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    return c.UpdateAuditWithContext(context.Background(), request)
}

// UpdateAudit
// Parameter requirements:
//
// 1. If the value of `IsCreateNewBucket` exists, `cosRegion` and `osBucketName` are required.
//
// 2. If the value of `IsEnableCmqNotify` is 1, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` are required.
//
// 3. If the value of `IsEnableCmqNotify` is 0, `IsCreateNewQueue`, `CmqRegion`, and `CmqQueueName` cannot be passed in.
//
// 4. If the value of `IsEnableKmsEncry` is 1, `KmsRegion` and `KeyId` are required.
//
// error code that may be returned:
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_UPDATEAUDITERROR = "InternalError.UpdateAuditError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) UpdateAuditWithContext(ctx context.Context, request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    if request == nil {
        request = NewUpdateAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAuditResponse()
    err = c.Send(request, response)
    return
}
