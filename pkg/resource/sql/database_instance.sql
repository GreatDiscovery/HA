CREATE TABLE IF NOT EXISTS database_instance
(
    hostname              varchar(128) CHARACTER SET ascii NOT NULL,
    port                  smallint(5) unsigned             NOT NULL,
    last_checked          timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_seen             timestamp                        NULL     DEFAULT NULL,
    server_id             int(10) unsigned                 NOT NULL,
    version               varchar(128) CHARACTER SET ascii NOT NULL,
    binlog_format         varchar(16) CHARACTER SET ascii  NOT NULL,
    log_bin               tinyint(3) unsigned              NOT NULL,
    log_slave_updates     tinyint(3) unsigned              NOT NULL,
    binary_log_file       varchar(128) CHARACTER SET ascii NOT NULL,
    binary_log_pos        bigint(20) unsigned              NOT NULL,
    master_host           varchar(128) CHARACTER SET ascii NOT NULL,
    master_port           smallint(5) unsigned             NOT NULL,
    slave_sql_running     tinyint(3) unsigned              NOT NULL,
    slave_io_running      tinyint(3) unsigned              NOT NULL,
    master_log_file       varchar(128) CHARACTER SET ascii NOT NULL,
    read_master_log_pos   bigint(20) unsigned              NOT NULL,
    relay_master_log_file varchar(128) CHARACTER SET ascii NOT NULL,
    exec_master_log_pos   bigint(20) unsigned              NOT NULL,
    seconds_behind_master bigint(20) unsigned                       DEFAULT NULL,
    slave_lag_seconds     bigint(20) unsigned                       DEFAULT NULL,
    num_slave_hosts       int(10) unsigned                 NOT NULL,
    slave_hosts           text CHARACTER SET ascii         NOT NULL,
    cluster_name          varchar(128) CHARACTER SET ascii NOT NULL,
    PRIMARY KEY (hostname, port)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii;

# DROP INDEX cluster_name_idx ON database_instance;
CREATE INDEX cluster_name_idx_database_instance ON database_instance (cluster_name);
# DROP INDEX last_checked_idx ON database_instance;
CREATE INDEX last_checked_idx_database_instance ON database_instance (last_checked);
# DROP INDEX last_seen_idx ON database_instance;
CREATE INDEX last_seen_idx_database_instance ON database_instance (last_seen);