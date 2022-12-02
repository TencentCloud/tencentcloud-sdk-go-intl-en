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

package v20180709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewDescribeAccountBalanceRequest() (request *DescribeAccountBalanceRequest) {
    request = &DescribeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAccountBalance")
    
    
    return
}

func NewDescribeAccountBalanceResponse() (response *DescribeAccountBalanceResponse) {
    response = &DescribeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountBalance
// This API is used to check the Tencent Cloud account balance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    return c.DescribeAccountBalanceWithContext(context.Background(), request)
}

// DescribeAccountBalance
// This API is used to check the Tencent Cloud account balance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalanceWithContext(ctx context.Context, request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillDetail
// This API is used to query bill details.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// This API is used to query bill details.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryRequest() (request *DescribeBillResourceSummaryRequest) {
    request = &DescribeBillResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummary")
    
    
    return
}

func NewDescribeBillResourceSummaryResponse() (response *DescribeBillResourceSummaryResponse) {
    response = &DescribeBillResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillResourceSummary
// This API is used to query bill resources summary. 
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeBillResourceSummary(request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    return c.DescribeBillResourceSummaryWithContext(context.Background(), request)
}

// DescribeBillResourceSummary
// This API is used to query bill resources summary. 
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeBillResourceSummaryWithContext(ctx context.Context, request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByPayMode
// Gets the bill summarized according to billing mode
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// Gets the bill summarized according to billing mode
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByProduct
// Gets the bill summarized according to product
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// Gets the bill summarized according to product
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProject")
    
    
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByProject
// Gets the bill summarized according to project
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

// DescribeBillSummaryByProject
// Gets the bill summarized according to project
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByRegion
// Gets the bill summarized according to region
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// Gets the bill summarized according to region
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
    request = &DescribeBillSummaryByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByTag")
    
    
    return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
    response = &DescribeBillSummaryByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByTag
// This API is used to get the cost distribution over different tags.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    return c.DescribeBillSummaryByTagWithContext(context.Background(), request)
}

// DescribeBillSummaryByTag
// This API is used to get the cost distribution over different tags.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTagWithContext(ctx context.Context, request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageCosDetailByDateRequest() (request *DescribeDosageCosDetailByDateRequest) {
    request = &DescribeDosageCosDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    
    return
}

func NewDescribeDosageCosDetailByDateResponse() (response *DescribeDosageCosDetailByDateResponse) {
    response = &DescribeDosageCosDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDosageCosDetailByDate
// This API is used to query COS usage details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDate(request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    return c.DescribeDosageCosDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageCosDetailByDate
// This API is used to query COS usage details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDateWithContext(ctx context.Context, request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageCosDetailByDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageCosDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageCosDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherInfoRequest() (request *DescribeVoucherInfoRequest) {
    request = &DescribeVoucherInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherInfo")
    
    
    return
}

func NewDescribeVoucherInfoResponse() (response *DescribeVoucherInfoResponse) {
    response = &DescribeVoucherInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVoucherInfo
// This API is used to query vouchers.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfo(request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    return c.DescribeVoucherInfoWithContext(context.Background(), request)
}

// DescribeVoucherInfo
// This API is used to query vouchers.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfoWithContext(ctx context.Context, request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherUsageDetailsRequest() (request *DescribeVoucherUsageDetailsRequest) {
    request = &DescribeVoucherUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherUsageDetails")
    
    
    return
}

func NewDescribeVoucherUsageDetailsResponse() (response *DescribeVoucherUsageDetailsResponse) {
    response = &DescribeVoucherUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVoucherUsageDetails
// This API is used to query voucher usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetails(request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    return c.DescribeVoucherUsageDetailsWithContext(context.Background(), request)
}

// DescribeVoucherUsageDetails
// This API is used to query voucher usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetailsWithContext(ctx context.Context, request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherUsageDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherUsageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherUsageDetailsResponse()
    err = c.Send(request, response)
    return
}
