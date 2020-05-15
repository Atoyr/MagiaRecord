begin tran
insert into Types values(1,'balance','バランス');
insert into Types values(2,'attack','アタック');
insert into Types values(3,'defence','ディフェンス');
insert into Types values(4,'support','サポート');
insert into Types values(5,'heal','ヒール');
insert into Types values(6,'magia','マギア');
insert into Types values(7,'ultimet','アルティメット');
insert into Types values(8,'ringmagia','円環マギア');
insert into Types values(9,'ringmagia','円環サポート');

update SystemVersion set Mejor = 0, Minor = 0,  Revison = 1 where TableName = 'Types';

commit tran
