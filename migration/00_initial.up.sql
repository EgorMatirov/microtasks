create table if not exists notes
(
    id          uuid not null
        constraint notes_pk
            primary key,
    note_header text,
    note_text   text,
    user_id     bigint
);

create unique index if not exists notes_id_uindex
    on notes (id);

create table if not exists note_tags
(
    note_id uuid not null,
    tag     text not null,
    constraint note_tags_pk
        unique (note_id, tag)
);

alter table note_tags
    add constraint note_tags_notes_id_fk
        foreign key (note_id) references notes
            on delete cascade;