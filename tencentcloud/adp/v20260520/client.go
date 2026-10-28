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

package v20260520

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-05-20"

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


func NewSearchKnowledgeRequest() (request *SearchKnowledgeRequest) {
    request = &SearchKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("adp", APIVersion, "SearchKnowledge")
    
    
    return
}

func NewSearchKnowledgeResponse() (response *SearchKnowledgeResponse) {
    response = &SearchKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchKnowledge
// Knowledge retrieval
func (c *Client) SearchKnowledge(request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    return c.SearchKnowledgeWithContext(context.Background(), request)
}

// SearchKnowledge
// Knowledge retrieval
func (c *Client) SearchKnowledgeWithContext(ctx context.Context, request *SearchKnowledgeRequest) (response *SearchKnowledgeResponse, err error) {
    if request == nil {
        request = NewSearchKnowledgeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "adp", APIVersion, "SearchKnowledge")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchKnowledgeResponse()
    err = c.Send(request, response)
    return
}
