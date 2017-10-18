// Package cloudresourcemanager provides access to the Google Cloud Resource Manager API.
//
// See https://cloud.google.com/resource-manager
//
// Usage example:
//
//   import "google.golang.org/api/cloudresourcemanager/v1beta1"
//   ...
//   cloudresourcemanagerService, err := cloudresourcemanager.New(oauthHttpClient)
package cloudresourcemanager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "cloudresourcemanager:v1beta1"
const apiName = "cloudresourcemanager"
const apiVersion = "v1beta1"
const basePath = "https://cloudresourcemanager.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

type Empty struct {
}

type ListProjectsResponse struct {
	// NextPageToken: Pagination token. If the result set is too large to
	// fit in a single response, this token is returned. It encodes the
	// position of the current result cursor. Feeding this value into a new
	// list request with the `page_token` parameter gives the next page of
	// the results. When `next_page_token` is not filled in, there is no
	// next page and the list returned is the last page in the result set.
	// Pagination tokens have a limited lifetime. Note: pagination is not
	// yet supported; the server will not set this field.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Projects: The list of projects that matched the list filter. This
	// list can be paginated.
	Projects []*Project `json:"projects,omitempty"`
}

type Project struct {
	// CreateTime: Creation time. Read-only.
	CreateTime string `json:"createTime,omitempty"`

	// Labels: The labels associated with this project. Label keys must be
	// between 1 and 63 characters long and must conform to the following
	// regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be
	// between 0 and 63 characters long and must conform to the regular
	// expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 256 labels can
	// be associated with a given resource. Note that additional character
	// may be included in labels in the future. Clients should store labels
	// in a representation such as JSON that does not depend on specific
	// characters being disallowed. Example: "environment" : "dev"
	// Read-write.
	Labels map[string]string `json:"labels,omitempty"`

	// LifecycleState: The project lifecycle state. Read-only.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED"
	//   "ACTIVE"
	//   "DELETE_REQUESTED"
	//   "DELETE_IN_PROGRESS"
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: The user-assigned name of the project. This field is optional
	// and can remain unset. Allowed characters are: lower- and upper-case
	// letters, numbers, hyphen, single-quote, double-quotes, space, and
	// exclamation point. Example: "My Project" Read-write.
	Name string `json:"name,omitempty"`

	// ProjectId: The unique, user-assigned id of the project. It must be 6
	// to 30 lowercase letters, digits, or hyphens. It must start with a
	// letter. Trailing hyphens are prohibited. Example: "tokyo-rain-123"
	// Read-only after creation.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: The number uniquely identifying the project. Example:
	// 415104041262. Read-only.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

// method id "cloudresourcemanager.projects.create":

type ProjectsCreateCall struct {
	s       *Service
	project *Project
	opt_    map[string]interface{}
}

// Create: Creates a project resource. Initially, the project resource
// is owned by its creator exclusively. The creator can later grant
// permission to others to read or update the project. Several APIs are
// activated automatically for the project, including Google Cloud
// Storage.
func (r *ProjectsService) Create(project *Project) *ProjectsCreateCall {
	c := &ProjectsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCreateCall) Fields(s ...googleapi.Field) *ProjectsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsCreateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a project resource. Initially, the project resource is owned by its creator exclusively. The creator can later grant permission to others to read or update the project. Several APIs are activated automatically for the project, including Google Cloud Storage.",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.create",
	//   "path": "v1beta1/projects",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.delete":

type ProjectsDeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Delete: Marks the project identified by the specified `project_id`
// (for example, `my-project-123`) for deletion. This method will only
// affect the project if it has a lifecycle state of
// [ACTIVE][cloudresourcemanager.projects.v1beta2.LifecycleState.ACTIVE]
// when this method is called. Otherwise this method does nothing (since
// all other states are phases of deletion). This method changes the
// project's lifecycle state from
// [ACTIVE][cloudresourcemanager.projects.v1beta2.LifecycleState.ACTIVE]
// to [DELETE_REQUESTED]
// [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_REQUESTED
// ]. The deletion starts at an unspecified time, at which point the
// lifecycle state changes to [DELETE_IN_PROGRESS]
// [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_IN_PROGRE
// SS]. Until the deletion completes, you can check the lifecycle state
// checked by retrieving the project with [GetProject]
// [cloudresourcemanager.projects.v1beta2.Projects.GetProject], and the
// project remains visible to [ListProjects]
// [cloudresourcemanager.projects.v1beta2.Projects.ListProjects].
// However, you cannot update the project. After the deletion completes,
// the project is not retrievable by the [GetProject]
// [cloudresourcemanager.projects.v1beta2.Projects.GetProject] and
// [ListProjects]
// [cloudresourcemanager.projects.v1beta2.Projects.ListProjects]
// methods. The caller must have modify permissions for this project.
func (r *ProjectsService) Delete(projectId string) *ProjectsDeleteCall {
	c := &ProjectsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks the project identified by the specified `project_id` (for example, `my-project-123`) for deletion. This method will only affect the project if it has a lifecycle state of [ACTIVE][cloudresourcemanager.projects.v1beta2.LifecycleState.ACTIVE] when this method is called. Otherwise this method does nothing (since all other states are phases of deletion). This method changes the project's lifecycle state from [ACTIVE][cloudresourcemanager.projects.v1beta2.LifecycleState.ACTIVE] to [DELETE_REQUESTED] [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_REQUESTED]. The deletion starts at an unspecified time, at which point the lifecycle state changes to [DELETE_IN_PROGRESS] [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_IN_PROGRESS]. Until the deletion completes, you can check the lifecycle state checked by retrieving the project with [GetProject] [cloudresourcemanager.projects.v1beta2.Projects.GetProject], and the project remains visible to [ListProjects] [cloudresourcemanager.projects.v1beta2.Projects.ListProjects]. However, you cannot update the project. After the deletion completes, the project is not retrievable by the [GetProject] [cloudresourcemanager.projects.v1beta2.Projects.GetProject] and [ListProjects] [cloudresourcemanager.projects.v1beta2.Projects.ListProjects] methods. The caller must have modify permissions for this project.",
	//   "httpMethod": "DELETE",
	//   "id": "cloudresourcemanager.projects.delete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `foo-bar-123`). Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.get":

type ProjectsGetCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Get: Retrieves the project identified by the specified `project_id`
// (for example, `my-project-123`). The caller must have read
// permissions for this project.
func (r *ProjectsService) Get(projectId string) *ProjectsGetCall {
	c := &ProjectsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsGetCall) Do() (*Project, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the project identified by the specified `project_id` (for example, `my-project-123`). The caller must have read permissions for this project.",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.get",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `my-project-123`). Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.list":

type ProjectsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists projects that are visible to the user and satisfy the
// specified filter. This method returns projects in an unspecified
// order. New projects do not necessarily appear at the end of the list.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Filter sets the optional parameter "filter": An expression for
// filtering the results of the request. Filter rules are case
// insensitive. The fields eligible for filtering are: name id labels.
// where  is a the name of a label Examples: name:* ==> The project has
// a name. name:Howl ==> The project’s name is `Howl` or 'howl'.
// name:HOWL ==> Equivalent to above. NAME:howl ==> Equivalent to above.
// labels.color:* ==> The project has the label "color".
// labels.color:red ==> The project’s label `color` has the value
// `red`. labels.color:red label.size:big ==> The project's label
// `color` has the value `red` and its label `size` has the value `big`.
func (c *ProjectsListCall) Filter(filter string) *ProjectsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Projects to return in the response. The server can return fewer
// projects than requested. If unspecified, server picks an appropriate
// default. Note: pagination is not yet supported; the server ignores
// this field.
func (c *ProjectsListCall) PageSize(pageSize int64) *ProjectsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// returned from a previous call to ListProject that indicates from
// where listing should continue. Note: pagination is not yet supported;
// the server ignores this field.
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsListCall) Do() (*ListProjectsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListProjectsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists projects that are visible to the user and satisfy the specified filter. This method returns projects in an unspecified order. New projects do not necessarily appear at the end of the list.",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.list",
	//   "parameters": {
	//     "filter": {
	//       "description": "An expression for filtering the results of the request. Filter rules are case insensitive. The fields eligible for filtering are: name id labels. where  is a the name of a label Examples: name:* ==\u003e The project has a name. name:Howl ==\u003e The project’s name is `Howl` or 'howl'. name:HOWL ==\u003e Equivalent to above. NAME:howl ==\u003e Equivalent to above. labels.color:* ==\u003e The project has the label \"color\". labels.color:red ==\u003e The project’s label `color` has the value `red`. labels.color:red label.size:big ==\u003e The project's label `color` has the value `red` and its label `size` has the value `big`. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Projects to return in the response. The server can return fewer projects than requested. If unspecified, server picks an appropriate default. Note: pagination is not yet supported; the server ignores this field. Optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token returned from a previous call to ListProject that indicates from where listing should continue. Note: pagination is not yet supported; the server ignores this field. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects",
	//   "response": {
	//     "$ref": "ListProjectsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.undelete":

type ProjectsUndeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Undelete: Restores the project identified by the specified
// `project_id` (for example, `my-project-123`). You can only use this
// method for a project that has a lifecycle state of [DELETE_REQUESTED]
// [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_REQUESTED
// ]. After deletion starts, as indicated by a lifecycle state of
// [DELETE_IN_PROGRESS]
// [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_IN_PROGRE
// SS], the project cannot be restored. The caller must have modify
// permissions for this project.
func (r *ProjectsService) Undelete(projectId string) *ProjectsUndeleteCall {
	c := &ProjectsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUndeleteCall) Fields(s ...googleapi.Field) *ProjectsUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUndeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}:undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores the project identified by the specified `project_id` (for example, `my-project-123`). You can only use this method for a project that has a lifecycle state of [DELETE_REQUESTED] [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_REQUESTED]. After deletion starts, as indicated by a lifecycle state of [DELETE_IN_PROGRESS] [cloudresourcemanager.projects.v1beta2.LifecycleState.DELETE_IN_PROGRESS], the project cannot be restored. The caller must have modify permissions for this project.",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.undelete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `foo-bar-123`). Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}:undelete",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.update":

type ProjectsUpdateCall struct {
	s         *Service
	projectId string
	project   *Project
	opt_      map[string]interface{}
}

// Update: Updates the attributes of the project identified by the
// specified `project_id` (for example, `my-project-123`). The caller
// must have modify permissions for this project.
func (r *ProjectsService) Update(projectId string, project *Project) *ProjectsUpdateCall {
	c := &ProjectsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateCall) Fields(s ...googleapi.Field) *ProjectsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUpdateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the attributes of the project identified by the specified `project_id` (for example, `my-project-123`). The caller must have modify permissions for this project.",
	//   "httpMethod": "PUT",
	//   "id": "cloudresourcemanager.projects.update",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `my-project-123`). Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
