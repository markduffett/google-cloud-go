// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package admin

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	adminpb "google.golang.org/genproto/googleapis/datastore/admin/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newDatastoreAdminClientHook clientHook

// DatastoreAdminCallOptions contains the retry settings for each method of DatastoreAdminClient.
type DatastoreAdminCallOptions struct {
	ExportEntities []gax.CallOption
	ImportEntities []gax.CallOption
	CreateIndex    []gax.CallOption
	DeleteIndex    []gax.CallOption
	GetIndex       []gax.CallOption
	ListIndexes    []gax.CallOption
}

func defaultDatastoreAdminGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("datastore.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("datastore.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://datastore.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultDatastoreAdminCallOptions() *DatastoreAdminCallOptions {
	return &DatastoreAdminCallOptions{
		ExportEntities: []gax.CallOption{},
		ImportEntities: []gax.CallOption{},
		CreateIndex:    []gax.CallOption{},
		DeleteIndex:    []gax.CallOption{},
		GetIndex: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListIndexes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalDatastoreAdminClient is an interface that defines the methods availaible from Cloud Datastore API.
type internalDatastoreAdminClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ExportEntities(context.Context, *adminpb.ExportEntitiesRequest, ...gax.CallOption) (*ExportEntitiesOperation, error)
	ExportEntitiesOperation(name string) *ExportEntitiesOperation
	ImportEntities(context.Context, *adminpb.ImportEntitiesRequest, ...gax.CallOption) (*ImportEntitiesOperation, error)
	ImportEntitiesOperation(name string) *ImportEntitiesOperation
	CreateIndex(context.Context, *adminpb.CreateIndexRequest, ...gax.CallOption) (*CreateIndexOperation, error)
	CreateIndexOperation(name string) *CreateIndexOperation
	DeleteIndex(context.Context, *adminpb.DeleteIndexRequest, ...gax.CallOption) (*DeleteIndexOperation, error)
	DeleteIndexOperation(name string) *DeleteIndexOperation
	GetIndex(context.Context, *adminpb.GetIndexRequest, ...gax.CallOption) (*adminpb.Index, error)
	ListIndexes(context.Context, *adminpb.ListIndexesRequest, ...gax.CallOption) *IndexIterator
}

// DatastoreAdminClient is a client for interacting with Cloud Datastore API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Google Cloud Datastore Admin API
//
// The Datastore Admin API provides several admin services for Cloud Datastore.
//
// ConceptsProject, namespace, kind, and entity as defined in the Google Cloud Datastore
// API.
//
// Operation: An Operation represents work being performed in the background.
//
// EntityFilter: Allows specifying a subset of entities in a project. This is
// specified as a combination of kinds and namespaces (either or both of which
// may be all).
//
// ServicesExport/ImportThe Export/Import service provides the ability to copy all or a subset of
// entities to/from Google Cloud Storage.
//
// Exported data may be imported into Cloud Datastore for any Google Cloud
// Platform project. It is not restricted to the export source project. It is
// possible to export from one project and then import into another.
//
// Exported data can also be loaded into Google BigQuery for analysis.
//
// Exports and imports are performed asynchronously. An Operation resource is
// created for each export/import. The state (including any errors encountered)
// of the export/import may be queried via the Operation resource.
//
// IndexThe index service manages Cloud Datastore composite indexes.
//
// Index creation and deletion are performed asynchronously.
// An Operation resource is created for each such asynchronous operation.
// The state of the operation (including any errors encountered)
// may be queried via the Operation resource.
//
// OperationThe Operations collection provides a record of actions performed for the
// specified project (including any operations in progress). Operations are not
// created directly but through calls on other collections or resources.
//
// An operation that is not yet done may be cancelled. The request to cancel is
// asynchronous and the operation may continue to run for some time after the
// request to cancel is made.
//
// An operation that is done may be deleted so that it is no longer listed as
// part of the Operation collection.
//
// ListOperations returns all pending operations, but not completed operations.
//
// Operations are created by service DatastoreAdmin,
// but are accessed via service google.longrunning.Operations.
type DatastoreAdminClient struct {
	// The internal transport-dependent client.
	internalClient internalDatastoreAdminClient

	// The call options for this service.
	CallOptions *DatastoreAdminCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DatastoreAdminClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DatastoreAdminClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *DatastoreAdminClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ExportEntities exports a copy of all or a subset of entities from Google Cloud Datastore
// to another storage system, such as Google Cloud Storage. Recent updates to
// entities may not be reflected in the export. The export occurs in the
// background and its progress can be monitored and managed via the
// Operation resource that is created. The output of an export may only be
// used once the associated operation is done. If an export operation is
// cancelled before completion it may leave partial data behind in Google
// Cloud Storage.
func (c *DatastoreAdminClient) ExportEntities(ctx context.Context, req *adminpb.ExportEntitiesRequest, opts ...gax.CallOption) (*ExportEntitiesOperation, error) {
	return c.internalClient.ExportEntities(ctx, req, opts...)
}

// ExportEntitiesOperation returns a new ExportEntitiesOperation from a given name.
// The name must be that of a previously created ExportEntitiesOperation, possibly from a different process.
func (c *DatastoreAdminClient) ExportEntitiesOperation(name string) *ExportEntitiesOperation {
	return c.internalClient.ExportEntitiesOperation(name)
}

// ImportEntities imports entities into Google Cloud Datastore. Existing entities with the
// same key are overwritten. The import occurs in the background and its
// progress can be monitored and managed via the Operation resource that is
// created. If an ImportEntities operation is cancelled, it is possible
// that a subset of the data has already been imported to Cloud Datastore.
func (c *DatastoreAdminClient) ImportEntities(ctx context.Context, req *adminpb.ImportEntitiesRequest, opts ...gax.CallOption) (*ImportEntitiesOperation, error) {
	return c.internalClient.ImportEntities(ctx, req, opts...)
}

// ImportEntitiesOperation returns a new ImportEntitiesOperation from a given name.
// The name must be that of a previously created ImportEntitiesOperation, possibly from a different process.
func (c *DatastoreAdminClient) ImportEntitiesOperation(name string) *ImportEntitiesOperation {
	return c.internalClient.ImportEntitiesOperation(name)
}

// CreateIndex creates the specified index.
// A newly created index’s initial state is CREATING. On completion of the
// returned google.longrunning.Operation, the state will be READY.
// If the index already exists, the call will return an ALREADY_EXISTS
// status.
//
// During index creation, the process could result in an error, in which
// case the index will move to the ERROR state. The process can be recovered
// by fixing the data that caused the error, removing the index with
// delete, then
// re-creating the index with [create]
// [google.datastore.admin.v1.DatastoreAdmin.CreateIndex].
//
// Indexes with a single property cannot be created.
func (c *DatastoreAdminClient) CreateIndex(ctx context.Context, req *adminpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	return c.internalClient.CreateIndex(ctx, req, opts...)
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *DatastoreAdminClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return c.internalClient.CreateIndexOperation(name)
}

// DeleteIndex deletes an existing index.
// An index can only be deleted if it is in a READY or ERROR state. On
// successful execution of the request, the index will be in a DELETING
// state. And on completion of the
// returned google.longrunning.Operation, the index will be removed.
//
// During index deletion, the process could result in an error, in which
// case the index will move to the ERROR state. The process can be recovered
// by fixing the data that caused the error, followed by calling
// delete again.
func (c *DatastoreAdminClient) DeleteIndex(ctx context.Context, req *adminpb.DeleteIndexRequest, opts ...gax.CallOption) (*DeleteIndexOperation, error) {
	return c.internalClient.DeleteIndex(ctx, req, opts...)
}

// DeleteIndexOperation returns a new DeleteIndexOperation from a given name.
// The name must be that of a previously created DeleteIndexOperation, possibly from a different process.
func (c *DatastoreAdminClient) DeleteIndexOperation(name string) *DeleteIndexOperation {
	return c.internalClient.DeleteIndexOperation(name)
}

// GetIndex gets an index.
func (c *DatastoreAdminClient) GetIndex(ctx context.Context, req *adminpb.GetIndexRequest, opts ...gax.CallOption) (*adminpb.Index, error) {
	return c.internalClient.GetIndex(ctx, req, opts...)
}

// ListIndexes lists the indexes that match the specified filters.  Datastore uses an
// eventually consistent query to fetch the list of indexes and may
// occasionally return stale results.
func (c *DatastoreAdminClient) ListIndexes(ctx context.Context, req *adminpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	return c.internalClient.ListIndexes(ctx, req, opts...)
}

// datastoreAdminGRPCClient is a client for interacting with Cloud Datastore API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type datastoreAdminGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing DatastoreAdminClient
	CallOptions **DatastoreAdminCallOptions

	// The gRPC API client.
	datastoreAdminClient adminpb.DatastoreAdminClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDatastoreAdminClient creates a new datastore admin client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Google Cloud Datastore Admin API
//
// The Datastore Admin API provides several admin services for Cloud Datastore.
//
// ConceptsProject, namespace, kind, and entity as defined in the Google Cloud Datastore
// API.
//
// Operation: An Operation represents work being performed in the background.
//
// EntityFilter: Allows specifying a subset of entities in a project. This is
// specified as a combination of kinds and namespaces (either or both of which
// may be all).
//
// ServicesExport/ImportThe Export/Import service provides the ability to copy all or a subset of
// entities to/from Google Cloud Storage.
//
// Exported data may be imported into Cloud Datastore for any Google Cloud
// Platform project. It is not restricted to the export source project. It is
// possible to export from one project and then import into another.
//
// Exported data can also be loaded into Google BigQuery for analysis.
//
// Exports and imports are performed asynchronously. An Operation resource is
// created for each export/import. The state (including any errors encountered)
// of the export/import may be queried via the Operation resource.
//
// IndexThe index service manages Cloud Datastore composite indexes.
//
// Index creation and deletion are performed asynchronously.
// An Operation resource is created for each such asynchronous operation.
// The state of the operation (including any errors encountered)
// may be queried via the Operation resource.
//
// OperationThe Operations collection provides a record of actions performed for the
// specified project (including any operations in progress). Operations are not
// created directly but through calls on other collections or resources.
//
// An operation that is not yet done may be cancelled. The request to cancel is
// asynchronous and the operation may continue to run for some time after the
// request to cancel is made.
//
// An operation that is done may be deleted so that it is no longer listed as
// part of the Operation collection.
//
// ListOperations returns all pending operations, but not completed operations.
//
// Operations are created by service DatastoreAdmin,
// but are accessed via service google.longrunning.Operations.
func NewDatastoreAdminClient(ctx context.Context, opts ...option.ClientOption) (*DatastoreAdminClient, error) {
	clientOpts := defaultDatastoreAdminGRPCClientOptions()
	if newDatastoreAdminClientHook != nil {
		hookOpts, err := newDatastoreAdminClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DatastoreAdminClient{CallOptions: defaultDatastoreAdminCallOptions()}

	c := &datastoreAdminGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		datastoreAdminClient: adminpb.NewDatastoreAdminClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *datastoreAdminGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *datastoreAdminGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *datastoreAdminGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *datastoreAdminGRPCClient) ExportEntities(ctx context.Context, req *adminpb.ExportEntitiesRequest, opts ...gax.CallOption) (*ExportEntitiesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExportEntities[0:len((*c.CallOptions).ExportEntities):len((*c.CallOptions).ExportEntities)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.ExportEntities(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportEntitiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *datastoreAdminGRPCClient) ImportEntities(ctx context.Context, req *adminpb.ImportEntitiesRequest, opts ...gax.CallOption) (*ImportEntitiesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportEntities[0:len((*c.CallOptions).ImportEntities):len((*c.CallOptions).ImportEntities)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.ImportEntities(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportEntitiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *datastoreAdminGRPCClient) CreateIndex(ctx context.Context, req *adminpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIndex[0:len((*c.CallOptions).CreateIndex):len((*c.CallOptions).CreateIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.CreateIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *datastoreAdminGRPCClient) DeleteIndex(ctx context.Context, req *adminpb.DeleteIndexRequest, opts ...gax.CallOption) (*DeleteIndexOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "index_id", url.QueryEscape(req.GetIndexId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIndex[0:len((*c.CallOptions).DeleteIndex):len((*c.CallOptions).DeleteIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.DeleteIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *datastoreAdminGRPCClient) GetIndex(ctx context.Context, req *adminpb.GetIndexRequest, opts ...gax.CallOption) (*adminpb.Index, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "index_id", url.QueryEscape(req.GetIndexId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIndex[0:len((*c.CallOptions).GetIndex):len((*c.CallOptions).GetIndex)], opts...)
	var resp *adminpb.Index
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.GetIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreAdminGRPCClient) ListIndexes(ctx context.Context, req *adminpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIndexes[0:len((*c.CallOptions).ListIndexes):len((*c.CallOptions).ListIndexes)], opts...)
	it := &IndexIterator{}
	req = proto.Clone(req).(*adminpb.ListIndexesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*adminpb.Index, string, error) {
		resp := &adminpb.ListIndexesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.datastoreAdminClient.ListIndexes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIndexes(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateIndexOperation manages a long-running operation from CreateIndex.
type CreateIndexOperation struct {
	lro *longrunning.Operation
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *datastoreAdminGRPCClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateIndexOperation) Metadata() (*adminpb.IndexOperationMetadata, error) {
	var meta adminpb.IndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateIndexOperation) Name() string {
	return op.lro.Name()
}

// DeleteIndexOperation manages a long-running operation from DeleteIndex.
type DeleteIndexOperation struct {
	lro *longrunning.Operation
}

// DeleteIndexOperation returns a new DeleteIndexOperation from a given name.
// The name must be that of a previously created DeleteIndexOperation, possibly from a different process.
func (c *datastoreAdminGRPCClient) DeleteIndexOperation(name string) *DeleteIndexOperation {
	return &DeleteIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteIndexOperation) Metadata() (*adminpb.IndexOperationMetadata, error) {
	var meta adminpb.IndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteIndexOperation) Name() string {
	return op.lro.Name()
}

// ExportEntitiesOperation manages a long-running operation from ExportEntities.
type ExportEntitiesOperation struct {
	lro *longrunning.Operation
}

// ExportEntitiesOperation returns a new ExportEntitiesOperation from a given name.
// The name must be that of a previously created ExportEntitiesOperation, possibly from a different process.
func (c *datastoreAdminGRPCClient) ExportEntitiesOperation(name string) *ExportEntitiesOperation {
	return &ExportEntitiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportEntitiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportEntitiesResponse, error) {
	var resp adminpb.ExportEntitiesResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportEntitiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportEntitiesResponse, error) {
	var resp adminpb.ExportEntitiesResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportEntitiesOperation) Metadata() (*adminpb.ExportEntitiesMetadata, error) {
	var meta adminpb.ExportEntitiesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportEntitiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportEntitiesOperation) Name() string {
	return op.lro.Name()
}

// ImportEntitiesOperation manages a long-running operation from ImportEntities.
type ImportEntitiesOperation struct {
	lro *longrunning.Operation
}

// ImportEntitiesOperation returns a new ImportEntitiesOperation from a given name.
// The name must be that of a previously created ImportEntitiesOperation, possibly from a different process.
func (c *datastoreAdminGRPCClient) ImportEntitiesOperation(name string) *ImportEntitiesOperation {
	return &ImportEntitiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportEntitiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ImportEntitiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ImportEntitiesOperation) Metadata() (*adminpb.ImportEntitiesMetadata, error) {
	var meta adminpb.ImportEntitiesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportEntitiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportEntitiesOperation) Name() string {
	return op.lro.Name()
}

// IndexIterator manages a stream of *adminpb.Index.
type IndexIterator struct {
	items    []*adminpb.Index
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*adminpb.Index, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IndexIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IndexIterator) Next() (*adminpb.Index, error) {
	var item *adminpb.Index
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IndexIterator) bufLen() int {
	return len(it.items)
}

func (it *IndexIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
