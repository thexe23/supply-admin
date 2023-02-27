CREATE
    DATABASE IF NOT EXISTS `supply`;
USE
    `supply`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`        BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `username`  VARCHAR(255)    NOT NULL COMMENT '用户名',
    `password`  VARCHAR(255)    NOT NULL COMMENT '密码',
    `phone`     CHAR(11)        NOT NULL COMMENT '手机号',
    `role`      TINYINT         NOT NULL COMMENT '用户类型: 1居民 2社区负责人 3商超 4供应商',
    `org_id`    INT             NOT NULL COMMENT '社区ID',
    `market_id` BIGINT UNSIGNED NULL COMMENT '超市ID',
    `img_url`   varchar(255)    NOT NULL DEFAULT '' COMMENT '图片链接'
) ENGINE = innodb
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '用户表';

INSERT INTO `user` (`username`, `password`, `phone`, `role`, `org_id`, `market_id`, `img_url`)
VALUES ('resident1', '123456', '19121693709', 1, 1001, 10008, ''),
       ('resident2', '123456', '19121693709', 1, 1002, 10008, ''),
       ('resident3', '123456', '19121693709', 1, 1003, 10009, ''),
       ('resident4', '123456', '19121693709', 1, 1004, 10008, ''),
       ('deliver1', '123456', '19121693709', 2, 1001, 0, ''),
       ('deliver2', '123456', '19121693709', 2, 1002, 0, ''),
       ('deliver3', '123456', '19121693709', 2, 1003, 0, ''),
       ('deliver4', '123456', '19121693709', 2, 1004, 0, ''),
       ('supermarket1', '123456', '19121693709', 3, 2001, 0, ''),
       ('supermarket2', '123456', '19121693709', 3, 2002, 0, ''),
       ('admin', '123456', '19121693709', 4, 3001, 0, '');

DROP TABLE IF EXISTS `item`;
CREATE TABLE `item`
(
    `id`        BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `market_id` BIGINT UNSIGNED NOT NULL COMMENT '超市ID',
    `name`      varchar(50)     NOT NULL DEFAULT '' COMMENT '商品名',
    `price`     INT             NOT NULL DEFAULT 0 COMMENT '价格',
    `stock`     INT             NOT NULL DEFAULT 0 COMMENT '库存',
    `category`  INT             NOT NULL DEFAULT 0 COMMENT '类别',
    `img_url`   varchar(255)    NOT NULL DEFAULT '' COMMENT '图片链接',
    `on_sale`   TINYINT(1)      NOT NULL DEFAULT 1 COMMENT '上架：0 否 1是'
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '商品表';

INSERT INTO `item` (`market_id`, `name`, `price`, `stock`, `category`, `img_url`)
VALUES (10008, '口罩', 1, 200, 1, ''),
       (10009, '口罩', 1, 300, 1, ''),
       (10008, '饮用水', 10, 300, 2, ''),
       (10009, '饮用水', 10, 200, 2, '');

DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`
(
    `id`           BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `user_id`      BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `market_id`    BIGINT UNSIGNED NOT NULL COMMENT '超市ID',
    `org_id`       INT             NOT NULL COMMENT '社区ID',
    `items`        JSON COMMENT '商品快照',
    `amount`       INT UNSIGNED    NOT NULL COMMENT '数量',
    `created_at`   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
    `order_status` TINYINT         NOT NULL DEFAULT 10 COMMENT '订单状态：10待支付 20待发货 30配送中 40已完成'
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '订单表';


DROP TABLE IF EXISTS `transport`;
CREATE TABLE `transport`
(
    `id`         BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `source_id`  BIGINT UNSIGNED NOT NULL COMMENT '来源ID',
    `target_id`  BIGINT UNSIGNED NOT NULL COMMENT '目标ID',
    `item`       VARCHAR(255)    NOT NULL COMMENT '商品',
    `quantity`   INT UNSIGNED    NOT NULL COMMENT '数量',
    `created_at` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
    `status`     TINYINT         NOT NULL DEFAULT 10 COMMENT '调配单状态：10待发货 30配送中 30已完成'
) ENGINE = InnoDB
  AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4 COMMENT '调配表';