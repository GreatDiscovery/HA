CREATE TABLE IF NOT EXISTS host_attributes
(
    hostname         varchar(128) NOT NULL,
    attribute_name   varchar(128) NOT NULL,
    attribute_value  varchar(128) NOT NULL,
    submit_timestamp timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expire_timestamp timestamp    NULL     DEFAULT NULL,
    PRIMARY KEY (hostname, attribute_name)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii
;

CREATE INDEX attribute_name_idx_host_attributes ON host_attributes (attribute_name)
;

CREATE INDEX attribute_value_idx_host_attributes ON host_attributes (attribute_value)
;

CREATE INDEX submit_timestamp_idx_host_attributes ON host_attributes (submit_timestamp)
;

CREATE INDEX expire_timestamp_idx_host_attributes ON host_attributes (expire_timestamp)
;
