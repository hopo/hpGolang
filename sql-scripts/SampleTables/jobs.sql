USE foo;

/*
CREATE TABLE jobs ( 
	job_id VARCHAR(10),
     job_title VARCHAR(80),
     min_salary int(8),
     max_salary int(8),
     create_date DATE DEFAULT 0, 
     update_date DATE DEFAULT 0
); 
*/
    
select * from jobs;
 
insert into jobs (create_date) values ( date('2014-01-08') ) where job_id='AD_PRES' ;
 
select ROW_NUMBER from dual;
 

     
     
Insert into jobs (job_id,job_title,min_salary,max_salary,create_date,update_date) values     ('AD_PRES2','Slave',20080,40000,date('2014-01-08'),null);
--  to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2    014/01/08 13:52','YYYY/MM/DD HH24:MI'));   ) ;
    
select * from jobs;

commit ;


SELECT * FROM foo.programlang;

/*
 * TABLE jobs
 */

CREATE TABLE jobs (
  	JobID VARCHAR(10), 
	JobTitle VARCHAR(80), 
	MinSalary INT(8) DEFAULT 0, 
	MaxSalary INT(8) DEFAULT 0, 
	CreateDate DATE,
	UpdateDate DATE
) ;


Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('AD_PRES','President',20080,40000,date('2014/01/08'),date('2014/01/08'));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('AD_VP','Administration Vice President',15000,30000,date('2014/01/08'),date('2014/01/08'));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('AD_ASST','Administration Assistant',3000,6000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('FI_MGR','Finance Manager',8200,16000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('FI_ACCOUNT','Accountant',4200,9000,date('2014/01/08  '),date('2014/01/08  '));

Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('AC_MGR','Accounting Manager',8200,16000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('AC_ACCOUNT','Public Accountant',4200,9000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('SA_MAN','Sales Manager',10000,20080,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('SA_REP','Sales Representative',6000,12008,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('PU_MAN','Purchasing Manager',8000,15000,date('2014/01/08  '),date('2014/01/08  '));

Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('PU_CLERK','Purchasing Clerk',2500,5500,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('ST_MAN','Stock Manager',5500,8500,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('ST_CLERK','Stock Clerk',2008,5000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('SH_CLERK','Shipping Clerk',2500,5500,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('IT_PROG','Programmer',4000,10000,date('2014/01/08  '),date('2014/01/08  '));

Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('MK_MAN','Marketing Manager',9000,15000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('MK_REP','Marketing Representative',4000,9000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('HR_REP','Human Resources Representative',4000,9000,date('2014/01/08  '),date('2014/01/08  '));
Insert into jobs (JobID,JobTitle,MinSalary,MaxSalary,CreateDate,UpdateDate) values ('PR_REP','Public Relations Representative',4500,10500,date('2014/01/08  '),date('2014/01/08  '));

ALTER TABLE jobs ADD PRIMARY KEY (JobID);


SELECT * FROM jobs;

COMMIT ;
