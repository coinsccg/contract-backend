create database if not exists `auction` charset = 'utf8';
use `auction`;

create table if not exists holders
(
	id bigint auto_increment
		primary key,
		stage int null,
	holder_address char(42) not null,
    last_reward decimal(38) null,
	last_quantity decimal(38) null,
    current_quantity decimal(38) null,
    refresh_time int null
)
charset=utf8mb4;

create table if not exists recomms
(
    id bigint auto_increment
    primary key,
    holder_address char(42) not null,
    recomm_address char(42) not null
)
charset=utf8mb4;

create table if not exists refresh_times
(
    stage int not null,
    refresh_time int not null
)
charset=utf8mb4;