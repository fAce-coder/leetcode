# 表连接
/*
    两表内连接，连接条件是日期差值1天：
        ON datediff(W2.recordDate,W1.recordDate)=1
    筛选出这些数据中，第2天比第1天温度高的：
        WHERE W2.temperature>W1.temperature
*/
SELECT W2.id
FROM Weather AS W1
INNER JOIN Weather AS W2
ON datediff(W2.recordDate,W1.recordDate)=1
WHERE W2.temperature>W1.temperature;