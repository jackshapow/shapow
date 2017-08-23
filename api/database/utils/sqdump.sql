PRAGMA synchronous = OFF;
PRAGMA journal_mode = MEMORY;
BEGIN TRANSACTION;
CREATE TABLE `albums` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `artist_id` integer  NOT NULL
,  `name` varchar(191) NOT NULL
,  `cover` varchar(191) NOT NULL DEFAULT ''
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  CONSTRAINT `albums_artist_id_foreign` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO `albums` VALUES (1,1,'Unknown Album','unknown-album.png','2017-08-17 21:18:38','2017-08-17 21:18:38');
CREATE TABLE `artists` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `name` varchar(191) NOT NULL
,  `image` varchar(191) DEFAULT NULL
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  UNIQUE (`name`)
);
INSERT INTO `artists` VALUES (1,'Unknown Artist',NULL,'2017-08-17 21:18:38','2017-08-17 21:18:38'),(2,'Various Artists',NULL,'2017-08-17 21:18:38','2017-08-17 21:18:38');
CREATE TABLE `interactions` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `user_id` integer  NOT NULL
,  `song_id` varchar(32) NOT NULL
,  `liked` integer NOT NULL DEFAULT '0'
,  `play_count` integer NOT NULL DEFAULT '0'
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  CONSTRAINT `interactions_song_id_foreign` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`) ON DELETE CASCADE
,  CONSTRAINT `interactions_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);
CREATE TABLE `migrations` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `migration` varchar(191) NOT NULL
,  `batch` integer NOT NULL
);
INSERT INTO `migrations` VALUES (1,'2014_10_12_000000_create_users_table',1),(2,'2014_10_12_100000_create_password_resets_table',1),(3,'2015_11_23_074600_create_artists_table',1),(4,'2015_11_23_074709_create_albums_table',1),(5,'2015_11_23_074713_create_songs_table',1),(6,'2015_11_23_074723_create_playlists_table',1),(7,'2015_11_23_074733_create_interactions_table',1),(8,'2015_11_23_082854_create_playlist_song_table',1),(9,'2015_11_25_033351_create_settings_table',1),(10,'2015_12_18_072523_add_preferences_to_users_table',1),(11,'2015_12_22_092542_add_image_to_artists_table',1),(12,'2016_03_20_134512_add_track_into_songs',1),(13,'2016_04_15_121215_add_is_complilation_into_albums',1),(14,'2016_04_15_125237_add_contributing_artist_id_into_songs',1),(15,'2016_04_16_082627_create_various_artists',1),(16,'2016_06_16_134516_cascade_delete_user',1),(17,'2016_07_09_054503_fix_artist_autoindex_value',1),(18,'2017_04_21_092159_copy_artist_to_contributing_artist',1),(19,'2017_04_22_161504_drop_is_complication_from_albums',1),(20,'2017_04_29_025836_rename_contributing_artist_id',1);
CREATE TABLE `password_resets` (
  `email` varchar(191) NOT NULL
,  `token` varchar(191) NOT NULL
,  `created_at` timestamp NOT NULL 
);
CREATE TABLE `playlist_song` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `playlist_id` integer  NOT NULL
,  `song_id` varchar(32) NOT NULL
,  CONSTRAINT `playlist_song_playlist_id_foreign` FOREIGN KEY (`playlist_id`) REFERENCES `playlists` (`id`) ON DELETE CASCADE
,  CONSTRAINT `playlist_song_song_id_foreign` FOREIGN KEY (`song_id`) REFERENCES `songs` (`id`) ON DELETE CASCADE
);
CREATE TABLE `playlists` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `user_id` integer  NOT NULL
,  `name` varchar(191) NOT NULL
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  CONSTRAINT `playlists_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE `settings` (
  `key` varchar(191) NOT NULL
,  `value` text NOT NULL
,  PRIMARY KEY (`key`)
);
INSERT INTO `settings` VALUES ('media_path','s:0:"";');
CREATE TABLE `songs` (
  `id` varchar(32) NOT NULL
,  `album_id` integer  NOT NULL
,  `artist_id` integer  DEFAULT NULL
,  `title` varchar(191) NOT NULL
,  `length` double(8,2) NOT NULL
,  `track` integer DEFAULT NULL
,  `lyrics` text NOT NULL
,  `path` text NOT NULL
,  `mtime` integer NOT NULL
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  PRIMARY KEY (`id`)
,  CONSTRAINT `songs_album_id_foreign` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`)
,  CONSTRAINT `songs_contributing_artist_id_foreign` FOREIGN KEY (`artist_id`) REFERENCES `artists` (`id`) ON DELETE CASCADE
);
CREATE TABLE `users` (
  `id` integer  NOT NULL PRIMARY KEY AUTOINCREMENT
,  `name` varchar(191) NOT NULL
,  `email` varchar(191) NOT NULL
,  `password` varchar(60) NOT NULL
,  `is_admin` integer NOT NULL DEFAULT '0'
,  `preferences` text COLLATE BINARY
,  `remember_token` varchar(100) DEFAULT NULL
,  `created_at` timestamp NULL DEFAULT NULL
,  `updated_at` timestamp NULL DEFAULT NULL
,  UNIQUE (`email`)
);
INSERT INTO `users` VALUES (1,'herp derp','herp@derp.com','$2y$10$6KWTX8o/kZPxCuU0azS3qOkrkgTAUmdzs5T6ql0soeC0sRbRDsuvG',1,NULL,NULL,'2017-08-17 21:18:38','2017-08-17 21:18:38');
CREATE INDEX "idx_albums_albums_artist_id_foreign" ON "albums" (`artist_id`);
CREATE INDEX "idx_playlists_playlists_user_id_foreign" ON "playlists" (`user_id`);
CREATE INDEX "idx_playlist_song_playlist_song_playlist_id_foreign" ON "playlist_song" (`playlist_id`);
CREATE INDEX "idx_playlist_song_playlist_song_song_id_foreign" ON "playlist_song" (`song_id`);
CREATE INDEX "idx_songs_songs_album_id_foreign" ON "songs" (`album_id`);
CREATE INDEX "idx_songs_songs_contributing_artist_id_foreign" ON "songs" (`artist_id`);
CREATE INDEX "idx_password_resets_password_resets_email_index" ON "password_resets" (`email`);
CREATE INDEX "idx_password_resets_password_resets_token_index" ON "password_resets" (`token`);
CREATE INDEX "idx_interactions_interactions_user_id_foreign" ON "interactions" (`user_id`);
CREATE INDEX "idx_interactions_interactions_song_id_foreign" ON "interactions" (`song_id`);
END TRANSACTION;
