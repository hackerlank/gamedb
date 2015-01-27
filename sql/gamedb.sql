DROP TABLE IF EXISTS game;
CREATE TABLE game (
	id int(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	url_name VARCHAR(255) NOT NULL,
	pic VARCHAR(255) NOT NULL,
	simple_desc VARCHAR(512) NOT NULL,
	detail_desc TEXT NOT NULL,
	tags VARCHAR(255) NOT NULL,
	FULLTEXT (simple_desc)
) engine=myisam;

insert into game values(1, '决战喵星', 'miao', '1.png', '《决战喵星》是一款走萌系路线的酷跑类游戏，但这款游戏还加入了类似于塔防的元素，被成为酷跑式塔防游戏。这是一个现实版喵星大战的场面，你的任务就是带领他们突破敌人包围杀出一条血路来，只有幸存下去才能保存喵星球最后的希望。', '![icon](http://lorempixel.com/g/220/220/)

## Summary

《决战喵星》是一款走萌系路线的酷跑类游戏，但这款游戏还加入了类似于塔防的元素，被成为酷跑式塔防游戏。这是一个现实版喵星大战的场面，你的任务就是带领他们突破敌人包围杀出一条血路来，只有幸存下去才能保存喵星球最后的希望。', '塔防');
insert into game values(2, '堡垒', 'bastion', '2.png', '《堡垒》是一款2D欧洲童话风格的ARPG游戏，本作由曾经担任过《命令与征服》系列作品的设计师所组成的独立工作室Supergiant Games制作，游戏中共有超过40个精美手绘场景，同时拥有全新叙述方式贯穿游戏整体。在动作，对战方面，本作描绘的也十分细腻。','![icon](http://lorempixel.com/g/220/220/)

## Summary

《堡垒》是一款2D欧洲童话风格的ARPG游戏，本作由曾经担任过《命令与征服》系列作品的设计师所组成的独立工作室Supergiant Games制作，游戏中共有超过40个精美手绘场景，同时拥有全新叙述方式贯穿游戏整体。在动作，对战方面，本作描绘的也十分细腻。

', '角色扮演');
insert into game values(3, '火箭飞车2', 'jet_car_stunts_2', '3.png', '火箭飞车2（Jet Car Stunts 2）是《火箭飞车》经典续作，而该游戏最大的可玩性就在于超长距离的飞跃、跳跃穿过空中的圆环、浮动平台以及螺旋跑道等多种游戏元素，这些都是在以前的游戏中无法体验到的。','![icon](http://lorempixel.com/g/220/220/)

## Summary

火箭飞车2（Jet Car Stunts 2）是《火箭飞车》经典续作，而该游戏最大的可玩性就在于超长距离的飞跃、跳跃穿过空中的圆环、浮动平台以及螺旋跑道等多种游戏元素，这些都是在以前的游戏中无法体验到的。

', '竞速');
insert into game values(4, '割绳子2', 'cut_the_rope_2', '4.png', '割绳子2（Cut the Rope 2）是益智游戏《割绳子》的官方正统续作，新增的5个小怪兽分别是：Roto能把Om Nom带到抓取糖果的最佳位置；Lick是个厉害的小家伙，他的长舌头能帮助Om Nom去往新地点；Blue能把Om Nom升高到更高的高度；Toss可以投掷各种物品；而坏坏的Boo通过惊吓Om Nom使他跳的更高。', '![icon](http://lorempixel.com/g/220/220/)

## Summary

割绳子2（Cut the Rope 2）是益智游戏《割绳子》的官方正统续作，新增的5个小怪兽分别是：Roto能把Om Nom带到抓取糖果的最佳位置；Lick是个厉害的小家伙，他的长舌头能帮助Om Nom去往新地点；Blue能把Om Nom升高到更高的高度；Toss可以投掷各种物品；而坏坏的Boo通过惊吓Om Nom使他跳的更高。


', '休闲益智');
insert into game values(5, '永恒战士3', 'eternity_warriors_3', '5.png', '永恒战士3（Eternity Warriors 3）是游戏厂商Glu Mobile公司风靡全球的火爆动作手游 《永恒战士》系列的最新作品。在Glu重金打造的游戏引擎支持下，游戏品质得到大幅提升。高清精致的人物建模，美轮美奂的丰富场景，精彩绝伦的动态光效，全新强大的敌兵AI,将带给玩家更加真实、爽快刺激、极具挑战的超绝游戏体验。', '![icon](http://lorempixel.com/g/220/220/)

## Summary

永恒战士3（Eternity Warriors 3）是游戏厂商Glu Mobile公司风靡全球的火爆动作手游 《永恒战士》系列的最新作品。在Glu重金打造的游戏引擎支持下，游戏品质得到大幅提升。高清精致的人物建模，美轮美奂的丰富场景，精彩绝伦的动态光效，全新强大的敌兵AI,将带给玩家更加真实、爽快刺激、极具挑战的超绝游戏体验。


', '角色扮演');