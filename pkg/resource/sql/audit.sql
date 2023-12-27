CREATE TABLE IF NOT EXISTS audit
(
    audit_id        bigint(20) unsigned              NOT NULL AUTO_INCREMENT,
    audit_timestamp timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    audit_type      varchar(128) CHARACTER SET ascii NOT NULL,
    hostname        varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',
    port            smallint(5) unsigned             NOT NULL,
    message         text CHARACTER SET utf8          NOT NULL,
    PRIMARY KEY (audit_id)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = latin1;
CREATE INDEX audit_timestamp_idx_audit ON audit (audit_timestamp);
CREATE INDEX host_port_idx_audit ON audit (hostname, port, audit_timestamp);