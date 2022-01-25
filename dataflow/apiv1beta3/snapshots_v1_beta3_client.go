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

package dataflow

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dataflowpb "google.golang.org/genproto/googleapis/dataflow/v1beta3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newSnapshotsV1Beta3ClientHook clientHook

// SnapshotsV1Beta3CallOptions contains the retry settings for each method of SnapshotsV1Beta3Client.
type SnapshotsV1Beta3CallOptions struct {
	GetSnapshot    []gax.CallOption
	DeleteSnapshot []gax.CallOption
	ListSnapshots  []gax.CallOption
}

func defaultSnapshotsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSnapshotsV1Beta3CallOptions() *SnapshotsV1Beta3CallOptions {
	return &SnapshotsV1Beta3CallOptions{
		GetSnapshot:    []gax.CallOption{},
		DeleteSnapshot: []gax.CallOption{},
		ListSnapshots:  []gax.CallOption{},
	}
}

// internalSnapshotsV1Beta3Client is an interface that defines the methods availaible from Dataflow API.
type internalSnapshotsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetSnapshot(context.Context, *dataflowpb.GetSnapshotRequest, ...gax.CallOption) (*dataflowpb.Snapshot, error)
	DeleteSnapshot(context.Context, *dataflowpb.DeleteSnapshotRequest, ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error)
	ListSnapshots(context.Context, *dataflowpb.ListSnapshotsRequest, ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error)
}

// SnapshotsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides methods to manage snapshots of Google Cloud Dataflow jobs.
type SnapshotsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalSnapshotsV1Beta3Client

	// The call options for this service.
	CallOptions *SnapshotsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SnapshotsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SnapshotsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SnapshotsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetSnapshot gets information about a snapshot.
func (c *SnapshotsV1Beta3Client) GetSnapshot(ctx context.Context, req *dataflowpb.GetSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	return c.internalClient.GetSnapshot(ctx, req, opts...)
}

// DeleteSnapshot deletes a snapshot.
func (c *SnapshotsV1Beta3Client) DeleteSnapshot(ctx context.Context, req *dataflowpb.DeleteSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error) {
	return c.internalClient.DeleteSnapshot(ctx, req, opts...)
}

// ListSnapshots lists snapshots.
func (c *SnapshotsV1Beta3Client) ListSnapshots(ctx context.Context, req *dataflowpb.ListSnapshotsRequest, opts ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error) {
	return c.internalClient.ListSnapshots(ctx, req, opts...)
}

// snapshotsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type snapshotsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SnapshotsV1Beta3Client
	CallOptions **SnapshotsV1Beta3CallOptions

	// The gRPC API client.
	snapshotsV1Beta3Client dataflowpb.SnapshotsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSnapshotsV1Beta3Client creates a new snapshots v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides methods to manage snapshots of Google Cloud Dataflow jobs.
func NewSnapshotsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*SnapshotsV1Beta3Client, error) {
	clientOpts := defaultSnapshotsV1Beta3GRPCClientOptions()
	if newSnapshotsV1Beta3ClientHook != nil {
		hookOpts, err := newSnapshotsV1Beta3ClientHook(ctx, clientHookParams{})
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
	client := SnapshotsV1Beta3Client{CallOptions: defaultSnapshotsV1Beta3CallOptions()}

	c := &snapshotsV1Beta3GRPCClient{
		connPool:               connPool,
		disableDeadlines:       disableDeadlines,
		snapshotsV1Beta3Client: dataflowpb.NewSnapshotsV1Beta3Client(connPool),
		CallOptions:            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *snapshotsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *snapshotsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *snapshotsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *snapshotsV1Beta3GRPCClient) GetSnapshot(ctx context.Context, req *dataflowpb.GetSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetSnapshot[0:len((*c.CallOptions).GetSnapshot):len((*c.CallOptions).GetSnapshot)], opts...)
	var resp *dataflowpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.GetSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snapshotsV1Beta3GRPCClient) DeleteSnapshot(ctx context.Context, req *dataflowpb.DeleteSnapshotRequest, opts ...gax.CallOption) (*dataflowpb.DeleteSnapshotResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).DeleteSnapshot[0:len((*c.CallOptions).DeleteSnapshot):len((*c.CallOptions).DeleteSnapshot)], opts...)
	var resp *dataflowpb.DeleteSnapshotResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.DeleteSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snapshotsV1Beta3GRPCClient) ListSnapshots(ctx context.Context, req *dataflowpb.ListSnapshotsRequest, opts ...gax.CallOption) (*dataflowpb.ListSnapshotsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListSnapshots[0:len((*c.CallOptions).ListSnapshots):len((*c.CallOptions).ListSnapshots)], opts...)
	var resp *dataflowpb.ListSnapshotsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snapshotsV1Beta3Client.ListSnapshots(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
