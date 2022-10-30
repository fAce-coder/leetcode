# 外连接

/*
    1.两表链接查询，通过personId进行连通：
        Person.personId = Address.personId
    2.当person表中有但address表中没有时，显示address为null：
        左外连接，显示左边表中有但右边表中没有的数据
        FROM Person LEFT JOIN Address
*/

SELECT firstName, lastName, city, state
FROM Person
LEFT JOIN Address
ON Person.personId = Address.personId;