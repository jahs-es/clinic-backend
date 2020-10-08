create table patient
(
id varchar(36),
name varchar(150) not null,
avatar_path varchar(200) null,
address varchar(200),
email varchar(50),
phone varchar(50) not null,
active BOOL default 1,
created_at datetime not null,
created_by varchar(50) not null,
updated_at datetime,
updated_by varchar(50),
PRIMARY KEY (`id`))
ENGINE=InnoDB DEFAULT CHARSET=latin1;
