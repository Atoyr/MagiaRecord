begin tran
insert into MagicalGirls values (1,N'環いろは',N'',N'tamaki iroha', 1,5, 4/*光*/,5/*ヒール*/,2,0,2,1,N'一緒にいるから大丈夫だよ',N'ストラーダ・フトゥーロ',N'沈黙のドッペル');
insert into MagicalGirls values (2,N'七海やちよ',N'',N'nanami yachiyo', 2,5, 2/*水*/,1/*バランス*/,1,0,3,1,N'道は照らしておくから',N'アブソリュート・レイン',N'モギリのドッペル');
insert into MagicalGirls values (3,N'由比鶴乃',N'',N'yui tsuruno', 2,5, 1/*火*/,6/*マギア*/,2,2,0,1,N'出たとこ勝負で行こうよ！',N'炎扇斬舞',N'団欒のドッペル');
insert into MagicalGirls values (4,N'深月フェリシア',N'',N'mitsuki ferisia', 2,5, 4/*闇*/,2/*アタック*/,1,1,1,2,N'ガツンと任せとけって！',N'ウルトラグレートビッグハンマー',N'攪拌のドッペル');
insert into MagicalGirls values (5,N'二葉さな',N'',N'futaba sana', 2,5, 3/*木*/,3/*ディフェンス*/,1,1,1,2,N'私が必要なら、いくらでも',N'フォルターゲフェングニス',N'無色透明のドッペル');

insert into MagicalGirls values (6,N'鹿目まどか',N'',N'kaname madoka', 4,5, 4, 5, 2,1,1,1, N'私が受け止めてあげるから',N'プルウィア☆マギカ',N'慈悲のドッペル');
insert into MagicalGirls values (7,N'暁美ほむら',N'',N'akemi homura', 4,5, 5, 5, 1,1,2,1, N'あなたの獲物よ',N'ミサイルによる集中砲火',N'業因のドッペル');
insert into MagicalGirls values (8,N'巴マミ',N'',N'tomoe mami', 4,5, 3, 6, 2,0,1,2, N'ティロ・デュエット！',N'ティロ・フィナーレ',N'ご招待のドッペル');
insert into MagicalGirls values (9,N'美樹さやか',N'',N'miki sayaka', 4,5, 2, 3, 2,0,1,2, N'ここであたしの出番でしょ！',N'プレスティッシモ・アジタート',N'恋慕のドッペル');
insert into MagicalGirls values (10,N'佐倉杏子',N'',N'sakura kyouko', 4,5, 1, 2, 1,1,2,1, N'ったく、見てらんねーな',N'盟神抉槍',N'自棄のドッペル');


update SystemVersion set Mejor = 0, Miner = 0,  Revison = 1 where TableName = 'MagicalGirls';

commit tran
