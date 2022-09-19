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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewDescribeAvailableCvmInstanceTypesRequest() (request *DescribeAvailableCvmInstanceTypesRequest) {
    request = &DescribeAvailableCvmInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeAvailableCvmInstanceTypes")
    
    
    return
}

func NewDescribeAvailableCvmInstanceTypesResponse() (response *DescribeAvailableCvmInstanceTypesResponse) {
    response = &DescribeAvailableCvmInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableCvmInstanceTypes
// This API is used to view the information of available CVM model configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypes(request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    return c.DescribeAvailableCvmInstanceTypesWithContext(context.Background(), request)
}

// DescribeAvailableCvmInstanceTypes
// This API is used to view the information of available CVM model configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypesWithContext(ctx context.Context, request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCvmInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableCvmInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableCvmInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmZoneInstanceConfigInfosRequest() (request *DescribeCvmZoneInstanceConfigInfosRequest) {
    request = &DescribeCvmZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeCvmZoneInstanceConfigInfos")
    
    
    return
}

func NewDescribeCvmZoneInstanceConfigInfosResponse() (response *DescribeCvmZoneInstanceConfigInfosResponse) {
    response = &DescribeCvmZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCvmZoneInstanceConfigInfos
// This API is used to get the model configuration information of the availability zone of BatchCompute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfos(request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    return c.DescribeCvmZoneInstanceConfigInfosWithContext(context.Background(), request)
}

// DescribeCvmZoneInstanceConfigInfos
// This API is used to get the model configuration information of the availability zone of BatchCompute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfosWithContext(ctx context.Context, request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCvmZoneInstanceConfigInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCvmZoneInstanceConfigInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCvmZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCategoriesRequest() (request *DescribeInstanceCategoriesRequest) {
    request = &DescribeInstanceCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeInstanceCategories")
    
    
    return
}

func NewDescribeInstanceCategoriesResponse() (response *DescribeInstanceCategoriesResponse) {
    response = &DescribeInstanceCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceCategories
// Currently, CVM instance families are classified into different category, and each category contains several instance families. This API is used to query the instance category information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategories(request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    return c.DescribeInstanceCategoriesWithContext(context.Background(), request)
}

// DescribeInstanceCategories
// Currently, CVM instance families are classified into different category, and each category contains several instance families. This API is used to query the instance category information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategoriesWithContext(ctx context.Context, request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCategoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCategories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCategoriesResponse()
    err = c.Send(request, response)
    return
}
