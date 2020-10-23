create table user
(
id varchar(36),
email varchar(255),
password varchar(255) not null,
name varchar(100) not null ,
rol varchar(10) not null ,
active BOOL default 1,
created_at datetime not null,
updated_at datetime,
PRIMARY KEY (`id`))
ENGINE=InnoDB DEFAULT CHARSET=latin1;

