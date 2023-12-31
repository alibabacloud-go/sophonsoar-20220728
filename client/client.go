// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchModifyInstanceStatusRequest struct {
	Active       *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s BatchModifyInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusRequest) SetActive(v int32) *BatchModifyInstanceStatusRequest {
	s.Active = &v
	return s
}

func (s *BatchModifyInstanceStatusRequest) SetLang(v string) *BatchModifyInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *BatchModifyInstanceStatusRequest) SetPlaybookUuid(v string) *BatchModifyInstanceStatusRequest {
	s.PlaybookUuid = &v
	return s
}

type BatchModifyInstanceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchModifyInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusResponseBody) SetRequestId(v string) *BatchModifyInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type BatchModifyInstanceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchModifyInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchModifyInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusResponse) SetHeaders(v map[string]*string) *BatchModifyInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyInstanceStatusResponse) SetStatusCode(v int32) *BatchModifyInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchModifyInstanceStatusResponse) SetBody(v *BatchModifyInstanceStatusResponseBody) *BatchModifyInstanceStatusResponse {
	s.Body = v
	return s
}

type ComparePlaybooksRequest struct {
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewPlaybookReleaseId *int32  `json:"NewPlaybookReleaseId,omitempty" xml:"NewPlaybookReleaseId,omitempty"`
	OldPlaybookReleaseId *int32  `json:"OldPlaybookReleaseId,omitempty" xml:"OldPlaybookReleaseId,omitempty"`
	PlaybookUuid         *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ComparePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksRequest) SetLang(v string) *ComparePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *ComparePlaybooksRequest) SetNewPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.NewPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetOldPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.OldPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetPlaybookUuid(v string) *ComparePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

type ComparePlaybooksResponseBody struct {
	CompareResult *ComparePlaybooksResponseBodyCompareResult `json:"CompareResult,omitempty" xml:"CompareResult,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ComparePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBody) SetCompareResult(v *ComparePlaybooksResponseBodyCompareResult) *ComparePlaybooksResponseBody {
	s.CompareResult = v
	return s
}

func (s *ComparePlaybooksResponseBody) SetRequestId(v string) *ComparePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type ComparePlaybooksResponseBodyCompareResult struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	New         *bool   `json:"New,omitempty" xml:"New,omitempty"`
	Same        *bool   `json:"Same,omitempty" xml:"Same,omitempty"`
}

func (s ComparePlaybooksResponseBodyCompareResult) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponseBodyCompareResult) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetDescription(v string) *ComparePlaybooksResponseBodyCompareResult {
	s.Description = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetNew(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.New = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetSame(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.Same = &v
	return s
}

type ComparePlaybooksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ComparePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ComparePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponse) SetHeaders(v map[string]*string) *ComparePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *ComparePlaybooksResponse) SetStatusCode(v int32) *ComparePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ComparePlaybooksResponse) SetBody(v *ComparePlaybooksResponseBody) *ComparePlaybooksResponse {
	s.Body = v
	return s
}

type CreatePlaybookRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CreatePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookRequest) GoString() string {
	return s.String()
}

func (s *CreatePlaybookRequest) SetDescription(v string) *CreatePlaybookRequest {
	s.Description = &v
	return s
}

func (s *CreatePlaybookRequest) SetDisplayName(v string) *CreatePlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *CreatePlaybookRequest) SetLang(v string) *CreatePlaybookRequest {
	s.Lang = &v
	return s
}

type CreatePlaybookResponseBody struct {
	Data      *CreatePlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponseBody) SetData(v *CreatePlaybookResponseBodyData) *CreatePlaybookResponseBody {
	s.Data = v
	return s
}

func (s *CreatePlaybookResponseBody) SetRequestId(v string) *CreatePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type CreatePlaybookResponseBodyData struct {
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s CreatePlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponseBodyData) SetPlaybookUuid(v string) *CreatePlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

type CreatePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponse) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponse) SetHeaders(v map[string]*string) *CreatePlaybookResponse {
	s.Headers = v
	return s
}

func (s *CreatePlaybookResponse) SetStatusCode(v int32) *CreatePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlaybookResponse) SetBody(v *CreatePlaybookResponseBody) *CreatePlaybookResponse {
	s.Body = v
	return s
}

type DebugPlaybookRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	Record       *string `json:"Record,omitempty" xml:"Record,omitempty"`
	Taskflow     *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s DebugPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DebugPlaybookRequest) SetLang(v string) *DebugPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DebugPlaybookRequest) SetPlaybookUuid(v string) *DebugPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DebugPlaybookRequest) SetRecord(v string) *DebugPlaybookRequest {
	s.Record = &v
	return s
}

func (s *DebugPlaybookRequest) SetTaskflow(v string) *DebugPlaybookRequest {
	s.Taskflow = &v
	return s
}

type DebugPlaybookResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DebugPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponseBody) SetRequestId(v string) *DebugPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugPlaybookResponseBody) SetRequestUuid(v string) *DebugPlaybookResponseBody {
	s.RequestUuid = &v
	return s
}

type DebugPlaybookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DebugPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DebugPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponse) SetHeaders(v map[string]*string) *DebugPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DebugPlaybookResponse) SetStatusCode(v int32) *DebugPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugPlaybookResponse) SetBody(v *DebugPlaybookResponseBody) *DebugPlaybookResponse {
	s.Body = v
	return s
}

type DeleteComponentAssetRequest struct {
	AssetId *int64  `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteComponentAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetRequest) SetAssetId(v int64) *DeleteComponentAssetRequest {
	s.AssetId = &v
	return s
}

func (s *DeleteComponentAssetRequest) SetLang(v string) *DeleteComponentAssetRequest {
	s.Lang = &v
	return s
}

type DeleteComponentAssetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComponentAssetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponseBody) SetRequestId(v string) *DeleteComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteComponentAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponse) SetHeaders(v map[string]*string) *DeleteComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentAssetResponse) SetStatusCode(v int32) *DeleteComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentAssetResponse) SetBody(v *DeleteComponentAssetResponseBody) *DeleteComponentAssetResponse {
	s.Body = v
	return s
}

type DeletePlaybookRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DeletePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaybookRequest) SetLang(v string) *DeletePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DeletePlaybookRequest) SetPlaybookUuid(v string) *DeletePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type DeletePlaybookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlaybookResponseBody) SetRequestId(v string) *DeletePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DeletePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DeletePlaybookResponse) SetHeaders(v map[string]*string) *DeletePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DeletePlaybookResponse) SetStatusCode(v int32) *DeletePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlaybookResponse) SetBody(v *DeletePlaybookResponseBody) *DeletePlaybookResponse {
	s.Body = v
	return s
}

type DescribeApiListRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeApiListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiListRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiListRequest) SetLang(v string) *DescribeApiListRequest {
	s.Lang = &v
	return s
}

type DescribeApiListResponseBody struct {
	ApiList   []*DescribeApiListResponseBodyApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiListResponseBody) SetApiList(v []*DescribeApiListResponseBodyApiList) *DescribeApiListResponseBody {
	s.ApiList = v
	return s
}

func (s *DescribeApiListResponseBody) SetRequestId(v string) *DescribeApiListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiListResponseBodyApiList struct {
	DocUrl      *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	PopCode     *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s DescribeApiListResponseBodyApiList) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiListResponseBodyApiList) GoString() string {
	return s.String()
}

func (s *DescribeApiListResponseBodyApiList) SetDocUrl(v string) *DescribeApiListResponseBodyApiList {
	s.DocUrl = &v
	return s
}

func (s *DescribeApiListResponseBodyApiList) SetPopCode(v string) *DescribeApiListResponseBodyApiList {
	s.PopCode = &v
	return s
}

func (s *DescribeApiListResponseBodyApiList) SetProductName(v string) *DescribeApiListResponseBodyApiList {
	s.ProductName = &v
	return s
}

type DescribeApiListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiListResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiListResponse) SetHeaders(v map[string]*string) *DescribeApiListResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiListResponse) SetStatusCode(v int32) *DescribeApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiListResponse) SetBody(v *DescribeApiListResponseBody) *DescribeApiListResponse {
	s.Body = v
	return s
}

type DescribeComponentAssetFormRequest struct {
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentAssetFormRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormRequest) SetComponentName(v string) *DescribeComponentAssetFormRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetFormRequest) SetLang(v string) *DescribeComponentAssetFormRequest {
	s.Lang = &v
	return s
}

type DescribeComponentAssetFormResponseBody struct {
	ComponentAssetForm *string `json:"ComponentAssetForm,omitempty" xml:"ComponentAssetForm,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponseBody) SetComponentAssetForm(v string) *DescribeComponentAssetFormResponseBody {
	s.ComponentAssetForm = &v
	return s
}

func (s *DescribeComponentAssetFormResponseBody) SetRequestId(v string) *DescribeComponentAssetFormResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentAssetFormResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentAssetFormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentAssetFormResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetFormResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetStatusCode(v int32) *DescribeComponentAssetFormResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetBody(v *DescribeComponentAssetFormResponseBody) *DescribeComponentAssetFormResponse {
	s.Body = v
	return s
}

type DescribeComponentAssetsRequest struct {
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsRequest) SetComponentName(v string) *DescribeComponentAssetsRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetsRequest) SetLang(v string) *DescribeComponentAssetsRequest {
	s.Lang = &v
	return s
}

type DescribeComponentAssetsResponseBody struct {
	ComponentAssets []*DescribeComponentAssetsResponseBodyComponentAssets `json:"ComponentAssets,omitempty" xml:"ComponentAssets,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBody) SetComponentAssets(v []*DescribeComponentAssetsResponseBodyComponentAssets) *DescribeComponentAssetsResponseBody {
	s.ComponentAssets = v
	return s
}

func (s *DescribeComponentAssetsResponseBody) SetRequestId(v string) *DescribeComponentAssetsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentAssetsResponseBodyComponentAssets struct {
	AssetUuid     *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	Componentname *string `json:"Componentname,omitempty" xml:"Componentname,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Params        *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetAssetUuid(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.AssetUuid = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetComponentname(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Componentname = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtCreate(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtCreate = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtModified(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtModified = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetId(v int64) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Id = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetName(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Name = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetParams(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Params = &v
	return s
}

type DescribeComponentAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetsResponse) SetStatusCode(v int32) *DescribeComponentAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetsResponse) SetBody(v *DescribeComponentAssetsResponseBody) *DescribeComponentAssetsResponse {
	s.Body = v
	return s
}

type DescribeComponentListRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentListRequest) SetLang(v string) *DescribeComponentListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentListRequest) SetPlaybookUuid(v string) *DescribeComponentListRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeComponentListResponseBody struct {
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponseBody) SetComponents(v string) *DescribeComponentListResponseBody {
	s.Components = &v
	return s
}

func (s *DescribeComponentListResponseBody) SetRequestId(v string) *DescribeComponentListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponse) SetHeaders(v map[string]*string) *DescribeComponentListResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentListResponse) SetStatusCode(v int32) *DescribeComponentListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentListResponse) SetBody(v *DescribeComponentListResponseBody) *DescribeComponentListResponse {
	s.Body = v
	return s
}

type DescribeComponentPlaybookRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookRequest) SetLang(v string) *DescribeComponentPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentPlaybookRequest) SetPlaybookUuid(v string) *DescribeComponentPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeComponentPlaybookResponseBody struct {
	Playbooks []*DescribeComponentPlaybookResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBody) SetPlaybooks(v []*DescribeComponentPlaybookResponseBodyPlaybooks) *DescribeComponentPlaybookResponseBody {
	s.Playbooks = v
	return s
}

func (s *DescribeComponentPlaybookResponseBody) SetRequestId(v string) *DescribeComponentPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentPlaybookResponseBodyPlaybooks struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDescription(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.Description = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDisplayName(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.DisplayName = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetInputParams(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.InputParams = &v
	return s
}

type DescribeComponentPlaybookResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponse) SetHeaders(v map[string]*string) *DescribeComponentPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetStatusCode(v int32) *DescribeComponentPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetBody(v *DescribeComponentPlaybookResponseBody) *DescribeComponentPlaybookResponse {
	s.Body = v
	return s
}

type DescribeComponentsJsRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentsJsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsRequest) SetLang(v string) *DescribeComponentsJsRequest {
	s.Lang = &v
	return s
}

type DescribeComponentsJsResponseBody struct {
	ComponentsJs *string `json:"ComponentsJs,omitempty" xml:"ComponentsJs,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentsJsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponseBody) SetComponentsJs(v string) *DescribeComponentsJsResponseBody {
	s.ComponentsJs = &v
	return s
}

func (s *DescribeComponentsJsResponseBody) SetRequestId(v string) *DescribeComponentsJsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentsJsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentsJsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentsJsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponse) SetHeaders(v map[string]*string) *DescribeComponentsJsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentsJsResponse) SetStatusCode(v int32) *DescribeComponentsJsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentsJsResponse) SetBody(v *DescribeComponentsJsResponseBody) *DescribeComponentsJsResponse {
	s.Body = v
	return s
}

type DescribeDistinctReleasesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	TaskflowMd5  *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribeDistinctReleasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesRequest) SetLang(v string) *DescribeDistinctReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetPlaybookUuid(v string) *DescribeDistinctReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetTaskflowMd5(v string) *DescribeDistinctReleasesRequest {
	s.TaskflowMd5 = &v
	return s
}

type DescribeDistinctReleasesResponseBody struct {
	Records   []*DescribeDistinctReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDistinctReleasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBody) SetRecords(v []*DescribeDistinctReleasesResponseBodyRecords) *DescribeDistinctReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDistinctReleasesResponseBody) SetRequestId(v string) *DescribeDistinctReleasesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDistinctReleasesResponseBodyRecords struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribeDistinctReleasesResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetDescription(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

type DescribeDistinctReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDistinctReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDistinctReleasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponse) SetHeaders(v map[string]*string) *DescribeDistinctReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetStatusCode(v int32) *DescribeDistinctReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetBody(v *DescribeDistinctReleasesResponseBody) *DescribeDistinctReleasesResponse {
	s.Body = v
	return s
}

type DescribeEnumItemsRequest struct {
	EnumType *string `json:"EnumType,omitempty" xml:"EnumType,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEnumItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsRequest) SetEnumType(v string) *DescribeEnumItemsRequest {
	s.EnumType = &v
	return s
}

func (s *DescribeEnumItemsRequest) SetLang(v string) *DescribeEnumItemsRequest {
	s.Lang = &v
	return s
}

type DescribeEnumItemsResponseBody struct {
	Data      []*DescribeEnumItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnumItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBody) SetData(v []*DescribeEnumItemsResponseBodyData) *DescribeEnumItemsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnumItemsResponseBody) SetRequestId(v string) *DescribeEnumItemsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnumItemsResponseBodyData struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnumItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBodyData) SetKey(v string) *DescribeEnumItemsResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeEnumItemsResponseBodyData) SetValue(v string) *DescribeEnumItemsResponseBodyData {
	s.Value = &v
	return s
}

type DescribeEnumItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEnumItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEnumItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponse) SetHeaders(v map[string]*string) *DescribeEnumItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnumItemsResponse) SetStatusCode(v int32) *DescribeEnumItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnumItemsResponse) SetBody(v *DescribeEnumItemsResponseBody) *DescribeEnumItemsResponse {
	s.Body = v
	return s
}

type DescribeExecutePlaybooksRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ParamType    *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksRequest) SetLang(v string) *DescribeExecutePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetParamType(v string) *DescribeExecutePlaybooksRequest {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetPlaybookName(v string) *DescribeExecutePlaybooksRequest {
	s.PlaybookName = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetUuid(v string) *DescribeExecutePlaybooksRequest {
	s.Uuid = &v
	return s
}

type DescribeExecutePlaybooksResponseBody struct {
	PlaybookMetrics []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics `json:"PlaybookMetrics,omitempty" xml:"PlaybookMetrics,omitempty" type:"Repeated"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBody) SetPlaybookMetrics(v []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics) *DescribeExecutePlaybooksResponseBody {
	s.PlaybookMetrics = v
	return s
}

func (s *DescribeExecutePlaybooksResponseBody) SetRequestId(v string) *DescribeExecutePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExecutePlaybooksResponseBodyPlaybookMetrics struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ParamConfig *string `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDescription(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Description = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDisplayName(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamConfig(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamConfig = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamType(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetUuid(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Uuid = &v
	return s
}

type DescribeExecutePlaybooksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExecutePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExecutePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponse) SetHeaders(v map[string]*string) *DescribeExecutePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetStatusCode(v int32) *DescribeExecutePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetBody(v *DescribeExecutePlaybooksResponseBody) *DescribeExecutePlaybooksResponse {
	s.Body = v
	return s
}

type DescribeFieldRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	QueryKey *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
}

func (s DescribeFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldRequest) SetLang(v string) *DescribeFieldRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldRequest) SetQueryKey(v string) *DescribeFieldRequest {
	s.QueryKey = &v
	return s
}

type DescribeFieldResponseBody struct {
	Fields    *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponseBody) SetFields(v string) *DescribeFieldResponseBody {
	s.Fields = &v
	return s
}

func (s *DescribeFieldResponseBody) SetName(v string) *DescribeFieldResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFieldResponseBody) SetRequestId(v string) *DescribeFieldResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFieldResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponse) SetHeaders(v map[string]*string) *DescribeFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldResponse) SetStatusCode(v int32) *DescribeFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldResponse) SetBody(v *DescribeFieldResponseBody) *DescribeFieldResponse {
	s.Body = v
	return s
}

type DescribeLatestRecordSchemaRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeLatestRecordSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaRequest) SetLang(v string) *DescribeLatestRecordSchemaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLatestRecordSchemaRequest) SetPlaybookUuid(v string) *DescribeLatestRecordSchemaRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeLatestRecordSchemaResponseBody struct {
	PlaybookNodeSchema *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema `json:"PlaybookNodeSchema,omitempty" xml:"PlaybookNodeSchema,omitempty" type:"Struct"`
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLatestRecordSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBody) SetPlaybookNodeSchema(v *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) *DescribeLatestRecordSchemaResponseBody {
	s.PlaybookNodeSchema = v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBody) SetRequestId(v string) *DescribeLatestRecordSchemaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema struct {
	NodeSchema []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema `json:"NodeSchema,omitempty" xml:"NodeSchema,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) SetNodeSchema(v []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema {
	s.NodeSchema = v
	return s
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema struct {
	ActionName    *string   `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ComponentName *string   `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	NodeName      *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OutputFields  []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetActionName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ActionName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetComponentName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ComponentName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetNodeName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.NodeName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetOutputFields(v []*string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.OutputFields = v
	return s
}

type DescribeLatestRecordSchemaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLatestRecordSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLatestRecordSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponse) SetHeaders(v map[string]*string) *DescribeLatestRecordSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetStatusCode(v int32) *DescribeLatestRecordSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetBody(v *DescribeLatestRecordSchemaResponseBody) *DescribeLatestRecordSchemaResponse {
	s.Body = v
	return s
}

type DescribeNodeParamTagsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeNodeParamTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsRequest) SetLang(v string) *DescribeNodeParamTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetNodeName(v string) *DescribeNodeParamTagsRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetPlaybookUuid(v string) *DescribeNodeParamTagsRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeNodeParamTagsResponseBody struct {
	ParamReferredPaths []*DescribeNodeParamTagsResponseBodyParamReferredPaths `json:"ParamReferredPaths,omitempty" xml:"ParamReferredPaths,omitempty" type:"Repeated"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeParamTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBody) SetParamReferredPaths(v []*DescribeNodeParamTagsResponseBodyParamReferredPaths) *DescribeNodeParamTagsResponseBody {
	s.ParamReferredPaths = v
	return s
}

func (s *DescribeNodeParamTagsResponseBody) SetRequestId(v string) *DescribeNodeParamTagsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodeParamTagsResponseBodyParamReferredPaths struct {
	ParamName    *string   `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ReferredPath []*string `json:"ReferredPath,omitempty" xml:"ReferredPath,omitempty" type:"Repeated"`
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetParamName(v string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ParamName = &v
	return s
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetReferredPath(v []*string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ReferredPath = v
	return s
}

type DescribeNodeParamTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodeParamTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeParamTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponse) SetHeaders(v map[string]*string) *DescribeNodeParamTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetStatusCode(v int32) *DescribeNodeParamTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetBody(v *DescribeNodeParamTagsResponseBody) *DescribeNodeParamTagsResponse {
	s.Body = v
	return s
}

type DescribeNodeUsedInfosRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeNodeUsedInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosRequest) SetLang(v string) *DescribeNodeUsedInfosRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNodeUsedInfosRequest) SetNodeName(v string) *DescribeNodeUsedInfosRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeUsedInfosRequest) SetPlaybookUuid(v string) *DescribeNodeUsedInfosRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeNodeUsedInfosResponseBody struct {
	NodeUsedInfos *string `json:"NodeUsedInfos,omitempty" xml:"NodeUsedInfos,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeUsedInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosResponseBody) SetNodeUsedInfos(v string) *DescribeNodeUsedInfosResponseBody {
	s.NodeUsedInfos = &v
	return s
}

func (s *DescribeNodeUsedInfosResponseBody) SetRequestId(v string) *DescribeNodeUsedInfosResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodeUsedInfosResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodeUsedInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeUsedInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosResponse) SetHeaders(v map[string]*string) *DescribeNodeUsedInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeUsedInfosResponse) SetStatusCode(v int32) *DescribeNodeUsedInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeUsedInfosResponse) SetBody(v *DescribeNodeUsedInfosResponseBody) *DescribeNodeUsedInfosResponse {
	s.Body = v
	return s
}

type DescribePlaybookRequest struct {
	DebugFlag    *int32  `json:"DebugFlag,omitempty" xml:"DebugFlag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	TaskflowMd5  *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookRequest) SetDebugFlag(v int32) *DescribePlaybookRequest {
	s.DebugFlag = &v
	return s
}

func (s *DescribePlaybookRequest) SetLang(v string) *DescribePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookRequest) SetPlaybookUuid(v string) *DescribePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookRequest) SetTaskflowMd5(v string) *DescribePlaybookRequest {
	s.TaskflowMd5 = &v
	return s
}

type DescribePlaybookResponseBody struct {
	Playbook  *DescribePlaybookResponseBodyPlaybook `json:"Playbook,omitempty" xml:"Playbook,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBody) SetPlaybook(v *DescribePlaybookResponseBodyPlaybook) *DescribePlaybookResponseBody {
	s.Playbook = v
	return s
}

func (s *DescribePlaybookResponseBody) SetRequestId(v string) *DescribePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookResponseBodyPlaybook struct {
	Creator                  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName              *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FailExeNum               *int32  `json:"FailExeNum,omitempty" xml:"FailExeNum,omitempty"`
	GmtCreate                *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified              *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InputParams              *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	LastExeTime              *int64  `json:"LastExeTime,omitempty" xml:"LastExeTime,omitempty"`
	Modifier                 *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	OnlineActive             *bool   `json:"OnlineActive,omitempty" xml:"OnlineActive,omitempty"`
	OnlineReleaseTaskflowMd5 *string `json:"OnlineReleaseTaskflowMd5,omitempty" xml:"OnlineReleaseTaskflowMd5,omitempty"`
	OwnType                  *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PlaybookUuid             *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	SuccessExeNum            *int32  `json:"SuccessExeNum,omitempty" xml:"SuccessExeNum,omitempty"`
	Taskflow                 *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s DescribePlaybookResponseBodyPlaybook) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponseBodyPlaybook) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBodyPlaybook) SetCreator(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Creator = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDescription(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Description = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDisplayName(v string) *DescribePlaybookResponseBodyPlaybook {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetFailExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.FailExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtCreate(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtModified(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetInputParams(v string) *DescribePlaybookResponseBodyPlaybook {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetLastExeTime(v int64) *DescribePlaybookResponseBodyPlaybook {
	s.LastExeTime = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetModifier(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Modifier = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineActive(v bool) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineActive = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineReleaseTaskflowMd5(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineReleaseTaskflowMd5 = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOwnType(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetPlaybookUuid(v string) *DescribePlaybookResponseBodyPlaybook {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetSuccessExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.SuccessExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetTaskflow(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Taskflow = &v
	return s
}

type DescribePlaybookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponse) SetHeaders(v map[string]*string) *DescribePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookResponse) SetStatusCode(v int32) *DescribePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookResponse) SetBody(v *DescribePlaybookResponseBody) *DescribePlaybookResponse {
	s.Body = v
	return s
}

type DescribePlaybookInputOutputRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputRequest) SetLang(v string) *DescribePlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookInputOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookInputOutputResponseBody struct {
	Config    *DescribePlaybookInputOutputResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBody) SetConfig(v *DescribePlaybookInputOutputResponseBodyConfig) *DescribePlaybookInputOutputResponseBody {
	s.Config = v
	return s
}

func (s *DescribePlaybookInputOutputResponseBody) SetRequestId(v string) *DescribePlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookInputOutputResponseBodyConfig struct {
	InputParams  *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	ParamType    *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetInputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetOutputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.OutputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetParamType(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.ParamType = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetPlaybookUuid(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookInputOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookInputOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetStatusCode(v int32) *DescribePlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetBody(v *DescribePlaybookInputOutputResponseBody) *DescribePlaybookInputOutputResponse {
	s.Body = v
	return s
}

type DescribePlaybookMetricsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsRequest) SetLang(v string) *DescribePlaybookMetricsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookMetricsRequest) SetPlaybookUuid(v string) *DescribePlaybookMetricsRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookMetricsResponseBody struct {
	Metrics   *DescribePlaybookMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBody) SetMetrics(v *DescribePlaybookMetricsResponseBodyMetrics) *DescribePlaybookMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookMetricsResponseBody) SetRequestId(v string) *DescribePlaybookMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookMetricsResponseBodyMetrics struct {
	Active       *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FailNum      *int32  `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	HistoryMd5   *int32  `json:"HistoryMd5,omitempty" xml:"HistoryMd5,omitempty"`
	LastRuntime  *int64  `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	OwnType      *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	SuccNum      *int32  `json:"SuccNum,omitempty" xml:"SuccNum,omitempty"`
}

func (s DescribePlaybookMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetActive(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Active = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDescription(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDisplayName(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetFailNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.FailNum = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetGmtCreate(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetHistoryMd5(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.HistoryMd5 = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetLastRuntime(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.LastRuntime = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetOwnType(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetPlaybookUuid(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetSuccNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.SuccNum = &v
	return s
}

type DescribePlaybookMetricsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetStatusCode(v int32) *DescribePlaybookMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetBody(v *DescribePlaybookMetricsResponseBody) *DescribePlaybookMetricsResponse {
	s.Body = v
	return s
}

type DescribePlaybookNodesOutputRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookNodesOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputRequest) SetLang(v string) *DescribePlaybookNodesOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetNodeName(v string) *DescribePlaybookNodesOutputRequest {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookNodesOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookNodesOutputResponseBody struct {
	PlaybookNodesOutput *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput `json:"PlaybookNodesOutput,omitempty" xml:"PlaybookNodesOutput,omitempty" type:"Struct"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBody) SetPlaybookNodesOutput(v *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) *DescribePlaybookNodesOutputResponseBody {
	s.PlaybookNodesOutput = v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBody) SetRequestId(v string) *DescribePlaybookNodesOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput struct {
	NodeName   *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeName(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeOutput(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeOutput = &v
	return s
}

type DescribePlaybookNodesOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookNodesOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookNodesOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookNodesOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetStatusCode(v int32) *DescribePlaybookNodesOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetBody(v *DescribePlaybookNodesOutputResponseBody) *DescribePlaybookNodesOutputResponse {
	s.Body = v
	return s
}

type DescribePlaybookNumberMetricsRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePlaybookNumberMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsRequest) SetLang(v string) *DescribePlaybookNumberMetricsRequest {
	s.Lang = &v
	return s
}

type DescribePlaybookNumberMetricsResponseBody struct {
	Metrics   *DescribePlaybookNumberMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetMetrics(v *DescribePlaybookNumberMetricsResponseBodyMetrics) *DescribePlaybookNumberMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetRequestId(v string) *DescribePlaybookNumberMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookNumberMetricsResponseBodyMetrics struct {
	StartUpNum *int32 `json:"StartUpNum,omitempty" xml:"StartUpNum,omitempty"`
	TotalNum   *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetStartUpNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.StartUpNum = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetTotalNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.TotalNum = &v
	return s
}

type DescribePlaybookNumberMetricsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookNumberMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookNumberMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookNumberMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetStatusCode(v int32) *DescribePlaybookNumberMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetBody(v *DescribePlaybookNumberMetricsResponseBody) *DescribePlaybookNumberMetricsResponse {
	s.Body = v
	return s
}

type DescribePlaybookReleasesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookReleasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesRequest) SetLang(v string) *DescribePlaybookReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageNumber(v int32) *DescribePlaybookReleasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageSize(v int32) *DescribePlaybookReleasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPlaybookUuid(v string) *DescribePlaybookReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookReleasesResponseBody struct {
	Page      *DescribePlaybookReleasesResponseBodyPage      `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	Records   []*DescribePlaybookReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookReleasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBody) SetPage(v *DescribePlaybookReleasesResponseBodyPage) *DescribePlaybookReleasesResponseBody {
	s.Page = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRecords(v []*DescribePlaybookReleasesResponseBodyRecords) *DescribePlaybookReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRequestId(v string) *DescribePlaybookReleasesResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookReleasesResponseBodyPage struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageNumber(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageSize(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetTotalCount(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribePlaybookReleasesResponseBodyRecords struct {
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetCreator(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Creator = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetDescription(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtCreate(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtModified(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetId(v int32) *DescribePlaybookReleasesResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

type DescribePlaybookReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybookReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybookReleasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponse) SetHeaders(v map[string]*string) *DescribePlaybookReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetStatusCode(v int32) *DescribePlaybookReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetBody(v *DescribePlaybookReleasesResponseBody) *DescribePlaybookReleasesResponse {
	s.Body = v
	return s
}

type DescribePlaybooksRequest struct {
	Active       *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	EndMillis    *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnType      *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	StartMillis  *int64  `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
}

func (s DescribePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksRequest) SetActive(v int32) *DescribePlaybooksRequest {
	s.Active = &v
	return s
}

func (s *DescribePlaybooksRequest) SetEndMillis(v int64) *DescribePlaybooksRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribePlaybooksRequest) SetLang(v string) *DescribePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybooksRequest) SetName(v string) *DescribePlaybooksRequest {
	s.Name = &v
	return s
}

func (s *DescribePlaybooksRequest) SetOwnType(v string) *DescribePlaybooksRequest {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageNumber(v string) *DescribePlaybooksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageSize(v string) *DescribePlaybooksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPlaybookUuid(v string) *DescribePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybooksRequest) SetStartMillis(v int64) *DescribePlaybooksRequest {
	s.StartMillis = &v
	return s
}

type DescribePlaybooksResponseBody struct {
	Page      *DescribePlaybooksResponseBodyPage        `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	Playbooks []*DescribePlaybooksResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBody) SetPage(v *DescribePlaybooksResponseBodyPage) *DescribePlaybooksResponseBody {
	s.Page = v
	return s
}

func (s *DescribePlaybooksResponseBody) SetPlaybooks(v []*DescribePlaybooksResponseBodyPlaybooks) *DescribePlaybooksResponseBody {
	s.Playbooks = v
	return s
}

func (s *DescribePlaybooksResponseBody) SetRequestId(v string) *DescribePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybooksResponseBodyPage struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlaybooksResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBodyPage) SetPageNumber(v int32) *DescribePlaybooksResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPage) SetPageSize(v int32) *DescribePlaybooksResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPage) SetTotalCount(v int32) *DescribePlaybooksResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribePlaybooksResponseBodyPlaybooks struct {
	Active       *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	LastRuntime  *int64  `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	OwnType      *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybooksResponseBodyPlaybooks) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBodyPlaybooks) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetActive(v int32) *DescribePlaybooksResponseBodyPlaybooks {
	s.Active = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetDisplayName(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetGmtCreate(v int64) *DescribePlaybooksResponseBodyPlaybooks {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetLastRuntime(v int64) *DescribePlaybooksResponseBodyPlaybooks {
	s.LastRuntime = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetOwnType(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetPlaybookUuid(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybooksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponse) SetHeaders(v map[string]*string) *DescribePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybooksResponse) SetStatusCode(v int32) *DescribePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybooksResponse) SetBody(v *DescribePlaybooksResponseBody) *DescribePlaybooksResponse {
	s.Body = v
	return s
}

type DescribePopApiRequest struct {
	ApiName    *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Env        *string `json:"Env,omitempty" xml:"Env,omitempty"`
	PopCode    *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiRequest) SetApiName(v string) *DescribePopApiRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiRequest) SetApiVersion(v string) *DescribePopApiRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribePopApiRequest) SetEnv(v string) *DescribePopApiRequest {
	s.Env = &v
	return s
}

func (s *DescribePopApiRequest) SetPopCode(v string) *DescribePopApiRequest {
	s.PopCode = &v
	return s
}

type DescribePopApiResponseBody struct {
	ApiName         *string                                      `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	OpenApiMetaList []*DescribePopApiResponseBodyOpenApiMetaList `json:"OpenApiMetaList,omitempty" xml:"OpenApiMetaList,omitempty" type:"Repeated"`
	PopCode         *string                                      `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version         *string                                      `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBody) SetApiName(v string) *DescribePopApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiResponseBody) SetOpenApiMetaList(v []*DescribePopApiResponseBodyOpenApiMetaList) *DescribePopApiResponseBody {
	s.OpenApiMetaList = v
	return s
}

func (s *DescribePopApiResponseBody) SetPopCode(v string) *DescribePopApiResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiResponseBody) SetRequestId(v string) *DescribePopApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiResponseBody) SetVersion(v string) *DescribePopApiResponseBody {
	s.Version = &v
	return s
}

type DescribePopApiResponseBodyOpenApiMetaList struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required     *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePopApiResponseBodyOpenApiMetaList) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponseBodyOpenApiMetaList) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetDescription(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Description = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetExampleValue(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.ExampleValue = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetName(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Name = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetRequired(v bool) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Required = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetType(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Type = &v
	return s
}

type DescribePopApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePopApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePopApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponse) SetHeaders(v map[string]*string) *DescribePopApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiResponse) SetStatusCode(v int32) *DescribePopApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiResponse) SetBody(v *DescribePopApiResponseBody) *DescribePopApiResponse {
	s.Body = v
	return s
}

type DescribePopApiItemListRequest struct {
	ApiName    *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Env        *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PopCode    *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListRequest) SetApiName(v string) *DescribePopApiItemListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetApiVersion(v string) *DescribePopApiItemListRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetEnv(v string) *DescribePopApiItemListRequest {
	s.Env = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetLang(v string) *DescribePopApiItemListRequest {
	s.Lang = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetPopCode(v string) *DescribePopApiItemListRequest {
	s.PopCode = &v
	return s
}

type DescribePopApiItemListResponseBody struct {
	Names     []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	PopCode   *string   `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64    `json:"Total,omitempty" xml:"Total,omitempty"`
	Version   *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListResponseBody) SetNames(v []*string) *DescribePopApiItemListResponseBody {
	s.Names = v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetPopCode(v string) *DescribePopApiItemListResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetRequestId(v string) *DescribePopApiItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetTotal(v int64) *DescribePopApiItemListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetVersion(v string) *DescribePopApiItemListResponseBody {
	s.Version = &v
	return s
}

type DescribePopApiItemListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePopApiItemListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePopApiItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListResponse) SetHeaders(v map[string]*string) *DescribePopApiItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiItemListResponse) SetStatusCode(v int32) *DescribePopApiItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiItemListResponse) SetBody(v *DescribePopApiItemListResponseBody) *DescribePopApiItemListResponse {
	s.Body = v
	return s
}

type DescribePopApiVersionListRequest struct {
	Env     *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiVersionListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiVersionListRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiVersionListRequest) SetEnv(v string) *DescribePopApiVersionListRequest {
	s.Env = &v
	return s
}

func (s *DescribePopApiVersionListRequest) SetLang(v string) *DescribePopApiVersionListRequest {
	s.Lang = &v
	return s
}

func (s *DescribePopApiVersionListRequest) SetPopCode(v string) *DescribePopApiVersionListRequest {
	s.PopCode = &v
	return s
}

type DescribePopApiVersionListResponseBody struct {
	PopCode     *string                                             `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	RequestId   *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	VersionList []*DescribePopApiVersionListResponseBodyVersionList `json:"VersionList,omitempty" xml:"VersionList,omitempty" type:"Repeated"`
}

func (s DescribePopApiVersionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiVersionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiVersionListResponseBody) SetPopCode(v string) *DescribePopApiVersionListResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiVersionListResponseBody) SetRequestId(v string) *DescribePopApiVersionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiVersionListResponseBody) SetTotal(v int32) *DescribePopApiVersionListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePopApiVersionListResponseBody) SetVersionList(v []*DescribePopApiVersionListResponseBodyVersionList) *DescribePopApiVersionListResponseBody {
	s.VersionList = v
	return s
}

type DescribePopApiVersionListResponseBodyVersionList struct {
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiVersionListResponseBodyVersionList) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiVersionListResponseBodyVersionList) GoString() string {
	return s.String()
}

func (s *DescribePopApiVersionListResponseBodyVersionList) SetApiName(v string) *DescribePopApiVersionListResponseBodyVersionList {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiVersionListResponseBodyVersionList) SetPopCode(v string) *DescribePopApiVersionListResponseBodyVersionList {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiVersionListResponseBodyVersionList) SetVersion(v string) *DescribePopApiVersionListResponseBodyVersionList {
	s.Version = &v
	return s
}

type DescribePopApiVersionListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePopApiVersionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePopApiVersionListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiVersionListResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiVersionListResponse) SetHeaders(v map[string]*string) *DescribePopApiVersionListResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiVersionListResponse) SetStatusCode(v int32) *DescribePopApiVersionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiVersionListResponse) SetBody(v *DescribePopApiVersionListResponseBody) *DescribePopApiVersionListResponse {
	s.Body = v
	return s
}

type DescribeProcessTasksRequest struct {
	Direction           *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EntityName          *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType          *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OrderField          *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	PageNumber          *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParamContent        *string `json:"ParamContent,omitempty" xml:"ParamContent,omitempty"`
	ProcessActionEnd    *int64  `json:"ProcessActionEnd,omitempty" xml:"ProcessActionEnd,omitempty"`
	ProcessActionStart  *int64  `json:"ProcessActionStart,omitempty" xml:"ProcessActionStart,omitempty"`
	ProcessRemoveEnd    *int64  `json:"ProcessRemoveEnd,omitempty" xml:"ProcessRemoveEnd,omitempty"`
	ProcessRemoveStart  *int64  `json:"ProcessRemoveStart,omitempty" xml:"ProcessRemoveStart,omitempty"`
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	Scope               *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	YunCode             *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksRequest) SetDirection(v string) *DescribeProcessTasksRequest {
	s.Direction = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityName(v string) *DescribeProcessTasksRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityType(v string) *DescribeProcessTasksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetOrderField(v string) *DescribeProcessTasksRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageNumber(v string) *DescribeProcessTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageSize(v int32) *DescribeProcessTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetParamContent(v string) *DescribeProcessTasksRequest {
	s.ParamContent = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessStrategyUuid(v string) *DescribeProcessTasksRequest {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSceneCode(v string) *DescribeProcessTasksRequest {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetScope(v string) *DescribeProcessTasksRequest {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSource(v string) *DescribeProcessTasksRequest {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskId(v string) *DescribeProcessTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskStatus(v string) *DescribeProcessTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetYunCode(v string) *DescribeProcessTasksRequest {
	s.YunCode = &v
	return s
}

type DescribeProcessTasksResponseBody struct {
	Page         *DescribeProcessTasksResponseBodyPage           `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ProcessTasks []*DescribeProcessTasksResponseBodyProcessTasks `json:"ProcessTasks,omitempty" xml:"ProcessTasks,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBody) SetPage(v *DescribeProcessTasksResponseBodyPage) *DescribeProcessTasksResponseBody {
	s.Page = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetProcessTasks(v []*DescribeProcessTasksResponseBodyProcessTasks) *DescribeProcessTasksResponseBody {
	s.ProcessTasks = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetRequestId(v string) *DescribeProcessTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessTasksResponseBodyPage struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessTasksResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageNumber(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageSize(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetTotalCount(v int32) *DescribeProcessTasksResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeProcessTasksResponseBodyProcessTasks struct {
	Creator             *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	EntityName          *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityType          *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	GmtCreateMillis     *int64  `json:"GmtCreateMillis,omitempty" xml:"GmtCreateMillis,omitempty"`
	GmtModifiedMillis   *int64  `json:"GmtModifiedMillis,omitempty" xml:"GmtModifiedMillis,omitempty"`
	InputParams         *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	ProcessTime         *int64  `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	RemoveTime          *int64  `json:"RemoveTime,omitempty" xml:"RemoveTime,omitempty"`
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SceneName           *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	Scope               *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus          *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	YunCode             *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksResponseBodyProcessTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyProcessTasks) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetCreator(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Creator = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityType(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtCreateMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtCreateMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtModifiedMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtModifiedMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetInputParams(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.InputParams = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessStrategyUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetRemoveTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.RemoveTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetScope(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSource(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskId(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskStatus(v int32) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetYunCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.YunCode = &v
	return s
}

type DescribeProcessTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProcessTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProcessTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponse) SetHeaders(v map[string]*string) *DescribeProcessTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessTasksResponse) SetStatusCode(v int32) *DescribeProcessTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessTasksResponse) SetBody(v *DescribeProcessTasksResponseBody) *DescribeProcessTasksResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordActionOutputListRequest struct {
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSoarRecordActionOutputListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListRequest) SetActionUuid(v string) *DescribeSoarRecordActionOutputListRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetLang(v string) *DescribeSoarRecordActionOutputListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageSize(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageSize = &v
	return s
}

type DescribeSoarRecordActionOutputListResponseBody struct {
	ActionOutputs *string `json:"ActionOutputs,omitempty" xml:"ActionOutputs,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarRecordActionOutputListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetActionOutputs(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.ActionOutputs = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageSize(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetRequestId(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetTotalCount(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSoarRecordActionOutputListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSoarRecordActionOutputListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSoarRecordActionOutputListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordActionOutputListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetStatusCode(v int32) *DescribeSoarRecordActionOutputListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetBody(v *DescribeSoarRecordActionOutputListResponseBody) *DescribeSoarRecordActionOutputListResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordInOutputRequest struct {
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSoarRecordInOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputRequest) SetActionUuid(v string) *DescribeSoarRecordInOutputRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordInOutputRequest) SetLang(v string) *DescribeSoarRecordInOutputRequest {
	s.Lang = &v
	return s
}

type DescribeSoarRecordInOutputResponseBody struct {
	InOutputInfo *string `json:"InOutputInfo,omitempty" xml:"InOutputInfo,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarRecordInOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponseBody) SetInOutputInfo(v string) *DescribeSoarRecordInOutputResponseBody {
	s.InOutputInfo = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponseBody) SetRequestId(v string) *DescribeSoarRecordInOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSoarRecordInOutputResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSoarRecordInOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSoarRecordInOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordInOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetStatusCode(v int32) *DescribeSoarRecordInOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetBody(v *DescribeSoarRecordInOutputResponseBody) *DescribeSoarRecordInOutputResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordsRequest struct {
	EndMillis    *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	StartMillis  *int64  `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskflowMd5  *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	TriggerUser  *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsRequest) SetEndMillis(v int64) *DescribeSoarRecordsRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetLang(v string) *DescribeSoarRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageNumber(v int32) *DescribeSoarRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageSize(v int32) *DescribeSoarRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPlaybookUuid(v string) *DescribeSoarRecordsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetStartMillis(v int64) *DescribeSoarRecordsRequest {
	s.StartMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskStatus(v string) *DescribeSoarRecordsRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskflowMd5(v string) *DescribeSoarRecordsRequest {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTriggerUser(v string) *DescribeSoarRecordsRequest {
	s.TriggerUser = &v
	return s
}

type DescribeSoarRecordsResponseBody struct {
	Page               *DescribeSoarRecordsResponseBodyPage                 `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SoarExecuteRecords []*DescribeSoarRecordsResponseBodySoarExecuteRecords `json:"SoarExecuteRecords,omitempty" xml:"SoarExecuteRecords,omitempty" type:"Repeated"`
}

func (s DescribeSoarRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBody) SetPage(v *DescribeSoarRecordsResponseBodyPage) *DescribeSoarRecordsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeSoarRecordsResponseBody) SetRequestId(v string) *DescribeSoarRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordsResponseBody) SetSoarExecuteRecords(v []*DescribeSoarRecordsResponseBodySoarExecuteRecords) *DescribeSoarRecordsResponseBody {
	s.SoarExecuteRecords = v
	return s
}

type DescribeSoarRecordsResponseBodyPage struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarRecordsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBodyPage) SetPageNumber(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodyPage) SetPageSize(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodyPage) SetTotalCount(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeSoarRecordsResponseBodySoarExecuteRecords struct {
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMsg      *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RawEventReq   *string `json:"RawEventReq,omitempty" xml:"RawEventReq,omitempty"`
	RequestUuid   *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskType      *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskflowMd5   *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	TriggerType   *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	TriggerUser   *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarRecordsResponseBodySoarExecuteRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBodySoarExecuteRecords) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetEndTime(v int64) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetErrorMsg(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetRawEventReq(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.RawEventReq = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetRequestUuid(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetResultMessage(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetStartTime(v int64) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetStatus(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.Status = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskName(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskType(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskType = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskflowMd5(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTriggerType(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTriggerUser(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TriggerUser = &v
	return s
}

type DescribeSoarRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSoarRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSoarRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordsResponse) SetStatusCode(v int32) *DescribeSoarRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordsResponse) SetBody(v *DescribeSoarRecordsResponseBody) *DescribeSoarRecordsResponse {
	s.Body = v
	return s
}

type DescribeSoarTaskAndActionsRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescribeSoarTaskAndActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsRequest) SetLang(v string) *DescribeSoarTaskAndActionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetRequestUuid(v string) *DescribeSoarTaskAndActionsRequest {
	s.RequestUuid = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBody struct {
	Details   *DescribeSoarTaskAndActionsResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetDetails(v *DescribeSoarTaskAndActionsResponseBodyDetails) *DescribeSoarTaskAndActionsResponseBody {
	s.Details = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetRequestId(v string) *DescribeSoarTaskAndActionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBodyDetails struct {
	Actions       []*DescribeSoarTaskAndActionsResponseBodyDetailsActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	EndTime       *int64                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMsg      *string                                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RawEventReq   *string                                                 `json:"RawEventReq,omitempty" xml:"RawEventReq,omitempty"`
	RequestUuid   *string                                                 `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	ResultLevel   *string                                                 `json:"ResultLevel,omitempty" xml:"ResultLevel,omitempty"`
	ResultMessage *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	StartTime     *int64                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskFlowMd5   *string                                                 `json:"TaskFlowMd5,omitempty" xml:"TaskFlowMd5,omitempty"`
	TaskName      *string                                                 `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskTenantId  *string                                                 `json:"TaskTenantId,omitempty" xml:"TaskTenantId,omitempty"`
	TriggerType   *string                                                 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	TriggerUser   *string                                                 `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetActions(v []*DescribeSoarTaskAndActionsResponseBodyDetailsActions) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Actions = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetErrorMsg(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRawEventReq(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RawEventReq = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRequestUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetResultLevel(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ResultLevel = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetResultMessage(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ResultMessage = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskFlowMd5(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskFlowMd5 = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskName(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskTenantId(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskTenantId = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerType(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerUser(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerUser = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBodyDetailsActions struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionUuid  *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	AssetName   *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	Component   *string `json:"Component,omitempty" xml:"Component,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NodeName    *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus  *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAction(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Action = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetActionUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAssetName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.AssetName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetComponent(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Component = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetNodeName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.NodeName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetRequestUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTaskName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTaskStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTriggerUser(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TriggerUser = &v
	return s
}

type DescribeSoarTaskAndActionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSoarTaskAndActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSoarTaskAndActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponse) SetHeaders(v map[string]*string) *DescribeSoarTaskAndActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetStatusCode(v int32) *DescribeSoarTaskAndActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetBody(v *DescribeSoarTaskAndActionsResponseBody) *DescribeSoarTaskAndActionsResponse {
	s.Body = v
	return s
}

type DescribeSophonCommandsRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSophonCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsRequest) SetName(v string) *DescribeSophonCommandsRequest {
	s.Name = &v
	return s
}

type DescribeSophonCommandsResponseBody struct {
	Data      []*DescribeSophonCommandsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSophonCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBody) SetData(v []*DescribeSophonCommandsResponseBodyData) *DescribeSophonCommandsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSophonCommandsResponseBody) SetRequestId(v string) *DescribeSophonCommandsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSophonCommandsResponseBodyData struct {
	Description *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string                                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamConfig []*DescribeSophonCommandsResponseBodyDataParamConfig `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty" type:"Repeated"`
}

func (s DescribeSophonCommandsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyData) SetDescription(v string) *DescribeSophonCommandsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetDisplayName(v string) *DescribeSophonCommandsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetName(v string) *DescribeSophonCommandsResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetParamConfig(v []*DescribeSophonCommandsResponseBodyDataParamConfig) *DescribeSophonCommandsResponseBodyData {
	s.ParamConfig = v
	return s
}

type DescribeSophonCommandsResponseBodyDataParamConfig struct {
	CheckField *string `json:"CheckField,omitempty" xml:"CheckField,omitempty"`
	Field      *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Necessary  *bool   `json:"Necessary,omitempty" xml:"Necessary,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetCheckField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.CheckField = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Field = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetNecessary(v bool) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Necessary = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetValue(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Value = &v
	return s
}

type DescribeSophonCommandsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSophonCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSophonCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponse) SetHeaders(v map[string]*string) *DescribeSophonCommandsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSophonCommandsResponse) SetStatusCode(v int32) *DescribeSophonCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSophonCommandsResponse) SetBody(v *DescribeSophonCommandsResponseBody) *DescribeSophonCommandsResponse {
	s.Body = v
	return s
}

type DescriberPython3ScriptLogsRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescriberPython3ScriptLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsRequest) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsRequest) SetLang(v string) *DescriberPython3ScriptLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescriberPython3ScriptLogsRequest) SetRequestUuid(v string) *DescriberPython3ScriptLogsRequest {
	s.RequestUuid = &v
	return s
}

type DescriberPython3ScriptLogsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s DescriberPython3ScriptLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRequestId(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRunResult(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RunResult = &v
	return s
}

type DescriberPython3ScriptLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescriberPython3ScriptLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescriberPython3ScriptLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponse) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponse) SetHeaders(v map[string]*string) *DescriberPython3ScriptLogsResponse {
	s.Headers = v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetStatusCode(v int32) *DescriberPython3ScriptLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetBody(v *DescriberPython3ScriptLogsResponseBody) *DescriberPython3ScriptLogsResponse {
	s.Body = v
	return s
}

type ModifyComponentAssetRequest struct {
	AssetConfig *string `json:"AssetConfig,omitempty" xml:"AssetConfig,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyComponentAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetRequest) SetAssetConfig(v string) *ModifyComponentAssetRequest {
	s.AssetConfig = &v
	return s
}

func (s *ModifyComponentAssetRequest) SetLang(v string) *ModifyComponentAssetRequest {
	s.Lang = &v
	return s
}

type ModifyComponentAssetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyComponentAssetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponseBody) SetRequestId(v string) *ModifyComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyComponentAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponse) SetHeaders(v map[string]*string) *ModifyComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *ModifyComponentAssetResponse) SetStatusCode(v int32) *ModifyComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyComponentAssetResponse) SetBody(v *ModifyComponentAssetResponseBody) *ModifyComponentAssetResponse {
	s.Body = v
	return s
}

type ModifyPlaybookRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	Taskflow     *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s ModifyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookRequest) SetDescription(v string) *ModifyPlaybookRequest {
	s.Description = &v
	return s
}

func (s *ModifyPlaybookRequest) SetDisplayName(v string) *ModifyPlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyPlaybookRequest) SetLang(v string) *ModifyPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookRequest) SetPlaybookUuid(v string) *ModifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ModifyPlaybookRequest) SetTaskflow(v string) *ModifyPlaybookRequest {
	s.Taskflow = &v
	return s
}

type ModifyPlaybookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponseBody) SetRequestId(v string) *ModifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponse) SetHeaders(v map[string]*string) *ModifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookResponse) SetStatusCode(v int32) *ModifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookResponse) SetBody(v *ModifyPlaybookResponseBody) *ModifyPlaybookResponse {
	s.Body = v
	return s
}

type ModifyPlaybookInputOutputRequest struct {
	InputParams  *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	ParamType    *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInputOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputRequest) SetInputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.InputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetLang(v string) *ModifyPlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetOutputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.OutputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetParamType(v string) *ModifyPlaybookInputOutputRequest {
	s.ParamType = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type ModifyPlaybookInputOutputResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookInputOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponseBody) SetRequestId(v string) *ModifyPlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookInputOutputResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPlaybookInputOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponse) SetHeaders(v map[string]*string) *ModifyPlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetStatusCode(v int32) *ModifyPlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetBody(v *ModifyPlaybookInputOutputResponseBody) *ModifyPlaybookInputOutputResponse {
	s.Body = v
	return s
}

type ModifyPlaybookInstanceStatusRequest struct {
	Active       *int32  `json:"Active,omitempty" xml:"Active,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusRequest) SetActive(v int32) *ModifyPlaybookInstanceStatusRequest {
	s.Active = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusRequest) SetLang(v string) *ModifyPlaybookInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusRequest) SetPlaybookUuid(v string) *ModifyPlaybookInstanceStatusRequest {
	s.PlaybookUuid = &v
	return s
}

type ModifyPlaybookInstanceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusResponseBody) SetRequestId(v string) *ModifyPlaybookInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookInstanceStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPlaybookInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPlaybookInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusResponse) SetHeaders(v map[string]*string) *ModifyPlaybookInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookInstanceStatusResponse) SetStatusCode(v int32) *ModifyPlaybookInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusResponse) SetBody(v *ModifyPlaybookInstanceStatusResponseBody) *ModifyPlaybookInstanceStatusResponse {
	s.Body = v
	return s
}

type PublishPlaybookRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s PublishPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookRequest) GoString() string {
	return s.String()
}

func (s *PublishPlaybookRequest) SetDescription(v string) *PublishPlaybookRequest {
	s.Description = &v
	return s
}

func (s *PublishPlaybookRequest) SetPlaybookUuid(v string) *PublishPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type PublishPlaybookResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponseBody) SetRequestId(v string) *PublishPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type PublishPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookResponse) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponse) SetHeaders(v map[string]*string) *PublishPlaybookResponse {
	s.Headers = v
	return s
}

func (s *PublishPlaybookResponse) SetStatusCode(v int32) *PublishPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishPlaybookResponse) SetBody(v *PublishPlaybookResponseBody) *PublishPlaybookResponse {
	s.Body = v
	return s
}

type QueryTreeDataRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryTreeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTreeDataRequest) SetLang(v string) *QueryTreeDataRequest {
	s.Lang = &v
	return s
}

type QueryTreeDataResponseBody struct {
	Playbooks *string `json:"Playbooks,omitempty" xml:"Playbooks,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTreeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponseBody) SetPlaybooks(v string) *QueryTreeDataResponseBody {
	s.Playbooks = &v
	return s
}

func (s *QueryTreeDataResponseBody) SetRequestId(v string) *QueryTreeDataResponseBody {
	s.RequestId = &v
	return s
}

type QueryTreeDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTreeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTreeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponse) SetHeaders(v map[string]*string) *QueryTreeDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTreeDataResponse) SetStatusCode(v int32) *QueryTreeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTreeDataResponse) SetBody(v *QueryTreeDataResponseBody) *QueryTreeDataResponse {
	s.Body = v
	return s
}

type RenamePlaybookNodeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewNodeName  *string `json:"NewNodeName,omitempty" xml:"NewNodeName,omitempty"`
	OldNodeName  *string `json:"OldNodeName,omitempty" xml:"OldNodeName,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s RenamePlaybookNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeRequest) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeRequest) SetLang(v string) *RenamePlaybookNodeRequest {
	s.Lang = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetNewNodeName(v string) *RenamePlaybookNodeRequest {
	s.NewNodeName = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetOldNodeName(v string) *RenamePlaybookNodeRequest {
	s.OldNodeName = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetPlaybookUuid(v string) *RenamePlaybookNodeRequest {
	s.PlaybookUuid = &v
	return s
}

type RenamePlaybookNodeResponseBody struct {
	RenameResult *string `json:"RenameResult,omitempty" xml:"RenameResult,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenamePlaybookNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeResponseBody) SetRenameResult(v string) *RenamePlaybookNodeResponseBody {
	s.RenameResult = &v
	return s
}

func (s *RenamePlaybookNodeResponseBody) SetRequestId(v string) *RenamePlaybookNodeResponseBody {
	s.RequestId = &v
	return s
}

type RenamePlaybookNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenamePlaybookNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenamePlaybookNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeResponse) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeResponse) SetHeaders(v map[string]*string) *RenamePlaybookNodeResponse {
	s.Headers = v
	return s
}

func (s *RenamePlaybookNodeResponse) SetStatusCode(v int32) *RenamePlaybookNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenamePlaybookNodeResponse) SetBody(v *RenamePlaybookNodeResponseBody) *RenamePlaybookNodeResponse {
	s.Body = v
	return s
}

type RevertPlaybookReleaseRequest struct {
	IsPublish     *bool   `json:"IsPublish,omitempty" xml:"IsPublish,omitempty"`
	PlayReleaseId *int32  `json:"PlayReleaseId,omitempty" xml:"PlayReleaseId,omitempty"`
	PlaybookUuid  *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s RevertPlaybookReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseRequest) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseRequest) SetIsPublish(v bool) *RevertPlaybookReleaseRequest {
	s.IsPublish = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlayReleaseId(v int32) *RevertPlaybookReleaseRequest {
	s.PlayReleaseId = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlaybookUuid(v string) *RevertPlaybookReleaseRequest {
	s.PlaybookUuid = &v
	return s
}

type RevertPlaybookReleaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertPlaybookReleaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponseBody) SetRequestId(v string) *RevertPlaybookReleaseResponseBody {
	s.RequestId = &v
	return s
}

type RevertPlaybookReleaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevertPlaybookReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevertPlaybookReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseResponse) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponse) SetHeaders(v map[string]*string) *RevertPlaybookReleaseResponse {
	s.Headers = v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetStatusCode(v int32) *RevertPlaybookReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetBody(v *RevertPlaybookReleaseResponseBody) *RevertPlaybookReleaseResponse {
	s.Body = v
	return s
}

type RunPython3ScriptRequest struct {
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Params       *string `json:"Params,omitempty" xml:"Params,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	PythonScript *string `json:"PythonScript,omitempty" xml:"PythonScript,omitempty"`
}

func (s RunPython3ScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptRequest) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptRequest) SetNodeName(v string) *RunPython3ScriptRequest {
	s.NodeName = &v
	return s
}

func (s *RunPython3ScriptRequest) SetParams(v string) *RunPython3ScriptRequest {
	s.Params = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPlaybookUuid(v string) *RunPython3ScriptRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPythonScript(v string) *RunPython3ScriptRequest {
	s.PythonScript = &v
	return s
}

type RunPython3ScriptResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s RunPython3ScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptResponseBody) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponseBody) SetRequestId(v string) *RunPython3ScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPython3ScriptResponseBody) SetRunResult(v string) *RunPython3ScriptResponseBody {
	s.RunResult = &v
	return s
}

type RunPython3ScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunPython3ScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunPython3ScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptResponse) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponse) SetHeaders(v map[string]*string) *RunPython3ScriptResponse {
	s.Headers = v
	return s
}

func (s *RunPython3ScriptResponse) SetStatusCode(v int32) *RunPython3ScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPython3ScriptResponse) SetBody(v *RunPython3ScriptResponseBody) *RunPython3ScriptResponse {
	s.Body = v
	return s
}

type TriggerPlaybookRequest struct {
	InputParam   *string `json:"InputParam,omitempty" xml:"InputParam,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s TriggerPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookRequest) SetInputParam(v string) *TriggerPlaybookRequest {
	s.InputParam = &v
	return s
}

func (s *TriggerPlaybookRequest) SetPlaybookUuid(v string) *TriggerPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type TriggerPlaybookResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TriggerUuid *string `json:"TriggerUuid,omitempty" xml:"TriggerUuid,omitempty"`
}

func (s TriggerPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponseBody) SetRequestId(v string) *TriggerPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerPlaybookResponseBody) SetTriggerUuid(v string) *TriggerPlaybookResponseBody {
	s.TriggerUuid = &v
	return s
}

type TriggerPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponse) SetHeaders(v map[string]*string) *TriggerPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerPlaybookResponse) SetStatusCode(v int32) *TriggerPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerPlaybookResponse) SetBody(v *TriggerPlaybookResponseBody) *TriggerPlaybookResponse {
	s.Body = v
	return s
}

type TriggerProcessTaskRequest struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TriggerProcessTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskRequest) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskRequest) SetActionType(v string) *TriggerProcessTaskRequest {
	s.ActionType = &v
	return s
}

func (s *TriggerProcessTaskRequest) SetTaskId(v string) *TriggerProcessTaskRequest {
	s.TaskId = &v
	return s
}

type TriggerProcessTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerProcessTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponseBody) SetRequestId(v string) *TriggerProcessTaskResponseBody {
	s.RequestId = &v
	return s
}

type TriggerProcessTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerProcessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerProcessTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskResponse) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponse) SetHeaders(v map[string]*string) *TriggerProcessTaskResponse {
	s.Headers = v
	return s
}

func (s *TriggerProcessTaskResponse) SetStatusCode(v int32) *TriggerProcessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerProcessTaskResponse) SetBody(v *TriggerProcessTaskResponseBody) *TriggerProcessTaskResponse {
	s.Body = v
	return s
}

type TriggerSophonPlaybookRequest struct {
	CommandName  *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	InputParams  *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	TriggerType  *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s TriggerSophonPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookRequest) SetCommandName(v string) *TriggerSophonPlaybookRequest {
	s.CommandName = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetInputParams(v string) *TriggerSophonPlaybookRequest {
	s.InputParams = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetSophonTaskId(v string) *TriggerSophonPlaybookRequest {
	s.SophonTaskId = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetTriggerType(v string) *TriggerSophonPlaybookRequest {
	s.TriggerType = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetUuid(v string) *TriggerSophonPlaybookRequest {
	s.Uuid = &v
	return s
}

type TriggerSophonPlaybookResponseBody struct {
	Data      *TriggerSophonPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBody) SetData(v *TriggerSophonPlaybookResponseBodyData) *TriggerSophonPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *TriggerSophonPlaybookResponseBody) SetRequestId(v string) *TriggerSophonPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type TriggerSophonPlaybookResponseBodyData struct {
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBodyData) SetSophonTaskId(v string) *TriggerSophonPlaybookResponseBodyData {
	s.SophonTaskId = &v
	return s
}

type TriggerSophonPlaybookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerSophonPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerSophonPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponse) SetHeaders(v map[string]*string) *TriggerSophonPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetStatusCode(v int32) *TriggerSophonPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetBody(v *TriggerSophonPlaybookResponseBody) *TriggerSophonPlaybookResponse {
	s.Body = v
	return s
}

type VerifyPlaybookRequest struct {
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	TaskFlow     *string `json:"TaskFlow,omitempty" xml:"TaskFlow,omitempty"`
}

func (s VerifyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookRequest) SetPlaybookUuid(v string) *VerifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *VerifyPlaybookRequest) SetTaskFlow(v string) *VerifyPlaybookRequest {
	s.TaskFlow = &v
	return s
}

type VerifyPlaybookResponseBody struct {
	CheckTaskInfos []*VerifyPlaybookResponseBodyCheckTaskInfos `json:"CheckTaskInfos,omitempty" xml:"CheckTaskInfos,omitempty" type:"Repeated"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBody) SetCheckTaskInfos(v []*VerifyPlaybookResponseBodyCheckTaskInfos) *VerifyPlaybookResponseBody {
	s.CheckTaskInfos = v
	return s
}

func (s *VerifyPlaybookResponseBody) SetRequestId(v string) *VerifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type VerifyPlaybookResponseBodyCheckTaskInfos struct {
	Detail    *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	NodeName  *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetDetail(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.Detail = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetNodeName(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.NodeName = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetRiskLevel(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.RiskLevel = &v
	return s
}

type VerifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponse) SetHeaders(v map[string]*string) *VerifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *VerifyPlaybookResponse) SetStatusCode(v int32) *VerifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPlaybookResponse) SetBody(v *VerifyPlaybookResponseBody) *VerifyPlaybookResponse {
	s.Body = v
	return s
}

type VerifyPythonFileRequest struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s VerifyPythonFileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileRequest) SetContent(v string) *VerifyPythonFileRequest {
	s.Content = &v
	return s
}

type VerifyPythonFileResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Syntax    []*VerifyPythonFileResponseBodySyntax `json:"Syntax,omitempty" xml:"Syntax,omitempty" type:"Repeated"`
}

func (s VerifyPythonFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBody) SetRequestId(v string) *VerifyPythonFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPythonFileResponseBody) SetSyntax(v []*VerifyPythonFileResponseBodySyntax) *VerifyPythonFileResponseBody {
	s.Syntax = v
	return s
}

type VerifyPythonFileResponseBodySyntax struct {
	EndColumn       *int32  `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	EndLineNumber   *int32  `json:"EndLineNumber,omitempty" xml:"EndLineNumber,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Severity        *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartColumn     *int32  `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
	StartLineNumber *int32  `json:"StartLineNumber,omitempty" xml:"StartLineNumber,omitempty"`
}

func (s VerifyPythonFileResponseBodySyntax) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponseBodySyntax) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndLineNumber = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetMessage(v string) *VerifyPythonFileResponseBodySyntax {
	s.Message = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetSeverity(v int32) *VerifyPythonFileResponseBodySyntax {
	s.Severity = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartLineNumber = &v
	return s
}

type VerifyPythonFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyPythonFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyPythonFileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponse) SetHeaders(v map[string]*string) *VerifyPythonFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyPythonFileResponse) SetStatusCode(v int32) *VerifyPythonFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPythonFileResponse) SetBody(v *VerifyPythonFileResponseBody) *VerifyPythonFileResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sophonsoar"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchModifyInstanceStatusWithOptions(request *BatchModifyInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *BatchModifyInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchModifyInstanceStatus"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchModifyInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchModifyInstanceStatus(request *BatchModifyInstanceStatusRequest) (_result *BatchModifyInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchModifyInstanceStatusResponse{}
	_body, _err := client.BatchModifyInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ComparePlaybooksWithOptions(request *ComparePlaybooksRequest, runtime *util.RuntimeOptions) (_result *ComparePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewPlaybookReleaseId)) {
		query["NewPlaybookReleaseId"] = request.NewPlaybookReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPlaybookReleaseId)) {
		query["OldPlaybookReleaseId"] = request.OldPlaybookReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ComparePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ComparePlaybooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ComparePlaybooks(request *ComparePlaybooksRequest) (_result *ComparePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ComparePlaybooksResponse{}
	_body, _err := client.ComparePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePlaybookWithOptions(request *CreatePlaybookRequest, runtime *util.RuntimeOptions) (_result *CreatePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePlaybook(request *CreatePlaybookRequest) (_result *CreatePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePlaybookResponse{}
	_body, _err := client.CreatePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DebugPlaybookWithOptions(request *DebugPlaybookRequest, runtime *util.RuntimeOptions) (_result *DebugPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Record)) {
		body["Record"] = request.Record
	}

	if !tea.BoolValue(util.IsUnset(request.Taskflow)) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DebugPlaybook(request *DebugPlaybookRequest) (_result *DebugPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugPlaybookResponse{}
	_body, _err := client.DebugPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteComponentAssetWithOptions(request *DeleteComponentAssetRequest, runtime *util.RuntimeOptions) (_result *DeleteComponentAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		query["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComponentAsset"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteComponentAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComponentAsset(request *DeleteComponentAssetRequest) (_result *DeleteComponentAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteComponentAssetResponse{}
	_body, _err := client.DeleteComponentAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePlaybookWithOptions(request *DeletePlaybookRequest, runtime *util.RuntimeOptions) (_result *DeletePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePlaybook(request *DeletePlaybookRequest) (_result *DeletePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePlaybookResponse{}
	_body, _err := client.DeletePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiListWithOptions(request *DescribeApiListRequest, runtime *util.RuntimeOptions) (_result *DescribeApiListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiList(request *DescribeApiListRequest) (_result *DescribeApiListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiListResponse{}
	_body, _err := client.DescribeApiListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentAssetFormWithOptions(request *DescribeComponentAssetFormRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentAssetFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentAssetForm"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentAssetFormResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentAssetForm(request *DescribeComponentAssetFormRequest) (_result *DescribeComponentAssetFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentAssetFormResponse{}
	_body, _err := client.DescribeComponentAssetFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentAssetsWithOptions(request *DescribeComponentAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentAssets"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentAssets(request *DescribeComponentAssetsRequest) (_result *DescribeComponentAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentAssetsResponse{}
	_body, _err := client.DescribeComponentAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentListWithOptions(request *DescribeComponentListRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentList(request *DescribeComponentListRequest) (_result *DescribeComponentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentListResponse{}
	_body, _err := client.DescribeComponentListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentPlaybookWithOptions(request *DescribeComponentPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentPlaybook(request *DescribeComponentPlaybookRequest) (_result *DescribeComponentPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentPlaybookResponse{}
	_body, _err := client.DescribeComponentPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentsJsWithOptions(request *DescribeComponentsJsRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentsJsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentsJs"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentsJsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponentsJs(request *DescribeComponentsJsRequest) (_result *DescribeComponentsJsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentsJsResponse{}
	_body, _err := client.DescribeComponentsJsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDistinctReleasesWithOptions(request *DescribeDistinctReleasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDistinctReleasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDistinctReleases"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDistinctReleasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDistinctReleases(request *DescribeDistinctReleasesRequest) (_result *DescribeDistinctReleasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDistinctReleasesResponse{}
	_body, _err := client.DescribeDistinctReleasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEnumItemsWithOptions(request *DescribeEnumItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeEnumItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEnumItems"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEnumItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEnumItems(request *DescribeEnumItemsRequest) (_result *DescribeEnumItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnumItemsResponse{}
	_body, _err := client.DescribeEnumItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExecutePlaybooksWithOptions(request *DescribeExecutePlaybooksRequest, runtime *util.RuntimeOptions) (_result *DescribeExecutePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExecutePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExecutePlaybooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExecutePlaybooks(request *DescribeExecutePlaybooksRequest) (_result *DescribeExecutePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExecutePlaybooksResponse{}
	_body, _err := client.DescribeExecutePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFieldWithOptions(request *DescribeFieldRequest, runtime *util.RuntimeOptions) (_result *DescribeFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeField"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeField(request *DescribeFieldRequest) (_result *DescribeFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFieldResponse{}
	_body, _err := client.DescribeFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLatestRecordSchemaWithOptions(request *DescribeLatestRecordSchemaRequest, runtime *util.RuntimeOptions) (_result *DescribeLatestRecordSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLatestRecordSchema"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLatestRecordSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLatestRecordSchema(request *DescribeLatestRecordSchemaRequest) (_result *DescribeLatestRecordSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLatestRecordSchemaResponse{}
	_body, _err := client.DescribeLatestRecordSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeParamTagsWithOptions(request *DescribeNodeParamTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeParamTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeParamTags"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeParamTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeParamTags(request *DescribeNodeParamTagsRequest) (_result *DescribeNodeParamTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeParamTagsResponse{}
	_body, _err := client.DescribeNodeParamTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeUsedInfosWithOptions(request *DescribeNodeUsedInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeUsedInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeUsedInfos"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeUsedInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeUsedInfos(request *DescribeNodeUsedInfosRequest) (_result *DescribeNodeUsedInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeUsedInfosResponse{}
	_body, _err := client.DescribeNodeUsedInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookWithOptions(request *DescribePlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybook(request *DescribePlaybookRequest) (_result *DescribePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookResponse{}
	_body, _err := client.DescribePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookInputOutputWithOptions(request *DescribePlaybookInputOutputRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookInputOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookInputOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookInputOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybookInputOutput(request *DescribePlaybookInputOutputRequest) (_result *DescribePlaybookInputOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookInputOutputResponse{}
	_body, _err := client.DescribePlaybookInputOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookMetricsWithOptions(request *DescribePlaybookMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookMetrics"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybookMetrics(request *DescribePlaybookMetricsRequest) (_result *DescribePlaybookMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookMetricsResponse{}
	_body, _err := client.DescribePlaybookMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookNodesOutputWithOptions(request *DescribePlaybookNodesOutputRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookNodesOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookNodesOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookNodesOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybookNodesOutput(request *DescribePlaybookNodesOutputRequest) (_result *DescribePlaybookNodesOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookNodesOutputResponse{}
	_body, _err := client.DescribePlaybookNodesOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookNumberMetricsWithOptions(request *DescribePlaybookNumberMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookNumberMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookNumberMetrics"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookNumberMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybookNumberMetrics(request *DescribePlaybookNumberMetricsRequest) (_result *DescribePlaybookNumberMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookNumberMetricsResponse{}
	_body, _err := client.DescribePlaybookNumberMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybookReleasesWithOptions(request *DescribePlaybookReleasesRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookReleasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookReleases"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybookReleasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybookReleases(request *DescribePlaybookReleasesRequest) (_result *DescribePlaybookReleasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookReleasesResponse{}
	_body, _err := client.DescribePlaybookReleasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlaybooksWithOptions(request *DescribePlaybooksRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlaybooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlaybooks(request *DescribePlaybooksRequest) (_result *DescribePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybooksResponse{}
	_body, _err := client.DescribePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePopApiWithOptions(request *DescribePopApiRequest, runtime *util.RuntimeOptions) (_result *DescribePopApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePopApi"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePopApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePopApi(request *DescribePopApiRequest) (_result *DescribePopApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePopApiResponse{}
	_body, _err := client.DescribePopApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePopApiItemListWithOptions(request *DescribePopApiItemListRequest, runtime *util.RuntimeOptions) (_result *DescribePopApiItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePopApiItemList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePopApiItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePopApiItemList(request *DescribePopApiItemListRequest) (_result *DescribePopApiItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePopApiItemListResponse{}
	_body, _err := client.DescribePopApiItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePopApiVersionListWithOptions(request *DescribePopApiVersionListRequest, runtime *util.RuntimeOptions) (_result *DescribePopApiVersionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePopApiVersionList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePopApiVersionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePopApiVersionList(request *DescribePopApiVersionListRequest) (_result *DescribePopApiVersionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePopApiVersionListResponse{}
	_body, _err := client.DescribePopApiVersionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProcessTasksWithOptions(request *DescribeProcessTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessTasks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProcessTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProcessTasks(request *DescribeProcessTasksRequest) (_result *DescribeProcessTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessTasksResponse{}
	_body, _err := client.DescribeProcessTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSoarRecordActionOutputListWithOptions(request *DescribeSoarRecordActionOutputListRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordActionOutputListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecordActionOutputList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSoarRecordActionOutputListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSoarRecordActionOutputList(request *DescribeSoarRecordActionOutputListRequest) (_result *DescribeSoarRecordActionOutputListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordActionOutputListResponse{}
	_body, _err := client.DescribeSoarRecordActionOutputListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSoarRecordInOutputWithOptions(request *DescribeSoarRecordInOutputRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordInOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecordInOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSoarRecordInOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSoarRecordInOutput(request *DescribeSoarRecordInOutputRequest) (_result *DescribeSoarRecordInOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordInOutputResponse{}
	_body, _err := client.DescribeSoarRecordInOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSoarRecordsWithOptions(request *DescribeSoarRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecords"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSoarRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSoarRecords(request *DescribeSoarRecordsRequest) (_result *DescribeSoarRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordsResponse{}
	_body, _err := client.DescribeSoarRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSoarTaskAndActionsWithOptions(request *DescribeSoarTaskAndActionsRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarTaskAndActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarTaskAndActions"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSoarTaskAndActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSoarTaskAndActions(request *DescribeSoarTaskAndActionsRequest) (_result *DescribeSoarTaskAndActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarTaskAndActionsResponse{}
	_body, _err := client.DescribeSoarTaskAndActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSophonCommandsWithOptions(request *DescribeSophonCommandsRequest, runtime *util.RuntimeOptions) (_result *DescribeSophonCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSophonCommands"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSophonCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSophonCommands(request *DescribeSophonCommandsRequest) (_result *DescribeSophonCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSophonCommandsResponse{}
	_body, _err := client.DescribeSophonCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescriberPython3ScriptLogsWithOptions(request *DescriberPython3ScriptLogsRequest, runtime *util.RuntimeOptions) (_result *DescriberPython3ScriptLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescriberPython3ScriptLogs"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescriberPython3ScriptLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescriberPython3ScriptLogs(request *DescriberPython3ScriptLogsRequest) (_result *DescriberPython3ScriptLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescriberPython3ScriptLogsResponse{}
	_body, _err := client.DescriberPython3ScriptLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyComponentAssetWithOptions(request *ModifyComponentAssetRequest, runtime *util.RuntimeOptions) (_result *ModifyComponentAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetConfig)) {
		query["AssetConfig"] = request.AssetConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyComponentAsset"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyComponentAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyComponentAsset(request *ModifyComponentAssetRequest) (_result *ModifyComponentAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyComponentAssetResponse{}
	_body, _err := client.ModifyComponentAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPlaybookWithOptions(request *ModifyPlaybookRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Taskflow)) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPlaybook(request *ModifyPlaybookRequest) (_result *ModifyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookResponse{}
	_body, _err := client.ModifyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPlaybookInputOutputWithOptions(request *ModifyPlaybookInputOutputRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookInputOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputParams)) {
		body["InputParams"] = request.InputParams
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OutputParams)) {
		body["OutputParams"] = request.OutputParams
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybookInputOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPlaybookInputOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPlaybookInputOutput(request *ModifyPlaybookInputOutputRequest) (_result *ModifyPlaybookInputOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookInputOutputResponse{}
	_body, _err := client.ModifyPlaybookInputOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPlaybookInstanceStatusWithOptions(request *ModifyPlaybookInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybookInstanceStatus"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPlaybookInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPlaybookInstanceStatus(request *ModifyPlaybookInstanceStatusRequest) (_result *ModifyPlaybookInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookInstanceStatusResponse{}
	_body, _err := client.ModifyPlaybookInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishPlaybookWithOptions(request *PublishPlaybookRequest, runtime *util.RuntimeOptions) (_result *PublishPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishPlaybook(request *PublishPlaybookRequest) (_result *PublishPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishPlaybookResponse{}
	_body, _err := client.PublishPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTreeDataWithOptions(request *QueryTreeDataRequest, runtime *util.RuntimeOptions) (_result *QueryTreeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTreeData"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTreeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTreeData(request *QueryTreeDataRequest) (_result *QueryTreeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTreeDataResponse{}
	_body, _err := client.QueryTreeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenamePlaybookNodeWithOptions(request *RenamePlaybookNodeRequest, runtime *util.RuntimeOptions) (_result *RenamePlaybookNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewNodeName)) {
		query["NewNodeName"] = request.NewNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.OldNodeName)) {
		query["OldNodeName"] = request.OldNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenamePlaybookNode"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenamePlaybookNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenamePlaybookNode(request *RenamePlaybookNodeRequest) (_result *RenamePlaybookNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenamePlaybookNodeResponse{}
	_body, _err := client.RenamePlaybookNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevertPlaybookReleaseWithOptions(request *RevertPlaybookReleaseRequest, runtime *util.RuntimeOptions) (_result *RevertPlaybookReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsPublish)) {
		body["IsPublish"] = request.IsPublish
	}

	if !tea.BoolValue(util.IsUnset(request.PlayReleaseId)) {
		body["PlayReleaseId"] = request.PlayReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertPlaybookRelease"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertPlaybookReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevertPlaybookRelease(request *RevertPlaybookReleaseRequest) (_result *RevertPlaybookReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertPlaybookReleaseResponse{}
	_body, _err := client.RevertPlaybookReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunPython3ScriptWithOptions(request *RunPython3ScriptRequest, runtime *util.RuntimeOptions) (_result *RunPython3ScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		body["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PythonScript)) {
		body["PythonScript"] = request.PythonScript
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunPython3Script"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunPython3ScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunPython3Script(request *RunPython3ScriptRequest) (_result *RunPython3ScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunPython3ScriptResponse{}
	_body, _err := client.RunPython3ScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerPlaybookWithOptions(request *TriggerPlaybookRequest, runtime *util.RuntimeOptions) (_result *TriggerPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputParam)) {
		body["InputParam"] = request.InputParam
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerPlaybook(request *TriggerPlaybookRequest) (_result *TriggerPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerPlaybookResponse{}
	_body, _err := client.TriggerPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerProcessTaskWithOptions(request *TriggerProcessTaskRequest, runtime *util.RuntimeOptions) (_result *TriggerProcessTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerProcessTask"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerProcessTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerProcessTask(request *TriggerProcessTaskRequest) (_result *TriggerProcessTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerProcessTaskResponse{}
	_body, _err := client.TriggerProcessTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerSophonPlaybookWithOptions(request *TriggerSophonPlaybookRequest, runtime *util.RuntimeOptions) (_result *TriggerSophonPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandName)) {
		query["CommandName"] = request.CommandName
	}

	if !tea.BoolValue(util.IsUnset(request.InputParams)) {
		query["InputParams"] = request.InputParams
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		query["SophonTaskId"] = request.SophonTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerType)) {
		query["TriggerType"] = request.TriggerType
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerSophonPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerSophonPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerSophonPlaybook(request *TriggerSophonPlaybookRequest) (_result *TriggerSophonPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerSophonPlaybookResponse{}
	_body, _err := client.TriggerSophonPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyPlaybookWithOptions(request *VerifyPlaybookRequest, runtime *util.RuntimeOptions) (_result *VerifyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlow)) {
		body["TaskFlow"] = request.TaskFlow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyPlaybook(request *VerifyPlaybookRequest) (_result *VerifyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPlaybookResponse{}
	_body, _err := client.VerifyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyPythonFileWithOptions(request *VerifyPythonFileRequest, runtime *util.RuntimeOptions) (_result *VerifyPythonFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPythonFile"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyPythonFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyPythonFile(request *VerifyPythonFileRequest) (_result *VerifyPythonFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPythonFileResponse{}
	_body, _err := client.VerifyPythonFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
