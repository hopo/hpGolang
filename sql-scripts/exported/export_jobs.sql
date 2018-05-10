--------------------------------------------------------
--  File created - Thursday-May-10-2018   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Table JOBS
--------------------------------------------------------

  CREATE TABLE "JAVA"."JOBS" 
   (	"JOB_ID" VARCHAR2(10 BYTE), 
	"JOB_TITLE" VARCHAR2(80 BYTE), 
	"MIN_SALARY" NUMBER(8,2) DEFAULT 0, 
	"MAX_SALARY" NUMBER(8,2) DEFAULT 0, 
	"CREATE_DATE" DATE DEFAULT SYSDATE, 
	"UPDATE_DATE" DATE DEFAULT SYSDATE
   ) SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "STUDY" ;

   COMMENT ON COLUMN "JAVA"."JOBS"."JOB_ID" IS 'JOB ID';
   COMMENT ON COLUMN "JAVA"."JOBS"."JOB_TITLE" IS 'JOB 명;';
   COMMENT ON COLUMN "JAVA"."JOBS"."MIN_SALARY" IS '최소급여';
   COMMENT ON COLUMN "JAVA"."JOBS"."MAX_SALARY" IS '최대급여';
   COMMENT ON COLUMN "JAVA"."JOBS"."CREATE_DATE" IS '생성일자';
   COMMENT ON COLUMN "JAVA"."JOBS"."UPDATE_DATE" IS '변경일자';
REM INSERTING into JAVA.JOBS
SET DEFINE OFF;
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('AD_PRES','President',20080,40000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('AD_VP','Administration Vice President',15000,30000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('AD_ASST','Administration Assistant',3000,6000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('FI_MGR','Finance Manager',8200,16000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('FI_ACCOUNT','Accountant',4200,9000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('AC_MGR','Accounting Manager',8200,16000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('AC_ACCOUNT','Public Accountant',4200,9000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('SA_MAN','Sales Manager',10000,20080,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('SA_REP','Sales Representative',6000,12008,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('PU_MAN','Purchasing Manager',8000,15000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('PU_CLERK','Purchasing Clerk',2500,5500,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('ST_MAN','Stock Manager',5500,8500,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('ST_CLERK','Stock Clerk',2008,5000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('SH_CLERK','Shipping Clerk',2500,5500,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('IT_PROG','Programmer',4000,10000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('MK_MAN','Marketing Manager',9000,15000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('MK_REP','Marketing Representative',4000,9000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('HR_REP','Human Resources Representative',4000,9000,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
Insert into JAVA.JOBS (JOB_ID,JOB_TITLE,MIN_SALARY,MAX_SALARY,CREATE_DATE,UPDATE_DATE) values ('PR_REP','Public Relations Representative',4500,10500,to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'),to_date('2014/01/08 13:52','YYYY/MM/DD HH24:MI'));
--------------------------------------------------------
--  DDL for Index PK_JOBS
--------------------------------------------------------

  CREATE UNIQUE INDEX "JAVA"."PK_JOBS" ON "JAVA"."JOBS" ("JOB_ID") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "STUDY" ;
--------------------------------------------------------
--  Constraints for Table JOBS
--------------------------------------------------------

  ALTER TABLE "JAVA"."JOBS" ADD CONSTRAINT "PK_JOBS" PRIMARY KEY ("JOB_ID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "STUDY"  ENABLE;
  ALTER TABLE "JAVA"."JOBS" MODIFY ("JOB_TITLE" NOT NULL ENABLE);
  ALTER TABLE "JAVA"."JOBS" MODIFY ("JOB_ID" NOT NULL ENABLE);
