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

package v20191016

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-10-16"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeDBDiagEventRequest() (request *DescribeDBDiagEventRequest) {
    request = &DescribeDBDiagEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagEvent")
    return
}

func NewDescribeDBDiagEventResponse() (response *DescribeDBDiagEventResponse) {
    response = &DescribeDBDiagEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the details of an instance exception diagnosis event.
func (c *Client) DescribeDBDiagEvent(request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagEventRequest()
    }
    response = NewDescribeDBDiagEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagHistoryRequest() (request *DescribeDBDiagHistoryRequest) {
    request = &DescribeDBDiagHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagHistory")
    return
}

func NewDescribeDBDiagHistoryResponse() (response *DescribeDBDiagHistoryResponse) {
    response = &DescribeDBDiagHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of instance diagnosis events.
func (c *Client) DescribeDBDiagHistory(request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagHistoryRequest()
    }
    response = NewDescribeDBDiagHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTimeSeriesStatsRequest() (request *DescribeSlowLogTimeSeriesStatsRequest) {
    request = &DescribeSlowLogTimeSeriesStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTimeSeriesStats")
    return
}

func NewDescribeSlowLogTimeSeriesStatsResponse() (response *DescribeSlowLogTimeSeriesStatsResponse) {
    response = &DescribeSlowLogTimeSeriesStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the slow log statistics histogram.
func (c *Client) DescribeSlowLogTimeSeriesStats(request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTimeSeriesStatsRequest()
    }
    response = NewDescribeSlowLogTimeSeriesStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTopSqlsRequest() (request *DescribeSlowLogTopSqlsRequest) {
    request = &DescribeSlowLogTopSqlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTopSqls")
    return
}

func NewDescribeSlowLogTopSqlsResponse() (response *DescribeSlowLogTopSqlsResponse) {
    response = &DescribeSlowLogTopSqlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get and sort the top slow SQL statements in a specified time period by the aggregation mode of SQL template plus schema.
func (c *Client) DescribeSlowLogTopSqls(request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTopSqlsRequest()
    }
    response = NewDescribeSlowLogTopSqlsResponse()
    err = c.Send(request, response)
    return
}
