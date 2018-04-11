--2일차

select * from departments;

create table ex2_2(
    col1 char(5),
    col2 varchar2(5),
    col3 varchar2(5 char)
) ;
-- char : 데이터사이즈가 정해진 경우 (우편번호 코드 데이블)
-- 나머지는 varchar2
-- DB에서는 " 거의 사용하지 않음

insert into ex2_2 values ('ab', 'ab', 'ab');
select * from ex2_2;

insert into ex2_2 values ('c', '동', '홍길동');
select col1, length(col1), lengthb(col1),
       col2, length(col2), lengthb(col2),
       col3, length(col3), lengthb(col3)
from ex2_2;

commit;

