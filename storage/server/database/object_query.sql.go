// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: object_query.sql

package database

import (
	"context"
	"time"
)

const createObject = `-- name: CreateObject :one
insert into storage.objects
    (bucket_id, name, content_type, size, public, metadata, upload_status)
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7) returning id
`

type CreateObjectParams struct {
	BucketID     string
	Name         string
	ContentType  *string
	Size         int64
	Public       bool
	Metadata     []byte
	UploadStatus string
}

func (q *Queries) CreateObject(ctx context.Context, arg *CreateObjectParams) (string, error) {
	row := q.db.QueryRow(ctx, createObject,
		arg.BucketID,
		arg.Name,
		arg.ContentType,
		arg.Size,
		arg.Public,
		arg.Metadata,
		arg.UploadStatus,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteObject = `-- name: DeleteObject :exec
delete
from storage.objects
where id = $1
`

func (q *Queries) DeleteObject(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteObject, id)
	return err
}

const getObjectByBucketIdAndName = `-- name: GetObjectByBucketIdAndName :one
select id,
       bucket_id,
       name,
       path_tokens,
       content_type,
       size,
       public,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where bucket_id = $1
  and name = $2
limit 1
`

type GetObjectByBucketIdAndNameParams struct {
	BucketID string
	Name     string
}

type GetObjectByBucketIdAndNameRow struct {
	ID             string
	BucketID       string
	Name           string
	PathTokens     []string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) GetObjectByBucketIdAndName(ctx context.Context, arg *GetObjectByBucketIdAndNameParams) (*GetObjectByBucketIdAndNameRow, error) {
	row := q.db.QueryRow(ctx, getObjectByBucketIdAndName, arg.BucketID, arg.Name)
	var i GetObjectByBucketIdAndNameRow
	err := row.Scan(
		&i.ID,
		&i.BucketID,
		&i.Name,
		&i.PathTokens,
		&i.ContentType,
		&i.Size,
		&i.Public,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getObjectById = `-- name: GetObjectById :one
select id,
       bucket_id,
       name,
       path_tokens,
       content_type,
       size,
       public,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where id = $1
limit 1
`

type GetObjectByIdRow struct {
	ID             string
	BucketID       string
	Name           string
	PathTokens     []string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) GetObjectById(ctx context.Context, id string) (*GetObjectByIdRow, error) {
	row := q.db.QueryRow(ctx, getObjectById, id)
	var i GetObjectByIdRow
	err := row.Scan(
		&i.ID,
		&i.BucketID,
		&i.Name,
		&i.PathTokens,
		&i.ContentType,
		&i.Size,
		&i.Public,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getObjectByIdWithBucketName = `-- name: GetObjectByIdWithBucketName :one
select o.id,
       o.bucket_id,
       b.name as bucket_name,
       o.name,
       o.path_tokens,
       o.content_type,
       o.size,
       o.public,
       o.metadata,
       o.upload_status,
       o.last_accessed_at,
       o.created_at,
       o.updated_at
from storage.objects as o
inner join storage.buckets as b on o.bucket_id = b.id
where o.id = $1
limit 1
`

type GetObjectByIdWithBucketNameRow struct {
	ID             string
	BucketID       string
	BucketName     string
	Name           string
	PathTokens     []string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) GetObjectByIdWithBucketName(ctx context.Context, id string) (*GetObjectByIdWithBucketNameRow, error) {
	row := q.db.QueryRow(ctx, getObjectByIdWithBucketName, id)
	var i GetObjectByIdWithBucketNameRow
	err := row.Scan(
		&i.ID,
		&i.BucketID,
		&i.BucketName,
		&i.Name,
		&i.PathTokens,
		&i.ContentType,
		&i.Size,
		&i.Public,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getObjectByName = `-- name: GetObjectByName :one
select id,
       bucket_id,
       name,
       path_tokens,
       content_type,
       size,
       public,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where name = $1
limit 1
`

type GetObjectByNameRow struct {
	ID             string
	BucketID       string
	Name           string
	PathTokens     []string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) GetObjectByName(ctx context.Context, name string) (*GetObjectByNameRow, error) {
	row := q.db.QueryRow(ctx, getObjectByName, name)
	var i GetObjectByNameRow
	err := row.Scan(
		&i.ID,
		&i.BucketID,
		&i.Name,
		&i.PathTokens,
		&i.ContentType,
		&i.Size,
		&i.Public,
		&i.Metadata,
		&i.UploadStatus,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listObjectsByBucketIdPaged = `-- name: ListObjectsByBucketIdPaged :many
select id,
       bucket_id,
       name,
       path_tokens,
       content_type,
       size,
       public,
       metadata,
       upload_status,
       last_accessed_at,
       created_at,
       updated_at
from storage.objects
where bucket_id = $1
limit $3 offset $2
`

type ListObjectsByBucketIdPagedParams struct {
	BucketID string
	Offset   int32
	Limit    int32
}

type ListObjectsByBucketIdPagedRow struct {
	ID             string
	BucketID       string
	Name           string
	PathTokens     []string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt *time.Time
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (q *Queries) ListObjectsByBucketIdPaged(ctx context.Context, arg *ListObjectsByBucketIdPagedParams) ([]*ListObjectsByBucketIdPagedRow, error) {
	rows, err := q.db.Query(ctx, listObjectsByBucketIdPaged, arg.BucketID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListObjectsByBucketIdPagedRow
	for rows.Next() {
		var i ListObjectsByBucketIdPagedRow
		if err := rows.Scan(
			&i.ID,
			&i.BucketID,
			&i.Name,
			&i.PathTokens,
			&i.ContentType,
			&i.Size,
			&i.Public,
			&i.Metadata,
			&i.UploadStatus,
			&i.LastAccessedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const makeObjectPrivate = `-- name: MakeObjectPrivate :exec
update storage.objects
set public = false
where id = $1
`

func (q *Queries) MakeObjectPrivate(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeObjectPrivate, id)
	return err
}

const makeObjectPublic = `-- name: MakeObjectPublic :exec
update storage.objects
set public = true
where id = $1
`

func (q *Queries) MakeObjectPublic(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeObjectPublic, id)
	return err
}

const mergeObjectMetadata = `-- name: MergeObjectMetadata :exec
update storage.objects
set metadata = metadata || $1
where id = $2
`

type MergeObjectMetadataParams struct {
	Metadata []byte
	ID       string
}

func (q *Queries) MergeObjectMetadata(ctx context.Context, arg *MergeObjectMetadataParams) error {
	_, err := q.db.Exec(ctx, mergeObjectMetadata, arg.Metadata, arg.ID)
	return err
}

const searchObjectsByPath = `-- name: SearchObjectsByPath :many
select id::text,
       version::int,
       name::text,
       bucket_id::text,
       bucket::text,
       content_type::text,
       size::bigint,
       public::boolean,
       metadata::jsonb,
       upload_status::text,
       last_accessed_at::timestamptz,
       created_at::timestamptz,
       updated_at::timestamptz
from storage.objects_search($1::text, $2::text, $3::int,
                            $4::int, $5::int)
`

type SearchObjectsByPathParams struct {
	BucketName string
	ObjectPath string
	Level      *int32
	Limit      *int32
	Offset     *int32
}

type SearchObjectsByPathRow struct {
	ID             string
	Version        int32
	Name           string
	BucketID       string
	Bucket         string
	ContentType    string
	Size           int64
	Public         bool
	Metadata       []byte
	UploadStatus   string
	LastAccessedAt time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (q *Queries) SearchObjectsByPath(ctx context.Context, arg *SearchObjectsByPathParams) ([]*SearchObjectsByPathRow, error) {
	rows, err := q.db.Query(ctx, searchObjectsByPath,
		arg.BucketName,
		arg.ObjectPath,
		arg.Level,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*SearchObjectsByPathRow
	for rows.Next() {
		var i SearchObjectsByPathRow
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.BucketID,
			&i.Bucket,
			&i.ContentType,
			&i.Size,
			&i.Public,
			&i.Metadata,
			&i.UploadStatus,
			&i.LastAccessedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateObject = `-- name: UpdateObject :exec
update storage.objects
set size         = coalesce($1, size),
    content_type = coalesce($2, content_type),
    metadata     = coalesce($3, metadata)
where id = $4
`

type UpdateObjectParams struct {
	Size        int64
	ContentType string
	Metadata    []byte
	ID          string
}

func (q *Queries) UpdateObject(ctx context.Context, arg *UpdateObjectParams) error {
	_, err := q.db.Exec(ctx, updateObject,
		arg.Size,
		arg.ContentType,
		arg.Metadata,
		arg.ID,
	)
	return err
}

const updateObjectLastAccessedAt = `-- name: UpdateObjectLastAccessedAt :exec
update storage.objects
set last_accessed_at = now()
where id = $1
`

func (q *Queries) UpdateObjectLastAccessedAt(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, updateObjectLastAccessedAt, id)
	return err
}

const updateObjectUploadStatus = `-- name: UpdateObjectUploadStatus :exec
update storage.objects
set upload_status = $1
where id = $2
`

type UpdateObjectUploadStatusParams struct {
	UploadStatus string
	ID           string
}

func (q *Queries) UpdateObjectUploadStatus(ctx context.Context, arg *UpdateObjectUploadStatusParams) error {
	_, err := q.db.Exec(ctx, updateObjectUploadStatus, arg.UploadStatus, arg.ID)
	return err
}
