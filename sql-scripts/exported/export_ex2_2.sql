--------------------------------------------------------
--  File created - Thursday-May-10-2018   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Table EX2_2
--------------------------------------------------------

  CREATE TABLE "JAVA"."EX2_2" 
   (	"COL1" CHAR(5 BYTE), 
	"COL2" VARCHAR2(5 BYTE), 
	"COL3" VARCHAR2(5 CHAR)
   ) SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "STUDY" ;
REM INSERTING into JAVA.EX2_2
SET DEFINE OFF;
Insert into JAVA.EX2_2 (COL1,COL2,COL3) values ('ab   ','ab','ab');
Insert into JAVA.EX2_2 (COL1,COL2,COL3) values ('c    ','동','홍길동');
