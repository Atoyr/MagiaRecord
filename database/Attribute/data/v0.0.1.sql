begin tran
insert into Attributes values(1,'Frame','火')
insert into Attributes values(1,'Aqua','水')
insert into Attributes values(1,'Forest','樹')
insert into Attributes values(1,'light','光')
insert into Attributes values(1,'Dark','闇')

update SystemVersion set Revison = 1 where TableName = 'Attributes'

commit tran

