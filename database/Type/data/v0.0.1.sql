begin tran
insert into Type values(1,'balance','バランス');
insert into Type values(2,'attack','アタック');
insert into Type values(3,'defence','ディフェンス');
insert into Type values(4,'support','サポート');
insert into Type values(5,'heal','ヒール');
insert into Type values(6,'magia','マギア');
insert into Type values(7,'ultimet','アルティメット');
insert into Type values(8,'ringmagia','円環マギア');
insert into Type values(9,'ringmagia','円環サポート');

update SystemVersion set Revison = 1 where TableName = 'Type'

commit tran
