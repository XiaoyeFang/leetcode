-- 176 第二高的薪水
-- 编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary）
select ifNull((select distinct Salary from Employee order by Salary DESC limit 1,1) ,null) as SecondHighestSalary

-- 177 第二高的薪水
-- 编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary）。
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT

BEGIN
  RETURN (
      # Write your MySQL query statement below.
    select if(count < N,null,min) from(
    select min(Salary) as min, count(1) as count from (
        select distinct Salary from Employee order by Salary DESC limit N
    ) as a)as b
  );
END

-- 178  编写一个 SQL 查询来实现分数排名。
--
-- 如果两个分数相同，则两个分数排名（Rank）相同。请注意，平分后的下一个名次应该是下一个连续的整数值。换句话说，名次之间不应该有“间隔”。

-- 此题涉及到三个窗口函数，
-- DENSE_RANK  合并排名 连续
-- RANK        合并排名 中断
-- ROW_NUMBER  不合并排名 连续
select Score,Dense_rank() over(Order by Score Desc) as `Rank` from  Scores
