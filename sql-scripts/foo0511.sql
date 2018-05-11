show tables;


select *
from (
    select row_number() over(order by HireDate desc) rn, EmpName, HireDate
    from employees
    where Salary >= 3000
) a
where rn between 11 and 20;