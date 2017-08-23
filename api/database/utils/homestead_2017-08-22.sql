# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.19-0ubuntu0.16.04.1)
# Database: homestead
# Generation Time: 2017-08-22 23:17:52 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table albums
# ------------------------------------------------------------

LOCK TABLES `albums` WRITE;
/*!40000 ALTER TABLE `albums` DISABLE KEYS */;

INSERT INTO `albums` (`id`, `artist_id`, `name`, `cover`, `created_at`, `updated_at`)
VALUES
	(1,1,'Unknown Album','unknown-album.png','2017-08-17 21:18:38','2017-08-17 21:18:38'),
	(2,2,'Modern Vampires Of The City','599cb5e12f5d89.97869087.jpeg','2017-08-22 22:53:21','2017-08-22 22:53:21'),
	(3,2,'Kaleidoscope EP','599cb7202544d1.46641926.jpeg','2017-08-22 22:58:40','2017-08-22 22:58:40');

/*!40000 ALTER TABLE `albums` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table artists
# ------------------------------------------------------------

LOCK TABLES `artists` WRITE;
/*!40000 ALTER TABLE `artists` DISABLE KEYS */;

INSERT INTO `artists` (`id`, `name`, `image`, `created_at`, `updated_at`)
VALUES
	(1,'Unknown Artist',NULL,'2017-08-17 21:18:38','2017-08-17 21:18:38'),
	(2,'Various Artists',NULL,'2017-08-17 21:18:38','2017-08-17 21:18:38'),
	(3,'Vampire Weekend',NULL,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	(4,'Coldplay',NULL,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	(5,'Coldplay & Big Sean',NULL,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	(6,'Coldplay & The Chainsmokers',NULL,'2017-08-22 22:58:40','2017-08-22 22:58:40');

/*!40000 ALTER TABLE `artists` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table songs
# ------------------------------------------------------------

LOCK TABLES `songs` WRITE;
/*!40000 ALTER TABLE `songs` DISABLE KEYS */;

INSERT INTO `songs` (`id`, `album_id`, `artist_id`, `title`, `length`, `track`, `lyrics`, `path`, `mtime`, `created_at`, `updated_at`)
VALUES
	('1806358967a29fe64b6822fafdd72e5f',3,4,'A L I E N S',282.50,3,'','/home/vagrant/Code/koel/music/Coldplay - Kaleidoscope EP (2017)/03. A L I E N S.mp3',1503442687,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	('19bfa1969dab011d47b447e12b17f63a',2,3,'Ya Hey',312.69,10,';Oh, sweet thing\r\nZion doesn\'t love you\r\nAnd Babylon don\'t love you\r\nBut you love everything\r\nOh, you saint\r\nAmerica don\'t love you\r\nSo I could never love you\r\nIn spite of everything\r\n\r\nIn the dark of this place\r\nThere\'s the glow of your face\r\nThere\'s the dust on the screen\r\nOf this broken machine\r\nAnd I can\'t help but feel\r\nThat I\'ve made some mistake\r\nBut I let it go\r\nYa Hey Ya Hey Ya Hey\r\n\r\nThrough the fire and through the flames\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nYou won\'t even say your name\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nThrough the fire and through the flames\r\nYou won\'t even say your name\r\nOnly I am that I am\r\nBut who could ever live that way?\r\n(Ya Hey Ya Hey)\r\nUt Deo, Ya Hey\r\nUt Deo, Deo\r\n\r\nOh, the motherland don\'t love you\r\nThe fatherland don\'t love you\r\nSo why love anything?\r\nOh, good God\r\nThe faithless they don\'t love you\r\nThe zealous hearts don\'t love you\r\nAnd that\'s not gonna change\r\n\r\nAll the cameras and files\r\nAll the paranoid styles\r\nAll the tension and fear\r\nOf a secret career\r\nAnd I think in your heart\r\nThat you\'ve seen the mistake\r\nBut you let it go\r\nYa Hey Ya Hey Ya Hey \r\n\r\nThrough the fire and through the flames\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nYou won\'t even say your name\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nThrough the fire and through the flames\r\nYou won\'t even say your name\r\nOnly I am that I am\r\nBut who could ever live that way?\r\n(Ya Hey Ya Hey)\r\nUt Deo, Ya Hey\r\nUt Deo, Deo\r\n\r\nOutside the tents, on the festival grounds\r\nAs the air began to cool, \r\nand the sun went down\r\nMy soul swooned, \r\nas I faintly heard the sound\r\nOf you spinning Israelites\r\nInto 19th Nervous Breakdown\r\n\r\nThrough the fire and through the flames\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nYou won\'t even say your name\r\n(Ya hey Ya hey, Ut Deo, ya hey ya hey )\r\nThrough the fire and through the flames\r\nYou won\'t even say your name\r\nOnly I am that I am\r\nBut who could ever live that way?\r\n(Ya Hey Ya Hey)\r\nUt Deo, Ya Hey\r\nUt Deo, Deo\r\n\r\nThrough the fire and through the flames\r\nYou won\'t even say your name\r\nOnly I am that I am\r\nBut who could ever live that way?\r\n(Ya Hey Ya Hey )\r\nUt Deo, Ya Hey\r\nUt Deo, Deo Annotate','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/10 Ya Hey.mp3',1502582121,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('1a0746a20d7e06393bc91d584616dd15',3,4,'All I Can Think About Is You',274.66,1,'','/home/vagrant/Code/koel/music/Coldplay - Kaleidoscope EP (2017)/01. All I Can Think About Is You.mp3',1503442692,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	('2d6fc57ed956048beb7e78ff8c5ddcb2',2,3,'Diane Young',160.10,4,';You torched a Saab like a pile of leaves\r\nI\'m gonna to find some better wheels\r\nFour-five meters running on around the bend\r\nWhen the government agents surround you again\r\nIf Diane Young won\'t change your mind\r\nBaby baby baby baby right on time\r\n\r\nOut of control but you\'re playing a role\r\nDo you think you can go to the eighteenth hole\r\nWell, you flip-flop the day of the championship\r\nTry to go it alone, you roll for a bit\r\nIf Diane Young won\'t change your mind\r\nBaby baby baby baby right on time\r\n\r\nBaby baby baby baby right on \r\nBaby baby baby baby right on \r\nBaby baby baby it\'s a light on\r\nBaby baby baby it\'s a lifetime\r\nBaby baby baby baby right on time \r\nTime time time time\r\nBaby baby baby baby right on time\r\n\r\nUh let \'em go\r\nIf Diane Young won\'t change your mind\r\nBaby baby baby baby right on time\r\n\r\nIrish and proud baby naturally\r\nBut you got the luck of a Kennedy\r\nSo grab the wheel, keep holding it tight\r\n\'Til you\'re driving off in the black of the night\r\nIf Diane Young won\'t change your mind\r\nBaby baby baby baby right on time\r\n\r\nBaby baby baby baby right on \r\nBaby baby baby baby right on \r\nBaby baby baby it\'s a light on\r\nBaby baby baby it\'s a lifetime\r\nBaby baby baby baby right on time \r\nTime time time time\r\nBaby baby baby baby right on time\r\n\r\nNobody knows what the future holds on\r\nSaid it\'s bad enough just getting old\r\nLive my life, they say it\'s too fast\r\nYou know I love the past, \'cause I used to spend \r\nIf Diane Young won\'t change your mind\r\nBaby baby baby baby right on time','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/04 Diane Young.mp3',1502582119,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('34f380415095420cdce93f7779e3de8d',2,3,'Everlasting Arms',183.35,7,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/07 Everlasting Arms.mp3',1502582123,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('38e5980d86c6af478995801d93eef032',2,3,'Young Lion',105.40,12,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/12 Young Lion.mp3',1502582121,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('3fa0fb6dc0d9391e42b1e97cb87e5c0c',2,3,'Hannah Hunt',238.00,6,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/06 Hannah Hunt.mp3',1502582120,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('4440a36cd128cd332e8a8c737e29dd3e',2,3,'Worship You',201.27,9,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/09 Worship You.mp3',1502582124,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('4f2083b7460591e956b635376c66c372',2,3,'Obvious Bicycle',251.30,1,';Morning’s come, you watch the red sun rise,\r\nThe LED still flickers in your eyes,\r\nOh you outta’ spare your face the razor,\r\nBecause no one’s gonna’ spare \r\nthe time for you\r\n\r\nNo one’s gonna’ watch you as you go,\r\nFrom a house you didn’t build \r\nand can’t control,\r\nOh, you outta’ spare \r\nyour face the razor,\r\nBecause no one’s gonna’ spare \r\nthe time for you\r\nWhy don’t you spare \r\nthe world your lab our,\r\nIt’s been twenty years \r\nand no one’s told the truth\r\n\r\nSo listen, Oh,\r\nSo listen, Oh,\r\nDon’t wait\r\nDon’t wait\r\n\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)\r\n\r\nSo keep that list of who to thank in mind,\r\nAnd don’t forget the rich ones \r\nwho were kind,\r\nOh, and don’t you spare \r\nyour face the razor,\r\nBecause no one’s gonna’ spare \r\nthe time for you\r\nWhy don’t you spare \r\ntheir word’s are ‘traitor’,\r\nTake your witch back \r\nand leave before you lose\r\n\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)\r\n\r\nSo while the sun’s coming down,\r\nCover ground,\r\nCover ground\r\n\r\nAnd if you found some love \r\nfor these clowns,\r\nTurn around,\r\nTurn around\r\n\r\nI’ll be asleep on the floor of our \r\nhigh school GYM,\r\nThinking of you \r\nand wondering if anyone else could begin,\r\n\r\nTo listen,\r\nListen\r\n\r\nDon’t wait\r\nDon’t wait\r\n\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nSo listen, Oh,\r\n(Listen, ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)\r\nDon’t wait\r\n(Don’t wait ay hey ay hey ay hey)','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/01 Obvious Bicycle.mp3',1502582122,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('691d25fba7dbbbefaa4c680a60b3854f',3,5,'Miracles (Someone Special)',276.97,2,'','/home/vagrant/Code/koel/music/Coldplay - Kaleidoscope EP (2017)/02. Miracles (Someone Special).mp3',1503442683,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	('700b16100d01df3ea872355f72174623',2,3,'Finger Back',206.00,8,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/08 Finger Back.mp3',1502582123,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('7eb70058d522ac3293c18608419d7aa9',3,6,'Something Just Like This (Tokyo Remix)',273.79,4,'','/home/vagrant/Code/koel/music/Coldplay - Kaleidoscope EP (2017)/04. Something Just Like This (Tokyo Remix).mp3',1503442683,'2017-08-22 22:58:40','2017-08-22 22:58:40'),
	('b3bf16db26886b4d571d5c76a1d259d7',2,3,'Hudson',254.85,11,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/11 Hudson.mp3',1502582121,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('b65ae046a159321cc1cd2c254c71fbe5',3,4,'Hypnotised (EP Mix)',391.45,5,'','/home/vagrant/Code/koel/music/Coldplay - Kaleidoscope EP (2017)/05. Hypnotised (EP Mix).mp3',1503442686,'2017-08-22 22:58:41','2017-08-22 22:58:41'),
	('d65df9004ad5303d4c8848eb732567f7',2,3,'Step',251.66,3,';Every time I see you in the world, \r\nyou always step to my girl\r\n\r\nBack back way back I used to \r\nfront like Angkor Wat\r\nMechanicsburg Anchorage \r\nand Dar es Salaam\r\nWhile home in New York \r\nwas champagne and disco\r\nTapes from L.A. slash San Francisco\r\nBut actually Oakland and not Alameda\r\nYour girl was in Berkeley \r\nwith her Communist reader\r\nMine was entombed within \r\nboom box and walkman\r\nI was a hoarder but girl that was back then\r\n\r\nThe gloves are off, the wisdom teeth are out\r\nWhat you on about?\r\nI feel it in my bones, I feel it in my bones\r\nI\'m stronger now, I\'m ready for the house\r\nSuch a modest mouse,\r\nI can\'t do it alone, I can\'t do it alone\r\n\r\nEvery time I see you in the world, \r\nyou always step to my girl\r\n\r\nAncestors told me that their girl was better\r\nShe\'s richer than Croesus, \r\nshe\'s tougher than leather\r\nI just ignored all the tales of a past life\r\nStale conversation deserves but a bread knife\r\nAnd punks who would laugh \r\nwhen they saw us together\r\nWell, they didn\'t know how to dress \r\nfor the weather\r\nI can still see them there huddled on Astor\r\nSnow falling slow to the sound of the master\r\n\r\nThe gloves are off, the wisdom teeth are out\r\nWhat you on about?\r\nI feel it in my bones, I feel it in my bones\r\nI\'m stronger now, I\'m ready for the house\r\nSuch a modest mouse,\r\nI can\'t do it alone, I can\'t do it alone\r\n\r\nWisdoms a gift, but you\'d trade it for youth\r\nAge is an honor, it\'s still not the truth\r\nWe saw the stars \r\nwhen they hid from the world\r\nYou cursed the sun when \r\nit stepped to your girl\r\nMaybe she\'s gone \r\nand I can\'t resurrect her\r\nThe truth is she doesn\'t \r\nneed me to protect her\r\nWe know the true death, \r\nthe true way of all flesh\r\nEveryone\'s dying, \r\nbut girl youre not old yet\r\n\r\nThe gloves are off, the wisdom teeth are out\r\nWhat you on about?\r\nI feel it in my bones, I feel it in my bones\r\nI\'m stronger now, I\'m ready for the house\r\nSuch a modest mouse,\r\nI can\'t do it alone, I can\'t do it alone\r\n\r\nThe gloves are off, the wisdom teeth are out\r\nWhat you on about?\r\nI feel it in my bones, I feel it in my bones\r\nI\'m stronger now, I\'m ready for the house\r\nSuch a modest mouse,\r\nI can\'t do it alone, I can\'t do it alone','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/03 Step.mp3',1502582120,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('e9f74609f248f1d6a67b840628a62e6a',2,3,'Don\'t Lie',213.45,5,'','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/05 Don\'t Lie.mp3',1502582120,'2017-08-22 22:53:21','2017-08-22 22:53:21'),
	('fbb13f9b59e6712e3bd5226335182c53',2,3,'Unbelievers',202.71,2,';Got a little soul\r\nThe world is a cold, cold place to be\r\nWant a little warmth\r\nBut who’s going to save a little warmth for me\r\n\r\nWe know the fire awaits unbelievers\r\nAll of the sinners the same\r\nGirl you and I will die unbelievers \r\nbound to the tracks of the train\r\n\r\nSee the snow come down\r\nIt’s coming on down from the highest peak\r\nWant a little leaf, \r\nbut who’s going to save a little leaf for me\r\n\r\nWe know the fire awaits unbelievers\r\nAll of the sinners the same\r\nGirl you and I will die unbelievers \r\nbound to the tracks of the train\r\n\r\nI’m not excited\r\nBut should I be\r\nIs this the fate that half of the world \r\nhas planned for me?\r\n\r\nI know I love you\r\nAnd you love the sea\r\nWonder if the water contains \r\na little drop little drop for me\r\n\r\nSee the sun go down\r\nIt’s going on down when the night is deep\r\nWant a little light \r\nbut who’s going to save a little light for me?\r\n\r\nWe know the fire awaits unbelievers\r\nAll of the sinners the same\r\nGirl you and I will die unbelievers \r\nbound to the tracks of the train\r\n\r\nIf I’m born again I know \r\nthat the world will disagree\r\nWant a little grace but who’s going to say \r\na little grace for me?\r\n\r\nWe know the fire awaits unbelievers\r\nAll of the sinners the same\r\nGirl you and I will die unbelievers \r\nbound to the tracks of the train\r\n\r\nI’m not excited\r\nBut should I be\r\nIs this the fate that half of the world \r\nhas planned for me?\r\n\r\nI know I love you\r\nAnd you love the sea\r\nWonder if the water contains \r\na little drop little drop for me\r\n\r\nI’m not excited\r\nBut should I be\r\nIs this the fate that half of the world \r\nhas planned for me?\r\n\r\nI know I love you\r\nAnd you love the sea\r\nWonder if the water contains \r\na little drop little drop for me','/home/vagrant/Code/koel/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/02 Unbelievers.mp3',1502582121,'2017-08-22 22:53:21','2017-08-22 22:53:21');

/*!40000 ALTER TABLE `songs` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
