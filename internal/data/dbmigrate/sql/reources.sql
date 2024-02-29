-- Version: 1.01
-- Description: Create table sources
CREATE TYPE doc_format AS ENUM ('DOC_FORMAT_JSON', 'DOC_FORMAT_YAML');

CREATE TYPE natural_languages AS ENUM ('EN', 'FR');

CREATE TYPE code_languages AS ENUM ('NOT_YET_KNOWN', 'FIND_OUT');


CREATE TABLE sources (
    id UUID,
    author_email TEXT,
    raw_url TEXT,
    scheme TEXT,
    host TEXT,
    path TEXT,
    query TEXT,
    active_domain boolean,
    "format" doc_format NOT NULL,
    created_at TIMESTAMP  WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP  WITH TIME ZONE NOT NULL,

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
    attempted_at TIMESTAMP  WITH TIME ZONE NOT NULL,
    downloaded_at TIMESTAMP  WITH TIME ZONE NOT NULL,
    status_code INTEGER,
    headers jsonb,
    body TEXT,

    PRIMARY KEY (id)
    CONSTRAINT fk_source
     FOREIGN KEY(source_id)
         REFERENCES sources(id)
);

CREATE TABLE documents (
    id UUID,
    source_id UUID,
    download_id UUID,
    "format" doc_format,
    indexed_at TIMESTAMP NOT NULL,
    min_chunk_size INTEGER,
    max_chunk_size INTEGER,
    published_at TIMESTAMP  WITH TIME ZONE NOT NULL,
    modified_at TIMESTAMP  WITH TIME ZONE NOT NULL,
    wp_version VARCHAR(10),

    PRIMARY KEY (id)

    CONSTRAINT fk_source
        FOREIGN KEY(source_id)
        REFERENCES sources(id)
    CONSTRAINT fk_download
        FOREIGN KEY(download_id)
        REFERENCES sources(downloads)


);

CREATE TABLE chunks (
    id UUID,
    document_id UUID,
    parent_chunk_id UUID,
    left_chunk_id UUID,
    right_chunk_id UUID,
    body TEXT,
    byte_size INTEGER,
    tokenizer VARCHAR(255),
    token_count INTEGER,
    natural_lang natural_languages,
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
    "key" VARCHAR(255),
    meta JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE embeddings (
    id UUID,
    embedding_1536 , -- Fill in type
    embedding_3072 , -- Fill in type
    model VARCHAR(255),
    embedded_at TIMESTAMP WITH TIME ZONE NOT NULL,
    "object_id" UUID,
    object_type VARCHAR(20),
    embedding_768 , -- Fill in type

    PRIMARY KEY (id)
);

CREATE TABLE requests (
    id UUID,
    "message" TEXT,
    meta JSONB,
    requested_at TIMESTAMP WITH TIME ZONE NOT NULL,
    result_chunks[] UUID,

    PRIMARY KEY (id)
);






