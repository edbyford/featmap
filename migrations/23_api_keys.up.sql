CREATE TABLE api_keys
(
    workspace_id  uuid                     NOT NULL,
    id            uuid                     NOT NULL,
    member_id     uuid                     NOT NULL,
    key_hash      varchar(255)             NOT NULL,
    name          varchar(100)             NOT NULL DEFAULT 'claude-code',
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    last_used_at  TIMESTAMP WITH TIME ZONE,
    CONSTRAINT "PK_api_keys_1" PRIMARY KEY (workspace_id, id),
    CONSTRAINT "FK_api_keys_1" FOREIGN KEY (workspace_id) REFERENCES workspaces (id) ON DELETE CASCADE,
    CONSTRAINT "FK_api_keys_2" FOREIGN KEY (workspace_id, member_id) REFERENCES members (workspace_id, id) ON DELETE CASCADE,
    CONSTRAINT "UN_api_keys_1" UNIQUE (key_hash)
)
    WITH (
        OIDS= FALSE
    );

CREATE INDEX "IX_api_keys_1" ON api_keys (key_hash);
