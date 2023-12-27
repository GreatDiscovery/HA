CREATE TABLE IF NOT EXISTS agent_seed
(
    agent_seed_id   int(10) unsigned    NOT NULL AUTO_INCREMENT,
    target_hostname varchar(128)        NOT NULL,
    source_hostname varchar(128)        NOT NULL,
    start_timestamp timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_timestamp   timestamp           NOT NULL DEFAULT '1971-01-01 00:00:00',
    is_complete     tinyint(3) unsigned NOT NULL DEFAULT '0',
    is_successful   tinyint(3) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (agent_seed_id)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii
;

CREATE INDEX target_hostname_idx_agent_seed ON agent_seed (target_hostname, is_complete)
;

CREATE INDEX source_hostname_idx_agent_seed ON agent_seed (source_hostname, is_complete)
;

CREATE INDEX start_timestamp_idx_agent_seed ON agent_seed (start_timestamp)
;

CREATE INDEX is_complete_idx_agent_seed ON agent_seed (is_complete, start_timestamp)
;

CREATE INDEX is_successful_idx_agent_seed ON agent_seed (is_successful, start_timestamp)
;
