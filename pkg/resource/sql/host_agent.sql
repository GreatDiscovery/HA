CREATE TABLE IF NOT EXISTS host_agent
(
    hostname              varchar(128)         NOT NULL,
    port                  smallint(5) unsigned NOT NULL,
    token                 varchar(128)         NOT NULL,
    last_submitted        timestamp            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_checked          timestamp            NULL     DEFAULT NULL,
    last_seen             timestamp            NULL     DEFAULT NULL,
    mysql_port            smallint(5) unsigned          DEFAULT NULL,
    count_mysql_snapshots smallint(5) unsigned NOT NULL,
    PRIMARY KEY (hostname)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii;

CREATE INDEX token_idx_host_agent ON host_agent (token);
CREATE INDEX last_submitted_idx_host_agent ON host_agent (last_submitted);
CREATE INDEX last_checked_idx_host_agent ON host_agent (last_checked);
CREATE INDEX last_seen_idx_host_agent ON host_agent (last_seen);
