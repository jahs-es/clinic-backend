create table treatment
(
id varchar(36),
name varchar(150) not null,
active BOOL default 1,
created_at datetime not null,
created_by varchar(50) not null,
updated_at datetime,
updated_by varchar(50),
PRIMARY KEY (`id`))
ENGINE=InnoDB DEFAULT CHARSET=latin1;
