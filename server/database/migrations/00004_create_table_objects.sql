-- +goose Up
-- +goose StatementBegin

create table if not exists storage.objects
(
    id               text                                           not null check ( storage.text_non_empty_trimmed_text(id) ),
    version          int         default 0                          not null check ( version >= 0 ),
    bucket_id        text                                           not null check ( storage.text_non_empty_trimmed_text(bucket_id) ),
    name             text                                           not null check ( storage.text_non_empty_trimmed_text(name) ),
    path_tokens      text[]                                         not null generated always as (string_to_array(name, '/')) stored,
    mime_type        text        default 'application/octet-stream' not null check ( storage.text_non_empty_trimmed_text(mime_type) ),
    size             bigint      default 0                          not null check ( size >= 0 ),
    metadata         jsonb                                          null,
    upload_status    text        default 'pending'                  not null check (
        upload_status in ('pending', 'processing', 'completed', 'failed')
        ),
    last_accessed_at timestamptz                                    null,
    locked_at        timestamptz                                    null,
    created_at       timestamptz default now()                      not null,
    updated_at       timestamptz                                    null,
    constraint objects_id_primary_key primary key (id),
    constraint objects_id_version_unique unique (id, version),
    constraint objects_bucket_id_foreign_key foreign key (bucket_id) references storage.buckets (id) on delete no action,
    constraint objects_name_unique unique (bucket_id, name)
);

create index if not exists objects_bucket_id_name_index on storage.objects using btree (bucket_id, name);

create or replace trigger objects_on_create
    before insert
    on storage.objects
    for each row
execute function storage.on_create();

create or replace trigger objects_on_update
    before update
    on storage.objects
    for each row
execute function storage.on_update();



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop trigger if exists objects_increment_version on storage.objects;

drop trigger if exists objects_set_updated_at on storage.objects;

drop index if exists storage.objects_bucket_id_name_idx;

drop table if exists storage.objects;

-- +goose StatementEnd