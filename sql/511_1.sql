# 分组+聚合函数

/*
    首先，要求获取每位用户的数据，因此根据用户id进行分组
    再要求求出每位用户的第一次登录平台的日期，即日期字段最早的时间，使用min聚合函数
*/

SELECT player_id, MIN(event_date) AS first_login
FROM Activity
GROUP BY player_id;