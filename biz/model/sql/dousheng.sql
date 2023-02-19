create schema gorm collate utf8_general_ci;

SET @@sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

create table `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `name` varchar(32) not null comment '姓名',
                        `password` varchar(32) not null comment '密码',
                        `follow_count` bigint(20) not null default 0 comment '关注数',
                        `follower_count` bigint(20) not null default 0 comment '被关注数',
                        `created_at` timestamp not null default current_timestamp comment '创建时间',
                        `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                        `deleted_at` timestamp null comment '删除时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table `video` (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `user_id` bigint(20) not null comment '用户id',
                         `title` varchar(128) not null comment '文章标题',
                         `play_url` varchar(128) not null comment '视频播放地址',
                         `cover_url` varchar(128) not null comment '视频封面地址',
                         `favorite_count` bigint(20) not null comment '视频的点赞总数',
                         `comment_count` bigint(20) not null comment '视频的评论总数',
                         `created_at` timestamp not null default current_timestamp comment '创建时间',
                         `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                         `deleted_at` timestamp null comment '删除时间',
                         PRIMARY KEY (`id`),
                         constraint fk_video_user
                             FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


create table `favorite` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `user_id` bigint(20) not null comment '用户id',
                            `video_id` bigint(20) not null comment '视频id',
                            `created_at` timestamp not null default current_timestamp comment '创建时间',
                            `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                            `deleted_at` timestamp null comment '删除时间',
                            PRIMARY KEY (`id`),
                            constraint fk_favorite_user
                                foreign key (user_id) references user(id),
                            constraint fk_favorite_video
                                foreign key (video_id) references video(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `comment` (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                           `user_id` bigint(20) not null comment '用户id',
                           `video_id` bigint(20) not null comment '视频id',
                           `contents` varchar(255) not null comment '用户评论',
                           `created_at` timestamp not null default current_timestamp comment '创建时间',
                           `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                           `deleted_at` timestamp null comment '删除时间',
                           PRIMARY KEY (`id`),
                           constraint fk_comment_user
                               foreign key (user_id) references user(id),
                           constraint fk_comment_video
                               foreign key (video_id) references video(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `relation` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `user_id` bigint(20) not null comment '用户id',
                            `to_user_id` bigint(20) not null comment '对方用户id',
                            `created_at` timestamp not null default current_timestamp comment '创建时间',
                            `updated_at` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
                            `deleted_at` timestamp null comment '删除时间',
                            PRIMARY KEY (`id`),
                            constraint fk_relation_user
                                FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;