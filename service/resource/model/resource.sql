-- auto-generated definition
create table resource_oss
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime                     null,
    updated_at  datetime                     null,
    deleted_at  datetime                     null,
    tenant_id   varchar(12) default '000000' null comment '租户ID',
    category    bigint                       null comment '分类',
    oss_code    varchar(32)                  null comment '资源编号',
    endpoint    varchar(255)                 null comment '资源地址',
    access_key  varchar(255)                 null comment 'accessKey',
    secret_key  varchar(255)                 null comment 'secretKey',
    bucket_name varchar(255)                 null comment '空间名',
    app_id      varchar(255)                 null comment '应用ID',
    region      varchar(255)                 null comment '地域简称',
    remark      varchar(255)                 null comment '备注',
    status      bigint                       null comment '状态',
    is_deleted  bigint                       null comment '是否已删除'
);