USE foo;

show tables;

CREATE TABLE foo.departments (
   	DeptID INT(6), 
	deptName VARCHAR(80), 
	ParentID INT(6), 
	ManagerID INT(6), 
	CreateDate DATE,
	UpdateDate DATE
);
   
   
   
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (10,'총무기획부',null,200,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (20,'마케팅',10,201,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (30,'구매/생산부',10,114,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (40,'인사부',10,203,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (50,'배송부',10,121,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (60,'IT',90,103,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (70,'홍보부',90,204,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (80,'영업부',10,145,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (90,'기획부',10,100,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (100,'자금부',90,108,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (110,'경리부',90,205,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (120,'재무팀',110,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (130,'세무팀',100,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (140,'신용관리팀',100,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (150,'주식관리팀',100,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (160,'수익관리팀',100,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (170,'생산팀',30,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (180,'건설팀',30,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (190,'계약팀',80,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (200,'운영팀',30,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (210,'IT 지원',30,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (220,'NOC',30,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (230,'IT 헬프데스크',60,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (240,'공공 판매사업팀',80,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (250,'판매팀',80,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (260,'채용팀',40,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));
Insert into foo.departments (DeptID,deptName,ParentID,ManagerID,CreateDate,UpdateDate) values (270,'급여팀',110,null,date('2014/01/08 13:49'),date('2014/01/08 13:49'));


ALTER TABLE foo.departments ADD PRIMARY KEY (DeptID);

 /*
--------------------------------------------------------
--  Constraints for Table DEPARTMENTS
--------------------------------------------------------

  ALTER TABLE "JAVA"."DEPARTMENTS" ADD CONSTRAINT "PK_DEPARTMENTS" PRIMARY KEY ("DeptID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "STUDY"  ENABLE;
  ALTER TABLE "JAVA"."DEPARTMENTS" MODIFY ("deptName" NOT NULL ENABLE);
  ALTER TABLE "JAVA"."DEPARTMENTS" MODIFY ("DeptID" NOT NULL ENABLE);
*/

select * from departments;

commit;