PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE employee_details ("name","id","salary","destignation");
INSERT INTO employee_details VALUES ("Ram","1","40000","chennai");
INSERT INTO employee_details VALUES ("Arun","2","30000","madurai");
INSERT INTO employee_details VALUES ("kumar","3","50000","madurai");
COMMIT;
