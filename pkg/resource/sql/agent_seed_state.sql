CREATE TABLE IF NOT EXISTS agent_seed_state
(
    agent_seed_state_id int(10) unsigned NOT NULL AUTO_INCREMENT,
    agent_seed_id       int(10) unsigned NOT NULL,
    state_timestamp     timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    state_action        varchar(127)     NOT NULL,
    error_message       varchar(255)     NOT NULL,
    PRIMARY KEY (agent_seed_state_id)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii
;

CREATE INDEX agent_seed_idx_agent_seed_state ON agent_seed_state (agent_seed_id, state_timestamp)
;