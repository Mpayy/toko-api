create table if not exists products(
    id int not null auto_increment primary key,
    name varchar(255) not null,
    price int not null,
    stock int not null
)engine=InnoDB;