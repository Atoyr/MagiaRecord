begin tran
insert into Attributes values(1,'Frame','火')
insert into Attributes values(2,'Aqua','水')
insert into Attributes values(3,'Forest','樹')
insert into Attributes values(4,'Light','光')
insert into Attributes values(5,'Dark','闇')
insert into Attributes values(6,'Void','無')

update SystemVersion set Mejor = 0, Miner = 0,  Revison = 1 where TableName = 'Attributes';

commit tran

