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

package gaming_test

import (
	"context"

	gaming "cloud.google.com/go/gaming/apiv1"
	"google.golang.org/api/iterator"
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
)

func ExampleNewGameServerClustersClient() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleGameServerClustersClient_ListGameServerClusters() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.ListGameServerClustersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#ListGameServerClustersRequest.
	}
	it := c.ListGameServerClusters(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleGameServerClustersClient_GetGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.GetGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#GetGameServerClusterRequest.
	}
	resp, err := c.GetGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerClustersClient_CreateGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.CreateGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#CreateGameServerClusterRequest.
	}
	op, err := c.CreateGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerClustersClient_PreviewCreateGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.PreviewCreateGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#PreviewCreateGameServerClusterRequest.
	}
	resp, err := c.PreviewCreateGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerClustersClient_DeleteGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.DeleteGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#DeleteGameServerClusterRequest.
	}
	op, err := c.DeleteGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleGameServerClustersClient_PreviewDeleteGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.PreviewDeleteGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#PreviewDeleteGameServerClusterRequest.
	}
	resp, err := c.PreviewDeleteGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerClustersClient_UpdateGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.UpdateGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#UpdateGameServerClusterRequest.
	}
	op, err := c.UpdateGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGameServerClustersClient_PreviewUpdateGameServerCluster() {
	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gamingpb.PreviewUpdateGameServerClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/gaming/v1#PreviewUpdateGameServerClusterRequest.
	}
	resp, err := c.PreviewUpdateGameServerCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
