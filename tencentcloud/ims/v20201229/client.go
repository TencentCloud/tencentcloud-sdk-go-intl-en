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

package v20201229

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-29"

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


func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ims", APIVersion, "ImageModeration")
    
    
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageModeration
// This API is used to submit an image for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate IMS](https://console.cloud.tencent.com/cms/image/package) and adjust the business configuration.
//
// ### API use instructions
//
// - Go to the "[CMS console - IMS](https://console.cloud.tencent.com/cms/image/package)" to activate IMS.
//
// - This API is a paid API. For its billing mode, see [IMS Pricing](https://www.tencentcloud.com/document/product/1122/43810).
//
// 
//
// ### API feature description
//
// - It can detect images or links and recognize content that may be offensive, unsafe, or inappropriate based on the deep learning technology;
//
// - It can capture frames from or split GIF/long images for detection;
//
// - It can recognize various non-compliant scenarios, including vulgarity, law or regulation violations, pornography, and advertising;
//
// - It can detect multiple types of objects (such as object, advertising logo, and QR code) and recognize text in images based on OCR;
//
// - It allows you to customize moderation policies based on different business scenarios;
//
// - You can select image risk libraries to filter non-compliant images of custom types (currently, only blocklist configuration is supported);
//
// - You can associate account or device information when moderating an image to recognize the account or device involved.
//
// 
//
// ### API call description
//
// - Supported image file size: **< 5 MB**
//
// - Supported image file resolution: **a resolution of 256x256 or higher** is recommended; otherwise, the recognition effect may be affected.
//
// - Supported image file formats: PNG, JPG, JPEG, BMP, GIF, and WEBP.
//
// - Supported image URL transfer protocols: HTTP and HTTPS.
//
// - If you pass in the access URL of an image, you need to **limit the image download time to 3 seconds**. To ensure the stability and reliability of the image to be detected, we recommend you use Tencent Cloud COS for storage or CDN for caching.
//
// - Default API request rate limit: **100 requests/sec**. If this limit is exceeded, an error will be reported.
//
// 
//
// <div class="rno-api-explorer" style="margin-bottom:20px">
//
//     <div class="rno-api-explorer-inner">
//
//         <div class="rno-api-explorer-hd">
//
//             <div class="rno-api-explorer-title">
//
//                 </p>
//
//             </div>
//
//         </div>
//
//     </div>
//
// </div>
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCEUNAVAILABLE_MODELCALLFAILED = "ResourceUnavailable.ModelCallFailed"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    return c.ImageModerationWithContext(context.Background(), request)
}

// ImageModeration
// This API is used to submit an image for smart moderation. Before using it, you need to log in to the console with the Tencent Cloud root account [to activate IMS](https://console.cloud.tencent.com/cms/image/package) and adjust the business configuration.
//
// ### API use instructions
//
// - Go to the "[CMS console - IMS](https://console.cloud.tencent.com/cms/image/package)" to activate IMS.
//
// - This API is a paid API. For its billing mode, see [IMS Pricing](https://www.tencentcloud.com/document/product/1122/43810).
//
// 
//
// ### API feature description
//
// - It can detect images or links and recognize content that may be offensive, unsafe, or inappropriate based on the deep learning technology;
//
// - It can capture frames from or split GIF/long images for detection;
//
// - It can recognize various non-compliant scenarios, including vulgarity, law or regulation violations, pornography, and advertising;
//
// - It can detect multiple types of objects (such as object, advertising logo, and QR code) and recognize text in images based on OCR;
//
// - It allows you to customize moderation policies based on different business scenarios;
//
// - You can select image risk libraries to filter non-compliant images of custom types (currently, only blocklist configuration is supported);
//
// - You can associate account or device information when moderating an image to recognize the account or device involved.
//
// 
//
// ### API call description
//
// - Supported image file size: **< 5 MB**
//
// - Supported image file resolution: **a resolution of 256x256 or higher** is recommended; otherwise, the recognition effect may be affected.
//
// - Supported image file formats: PNG, JPG, JPEG, BMP, GIF, and WEBP.
//
// - Supported image URL transfer protocols: HTTP and HTTPS.
//
// - If you pass in the access URL of an image, you need to **limit the image download time to 3 seconds**. To ensure the stability and reliability of the image to be detected, we recommend you use Tencent Cloud COS for storage or CDN for caching.
//
// - Default API request rate limit: **100 requests/sec**. If this limit is exceeded, an error will be reported.
//
// 
//
// <div class="rno-api-explorer" style="margin-bottom:20px">
//
//     <div class="rno-api-explorer-inner">
//
//         <div class="rno-api-explorer-hd">
//
//             <div class="rno-api-explorer-title">
//
//                 </p>
//
//             </div>
//
//         </div>
//
//     </div>
//
// </div>
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCEUNAVAILABLE_MODELCALLFAILED = "ResourceUnavailable.ModelCallFailed"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageModerationWithContext(ctx context.Context, request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageModerationResponse()
    err = c.Send(request, response)
    return
}
