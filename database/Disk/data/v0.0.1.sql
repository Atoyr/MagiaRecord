begin tran
insert into Disk values(0,0,0,0,0)
insert into Disk values(1,1,0,1,3)
insert into Disk values(2,1,1,0,3)
insert into Disk values(3,1,0,2,2)
insert into Disk values(4,1,1,1,2)
insert into Disk values(5,1,2,0,2)
insert into Disk values(6,1,0,3,1)
insert into Disk values(7,1,1,2,1)
insert into Disk values(8,1,2,1,1)
insert into Disk values(9,1,3,0,1)
insert into Disk values(10,2,0,2,1)
insert into Disk values(11,2,1,1,1)
insert into Disk values(12,2,2,0,1)
insert into Disk values(13,2,0,1,2)
insert into Disk values(14,2,1,0,2)
insert into Disk values(15,3,0,1,1)
insert into Disk values(16,3,1,0,1)

update SystemVersion set Revison = 1 where TableName = 'Disk'

commit tran
