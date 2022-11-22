# 条件查询

/*
推荐人的编号不是2可以拆分成两个条件：
    1.推荐人的编号是数字且不是2（referee_id != 2）
    2.推荐人的编号不是数字而是null，此时也不是2（isnull(referee_id)=TRUE）
*/

SELECT name
FROM Customer
WHERE referee_id != 2 OR isnull(referee_id)=TRUE;