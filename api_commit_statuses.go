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


// CommitStatusesApiService CommitStatusesApi service
type CommitStatusesApiService service

type CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest struct {
	ctx context.Context
	ApiService *CommitStatusesApiService
	commit string
	key string
	repoSlug string
	workspace string
}

func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest) Execute() (*Commitstatus, *http.Response, error) {
	return r.ApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetExecute(r)
}

/*
RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet Get a build status for a commit

Returns the specified build status for a commit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commit The commit's SHA1.
 @param key The build status' unique key
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
 @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
 @return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest
*/
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet(ctx context.Context, commit string, key string, repoSlug string, workspace string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest {
	return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest{
		ApiService: a,
		ctx: ctx,
		commit: commit,
		key: key,
		repoSlug: repoSlug,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return Commitstatus
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetExecute(r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGetRequest) (*Commitstatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Commitstatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommitStatusesApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", url.PathEscape(parameterToString(r.commit, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterToString(r.key, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", url.PathEscape(parameterToString(r.repoSlug, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterToString(r.workspace, "")), -1)

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

type CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest struct {
	ctx context.Context
	ApiService *CommitStatusesApiService
	commit string
	key string
	repoSlug string
	workspace string
	body *Commitstatus
}

// The updated build status object
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest) Body(body Commitstatus) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest {
	r.body = &body
	return r
}

func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest) Execute() (*Commitstatus, *http.Response, error) {
	return r.ApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutExecute(r)
}

/*
RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut Update a build status for a commit

Used to update the current status of a build status object on the
specific commit.

This operation can also be used to change other properties of the
build status:

* `state`
* `name`
* `description`
* `url`
* `refname`

The `key` cannot be changed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commit The commit's SHA1.
 @param key The build status' unique key
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
 @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
 @return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest
*/
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut(ctx context.Context, commit string, key string, repoSlug string, workspace string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest {
	return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest{
		ApiService: a,
		ctx: ctx,
		commit: commit,
		key: key,
		repoSlug: repoSlug,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return Commitstatus
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutExecute(r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutRequest) (*Commitstatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Commitstatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommitStatusesApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", url.PathEscape(parameterToString(r.commit, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterToString(r.key, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", url.PathEscape(parameterToString(r.repoSlug, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterToString(r.workspace, "")), -1)

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

type CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest struct {
	ctx context.Context
	ApiService *CommitStatusesApiService
	commit string
	repoSlug string
	workspace string
	body *Commitstatus
}

// The new commit status object.
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest) Body(body Commitstatus) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest {
	r.body = &body
	return r
}

func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest) Execute() (*Commitstatus, *http.Response, error) {
	return r.ApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostExecute(r)
}

/*
RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost Create a build status for a commit

Creates a new build status against the specified commit.

If the specified key already exists, the existing status object will
be overwritten.

Example:

```
curl https://api.bitbucket.org/2.0/repositories/my-workspace/my-repo/commit/e10dae226959c2194f2b07b077c07762d93821cf/statuses/build/           -X POST -u jdoe -H 'Content-Type: application/json'           -d '{
    "key": "MY-BUILD",
    "state": "SUCCESSFUL",
    "description": "42 tests passed",
    "url": "https://www.example.org/my-build-result"
  }'
```

When creating a new commit status, you can use a URI template for the URL.
Templates are URLs that contain variable names that Bitbucket will
evaluate at runtime whenever the URL is displayed anywhere similar to
parameter substitution in
[Bitbucket Connect](https://developer.atlassian.com/bitbucket/concepts/context-parameters.html).
For example, one could use `https://foo.com/builds/{repository.full_name}`
which Bitbucket will turn into `https://foo.com/builds/foo/bar` at render time.
The context variables available are `repository` and `commit`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commit The commit's SHA1.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
 @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
 @return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest
*/
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost(ctx context.Context, commit string, repoSlug string, workspace string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest {
	return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest{
		ApiService: a,
		ctx: ctx,
		commit: commit,
		repoSlug: repoSlug,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return Commitstatus
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostExecute(r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostRequest) (*Commitstatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Commitstatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommitStatusesApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build"
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", url.PathEscape(parameterToString(r.commit, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", url.PathEscape(parameterToString(r.repoSlug, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterToString(r.workspace, "")), -1)

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

type CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest struct {
	ctx context.Context
	ApiService *CommitStatusesApiService
	commit string
	repoSlug string
	workspace string
	q *string
	sort *string
}

// Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). 
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest) Q(q string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest {
	r.q = &q
	return r
}

// Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). Defaults to &#x60;created_on&#x60;. 
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest) Sort(sort string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest {
	r.sort = &sort
	return r
}

func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest) Execute() (*PaginatedCommitstatuses, *http.Response, error) {
	return r.ApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGetExecute(r)
}

/*
RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet List commit statuses for a commit

Returns all statuses (e.g. build results) for a specific commit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commit The commit's SHA1.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
 @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
 @return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest
*/
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet(ctx context.Context, commit string, repoSlug string, workspace string) CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest {
	return CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest{
		ApiService: a,
		ctx: ctx,
		commit: commit,
		repoSlug: repoSlug,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return PaginatedCommitstatuses
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugCommitCommitStatusesGetExecute(r CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetRequest) (*PaginatedCommitstatuses, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCommitstatuses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommitStatusesApiService.RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", url.PathEscape(parameterToString(r.commit, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", url.PathEscape(parameterToString(r.repoSlug, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterToString(r.workspace, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
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

type CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest struct {
	ctx context.Context
	ApiService *CommitStatusesApiService
	pullRequestId int32
	repoSlug string
	workspace string
	q *string
	sort *string
}

// Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). 
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest) Q(q string) CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest {
	r.q = &q
	return r
}

// Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). Defaults to &#x60;created_on&#x60;. 
func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest) Sort(sort string) CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest {
	r.sort = &sort
	return r
}

func (r CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest) Execute() (*PaginatedCommitstatuses, *http.Response, error) {
	return r.ApiService.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetExecute(r)
}

/*
RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet List commit statuses for a pull request

Returns all statuses (e.g. build results) for the given pull
request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pullRequestId The id of the pull request.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: `{repository UUID}`. 
 @param workspace This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`. 
 @return CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest
*/
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(ctx context.Context, pullRequestId int32, repoSlug string, workspace string) CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest {
	return CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest{
		ApiService: a,
		ctx: ctx,
		pullRequestId: pullRequestId,
		repoSlug: repoSlug,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return PaginatedCommitstatuses
func (a *CommitStatusesApiService) RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetExecute(r CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetRequest) (*PaginatedCommitstatuses, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCommitstatuses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommitStatusesApiService.RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"pull_request_id"+"}", url.PathEscape(parameterToString(r.pullRequestId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", url.PathEscape(parameterToString(r.repoSlug, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterToString(r.workspace, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
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
