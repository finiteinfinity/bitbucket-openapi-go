/*
Bitbucket API

Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.

API version: 2.0
Contact: support@bitbucket.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// SshApiService SshApi service
type SshApiService service

type SshApiUsersSelectedUserSshKeysGetRequest struct {
	ctx context.Context
	ApiService *SshApiService
	selectedUser string
}

func (r SshApiUsersSelectedUserSshKeysGetRequest) Execute() (*PaginatedSshUserKeys, *http.Response, error) {
	return r.ApiService.UsersSelectedUserSshKeysGetExecute(r)
}

/*
UsersSelectedUserSshKeysGet List SSH keys

Returns a paginated list of the user's SSH public keys.

Example:

```
$ curl https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys
{
    "page": 1,
    "pagelen": 10,
    "size": 1,
    "values": [
        {
            "comment": "user@myhost",
            "created_on": "2018-03-14T13:17:05.196003+00:00",
            "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
            "label": "",
            "last_used": "2018-03-20T13:18:05.196003+00:00",
            "links": {
                "self": {
                    "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
                }
            },
            "owner": {
                "display_name": "Mark Adams",
                "links": {
                    "avatar": {
                        "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
                    },
                    "html": {
                        "href": "https://bitbucket.org/markadams-atl/"
                    },
                    "self": {
                        "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}"
                    }
                },
                "type": "user",
                "username": "markadams-atl",
                "nickname": "markadams-atl",
                "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
            },
            "type": "ssh_key",
            "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
        }
    ]
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param selectedUser This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
 @return SshApiUsersSelectedUserSshKeysGetRequest
*/
func (a *SshApiService) UsersSelectedUserSshKeysGet(ctx context.Context, selectedUser string) SshApiUsersSelectedUserSshKeysGetRequest {
	return SshApiUsersSelectedUserSshKeysGetRequest{
		ApiService: a,
		ctx: ctx,
		selectedUser: selectedUser,
	}
}

// Execute executes the request
//  @return PaginatedSshUserKeys
func (a *SshApiService) UsersSelectedUserSshKeysGetExecute(r SshApiUsersSelectedUserSshKeysGetRequest) (*PaginatedSshUserKeys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSshUserKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SshApiService.UsersSelectedUserSshKeysGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{selected_user}/ssh-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", url.PathEscape(parameterToString(r.selectedUser, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest struct {
	ctx context.Context
	ApiService *SshApiService
	keyId string
	selectedUser string
}

func (r SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersSelectedUserSshKeysKeyIdDeleteExecute(r)
}

/*
UsersSelectedUserSshKeysKeyIdDelete Delete a SSH key

Deletes a specific SSH public key from a user's account

Example:
```
$ curl -X DELETE https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyId The SSH key's UUID value.
 @param selectedUser This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
 @return SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest
*/
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdDelete(ctx context.Context, keyId string, selectedUser string) SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest {
	return SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
		selectedUser: selectedUser,
	}
}

// Execute executes the request
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdDeleteExecute(r SshApiUsersSelectedUserSshKeysKeyIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SshApiService.UsersSelectedUserSshKeysKeyIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{selected_user}/ssh-keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", url.PathEscape(parameterToString(r.selectedUser, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SshApiUsersSelectedUserSshKeysKeyIdGetRequest struct {
	ctx context.Context
	ApiService *SshApiService
	keyId string
	selectedUser string
}

func (r SshApiUsersSelectedUserSshKeysKeyIdGetRequest) Execute() (*SshAccountKey, *http.Response, error) {
	return r.ApiService.UsersSelectedUserSshKeysKeyIdGetExecute(r)
}

/*
UsersSelectedUserSshKeysKeyIdGet Get a SSH key

Returns a specific SSH public key belonging to a user.

Example:
```
$ curl https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/{fbe4bbab-f6f7-4dde-956b-5c58323c54b3}

{
    "comment": "user@myhost",
    "created_on": "2018-03-14T13:17:05.196003+00:00",
    "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
    "label": "",
    "last_used": "2018-03-20T13:18:05.196003+00:00",
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
        }
    },
    "owner": {
        "display_name": "Mark Adams",
        "links": {
            "avatar": {
                "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
            },
            "html": {
                "href": "https://bitbucket.org/markadams-atl/"
            },
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}"
            }
        },
        "type": "user",
        "username": "markadams-atl",
        "nickname": "markadams-atl",
        "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
    },
    "type": "ssh_key",
    "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyId The SSH key's UUID value.
 @param selectedUser This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
 @return SshApiUsersSelectedUserSshKeysKeyIdGetRequest
*/
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdGet(ctx context.Context, keyId string, selectedUser string) SshApiUsersSelectedUserSshKeysKeyIdGetRequest {
	return SshApiUsersSelectedUserSshKeysKeyIdGetRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
		selectedUser: selectedUser,
	}
}

// Execute executes the request
//  @return SshAccountKey
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdGetExecute(r SshApiUsersSelectedUserSshKeysKeyIdGetRequest) (*SshAccountKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SshAccountKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SshApiService.UsersSelectedUserSshKeysKeyIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{selected_user}/ssh-keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", url.PathEscape(parameterToString(r.selectedUser, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SshApiUsersSelectedUserSshKeysKeyIdPutRequest struct {
	ctx context.Context
	ApiService *SshApiService
	keyId string
	selectedUser string
	body *SshAccountKey
}

// The updated SSH key object
func (r SshApiUsersSelectedUserSshKeysKeyIdPutRequest) Body(body SshAccountKey) SshApiUsersSelectedUserSshKeysKeyIdPutRequest {
	r.body = &body
	return r
}

func (r SshApiUsersSelectedUserSshKeysKeyIdPutRequest) Execute() (*SshAccountKey, *http.Response, error) {
	return r.ApiService.UsersSelectedUserSshKeysKeyIdPutExecute(r)
}

/*
UsersSelectedUserSshKeysKeyIdPut Update a SSH key

Updates a specific SSH public key on a user's account

Note: Only the 'comment' field can be updated using this API. To modify the key or comment values, you must delete and add the key again.

Example:
```
$ curl -X PUT -H "Content-Type: application/json" -d '{"label": "Work key"}' https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a}

{
    "comment": "",
    "created_on": "2018-03-14T13:17:05.196003+00:00",
    "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
    "label": "Work key",
    "last_used": "2018-03-20T13:18:05.196003+00:00",
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
        }
    },
    "owner": {
        "display_name": "Mark Adams",
        "links": {
            "avatar": {
                "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
            },
            "html": {
                "href": "https://bitbucket.org/markadams-atl/"
            },
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}"
            }
        },
        "type": "user",
        "username": "markadams-atl",
        "nickname": "markadams-atl",
        "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
    },
    "type": "ssh_key",
    "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyId The SSH key's UUID value.
 @param selectedUser This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
 @return SshApiUsersSelectedUserSshKeysKeyIdPutRequest
*/
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdPut(ctx context.Context, keyId string, selectedUser string) SshApiUsersSelectedUserSshKeysKeyIdPutRequest {
	return SshApiUsersSelectedUserSshKeysKeyIdPutRequest{
		ApiService: a,
		ctx: ctx,
		keyId: keyId,
		selectedUser: selectedUser,
	}
}

// Execute executes the request
//  @return SshAccountKey
func (a *SshApiService) UsersSelectedUserSshKeysKeyIdPutExecute(r SshApiUsersSelectedUserSshKeysKeyIdPutRequest) (*SshAccountKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SshAccountKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SshApiService.UsersSelectedUserSshKeysKeyIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{selected_user}/ssh-keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", url.PathEscape(parameterToString(r.selectedUser, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SshApiUsersSelectedUserSshKeysPostRequest struct {
	ctx context.Context
	ApiService *SshApiService
	selectedUser string
	body *SshAccountKey
}

// The new SSH key object. Note that the username property has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis).
func (r SshApiUsersSelectedUserSshKeysPostRequest) Body(body SshAccountKey) SshApiUsersSelectedUserSshKeysPostRequest {
	r.body = &body
	return r
}

func (r SshApiUsersSelectedUserSshKeysPostRequest) Execute() (*SshAccountKey, *http.Response, error) {
	return r.ApiService.UsersSelectedUserSshKeysPostExecute(r)
}

/*
UsersSelectedUserSshKeysPost Add a new SSH key

Adds a new SSH public key to the specified user account and returns the resulting key.

Example:
```
$ curl -X POST -H "Content-Type: application/json" -d '{"key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY user@myhost"}' https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys

{
    "comment": "user@myhost",
    "created_on": "2018-03-14T13:17:05.196003+00:00",
    "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
    "label": "",
    "last_used": "2018-03-20T13:18:05.196003+00:00",
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
        }
    },
    "owner": {
        "display_name": "Mark Adams",
        "links": {
            "avatar": {
                "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
            },
            "html": {
                "href": "https://bitbucket.org/markadams-atl/"
            },
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}"
            }
        },
        "type": "user",
        "username": "markadams-atl",
        "nickname": "markadams-atl",
        "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
    },
    "type": "ssh_key",
    "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param selectedUser This can either be the UUID of the account, surrounded by curly-braces, for example: `{account UUID}`, OR an Atlassian Account ID. 
 @return SshApiUsersSelectedUserSshKeysPostRequest
*/
func (a *SshApiService) UsersSelectedUserSshKeysPost(ctx context.Context, selectedUser string) SshApiUsersSelectedUserSshKeysPostRequest {
	return SshApiUsersSelectedUserSshKeysPostRequest{
		ApiService: a,
		ctx: ctx,
		selectedUser: selectedUser,
	}
}

// Execute executes the request
//  @return SshAccountKey
func (a *SshApiService) UsersSelectedUserSshKeysPostExecute(r SshApiUsersSelectedUserSshKeysPostRequest) (*SshAccountKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SshAccountKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SshApiService.UsersSelectedUserSshKeysPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{selected_user}/ssh-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", url.PathEscape(parameterToString(r.selectedUser, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}