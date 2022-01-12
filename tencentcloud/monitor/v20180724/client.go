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

package v20180724

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

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


func NewBindingPolicyObjectRequest() (request *BindingPolicyObjectRequest) {
    request = &BindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "BindingPolicyObject")
    
    
    return
}

func NewBindingPolicyObjectResponse() (response *BindingPolicyObjectResponse) {
    response = &BindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindingPolicyObject
// This API is used to bind an alarm policy to a specific object.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObject(request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewBindingPolicyObjectRequest()
    }
    
    response = NewBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

// BindingPolicyObject
// This API is used to bind an alarm policy to a specific object.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObjectWithContext(ctx context.Context, request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewBindingPolicyObjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
    request = &CreateAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmNotice")
    
    
    return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
    response = &CreateAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAlarmNotice
// This API is used to create a notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

// CreateAlarmNotice
// This API is used to create a notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
    request = &CreateAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
    
    
    return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
    response = &CreateAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAlarmPolicy
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    
    response = NewCreateAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

// CreateAlarmPolicy
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicyWithContext(ctx context.Context, request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyGroupRequest() (request *CreatePolicyGroupRequest) {
    request = &CreatePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePolicyGroup")
    
    
    return
}

func NewCreatePolicyGroupResponse() (response *CreatePolicyGroupResponse) {
    response = &CreatePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePolicyGroup
// This API is used to add a policy group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    if request == nil {
        request = NewCreatePolicyGroupRequest()
    }
    
    response = NewCreatePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

// CreatePolicyGroup
// This API is used to add a policy group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroupWithContext(ctx context.Context, request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    if request == nil {
        request = NewCreatePolicyGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmNoticesRequest() (request *DeleteAlarmNoticesRequest) {
    request = &DeleteAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmNotices")
    
    
    return
}

func NewDeleteAlarmNoticesResponse() (response *DeleteAlarmNoticesResponse) {
    response = &DeleteAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAlarmNotices
// This API is used to delete alarm notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAlarmNotices(request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticesRequest()
    }
    
    response = NewDeleteAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

// DeleteAlarmNotices
// This API is used to delete alarm notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAlarmNoticesWithContext(ctx context.Context, request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
    request = &DeleteAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
    
    
    return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
    response = &DeleteAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAlarmPolicy
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    
    response = NewDeleteAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

// DeleteAlarmPolicy
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyGroupRequest() (request *DeletePolicyGroupRequest) {
    request = &DeletePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePolicyGroup")
    
    
    return
}

func NewDeletePolicyGroupResponse() (response *DeletePolicyGroupResponse) {
    response = &DeletePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePolicyGroup
// This API is used to delete an alarm policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    if request == nil {
        request = NewDeletePolicyGroupRequest()
    }
    
    response = NewDeletePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

// DeletePolicyGroup
// This API is used to delete an alarm policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroupWithContext(ctx context.Context, request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    if request == nil {
        request = NewDeletePolicyGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccidentEventListRequest() (request *DescribeAccidentEventListRequest) {
    request = &DescribeAccidentEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAccidentEventList")
    
    
    return
}

func NewDescribeAccidentEventListResponse() (response *DescribeAccidentEventListResponse) {
    response = &DescribeAccidentEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccidentEventList
// This API is used to get the platform event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventList(request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentEventListRequest()
    }
    
    response = NewDescribeAccidentEventListResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccidentEventList
// This API is used to get the platform event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventListWithContext(ctx context.Context, request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentEventListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccidentEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmEventsRequest() (request *DescribeAlarmEventsRequest) {
    request = &DescribeAlarmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmEvents")
    
    
    return
}

func NewDescribeAlarmEventsResponse() (response *DescribeAlarmEventsResponse) {
    response = &DescribeAlarmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmEvents
// This API is used to query the list of alarm events.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEvents(request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmEventsRequest()
    }
    
    response = NewDescribeAlarmEventsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmEvents
// This API is used to query the list of alarm events.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEventsWithContext(ctx context.Context, request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmEventsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmHistoriesRequest() (request *DescribeAlarmHistoriesRequest) {
    request = &DescribeAlarmHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmHistories")
    
    
    return
}

func NewDescribeAlarmHistoriesResponse() (response *DescribeAlarmHistoriesResponse) {
    response = &DescribeAlarmHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmHistories
// This API is used to query the alarm records.
//
// 
//
// Note: **If you query the alarm records as a sub-user, you can only query those of authorized projects** or those of products which are not categorized by projects. For information on how to grant a sub-account the project permission, see [Project & Tag](https://intl.cloud.tencent.com/document/product/598/32738?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistories(request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmHistoriesRequest()
    }
    
    response = NewDescribeAlarmHistoriesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmHistories
// This API is used to query the alarm records.
//
// 
//
// Note: **If you query the alarm records as a sub-user, you can only query those of authorized projects** or those of products which are not categorized by projects. For information on how to grant a sub-account the project permission, see [Project & Tag](https://intl.cloud.tencent.com/document/product/598/32738?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistoriesWithContext(ctx context.Context, request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmHistoriesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmMetricsRequest() (request *DescribeAlarmMetricsRequest) {
    request = &DescribeAlarmMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmMetrics")
    
    
    return
}

func NewDescribeAlarmMetricsResponse() (response *DescribeAlarmMetricsResponse) {
    response = &DescribeAlarmMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmMetrics
// This API is used to query the list of alarm metrics.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetrics(request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmMetricsRequest()
    }
    
    response = NewDescribeAlarmMetricsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmMetrics
// This API is used to query the list of alarm metrics.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetricsWithContext(ctx context.Context, request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmMetricsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeRequest() (request *DescribeAlarmNoticeRequest) {
    request = &DescribeAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotice")
    
    
    return
}

func NewDescribeAlarmNoticeResponse() (response *DescribeAlarmNoticeResponse) {
    response = &DescribeAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmNotice
// This API is used to query the details of a single notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotice(request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeRequest()
    }
    
    response = NewDescribeAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmNotice
// This API is used to query the details of a single notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeWithContext(ctx context.Context, request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeCallbacksRequest() (request *DescribeAlarmNoticeCallbacksRequest) {
    request = &DescribeAlarmNoticeCallbacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNoticeCallbacks")
    
    
    return
}

func NewDescribeAlarmNoticeCallbacksResponse() (response *DescribeAlarmNoticeCallbacksResponse) {
    response = &DescribeAlarmNoticeCallbacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmNoticeCallbacks
// This API is used to get all the callback URLs of an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacks(request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeCallbacksRequest()
    }
    
    response = NewDescribeAlarmNoticeCallbacksResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmNoticeCallbacks
// This API is used to get all the callback URLs of an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacksWithContext(ctx context.Context, request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeCallbacksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeCallbacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
    request = &DescribeAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotices")
    
    
    return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
    response = &DescribeAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmNotices
// This API is used to query the list of notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmNotices
// This API is used to query the list of notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPoliciesRequest() (request *DescribeAlarmPoliciesRequest) {
    request = &DescribeAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicies")
    
    
    return
}

func NewDescribeAlarmPoliciesResponse() (response *DescribeAlarmPoliciesResponse) {
    response = &DescribeAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmPolicies
// This API is used to query the list of alarm policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    
    response = NewDescribeAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmPolicies
// This API is used to query the list of alarm policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPoliciesWithContext(ctx context.Context, request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
    request = &DescribeAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicy")
    
    
    return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
    response = &DescribeAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlarmPolicy
// This API is used to get the details of a single alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    
    response = NewDescribeAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

// DescribeAlarmPolicy
// This API is used to get the details of a single alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicyWithContext(ctx context.Context, request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllNamespacesRequest() (request *DescribeAllNamespacesRequest) {
    request = &DescribeAllNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAllNamespaces")
    
    
    return
}

func NewDescribeAllNamespacesResponse() (response *DescribeAllNamespacesResponse) {
    response = &DescribeAllNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllNamespaces
// This API is used to query all namespaces.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespaces(request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAllNamespacesRequest()
    }
    
    response = NewDescribeAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAllNamespaces
// This API is used to query all namespaces.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespacesWithContext(ctx context.Context, request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAllNamespacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
    request = &DescribeBaseMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
    
    
    return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
    response = &DescribeBaseMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBaseMetrics
// This API is used to get the attributes of basic metrics.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsRequest()
    }
    
    response = NewDescribeBaseMetricsResponse()
    err = c.Send(request, response)
    return
}

// DescribeBaseMetrics
// This API is used to get the attributes of basic metrics.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetricsWithContext(ctx context.Context, request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBaseMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicAlarmListRequest() (request *DescribeBasicAlarmListRequest) {
    request = &DescribeBasicAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBasicAlarmList")
    
    
    return
}

func NewDescribeBasicAlarmListResponse() (response *DescribeBasicAlarmListResponse) {
    response = &DescribeBasicAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicAlarmList
// This API is used to get the basic alarm list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmList(request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAlarmListRequest()
    }
    
    response = NewDescribeBasicAlarmListResponse()
    err = c.Send(request, response)
    return
}

// DescribeBasicAlarmList
// This API is used to get the basic alarm list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmListWithContext(ctx context.Context, request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAlarmListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBasicAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindingPolicyObjectListRequest() (request *DescribeBindingPolicyObjectListRequest) {
    request = &DescribeBindingPolicyObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBindingPolicyObjectList")
    
    
    return
}

func NewDescribeBindingPolicyObjectListResponse() (response *DescribeBindingPolicyObjectListResponse) {
    response = &DescribeBindingPolicyObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBindingPolicyObjectList
// This API is used to get the bound object list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectList(request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeBindingPolicyObjectListRequest()
    }
    
    response = NewDescribeBindingPolicyObjectListResponse()
    err = c.Send(request, response)
    return
}

// DescribeBindingPolicyObjectList
// This API is used to get the bound object list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectListWithContext(ctx context.Context, request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeBindingPolicyObjectListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBindingPolicyObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorTypesRequest() (request *DescribeMonitorTypesRequest) {
    request = &DescribeMonitorTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorTypes")
    
    
    return
}

func NewDescribeMonitorTypesResponse() (response *DescribeMonitorTypesResponse) {
    response = &DescribeMonitorTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMonitorTypes
// This API is used to list all the monitor types supported by CM.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypes(request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorTypesRequest()
    }
    
    response = NewDescribeMonitorTypesResponse()
    err = c.Send(request, response)
    return
}

// DescribeMonitorTypes
// This API is used to list all the monitor types supported by CM.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypesWithContext(ctx context.Context, request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorTypesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMonitorTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyConditionListRequest() (request *DescribePolicyConditionListRequest) {
    request = &DescribePolicyConditionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyConditionList")
    
    
    return
}

func NewDescribePolicyConditionListResponse() (response *DescribePolicyConditionListResponse) {
    response = &DescribePolicyConditionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePolicyConditionList
// This API is used to get basic alarm policy conditions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionList(request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConditionListRequest()
    }
    
    response = NewDescribePolicyConditionListResponse()
    err = c.Send(request, response)
    return
}

// DescribePolicyConditionList
// This API is used to get basic alarm policy conditions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionListWithContext(ctx context.Context, request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConditionListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePolicyConditionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupInfoRequest() (request *DescribePolicyGroupInfoRequest) {
    request = &DescribePolicyGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupInfo")
    
    
    return
}

func NewDescribePolicyGroupInfoResponse() (response *DescribePolicyGroupInfoResponse) {
    response = &DescribePolicyGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePolicyGroupInfo
// This API is used to get details of a basic policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfo(request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupInfoRequest()
    }
    
    response = NewDescribePolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribePolicyGroupInfo
// This API is used to get details of a basic policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfoWithContext(ctx context.Context, request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupListRequest() (request *DescribePolicyGroupListRequest) {
    request = &DescribePolicyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupList")
    
    
    return
}

func NewDescribePolicyGroupListResponse() (response *DescribePolicyGroupListResponse) {
    response = &DescribePolicyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePolicyGroupList
// This API is used to get the list of basic policy alarm groups.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupList(request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupListRequest()
    }
    
    response = NewDescribePolicyGroupListResponse()
    err = c.Send(request, response)
    return
}

// DescribePolicyGroupList
// This API is used to get the list of basic policy alarm groups.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupListWithContext(ctx context.Context, request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductEventListRequest() (request *DescribeProductEventListRequest) {
    request = &DescribeProductEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductEventList")
    
    
    return
}

func NewDescribeProductEventListResponse() (response *DescribeProductEventListResponse) {
    response = &DescribeProductEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductEventList
// This API is used to get the list of product events by page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventList(request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    if request == nil {
        request = NewDescribeProductEventListRequest()
    }
    
    response = NewDescribeProductEventListResponse()
    err = c.Send(request, response)
    return
}

// DescribeProductEventList
// This API is used to get the list of product events by page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventListWithContext(ctx context.Context, request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    if request == nil {
        request = NewDescribeProductEventListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeProductEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticDataRequest() (request *DescribeStatisticDataRequest) {
    request = &DescribeStatisticDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeStatisticData")
    
    
    return
}

func NewDescribeStatisticDataResponse() (response *DescribeStatisticDataResponse) {
    response = &DescribeStatisticDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStatisticData
// This API is used to query monitoring data by dimension conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticData(request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDataRequest()
    }
    
    response = NewDescribeStatisticDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeStatisticData
// This API is used to query monitoring data by dimension conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticDataWithContext(ctx context.Context, request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStatisticDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
    request = &GetMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
    
    
    return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
    response = &GetMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetMonitorData
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, please use the API [DescribeStatisticData](https://intl.cloud.tencent.com/document/product/248/51845?from_cn_redirect=1).
//
// You can get the monitoring data of a Tencent Cloud service by passing in its namespace, object dimension description, and monitoring metrics.
//
// API call rate limit: 20 calls/second (1,200 calls/minute). A single request can get the data of up to 10 instances for up to 1,440 data points.
//
// If you need to call a large number of APIs to pull metrics or objects at a time, some APIs may fail to be called due to the rate limit. We suggest you evenly arrange API calls at a time granularity.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataRequest()
    }
    
    response = NewGetMonitorDataResponse()
    err = c.Send(request, response)
    return
}

// GetMonitorData
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, please use the API [DescribeStatisticData](https://intl.cloud.tencent.com/document/product/248/51845?from_cn_redirect=1).
//
// You can get the monitoring data of a Tencent Cloud service by passing in its namespace, object dimension description, and monitoring metrics.
//
// API call rate limit: 20 calls/second (1,200 calls/minute). A single request can get the data of up to 10 instances for up to 1,440 data points.
//
// If you need to call a large number of APIs to pull metrics or objects at a time, some APIs may fail to be called due to the rate limit. We suggest you evenly arrange API calls at a time granularity.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMonitorDataWithContext(ctx context.Context, request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewGetMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
    request = &ModifyAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmNotice")
    
    
    return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
    response = &ModifyAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmNotice
// This API is used to edit an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmNotice
// This API is used to edit an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyConditionRequest() (request *ModifyAlarmPolicyConditionRequest) {
    request = &ModifyAlarmPolicyConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyCondition")
    
    
    return
}

func NewModifyAlarmPolicyConditionResponse() (response *ModifyAlarmPolicyConditionResponse) {
    response = &ModifyAlarmPolicyConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmPolicyCondition
// This API is used to modify the trigger condition of an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyCondition(request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyConditionRequest()
    }
    
    response = NewModifyAlarmPolicyConditionResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmPolicyCondition
// This API is used to modify the trigger condition of an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyConditionWithContext(ctx context.Context, request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyConditionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyConditionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyInfoRequest() (request *ModifyAlarmPolicyInfoRequest) {
    request = &ModifyAlarmPolicyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyInfo")
    
    
    return
}

func NewModifyAlarmPolicyInfoResponse() (response *ModifyAlarmPolicyInfoResponse) {
    response = &ModifyAlarmPolicyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmPolicyInfo
// This API is used to edit the basic information of a v2.0 alarm policy, including policy name and remarks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfo(request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyInfoRequest()
    }
    
    response = NewModifyAlarmPolicyInfoResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmPolicyInfo
// This API is used to edit the basic information of a v2.0 alarm policy, including policy name and remarks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfoWithContext(ctx context.Context, request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyNoticeRequest() (request *ModifyAlarmPolicyNoticeRequest) {
    request = &ModifyAlarmPolicyNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyNotice")
    
    
    return
}

func NewModifyAlarmPolicyNoticeResponse() (response *ModifyAlarmPolicyNoticeResponse) {
    response = &ModifyAlarmPolicyNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmPolicyNotice
// This API is used to modify the alarm notification template bound to an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyNotice(request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyNoticeRequest()
    }
    
    response = NewModifyAlarmPolicyNoticeResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmPolicyNotice
// This API is used to modify the alarm notification template bound to an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyNoticeWithContext(ctx context.Context, request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyNoticeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyStatusRequest() (request *ModifyAlarmPolicyStatusRequest) {
    request = &ModifyAlarmPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyStatus")
    
    
    return
}

func NewModifyAlarmPolicyStatusResponse() (response *ModifyAlarmPolicyStatusResponse) {
    response = &ModifyAlarmPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmPolicyStatus
// This API is used to enable/disable an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatus(request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyStatusRequest()
    }
    
    response = NewModifyAlarmPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmPolicyStatus
// This API is used to enable/disable an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatusWithContext(ctx context.Context, request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyTasksRequest() (request *ModifyAlarmPolicyTasksRequest) {
    request = &ModifyAlarmPolicyTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyTasks")
    
    
    return
}

func NewModifyAlarmPolicyTasksResponse() (response *ModifyAlarmPolicyTasksResponse) {
    response = &ModifyAlarmPolicyTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmPolicyTasks
// This API is used to modify the task triggered by an alarm policy. The `TriggerTasks` field contains the list of triggered tasks. If an empty array is passed in for `TriggerTasks`, it indicates to unbind all the triggered tasks from this policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasks(request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyTasksRequest()
    }
    
    response = NewModifyAlarmPolicyTasksResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmPolicyTasks
// This API is used to modify the task triggered by an alarm policy. The `TriggerTasks` field contains the list of triggered tasks. If an empty array is passed in for `TriggerTasks`, it indicates to unbind all the triggered tasks from this policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasksWithContext(ctx context.Context, request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmReceiversRequest() (request *ModifyAlarmReceiversRequest) {
    request = &ModifyAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmReceivers")
    
    
    return
}

func NewModifyAlarmReceiversResponse() (response *ModifyAlarmReceiversResponse) {
    response = &ModifyAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmReceivers
// This API is used to modify alarm recipients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceivers(request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewModifyAlarmReceiversRequest()
    }
    
    response = NewModifyAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

// ModifyAlarmReceivers
// This API is used to modify alarm recipients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceiversWithContext(ctx context.Context, request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewModifyAlarmReceiversRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyGroupRequest() (request *ModifyPolicyGroupRequest) {
    request = &ModifyPolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPolicyGroup")
    
    
    return
}

func NewModifyPolicyGroupResponse() (response *ModifyPolicyGroupResponse) {
    response = &ModifyPolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPolicyGroup
// This API is used to update policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupRequest()
    }
    
    response = NewModifyPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

// ModifyPolicyGroup
// This API is used to update policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroupWithContext(ctx context.Context, request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewPutMonitorDataRequest() (request *PutMonitorDataRequest) {
    request = &PutMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "PutMonitorData")
    
    
    return
}

func NewPutMonitorDataResponse() (response *PutMonitorDataResponse) {
    response = &PutMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutMonitorData
// The default API request rate limit is 50 requests/sec.
//
// The default upper limit on metrics of a single tenant is 100.
//
// A maximum of 30 metric/value pairs can be reported at a time. When an error is returned for a request, no metrics/values in the request will be saved.
//
// 
//
// The reporting timestamp is the timestamp when you want to save the data. We recommend that you construct a timestamp at integer minutes.
//
// The time range of a timestamp is from 300 seconds before the current time to the current time.
//
// The data of the same IP metric/value pair must be reported by minute in chronological order.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMonitorData(request *PutMonitorDataRequest) (response *PutMonitorDataResponse, err error) {
    if request == nil {
        request = NewPutMonitorDataRequest()
    }
    
    response = NewPutMonitorDataResponse()
    err = c.Send(request, response)
    return
}

// PutMonitorData
// The default API request rate limit is 50 requests/sec.
//
// The default upper limit on metrics of a single tenant is 100.
//
// A maximum of 30 metric/value pairs can be reported at a time. When an error is returned for a request, no metrics/values in the request will be saved.
//
// 
//
// The reporting timestamp is the timestamp when you want to save the data. We recommend that you construct a timestamp at integer minutes.
//
// The time range of a timestamp is from 300 seconds before the current time to the current time.
//
// The data of the same IP metric/value pair must be reported by minute in chronological order.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMonitorDataWithContext(ctx context.Context, request *PutMonitorDataRequest) (response *PutMonitorDataResponse, err error) {
    if request == nil {
        request = NewPutMonitorDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewPutMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewSendCustomAlarmMsgRequest() (request *SendCustomAlarmMsgRequest) {
    request = &SendCustomAlarmMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SendCustomAlarmMsg")
    
    
    return
}

func NewSendCustomAlarmMsgResponse() (response *SendCustomAlarmMsgResponse) {
    response = &SendCustomAlarmMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendCustomAlarmMsg
// This API is used to send a custom alarm notification.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCustomAlarmMsg(request *SendCustomAlarmMsgRequest) (response *SendCustomAlarmMsgResponse, err error) {
    if request == nil {
        request = NewSendCustomAlarmMsgRequest()
    }
    
    response = NewSendCustomAlarmMsgResponse()
    err = c.Send(request, response)
    return
}

// SendCustomAlarmMsg
// This API is used to send a custom alarm notification.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCustomAlarmMsgWithContext(ctx context.Context, request *SendCustomAlarmMsgRequest) (response *SendCustomAlarmMsgResponse, err error) {
    if request == nil {
        request = NewSendCustomAlarmMsgRequest()
    }
    request.SetContext(ctx)
    
    response = NewSendCustomAlarmMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSetDefaultAlarmPolicyRequest() (request *SetDefaultAlarmPolicyRequest) {
    request = &SetDefaultAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "SetDefaultAlarmPolicy")
    
    
    return
}

func NewSetDefaultAlarmPolicyResponse() (response *SetDefaultAlarmPolicyResponse) {
    response = &SetDefaultAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetDefaultAlarmPolicy
// This API is used to set an alarm policy as the default policy in the current policy type under the current project.
//
// Alarm policies in the same type under the project will be set as non-default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicy(request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewSetDefaultAlarmPolicyRequest()
    }
    
    response = NewSetDefaultAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

// SetDefaultAlarmPolicy
// This API is used to set an alarm policy as the default policy in the current policy type under the current project.
//
// Alarm policies in the same type under the project will be set as non-default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicyWithContext(ctx context.Context, request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewSetDefaultAlarmPolicyRequest()
    }
    request.SetContext(ctx)
    
    response = NewSetDefaultAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingAllPolicyObjectRequest() (request *UnBindingAllPolicyObjectRequest) {
    request = &UnBindingAllPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingAllPolicyObject")
    
    
    return
}

func NewUnBindingAllPolicyObjectResponse() (response *UnBindingAllPolicyObjectResponse) {
    response = &UnBindingAllPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindingAllPolicyObject
// This API is used to delete all bound objects.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObject(request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingAllPolicyObjectRequest()
    }
    
    response = NewUnBindingAllPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

// UnBindingAllPolicyObject
// This API is used to delete all bound objects.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObjectWithContext(ctx context.Context, request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingAllPolicyObjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewUnBindingAllPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingPolicyObjectRequest() (request *UnBindingPolicyObjectRequest) {
    request = &UnBindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingPolicyObject")
    
    
    return
}

func NewUnBindingPolicyObjectResponse() (response *UnBindingPolicyObjectResponse) {
    response = &UnBindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindingPolicyObject
// This API is used to delete an object that is bound to a policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObject(request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingPolicyObjectRequest()
    }
    
    response = NewUnBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

// UnBindingPolicyObject
// This API is used to delete an object that is bound to a policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObjectWithContext(ctx context.Context, request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingPolicyObjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewUnBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}
