// Package webmasters provides access to the Webmaster Tools API.
//
// See https://developers.google.com/webmaster-tools/
//
// Usage example:
//
//   import "google.golang.org/api/webmasters/v3"
//   ...
//   webmastersService, err := webmasters.New(oauthHttpClient)
package webmasters

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

const apiId = "webmasters:v3"
const apiName = "webmasters"
const apiVersion = "v3"
const basePath = "https://www.googleapis.com/webmasters/v3/"

// OAuth2 scopes used by this API.
const (
	// View and modify Webmaster Tools data for your verified sites
	WebmastersScope = "https://www.googleapis.com/auth/webmasters"

	// View Webmaster Tools data for your verified sites
	WebmastersReadonlyScope = "https://www.googleapis.com/auth/webmasters.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Searchanalytics = NewSearchanalyticsService(s)
	s.Sitemaps = NewSitemapsService(s)
	s.Sites = NewSitesService(s)
	s.Urlcrawlerrorscounts = NewUrlcrawlerrorscountsService(s)
	s.Urlcrawlerrorssamples = NewUrlcrawlerrorssamplesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Searchanalytics *SearchanalyticsService

	Sitemaps *SitemapsService

	Sites *SitesService

	Urlcrawlerrorscounts *UrlcrawlerrorscountsService

	Urlcrawlerrorssamples *UrlcrawlerrorssamplesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewSearchanalyticsService(s *Service) *SearchanalyticsService {
	rs := &SearchanalyticsService{s: s}
	return rs
}

type SearchanalyticsService struct {
	s *Service
}

func NewSitemapsService(s *Service) *SitemapsService {
	rs := &SitemapsService{s: s}
	return rs
}

type SitemapsService struct {
	s *Service
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	return rs
}

type SitesService struct {
	s *Service
}

func NewUrlcrawlerrorscountsService(s *Service) *UrlcrawlerrorscountsService {
	rs := &UrlcrawlerrorscountsService{s: s}
	return rs
}

type UrlcrawlerrorscountsService struct {
	s *Service
}

func NewUrlcrawlerrorssamplesService(s *Service) *UrlcrawlerrorssamplesService {
	rs := &UrlcrawlerrorssamplesService{s: s}
	return rs
}

type UrlcrawlerrorssamplesService struct {
	s *Service
}

type ApiDataRow struct {
	Clicks float64 `json:"clicks,omitempty"`

	Ctr float64 `json:"ctr,omitempty"`

	Impressions float64 `json:"impressions,omitempty"`

	Keys []string `json:"keys,omitempty"`

	Position float64 `json:"position,omitempty"`
}

type ApiDimensionFilter struct {
	Dimension string `json:"dimension,omitempty"`

	Expression string `json:"expression,omitempty"`

	Operator string `json:"operator,omitempty"`
}

type ApiDimensionFilterGroup struct {
	Filters []*ApiDimensionFilter `json:"filters,omitempty"`

	GroupType string `json:"groupType,omitempty"`
}

type SearchAnalyticsQueryRequest struct {
	// AggregationType: [Optional; Default is AUTO] How data is aggregated.
	// If aggregated by property, all data for the same property is
	// aggregated; if aggregated by page, all data is aggregated by
	// canonical URI. If you filter or group by page, choose AUTO; otherwise
	// you can aggregate either by property or by page, depending on how you
	// want your data calculated; see  the help documentation to learn how
	// data is calculated differently by site versus by page.
	//
	// Note: If you group or filter by page, you cannot aggregate by
	// property.
	//
	// If you specify any value other than AUTO, the aggregation type in the
	// result will match the requested type, or if you request an invalid
	// type, you will get an error. The API will never change your
	// aggregation type if the requested type is invalid.
	AggregationType string `json:"aggregationType,omitempty"`

	// DimensionFilterGroups: [Optional] Zero or more filters to apply to
	// the dimension grouping values; for example, 'Country CONTAINS
	// "Guinea"' to see only data where the country contains the substring
	// "Guinea". You can filter by a dimension without grouping by it.
	DimensionFilterGroups []*ApiDimensionFilterGroup `json:"dimensionFilterGroups,omitempty"`

	// Dimensions: [Optional] Zero or more dimensions to group results by.
	// Dimensions are the group-by values in the Search Analytics page.
	// Dimensions are combined to create a unique row key for each row.
	// Results are grouped in the order that you supply these dimensions.
	Dimensions []string `json:"dimensions,omitempty"`

	// EndDate: [Required] End date of the requested date range, in
	// YYYY-MM-DD format, in PST (UTC - 8:00). Must be greater than or equal
	// to the start date. This value is included in the range.
	EndDate string `json:"endDate,omitempty"`

	// RowLimit: [Optional; Default is 1000] The maximum number of rows to
	// return. Must be a number from 1 to 1,000 (inclusive).
	RowLimit int64 `json:"rowLimit,omitempty"`

	// SearchType: [Optional; Default is WEB] The search type to filter for.
	SearchType string `json:"searchType,omitempty"`

	// StartDate: [Required] Start date of the requested date range, in
	// YYYY-MM-DD format, in PST time (UTC - 8:00). Must be less than or
	// equal to the end date. This value is included in the range.
	StartDate string `json:"startDate,omitempty"`
}

type SearchAnalyticsQueryResponse struct {
	// ResponseAggregationType: How the results were aggregated.
	ResponseAggregationType string `json:"responseAggregationType,omitempty"`

	// Rows: A list of rows grouped by the key values in the order given in
	// the query.
	Rows []*ApiDataRow `json:"rows,omitempty"`
}

type SitemapsListResponse struct {
	// Sitemap: Contains detailed information about a specific URL submitted
	// as a sitemap.
	Sitemap []*WmxSitemap `json:"sitemap,omitempty"`
}

type SitesListResponse struct {
	// SiteEntry: Contains permission level information about a Webmaster
	// Tools site. For more information, see Permissions in Webmaster Tools.
	SiteEntry []*WmxSite `json:"siteEntry,omitempty"`
}

type UrlCrawlErrorCount struct {
	// Count: The error count at the given timestamp.
	Count int64 `json:"count,omitempty,string"`

	// Timestamp: The date and time when the crawl attempt took place, in
	// RFC 3339 format.
	Timestamp string `json:"timestamp,omitempty"`
}

type UrlCrawlErrorCountsPerType struct {
	// Category: The crawl error type.
	Category string `json:"category,omitempty"`

	// Entries: The error count entries time series.
	Entries []*UrlCrawlErrorCount `json:"entries,omitempty"`

	// Platform: The general type of Googlebot that made the request (see
	// list of Googlebot user-agents for the user-agents used).
	Platform string `json:"platform,omitempty"`
}

type UrlCrawlErrorsCountsQueryResponse struct {
	// CountPerTypes: The time series of the number of URL crawl errors per
	// error category and platform.
	CountPerTypes []*UrlCrawlErrorCountsPerType `json:"countPerTypes,omitempty"`
}

type UrlCrawlErrorsSample struct {
	// FirstDetected: The time the error was first detected, in RFC 3339
	// format.
	FirstDetected string `json:"first_detected,omitempty"`

	// LastCrawled: The time when the URL was last crawled, in RFC 3339
	// format.
	LastCrawled string `json:"last_crawled,omitempty"`

	// PageUrl: The URL of an error, relative to the site.
	PageUrl string `json:"pageUrl,omitempty"`

	// ResponseCode: The HTTP response code, if any.
	ResponseCode int64 `json:"responseCode,omitempty"`

	// UrlDetails: Additional details about the URL, set only when calling
	// get().
	UrlDetails *UrlSampleDetails `json:"urlDetails,omitempty"`
}

type UrlCrawlErrorsSamplesListResponse struct {
	// UrlCrawlErrorSample: Information about the sample URL and its crawl
	// error.
	UrlCrawlErrorSample []*UrlCrawlErrorsSample `json:"urlCrawlErrorSample,omitempty"`
}

type UrlSampleDetails struct {
	// ContainingSitemaps: List of sitemaps pointing at this URL.
	ContainingSitemaps []string `json:"containingSitemaps,omitempty"`

	// LinkedFromUrls: A sample set of URLs linking to this URL.
	LinkedFromUrls []string `json:"linkedFromUrls,omitempty"`
}

type WmxSite struct {
	// PermissionLevel: The user's permission level for the site.
	PermissionLevel string `json:"permissionLevel,omitempty"`

	// SiteUrl: The URL of the site.
	SiteUrl string `json:"siteUrl,omitempty"`
}

type WmxSitemap struct {
	// Contents: The various content types in the sitemap.
	Contents []*WmxSitemapContent `json:"contents,omitempty"`

	// Errors: Number of errors in the sitemap. These are issues with the
	// sitemap itself that need to be fixed before it can be processed
	// correctly.
	Errors int64 `json:"errors,omitempty,string"`

	// IsPending: If true, the sitemap has not been processed.
	IsPending bool `json:"isPending,omitempty"`

	// IsSitemapsIndex: If true, the sitemap is a collection of sitemaps.
	IsSitemapsIndex bool `json:"isSitemapsIndex,omitempty"`

	// LastDownloaded: Date & time in which this sitemap was last
	// downloaded. Date format is in RFC 3339 format (yyyy-mm-dd).
	LastDownloaded string `json:"lastDownloaded,omitempty"`

	// LastSubmitted: Date & time in which this sitemap was submitted. Date
	// format is in RFC 3339 format (yyyy-mm-dd).
	LastSubmitted string `json:"lastSubmitted,omitempty"`

	// Path: The url of the sitemap.
	Path string `json:"path,omitempty"`

	// Type: The type of the sitemap. For example: rssFeed.
	Type string `json:"type,omitempty"`

	// Warnings: Number of warnings for the sitemap. These are generally
	// non-critical issues with URLs in the sitemaps.
	Warnings int64 `json:"warnings,omitempty,string"`
}

type WmxSitemapContent struct {
	// Indexed: The number of URLs from the sitemap that were indexed (of
	// the content type).
	Indexed int64 `json:"indexed,omitempty,string"`

	// Submitted: The number of URLs in the sitemap (of the content type).
	Submitted int64 `json:"submitted,omitempty,string"`

	// Type: The specific type of content in this sitemap. For example: web.
	Type string `json:"type,omitempty"`
}

// method id "webmasters.searchanalytics.query":

type SearchanalyticsQueryCall struct {
	s                           *Service
	siteUrl                     string
	searchanalyticsqueryrequest *SearchAnalyticsQueryRequest
	opt_                        map[string]interface{}
}

// Query: [LIMITED ACCESS]
//
// Query your data with filters and parameters that you define. Returns
// zero or more rows grouped by the row keys that you define. You must
// define a date range of one or more days.
//
// When date is one of the group by values, any days without data are
// omitted from the result list. If you need to know which days have
// data, issue a broad date range query grouped by date for any metric,
// and see which day rows are returned.
func (r *SearchanalyticsService) Query(siteUrl string, searchanalyticsqueryrequest *SearchAnalyticsQueryRequest) *SearchanalyticsQueryCall {
	c := &SearchanalyticsQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.searchanalyticsqueryrequest = searchanalyticsqueryrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SearchanalyticsQueryCall) Fields(s ...googleapi.Field) *SearchanalyticsQueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SearchanalyticsQueryCall) Do() (*SearchAnalyticsQueryResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchanalyticsqueryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/searchAnalytics/query")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
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
	var ret *SearchAnalyticsQueryResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "[LIMITED ACCESS]\n\nQuery your data with filters and parameters that you define. Returns zero or more rows grouped by the row keys that you define. You must define a date range of one or more days.\n\nWhen date is one of the group by values, any days without data are omitted from the result list. If you need to know which days have data, issue a broad date range query grouped by date for any metric, and see which day rows are returned.",
	//   "httpMethod": "POST",
	//   "id": "webmasters.searchanalytics.query",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/searchAnalytics/query",
	//   "request": {
	//     "$ref": "SearchAnalyticsQueryRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchAnalyticsQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sitemaps.delete":

type SitemapsDeleteCall struct {
	s        *Service
	siteUrl  string
	feedpath string
	opt_     map[string]interface{}
}

// Delete: Deletes a sitemap from this site.
func (r *SitemapsService) Delete(siteUrl string, feedpath string) *SitemapsDeleteCall {
	c := &SitemapsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.feedpath = feedpath
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsDeleteCall) Fields(s ...googleapi.Field) *SitemapsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitemapsDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/sitemaps/{feedpath}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a sitemap from this site.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.sitemaps.delete",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the actual sitemap. For example: http://www.example.com/sitemap.xml",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sitemaps.get":

type SitemapsGetCall struct {
	s        *Service
	siteUrl  string
	feedpath string
	opt_     map[string]interface{}
}

// Get: Retrieves information about a specific sitemap.
func (r *SitemapsService) Get(siteUrl string, feedpath string) *SitemapsGetCall {
	c := &SitemapsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.feedpath = feedpath
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsGetCall) Fields(s ...googleapi.Field) *SitemapsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitemapsGetCall) Do() (*WmxSitemap, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/sitemaps/{feedpath}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
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
	var ret *WmxSitemap
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about a specific sitemap.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sitemaps.get",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the actual sitemap. For example: http://www.example.com/sitemap.xml",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "response": {
	//     "$ref": "WmxSitemap"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sitemaps.list":

type SitemapsListCall struct {
	s       *Service
	siteUrl string
	opt_    map[string]interface{}
}

// List: Lists the sitemaps-entries submitted for this site, or included
// in the sitemap index file (if sitemapIndex is specified in the
// request).
func (r *SitemapsService) List(siteUrl string) *SitemapsListCall {
	c := &SitemapsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	return c
}

// SitemapIndex sets the optional parameter "sitemapIndex": A URL of a
// site's sitemap index. For example:
// http://www.example.com/sitemapindex.xml
func (c *SitemapsListCall) SitemapIndex(sitemapIndex string) *SitemapsListCall {
	c.opt_["sitemapIndex"] = sitemapIndex
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsListCall) Fields(s ...googleapi.Field) *SitemapsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitemapsListCall) Do() (*SitemapsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["sitemapIndex"]; ok {
		params.Set("sitemapIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/sitemaps")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
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
	var ret *SitemapsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the sitemaps-entries submitted for this site, or included in the sitemap index file (if sitemapIndex is specified in the request).",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sitemaps.list",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sitemapIndex": {
	//       "description": "A URL of a site's sitemap index. For example: http://www.example.com/sitemapindex.xml",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps",
	//   "response": {
	//     "$ref": "SitemapsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sitemaps.submit":

type SitemapsSubmitCall struct {
	s        *Service
	siteUrl  string
	feedpath string
	opt_     map[string]interface{}
}

// Submit: Submits a sitemap for a site.
func (r *SitemapsService) Submit(siteUrl string, feedpath string) *SitemapsSubmitCall {
	c := &SitemapsSubmitCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.feedpath = feedpath
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsSubmitCall) Fields(s ...googleapi.Field) *SitemapsSubmitCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitemapsSubmitCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/sitemaps/{feedpath}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Submits a sitemap for a site.",
	//   "httpMethod": "PUT",
	//   "id": "webmasters.sitemaps.submit",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the sitemap to add. For example: http://www.example.com/sitemap.xml",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.add":

type SitesAddCall struct {
	s       *Service
	siteUrl string
	opt_    map[string]interface{}
}

// Add: Adds a site to the set of the user's sites in Webmaster Tools.
func (r *SitesService) Add(siteUrl string) *SitesAddCall {
	c := &SitesAddCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesAddCall) Fields(s ...googleapi.Field) *SitesAddCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitesAddCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Adds a site to the set of the user's sites in Webmaster Tools.",
	//   "httpMethod": "PUT",
	//   "id": "webmasters.sites.add",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The URL of the site to add.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.delete":

type SitesDeleteCall struct {
	s       *Service
	siteUrl string
	opt_    map[string]interface{}
}

// Delete: Removes a site from the set of the user's Webmaster Tools
// sites.
func (r *SitesService) Delete(siteUrl string) *SitesDeleteCall {
	c := &SitesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesDeleteCall) Fields(s ...googleapi.Field) *SitesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitesDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Removes a site from the set of the user's Webmaster Tools sites.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.sites.delete",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The URI of the property as defined in Search Console. Examples: http://www.example.com/ or android-app://com.example/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.get":

type SitesGetCall struct {
	s       *Service
	siteUrl string
	opt_    map[string]interface{}
}

// Get: Retrieves information about specific site.
func (r *SitesService) Get(siteUrl string) *SitesGetCall {
	c := &SitesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesGetCall) Fields(s ...googleapi.Field) *SitesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitesGetCall) Do() (*WmxSite, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
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
	var ret *WmxSite
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves information about specific site.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sites.get",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The URI of the property as defined in Search Console. Examples: http://www.example.com/ or android-app://com.example/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "response": {
	//     "$ref": "WmxSite"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sites.list":

type SitesListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists the user's Webmaster Tools sites.
func (r *SitesService) List() *SitesListCall {
	c := &SitesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesListCall) Fields(s ...googleapi.Field) *SitesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SitesListCall) Do() (*SitesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites")
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
	var ret *SitesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the user's Webmaster Tools sites.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sites.list",
	//   "path": "sites",
	//   "response": {
	//     "$ref": "SitesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorscounts.query":

type UrlcrawlerrorscountsQueryCall struct {
	s       *Service
	siteUrl string
	opt_    map[string]interface{}
}

// Query: Retrieves a time series of the number of URL crawl errors per
// error category and platform.
func (r *UrlcrawlerrorscountsService) Query(siteUrl string) *UrlcrawlerrorscountsQueryCall {
	c := &UrlcrawlerrorscountsQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	return c
}

// Category sets the optional parameter "category": The crawl error
// category. For example: serverError. If not specified, returns results
// for all categories.
//
// Possible values:
//   "authPermissions"
//   "manyToOneRedirect"
//   "notFollowed"
//   "notFound"
//   "other"
//   "roboted"
//   "serverError"
//   "soft404"
func (c *UrlcrawlerrorscountsQueryCall) Category(category string) *UrlcrawlerrorscountsQueryCall {
	c.opt_["category"] = category
	return c
}

// LatestCountsOnly sets the optional parameter "latestCountsOnly": If
// true, returns only the latest crawl error counts.
func (c *UrlcrawlerrorscountsQueryCall) LatestCountsOnly(latestCountsOnly bool) *UrlcrawlerrorscountsQueryCall {
	c.opt_["latestCountsOnly"] = latestCountsOnly
	return c
}

// Platform sets the optional parameter "platform": The user agent type
// (platform) that made the request. For example: web. If not specified,
// returns results for all platforms.
//
// Possible values:
//   "mobile"
//   "smartphoneOnly"
//   "web"
func (c *UrlcrawlerrorscountsQueryCall) Platform(platform string) *UrlcrawlerrorscountsQueryCall {
	c.opt_["platform"] = platform
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorscountsQueryCall) Fields(s ...googleapi.Field) *UrlcrawlerrorscountsQueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UrlcrawlerrorscountsQueryCall) Do() (*UrlCrawlErrorsCountsQueryResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["category"]; ok {
		params.Set("category", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["latestCountsOnly"]; ok {
		params.Set("latestCountsOnly", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["platform"]; ok {
		params.Set("platform", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/urlCrawlErrorsCounts/query")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
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
	var ret *UrlCrawlErrorsCountsQueryResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a time series of the number of URL crawl errors per error category and platform.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorscounts.query",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category. For example: serverError. If not specified, returns results for all categories.",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "latestCountsOnly": {
	//       "default": "true",
	//       "description": "If true, returns only the latest crawl error counts.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request. For example: web. If not specified, returns results for all platforms.",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsCounts/query",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsCountsQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.get":

type UrlcrawlerrorssamplesGetCall struct {
	s        *Service
	siteUrl  string
	url      string
	category string
	platform string
	opt_     map[string]interface{}
}

// Get: Retrieves details about crawl errors for a site's sample URL.
func (r *UrlcrawlerrorssamplesService) Get(siteUrl string, url string, category string, platform string) *UrlcrawlerrorssamplesGetCall {
	c := &UrlcrawlerrorssamplesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.url = url
	c.category = category
	c.platform = platform
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorssamplesGetCall) Fields(s ...googleapi.Field) *UrlcrawlerrorssamplesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UrlcrawlerrorssamplesGetCall) Do() (*UrlCrawlErrorsSample, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("category", fmt.Sprintf("%v", c.category))
	params.Set("platform", fmt.Sprintf("%v", c.platform))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/urlCrawlErrorsSamples/{url}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
		"url":     c.url,
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
	var ret *UrlCrawlErrorsSample
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves details about crawl errors for a site's sample URL.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorssamples.get",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "url",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category. For example: authPermissions",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request. For example: web",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The relative path (without the site) of the sample URL. It must be one of the URLs returned by list(). For example, for the URL https://www.example.com/pagename on the site https://www.example.com/, the url value is pagename",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsSample"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.list":

type UrlcrawlerrorssamplesListCall struct {
	s        *Service
	siteUrl  string
	category string
	platform string
	opt_     map[string]interface{}
}

// List: Lists a site's sample URLs for the specified crawl error
// category and platform.
func (r *UrlcrawlerrorssamplesService) List(siteUrl string, category string, platform string) *UrlcrawlerrorssamplesListCall {
	c := &UrlcrawlerrorssamplesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.category = category
	c.platform = platform
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorssamplesListCall) Fields(s ...googleapi.Field) *UrlcrawlerrorssamplesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UrlcrawlerrorssamplesListCall) Do() (*UrlCrawlErrorsSamplesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("category", fmt.Sprintf("%v", c.category))
	params.Set("platform", fmt.Sprintf("%v", c.platform))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/urlCrawlErrorsSamples")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
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
	var ret *UrlCrawlErrorsSamplesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists a site's sample URLs for the specified crawl error category and platform.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorssamples.list",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category. For example: authPermissions",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request. For example: web",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsSamplesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.markAsFixed":

type UrlcrawlerrorssamplesMarkAsFixedCall struct {
	s        *Service
	siteUrl  string
	url      string
	category string
	platform string
	opt_     map[string]interface{}
}

// MarkAsFixed: Marks the provided site's sample URL as fixed, and
// removes it from the samples list.
func (r *UrlcrawlerrorssamplesService) MarkAsFixed(siteUrl string, url string, category string, platform string) *UrlcrawlerrorssamplesMarkAsFixedCall {
	c := &UrlcrawlerrorssamplesMarkAsFixedCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteUrl = siteUrl
	c.url = url
	c.category = category
	c.platform = platform
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorssamplesMarkAsFixedCall) Fields(s ...googleapi.Field) *UrlcrawlerrorssamplesMarkAsFixedCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UrlcrawlerrorssamplesMarkAsFixedCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("category", fmt.Sprintf("%v", c.category))
	params.Set("platform", fmt.Sprintf("%v", c.platform))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "sites/{siteUrl}/urlCrawlErrorsSamples/{url}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"siteUrl": c.siteUrl,
		"url":     c.url,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Marks the provided site's sample URL as fixed, and removes it from the samples list.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.urlcrawlerrorssamples.markAsFixed",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "url",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category. For example: authPermissions",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request. For example: web",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol. For example: http://www.example.com/",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The relative path (without the site) of the sample URL. It must be one of the URLs returned by list(). For example, for the URL https://www.example.com/pagename on the site https://www.example.com/, the url value is pagename",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}
