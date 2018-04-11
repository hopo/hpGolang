create tablespace study datafile '/oracdata/sudy.dbf'
size 100M autoextend on next 10M;

create user java identified by oracle
default tablespace study;

grant connect to java;
grant resource to java;