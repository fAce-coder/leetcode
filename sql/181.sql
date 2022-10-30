# 子查询
select e1.name as Employee
from Employee as e1
where salary > (
    select salary
    from Employee as e2
    where e2.id = e1.managerId);

# 自连接 sql92语法
select e1.name AS Employee
from Employee AS e1,Employee AS e2
where e1.managerId = e2.id
and e1.salary > e2.salary;

# 自连接 sql99语法
select e1.name as Employee
from Employee as e1
inner join Employee as e2
on e1.managerId = e2.id
where e1.salary > e2.salary;
