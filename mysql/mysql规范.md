## mysql 规范

- 不在数据库做运算：CPU计算务必移至业务层；
- 控制单表数据量：含int型不超过1000w，含char则不超过500w；合理分表；限制单库表数量在300以内；
- 控制列数量：字段数控制在20以内；
- 平衡范式与冗余：为提高效率牺牲范式设计，冗余数据；
- 拒绝3B：拒绝大sql，大事务，大批量；

1. 用好数值类型

```sql
tinyint(1Byte)

smallint(2Byte)

mediumint(3Byte)

int(4Byte)

bigint(8Byte)

bad case：int(1)/int(11)
```

2.优先使用enum或set

```shell script
例如：sex enum (‘F’, ‘M’)
```
3. 避免使用NULL字段

```go
NULL字段很难查询优化

NULL字段的索引需要额外空间

NULL字段的复合索引无效

bad case：

name char(32) default null

age int not null

good case：

age int not null default 0
```

4. 少用text/blob

```go
varchar的性能会比text高很多

实在避免不了blob，请拆表
```
--

### 索引类

1. 谨慎合理使用索引

```sql
改善查询、减慢更新

索引一定不是越多越好（能不加就不加，要加的一定得加）

覆盖记录条数过多不适合建索引，例如“性别”
```

2. 不在索引做列运算

```go
bad case：

select id where age +1 = 10;
```

3. innodb主键推荐使用自增列

```go
主键建立聚簇索引

主键不应该被修改

字符串不应该做主键

如果不指定主键，innodb会使用唯一且非空值索引代替
```

### sql类

1. sql语句尽可能简单
```sql
一条sql只能在一个cpu运算

大语句拆小语句，减少锁时间

一条大sql可以堵死整个库
```

2.简单的事务
```sql
事务时间尽可能短

bad case：

上传图片事务
```

3. 不用select *
```sql
消耗cpu，io，内存，带宽

这种程序不具有扩展性

```

4. OR改写为IN()

```sql
or的效率是n级别

in的效率是log(n)级别

in的个数建议控制在200以内

select id from t where phone=’159′ or phone=’136′;

改写成：

select id from t where phone in (’159′, ’136′);
```

5. 避免负向%, 慎用count(*)

6. limit高效分页
```sql
limit越大，效率越低

select id from t limit 10000, 10;

改写成：

select id from t where id > 10000 limit 10;
```