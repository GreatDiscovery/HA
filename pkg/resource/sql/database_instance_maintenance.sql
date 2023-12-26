CREATE TABLE IF NOT EXISTS database_instance_maintenance
(
    database_instance_maintenance_id int(10) unsigned                NOT NULL AUTO_INCREMENT,
    hostname                         varchar(128)                    NOT NULL,
    port                             smallint(5) unsigned            NOT NULL,
    maintenance_active               tinyint(4)                           DEFAULT NULL,
    begin_timestamp                  timestamp                       NULL DEFAULT NULL,
    end_timestamp                    timestamp                       NULL DEFAULT NULL,
    owner                            varchar(128) CHARACTER SET utf8 NOT NULL,
    reason                           text CHARACTER SET utf8         NOT NULL,
    PRIMARY KEY (database_instance_maintenance_id)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = ascii;
CREATE UNIQUE INDEX maintenance_uidx_database_instance_maintenance ON database_instance_maintenance (maintenance_active, hostname, port)
