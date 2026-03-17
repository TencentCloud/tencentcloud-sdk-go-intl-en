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

package v20250106

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2025-01-06"

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


func NewDescribeMNGMAUDataDetailRequest() (request *DescribeMNGMAUDataDetailRequest) {
    request = &DescribeMNGMAUDataDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNGMAUDataDetail")
    
    
    return
}

func NewDescribeMNGMAUDataDetailResponse() (response *DescribeMNGMAUDataDetailResponse) {
    response = &DescribeMNGMAUDataDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNGMAUDataDetail
// This API is used to retrieve the detailed mini game monthly active user data.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNGMAUDataDetail(request *DescribeMNGMAUDataDetailRequest) (response *DescribeMNGMAUDataDetailResponse, err error) {
    return c.DescribeMNGMAUDataDetailWithContext(context.Background(), request)
}

// DescribeMNGMAUDataDetail
// This API is used to retrieve the detailed mini game monthly active user data.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNGMAUDataDetailWithContext(ctx context.Context, request *DescribeMNGMAUDataDetailRequest) (response *DescribeMNGMAUDataDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMNGMAUDataDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNGMAUDataDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNGMAUDataDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNGMAUDataDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNGMAULineChartRequest() (request *DescribeMNGMAULineChartRequest) {
    request = &DescribeMNGMAULineChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNGMAULineChart")
    
    
    return
}

func NewDescribeMNGMAULineChartResponse() (response *DescribeMNGMAULineChartResponse) {
    response = &DescribeMNGMAULineChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNGMAULineChart
// This API is used to retrieve mini game monthly active user data in a line chart format.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMNGMAULineChart(request *DescribeMNGMAULineChartRequest) (response *DescribeMNGMAULineChartResponse, err error) {
    return c.DescribeMNGMAULineChartWithContext(context.Background(), request)
}

// DescribeMNGMAULineChart
// This API is used to retrieve mini game monthly active user data in a line chart format.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMNGMAULineChartWithContext(ctx context.Context, request *DescribeMNGMAULineChartRequest) (response *DescribeMNGMAULineChartResponse, err error) {
    if request == nil {
        request = NewDescribeMNGMAULineChartRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNGMAULineChart")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNGMAULineChart require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNGMAULineChartResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNGMAUMonthlyComparisonMetricCardRequest() (request *DescribeMNGMAUMonthlyComparisonMetricCardRequest) {
    request = &DescribeMNGMAUMonthlyComparisonMetricCardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNGMAUMonthlyComparisonMetricCard")
    
    
    return
}

func NewDescribeMNGMAUMonthlyComparisonMetricCardResponse() (response *DescribeMNGMAUMonthlyComparisonMetricCardResponse) {
    response = &DescribeMNGMAUMonthlyComparisonMetricCardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNGMAUMonthlyComparisonMetricCard
// This API is used to retrieve MAU comparison data for a mini game between two months.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNGMAUMonthlyComparisonMetricCard(request *DescribeMNGMAUMonthlyComparisonMetricCardRequest) (response *DescribeMNGMAUMonthlyComparisonMetricCardResponse, err error) {
    return c.DescribeMNGMAUMonthlyComparisonMetricCardWithContext(context.Background(), request)
}

// DescribeMNGMAUMonthlyComparisonMetricCard
// This API is used to retrieve MAU comparison data for a mini game between two months.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNGMAUMonthlyComparisonMetricCardWithContext(ctx context.Context, request *DescribeMNGMAUMonthlyComparisonMetricCardRequest) (response *DescribeMNGMAUMonthlyComparisonMetricCardResponse, err error) {
    if request == nil {
        request = NewDescribeMNGMAUMonthlyComparisonMetricCardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNGMAUMonthlyComparisonMetricCard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNGMAUMonthlyComparisonMetricCard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNGMAUMonthlyComparisonMetricCardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPMAUDataDetailRequest() (request *DescribeMNPMAUDataDetailRequest) {
    request = &DescribeMNPMAUDataDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNPMAUDataDetail")
    
    
    return
}

func NewDescribeMNPMAUDataDetailResponse() (response *DescribeMNPMAUDataDetailResponse) {
    response = &DescribeMNPMAUDataDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPMAUDataDetail
// This API is used to retrieve the detailed mini program monthly active user data.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPMAUDataDetail(request *DescribeMNPMAUDataDetailRequest) (response *DescribeMNPMAUDataDetailResponse, err error) {
    return c.DescribeMNPMAUDataDetailWithContext(context.Background(), request)
}

// DescribeMNPMAUDataDetail
// This API is used to retrieve the detailed mini program monthly active user data.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPMAUDataDetailWithContext(ctx context.Context, request *DescribeMNPMAUDataDetailRequest) (response *DescribeMNPMAUDataDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMNPMAUDataDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNPMAUDataDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPMAUDataDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPMAUDataDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPMAULineChartRequest() (request *DescribeMNPMAULineChartRequest) {
    request = &DescribeMNPMAULineChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNPMAULineChart")
    
    
    return
}

func NewDescribeMNPMAULineChartResponse() (response *DescribeMNPMAULineChartResponse) {
    response = &DescribeMNPMAULineChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPMAULineChart
// This API is used to retrieve mini program monthly active user data in a line chart format.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPMAULineChart(request *DescribeMNPMAULineChartRequest) (response *DescribeMNPMAULineChartResponse, err error) {
    return c.DescribeMNPMAULineChartWithContext(context.Background(), request)
}

// DescribeMNPMAULineChart
// This API is used to retrieve mini program monthly active user data in a line chart format.
//
// error code that may be returned:
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMNPMAULineChartWithContext(ctx context.Context, request *DescribeMNPMAULineChartRequest) (response *DescribeMNPMAULineChartResponse, err error) {
    if request == nil {
        request = NewDescribeMNPMAULineChartRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNPMAULineChart")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPMAULineChart require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPMAULineChartResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMNPMAUMetricCardRequest() (request *DescribeMNPMAUMetricCardRequest) {
    request = &DescribeMNPMAUMetricCardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcsas", APIVersion, "DescribeMNPMAUMetricCard")
    
    
    return
}

func NewDescribeMNPMAUMetricCardResponse() (response *DescribeMNPMAUMetricCardResponse) {
    response = &DescribeMNPMAUMetricCardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMNPMAUMetricCard
// This API is used to retrieve MAU comparison data for a mini program between two months.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPMAUMetricCard(request *DescribeMNPMAUMetricCardRequest) (response *DescribeMNPMAUMetricCardResponse, err error) {
    return c.DescribeMNPMAUMetricCardWithContext(context.Background(), request)
}

// DescribeMNPMAUMetricCard
// This API is used to retrieve MAU comparison data for a mini program between two months.
//
// error code that may be returned:
//  FAILEDOPERATION_NOACCESSPERMISSION = "FailedOperation.NoAccessPermission"
//  FAILEDOPERATION_PACKAGEALREADYEXPIRED = "FailedOperation.PackageAlreadyExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMNPID = "InvalidParameterValue.InvalidMNPId"
//  INVALIDPARAMETERVALUE_INVALIDPLATFORMID = "InvalidParameterValue.InvalidPlatformId"
func (c *Client) DescribeMNPMAUMetricCardWithContext(ctx context.Context, request *DescribeMNPMAUMetricCardRequest) (response *DescribeMNPMAUMetricCardResponse, err error) {
    if request == nil {
        request = NewDescribeMNPMAUMetricCardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tcsas", APIVersion, "DescribeMNPMAUMetricCard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMNPMAUMetricCard require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMNPMAUMetricCardResponse()
    err = c.Send(request, response)
    return
}
