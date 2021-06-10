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

const (
	// error codes for specific actions

	// CAM signature/authentication error
	AUTHFAILURE = "AuthFailure"

	// Operation failed.
	FAILEDOPERATION = "FailedOperation"

	// An internal error occurred.
	INTERNALERROR = "InternalError"

	// Invalid parameter.
	INVALIDPARAMETER = "InvalidParameter"

	// Incorrect parameter value.
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Missing parameter.
	MISSINGPARAMETER = "MissingParameter"

	// Operation denied
	OPERATIONDENIED = "OperationDenied"

	// CAM authentication error
	OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"

	// Request exceeded the frequency limit
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// Unauthorized operation.
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// Unknown parameter.
	UNKNOWNPARAMETER = "UnknownParameter"

	// Unsupported operation
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
