--
-- 3일차
--

number
number(10)
number(10, 2)


-- *** ex2_3
create table ex2_3(
	col1 number(5),
	col2 number(5, 2)
	col3 date
) ;

insert into ex2_3 values( 123456, 123.547, sysdate ) ; -- error
insert into ex2_3 values( 12345, 123.547, sysdate ) ;
insert into ex2_3 values( 12345, 1004.547, sysdate ) ; -- error
insert into ex2_3 values( 12345, 450.547, '2018/5/8' ) ;
insert into ex2_3 values( 12345, 450.547, '2018/5/8 23:14' ) ; --erorr
select * from ex2_3;



-- *** ex2_4
-- p.63
create table ex2_4(
	col1 varchar2(10) ,
	col2 varchar2(10) not null ,
	col3 varchar2(10) unique ,
	col4 varchar2(10) ,
	constraint uk_col4 unique(con4)
) ;

insert into ex2_4(col1, col2, col3, col4) values ('hong', 'hong', 'hong', 'hong') ;
insert into ex2_4(col1, col2, col3, col4) values ('hong', 'hong', 'hong', 'hong') ; --unique error
insert into ex2_4(col1, col2, col3, col4) values ('hong', 'hong', 'dong', 'gil') ;

select * from ex2_4;

-- null insert at unique colomn
-- *오라클은 빈공백('')을 null로 인식. cf)타 DBMS는 빈공백으로 인식.
insert into ex2_4(col1, col2, col3, col4) values ('hong', 'hong', null, '') ;

-- null은 unique제약으로 제외 됩니다.
insert into ex2_4(col1, col2, col3, col4) values ('hong', 'hong', null, '') ;

-- col2(NOT NULL)
insert into ex2_4(col1, col2, col3, col4) values ('hong', '', null, '') ; -- NN error

-- col2 (NN)
insert into ex2_4(col1, col3, col4) values ('abc', 'def', 'hgj') ; -- NN error, default is NULL

-- col2
insert into ex2_4(col2) values ('수요일') ; -- NN error, default is NULL

commit;

-- CONSTRAINTS
select * from user_constraints ;
select * from user_constraints where table_name = 'EX2_4' ;


-- *PRIMARY KEY
-- p.64

-- ex2_5
create table ex2_5( 
	col1 number ,
	col2 varchar2(20),
	col3 varchar2(300) ,
	CONSTRAINTS pk_ex2_5 primary key(col1, col2) -- 복합 PK설정
) ;

insert into ex2_5 values(1, 'abc', '5월 8일은 임시공휴우우') ;
insert into ex2_5 values(1, 'abc', '고고고씨') ; -- error
insert into ex2_5 values(1, 'def', 'JAVA 힘들어요허허허') ;

select * from ex2_5 ;

insert into ex2_5 values(2, null, '고!') ; -- error


-- /* SELECT 구문
-- p.92
select       : 5. 필요한 정보(컬럼, 계산된값, 상수)
	from	 : 1. 조회를 위한 근간이 되는 데이터 (테이블, 뷰)
	where 	 : 2. 검색조건을 명시
	group by : 3. 데이터를 그루핑
	having   : 4. 그루핑 조건을 명시
	order by : 6. 결과를 정렬
;
-- */

select * from employees;
select employee_id, emp_name, salary from employees;
select employee_id, emp_name, salary, job_id
	from employees
	where salary > 6000
	and job_id = 'IT_PROG'
	order by salary desc -- 기본값 ASC
;


-- _EXAM)
-- 사원중에 급여가 5000이하 사원을 조회해 주세요
-- 사번, 사원명, 급여, 업무ID 를 출력
-- 정렬은 사번 DESC
select employee_id, emp_name, salary, job_id
	from employees
	where salary <= 5000
	order by employee_id desc
;

-- using as
select employee_id as 사번,
		emp_name as 사원명,
		salary + 100 as 급여, -- usually 본 column 명이 변경 될때 사용
		job_id as 업무
	from employees emp  -- as의 생략
	where salary <= 5000
	order by emp.employee_id desc
;


commit;