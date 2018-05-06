SHOW databases ;
create database foo ;
use foo ;
show tables ;
create table programlang(
	name varchar(20),
	side varchar(50),
	descr varchar(50)
) ;

insert into programlang values ('c', 'system', 'Acient Lang of programlang.') ;
insert into programlang values ('python', 'anyway', 'python2 and python3...') ;
insert into programlang values ('Go', 'no system side', 'GOPHER!!!') ;
insert into programlang values ('JAVA', 'anyway', 'OTL ...') ;


select * from foo.programlang ;