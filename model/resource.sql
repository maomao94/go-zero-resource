create table `t_oss`
(
    `id`          bigint auto_increment,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state`   tinyint(1) NOT NULL DEFAULT '0',
    `version`     bigint   NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
    `tenant_id`   varchar(12)       default '000000' null comment '租户ID',
    `category`    bigint null comment '分类',
    `oss_code`    varchar(32) null comment '资源编号',
    `endpoint`    varchar(255) null comment '资源地址',
    `access_key`  varchar(255) null comment 'accessKey',
    `secret_key`  varchar(255) null comment 'secretKey',
    `bucket_name` varchar(255) null comment '空间名',
    `app_id`      varchar(255) null comment '应用ID',
    `region`      varchar(255) null comment '地域简称',
    `remark`      varchar(255) null comment '备注',
    `status`      bigint null comment '状态',
    primary key (`id`)
);