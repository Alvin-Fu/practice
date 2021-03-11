CREATE  TABLE `student` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `uid` bigint(20) NOT NULL COMMENT '学号',
    `class_id` bigint(10) NOT NULL COMMENT '班级',
    `score` bigint(20) NOT NULL COMMENT '分数',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`, `class_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

对每个班级的所有学生按照分数进行降序排序
```sql
SELECT * FROM `fuli`.`student` order by class_id asc, score desc
```

