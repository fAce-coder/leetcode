# 子查询1
/*
    1.查找出有过订购记录的客户：将两个表以客户id作为连接条件，内连接（inner join）后的id
    2.因为本题要求没有订购记录的客户，因此从客户表名字中去除（not in）第一步找出的有订购记录的客户
*/
SELECT name as Customers
FROM Customers
WHERE id NOT IN (
    SELECT C.id
    FROM Customers as C
    INNER JOIN Orders as O
    ON C.id = O.customerId);

# 子查询2
/*
    1.找出有过订购记录的客户：找出订购单表中所有的客户id
    2.找出没有过订购记录的客户：在客户表中，将第一步查出来的id排除
*/
SELECT name as Customers
FROM Customers
WHERE id NOT IN (
    SELECT distinct CustomerId
    FROM Orders);