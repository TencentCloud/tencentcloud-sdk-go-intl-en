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

package v20181119

const (
	// error codes for specific actions

	// File download failed.
	FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"

	// Image decoding failed.
	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"

	// OCR failed. This error may be caused by unstable network connections,service anomalies or other issues.
	FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"

	// Unknown error.
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"

	// The service is not activated.
	FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"

	// Incorrect parameter value.
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"

	// The file is too large.
	LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"

	// The account is in arrears.
	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"

	// Exceptional billing status.
	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
)
