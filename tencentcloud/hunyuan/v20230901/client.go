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

package v20230901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-09-01"

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


func NewConvert3DFormatRequest() (request *Convert3DFormatRequest) {
    request = &Convert3DFormatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "Convert3DFormat")
    
    
    return
}

func NewConvert3DFormatResponse() (response *Convert3DFormatResponse) {
    response = &Convert3DFormatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Convert3DFormat
// After inputting a 3D model file, the 3D model file format can be switched.
func (c *Client) Convert3DFormat(request *Convert3DFormatRequest) (response *Convert3DFormatResponse, err error) {
    return c.Convert3DFormatWithContext(context.Background(), request)
}

// Convert3DFormat
// After inputting a 3D model file, the 3D model file format can be switched.
func (c *Client) Convert3DFormatWithContext(ctx context.Context, request *Convert3DFormatRequest) (response *Convert3DFormatResponse, err error) {
    if request == nil {
        request = NewConvert3DFormatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "Convert3DFormat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Convert3DFormat require credential")
    }

    request.SetContext(ctx)
    
    response = NewConvert3DFormatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribe3DSmartTopologyJobRequest() (request *Describe3DSmartTopologyJobRequest) {
    request = &Describe3DSmartTopologyJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "Describe3DSmartTopologyJob")
    
    
    return
}

func NewDescribe3DSmartTopologyJobResponse() (response *Describe3DSmartTopologyJobResponse) {
    response = &Describe3DSmartTopologyJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Describe3DSmartTopologyJob
// The SmartTopoly API uses the Polygon 1.5 model. After manually inputting a 3D high-poly model, it can generate a neat 3D model with lower polygon count.
//
// 1 concurrent is provided by default, which means 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) Describe3DSmartTopologyJob(request *Describe3DSmartTopologyJobRequest) (response *Describe3DSmartTopologyJobResponse, err error) {
    return c.Describe3DSmartTopologyJobWithContext(context.Background(), request)
}

// Describe3DSmartTopologyJob
// The SmartTopoly API uses the Polygon 1.5 model. After manually inputting a 3D high-poly model, it can generate a neat 3D model with lower polygon count.
//
// 1 concurrent is provided by default, which means 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) Describe3DSmartTopologyJobWithContext(ctx context.Context, request *Describe3DSmartTopologyJobRequest) (response *Describe3DSmartTopologyJobResponse, err error) {
    if request == nil {
        request = NewDescribe3DSmartTopologyJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "Describe3DSmartTopologyJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Describe3DSmartTopologyJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribe3DSmartTopologyJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuan3DPartJobRequest() (request *QueryHunyuan3DPartJobRequest) {
    request = &QueryHunyuan3DPartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuan3DPartJob")
    
    
    return
}

func NewQueryHunyuan3DPartJobResponse() (response *QueryHunyuan3DPartJobResponse) {
    response = &QueryHunyuan3DPartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuan3DPartJob
// This API is used to query the generation task of a component.
func (c *Client) QueryHunyuan3DPartJob(request *QueryHunyuan3DPartJobRequest) (response *QueryHunyuan3DPartJobResponse, err error) {
    return c.QueryHunyuan3DPartJobWithContext(context.Background(), request)
}

// QueryHunyuan3DPartJob
// This API is used to query the generation task of a component.
func (c *Client) QueryHunyuan3DPartJobWithContext(ctx context.Context, request *QueryHunyuan3DPartJobRequest) (response *QueryHunyuan3DPartJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuan3DPartJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "QueryHunyuan3DPartJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuan3DPartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuan3DPartJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanTo3DProJobRequest() (request *QueryHunyuanTo3DProJobRequest) {
    request = &QueryHunyuanTo3DProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanTo3DProJob")
    
    
    return
}

func NewQueryHunyuanTo3DProJobResponse() (response *QueryHunyuanTo3DProJobResponse) {
    response = &QueryHunyuanTo3DProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DProJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model and input text descriptions/images.
//
// This API is used to provide 3 concurrent tasks by default, which can process 3 submitted tasks simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) QueryHunyuanTo3DProJob(request *QueryHunyuanTo3DProJobRequest) (response *QueryHunyuanTo3DProJobResponse, err error) {
    return c.QueryHunyuanTo3DProJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DProJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model and input text descriptions/images.
//
// This API is used to provide 3 concurrent tasks by default, which can process 3 submitted tasks simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) QueryHunyuanTo3DProJobWithContext(ctx context.Context, request *QueryHunyuanTo3DProJobRequest) (response *QueryHunyuanTo3DProJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DProJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "QueryHunyuanTo3DProJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DProJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanTo3DRapidJobRequest() (request *QueryHunyuanTo3DRapidJobRequest) {
    request = &QueryHunyuanTo3DRapidJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "QueryHunyuanTo3DRapidJob")
    
    
    return
}

func NewQueryHunyuanTo3DRapidJobResponse() (response *QueryHunyuanTo3DRapidJobResponse) {
    response = &QueryHunyuanTo3DRapidJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DRapidJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model with input text descriptions or images.
//
// This API is used to provide 1 concurrent task by default, which means only 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) QueryHunyuanTo3DRapidJob(request *QueryHunyuanTo3DRapidJobRequest) (response *QueryHunyuanTo3DRapidJobResponse, err error) {
    return c.QueryHunyuanTo3DRapidJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DRapidJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model with input text descriptions or images.
//
// This API is used to provide 1 concurrent task by default, which means only 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) QueryHunyuanTo3DRapidJobWithContext(ctx context.Context, request *QueryHunyuanTo3DRapidJobRequest) (response *QueryHunyuanTo3DRapidJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DRapidJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "QueryHunyuanTo3DRapidJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DRapidJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DRapidJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmit3DSmartTopologyJobRequest() (request *Submit3DSmartTopologyJobRequest) {
    request = &Submit3DSmartTopologyJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "Submit3DSmartTopologyJob")
    
    
    return
}

func NewSubmit3DSmartTopologyJobResponse() (response *Submit3DSmartTopologyJobResponse) {
    response = &Submit3DSmartTopologyJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Submit3DSmartTopologyJob
// The SmartTopoly API uses the Polygon 1.5 model. After manually inputting a 3D high-poly model, it can generate a neat 3D model with lower polygon count.
//
// 1 concurrent is provided by default, which means 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) Submit3DSmartTopologyJob(request *Submit3DSmartTopologyJobRequest) (response *Submit3DSmartTopologyJobResponse, err error) {
    return c.Submit3DSmartTopologyJobWithContext(context.Background(), request)
}

// Submit3DSmartTopologyJob
// The SmartTopoly API uses the Polygon 1.5 model. After manually inputting a 3D high-poly model, it can generate a neat 3D model with lower polygon count.
//
// 1 concurrent is provided by default, which means 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) Submit3DSmartTopologyJobWithContext(ctx context.Context, request *Submit3DSmartTopologyJobRequest) (response *Submit3DSmartTopologyJobResponse, err error) {
    if request == nil {
        request = NewSubmit3DSmartTopologyJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "Submit3DSmartTopologyJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Submit3DSmartTopologyJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmit3DSmartTopologyJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuan3DPartJobRequest() (request *SubmitHunyuan3DPartJobRequest) {
    request = &SubmitHunyuan3DPartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuan3DPartJob")
    
    
    return
}

func NewSubmitHunyuan3DPartJobResponse() (response *SubmitHunyuan3DPartJobResponse) {
    response = &SubmitHunyuan3DPartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuan3DPartJob
// This API is used to automatically perform component identification and generation based on the model structure after inputting a 3D model file.
func (c *Client) SubmitHunyuan3DPartJob(request *SubmitHunyuan3DPartJobRequest) (response *SubmitHunyuan3DPartJobResponse, err error) {
    return c.SubmitHunyuan3DPartJobWithContext(context.Background(), request)
}

// SubmitHunyuan3DPartJob
// This API is used to automatically perform component identification and generation based on the model structure after inputting a 3D model file.
func (c *Client) SubmitHunyuan3DPartJobWithContext(ctx context.Context, request *SubmitHunyuan3DPartJobRequest) (response *SubmitHunyuan3DPartJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuan3DPartJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "SubmitHunyuan3DPartJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuan3DPartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuan3DPartJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanTo3DProJobRequest() (request *SubmitHunyuanTo3DProJobRequest) {
    request = &SubmitHunyuanTo3DProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanTo3DProJob")
    
    
    return
}

func NewSubmitHunyuanTo3DProJobResponse() (response *SubmitHunyuanTo3DProJobResponse) {
    response = &SubmitHunyuanTo3DProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DProJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model and input text descriptions/images.
//
// This API is used to provide 3 concurrent tasks by default. Up to 3 submitted tasks can be processed simultaneously. A new task can be processed only after the previous one is completed.
func (c *Client) SubmitHunyuanTo3DProJob(request *SubmitHunyuanTo3DProJobRequest) (response *SubmitHunyuanTo3DProJobResponse, err error) {
    return c.SubmitHunyuanTo3DProJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DProJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model and input text descriptions/images.
//
// This API is used to provide 3 concurrent tasks by default. Up to 3 submitted tasks can be processed simultaneously. A new task can be processed only after the previous one is completed.
func (c *Client) SubmitHunyuanTo3DProJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DProJobRequest) (response *SubmitHunyuanTo3DProJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DProJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "SubmitHunyuanTo3DProJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DProJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanTo3DRapidJobRequest() (request *SubmitHunyuanTo3DRapidJobRequest) {
    request = &SubmitHunyuanTo3DRapidJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hunyuan", APIVersion, "SubmitHunyuanTo3DRapidJob")
    
    
    return
}

func NewSubmitHunyuanTo3DRapidJobResponse() (response *SubmitHunyuanTo3DRapidJobResponse) {
    response = &SubmitHunyuanTo3DRapidJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DRapidJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model with input text descriptions or images.
//
// This API is used to provide 1 concurrent task by default, which means only 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) SubmitHunyuanTo3DRapidJob(request *SubmitHunyuanTo3DRapidJobRequest) (response *SubmitHunyuanTo3DRapidJobResponse, err error) {
    return c.SubmitHunyuanTo3DRapidJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DRapidJob
// This API is used to intelligently generate 3D content based on the HunYuan Large Model with input text descriptions or images.
//
// This API is used to provide 1 concurrent task by default, which means only 1 submitted task can be processed simultaneously. The next task can be processed only after the previous task is completed.
func (c *Client) SubmitHunyuanTo3DRapidJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DRapidJobRequest) (response *SubmitHunyuanTo3DRapidJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DRapidJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "hunyuan", APIVersion, "SubmitHunyuanTo3DRapidJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DRapidJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DRapidJobResponse()
    err = c.Send(request, response)
    return
}
