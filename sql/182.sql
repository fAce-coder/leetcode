# 分组查询

/*
    将email分组，找到数量大于1的email
*/

select email as Email
from Person
group by email
having count(email)>1;