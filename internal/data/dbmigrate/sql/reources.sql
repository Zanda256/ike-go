-- Version: 1.01
-- Description: Create resource tables
CREATE TYPE doc_format AS ENUM ('DOC_FORMAT_JSON', 'DOC_FORMAT_YAML');

CREATE TYPE natural_languages AS ENUM ('EN', 'FR');

CREATE TYPE code_languages AS ENUM ('NOT_YET_KNOWN', 'FIND_OUT');


CREATE TABLE sources (
    id UUID,
    author_email TEXT,
    raw_url TEXT NOT NULL,
    scheme TEXT NOT NULL,
    host TEXT NOT NULL,
    path TEXT NOT NULL,
    query TEXT,
    active_domain boolean NOT NULL,
    "format" doc_format NOT NULL,
    created_at TIMESTAMP  WITH TIME ZONE,
    updated_at TIMESTAMP  WITH TIME ZONE,

    PRIMARY KEY (id)
);

CREATE TABLE tags (
    id UUID,
    "name" VARCHAR(255),
    created_at TIMESTAMP  WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE downloads (
    id UUID,
    source_id UUID,
    attempted_at TIMESTAMP  WITH TIME ZONE,
    downloaded_at TIMESTAMP  WITH TIME ZONE,
    status_code INTEGER,
    headers jsonb NOT NULL,
    body TEXT,

    PRIMARY KEY (id),
    CONSTRAINT fk_source
     FOREIGN KEY(source_id)
         REFERENCES sources(id)
);

CREATE TABLE documents (
    id UUID,
    source_id UUID NOT NULL,
    download_id UUID,
    "format" doc_format NOT NULL,
    indexed_at TIMESTAMP,
    min_chunk_size INTEGER NOT NULL,
    max_chunk_size INTEGER NOT NULL,
    published_at TIMESTAMP  WITH TIME ZONE,
    modified_at TIMESTAMP  WITH TIME ZONE,
    wp_version VARCHAR(10),

    PRIMARY KEY (id),

    CONSTRAINT fk_source
        FOREIGN KEY(source_id)
        REFERENCES sources(id),
    CONSTRAINT fk_download
        FOREIGN KEY(download_id)
        REFERENCES downloads(id)


);

CREATE TABLE chunks (
    id UUID,
    document_id UUID,
    parent_chunk_id UUID,
    left_chunk_id UUID,
    right_chunk_id UUID,
    body TEXT NOT NULL,
    byte_size INTEGER NOT NULL,
    tokenizer VARCHAR(255) NOT NULL,
    token_count INTEGER NOT NULL,
    natural_lang natural_languages NOT NULL,
    code_lang code_languages,

    PRIMARY KEY (id)
);

CREATE TABLE document_tags (
    id UUID,
    document_id UUID,
    tag_id UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE document_meta (
    id UUID,
    document_id UUID,
    "key" VARCHAR(255) NOT NULL,
    meta JSONB,
    created_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY (id)
);

CREATE TABLE embeddings (
    id UUID,
    embedding_1536 , -- Fill in type
    embedding_3072 , -- Fill in type
    model VARCHAR(255),
    embedded_at TIMESTAMP WITH TIME ZONE,
    "object_id" UUID NOT NULL,
    object_type VARCHAR(20) NOT NULL,
    embedding_768 , -- Fill in type

    PRIMARY KEY (id)
);

CREATE TABLE requests (
    id UUID,
    "message" TEXT NOT NULL,
    meta JSONB,
    requested_at TIMESTAMP WITH TIME ZONE,
    result_chunks[] UUID,

    PRIMARY KEY (id)
);






