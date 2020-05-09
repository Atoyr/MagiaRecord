begin tran
insert into Disks values(0,0,0,0,0)
insert into Disks values(1,1,0,1,3)
insert into Disks values(2,1,1,0,3)
insert into Disks values(3,1,0,2,2)
insert into Disks values(4,1,1,1,2)
insert into Disks values(5,1,2,0,2)
insert into Disks values(6,1,0,3,1)
insert into Disks values(7,1,1,2,1)
insert into Disks values(8,1,2,1,1)
insert into Disks values(9,1,3,0,1)
insert into Disks values(10,2,0,2,1)
insert into Disks values(11,2,1,1,1)
insert into Disks values(12,2,2,0,1)
insert into Disks values(13,2,0,1,2)
insert into Disks values(14,2,1,0,2)
insert into Disks values(15,3,0,1,1)
insert into Disks values(16,3,1,0,1)

update SystemVersion set Revison = 1 where TableName = 'Disks'

commit tran
