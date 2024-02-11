// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"
)

type Querier interface {
	CountBuckets(ctx context.Context) (int64, error)
	CreateBucket(ctx context.Context, arg *CreateBucketParams) (string, error)
	CreateEvent(ctx context.Context, arg *CreateEventParams) (string, error)
	CreateObject(ctx context.Context, arg *CreateObjectParams) (string, error)
	DeleteBucket(ctx context.Context, id string) error
	DeleteObject(ctx context.Context, id string) error
	DisableBucket(ctx context.Context, id string) error
	EnableBucket(ctx context.Context, id string) error
	GetBucketById(ctx context.Context, id string) (*StorageBucket, error)
	GetBucketByName(ctx context.Context, name string) (*StorageBucket, error)
	GetBucketObjectCountById(ctx context.Context, id string) (int64, error)
	GetBucketSizeById(ctx context.Context, id string) (*GetBucketSizeByIdRow, error)
	GetObjectByBucketIdAndName(ctx context.Context, arg *GetObjectByBucketIdAndNameParams) (*GetObjectByBucketIdAndNameRow, error)
	GetObjectById(ctx context.Context, id string) (*GetObjectByIdRow, error)
	ListAllBuckets(ctx context.Context) ([]*StorageBucket, error)
	ListBucketsPaginated(ctx context.Context, arg *ListBucketsPaginatedParams) ([]*ListBucketsPaginatedRow, error)
	ListObjectsByBucketIdPaged(ctx context.Context, arg *ListObjectsByBucketIdPagedParams) ([]*ListObjectsByBucketIdPagedRow, error)
	LockBucket(ctx context.Context, arg *LockBucketParams) error
	MakeBucketPrivate(ctx context.Context, id string) error
	MakeBucketPublic(ctx context.Context, id string) error
	MakeObjectPrivate(ctx context.Context, id string) error
	MakeObjectPublic(ctx context.Context, id string) error
	MergeObjectMetadata(ctx context.Context, arg *MergeObjectMetadataParams) error
	SearchBucketsPaginated(ctx context.Context, arg *SearchBucketsPaginatedParams) ([]*SearchBucketsPaginatedRow, error)
	SearchObjectsByPath(ctx context.Context, arg *SearchObjectsByPathParams) ([]*SearchObjectsByPathRow, error)
	UnlockBucket(ctx context.Context, id string) error
	UpdateBucket(ctx context.Context, arg *UpdateBucketParams) error
	UpdateObject(ctx context.Context, arg *UpdateObjectParams) error
	UpdateObjectLastAccessedAt(ctx context.Context, id string) error
	UpdateObjectUploadStatus(ctx context.Context, arg *UpdateObjectUploadStatusParams) error
}

var _ Querier = (*Queries)(nil)