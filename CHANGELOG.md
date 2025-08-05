# Changelog

### [v3.0.1251] - 2025-07-29

- fix: Auto-complete Request object initialization

### [v3.0.1152] - 2025-03-07

- fix: Resolved an issue where an HTTP request body could not be retried after being consumed
- fix: SSE message parser (removes a leading U+0020 space character from the value if it exists)

### [v3.0.926] - 2024-03-27

- feat: Use OmitBehaviour to control the omitnil behavior of JSON serialization
- fix: Backward compatibility issues with omitnil and omitempty
- feat: OIDCRoleArnProvider now supports beforeRefresh
- fix: Removed omitempty from common/json (allows omitempty tags to be added in models.go while maintaining code
  compatibility)
- fix: Made STS credentials more consistent
- feat: OIDCRoleArnProvider supports endpoint

### [v3.0.854] - 2023-12-06

- feat: UnsafeRetryOnConnectionFailure

### [v3.0.810] - 2024-03-14

- feat: Supports apigw endpoint
- feat: Supports SSE

### [v3.0.777] - 2023-09-01

- feat: The json package now supports the omitnil tag
- feat: Supports custom loggers
- fix: Fixed a bug that caused the program to crash when an HTTP request failed with the debug flag enabled
- feat: Logs the response when debug is enabled
- fix: Supports httpmock
- feat: Supports custom HTTP clients
- feat: Supports Content-Encoding: gzip
- feat: OIDC role provider
- feat: Supports skipping signing

### [v3.0.510] - 2022-08-10

- feat: HTTP proxy
- Supports custom headers
- feat: common/client adds context support
- deprecated: ClientProfile.BackupEndPoint
- fix: BaseRequest.SetPath
- feat: Supports octet-stream
- fix: Use crypto/rand for clienttoken
- feat: Supports lz4 compression for uploading logs
- fix: Changed common request to common client
- feat: Supports regional circuit breaking
- SDK enhancement: client-side retries
- feat: Sends octet-stream via common request
- feat: Sends requests via common client
- Added support for SetScheme and SetRootDomain