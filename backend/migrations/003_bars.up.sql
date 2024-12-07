DROP TABLE IF EXISTS bars;

CREATE TABLE IF NOT EXISTS bars (
    id SERIAL PRIMARY KEY,
    lyrics VARCHAR NOT NULL,
    wins BIGINT DEFAULT 0 NOT NULL,
    song_name VARCHAR REFERENCES songs(title)
);

INSERT INTO albums (album_name, album_year) VALUES('Views', '2016');
INSERT INTO albums (album_name, album_year) VALUES('Take Care', '2011');
INSERT INTO albums (album_name, album_year) VALUES('Nothing was the Same', '2013');
INSERT INTO songs (title, album_name) VALUES('Feel No Ways', 'Views');
INSERT INTO songs (title, album_name) VALUES('Over My Dead Body', 'Take Care');
INSERT INTO songs (title, album_name) VALUES('Furthest Thing', 'Nothing was the Same');

INSERT INTO bars(lyrics, song_name) VALUES('I should be downtown, whipping on the way to you', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('You got something that belongs to me', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Your body language says it all', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Despite the things you said to me', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Who is it that''s got you all gassed up? (Yeah)', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Changing your opinion on me', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('I was only gone for the last few months', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('But you don''t have the time to wait on me (yeah)', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('I tried with you. There''s more to life than sleeping in and getting high with you', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('I had to let go of us, to show myself what I could do', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('And that just didn''t sit right with you (yeah)', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('And now you''re tryna make me feel a way, on purpose', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Now you''re throwing it back in my face, on purpose', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Now you''re talking down on my name (ah), on purpose (yeah)', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('And you don''t feel no way, you think I deserve it', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Feel a way, feel a way, young ni**a feel a way', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('I''ve stopped listening to things you say, ''Cause you don''t mean it anyway, yeah', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Feel a way, feel a way, young ni**a feel a way', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Maybe we just should have did things my way, Instead of the other way', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, yeah, yeah', 'Feel No Ways');
INSERT INTO bars(lyrics, song_name) VALUES('I stopped listening to things you say', 'Feel No Ways');


INSERT INTO bars(lyrics, song_name) VALUES('I think I killed everybody in the game last year, Man, fuck it, I was on though', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And I thought I found the girl of my dreams at the strip club, Mm-mm, fuck it, I was wrong though', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Shout out to all my ni**as livin'' tax free. Nowadays, it''s six figures when they tax me', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Oh well, guess you lose some and win some. Long as the outcome is income', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('You know I want it all and then some. Shout out to Asian girls, let the lights dim some', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Shots came, I don''t know where they was sent from. Probably some bad hoes I''m ''bout to take the hint from', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, you know me well, ni**a. Yeah, I mean you ain''t the only real ni**a', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('They got me on these white women like Seal, ni**a. Slave to the pussy but I''m just playin'' the field, ni**a, yeah', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Are these people really discussin'' my career again? Askin'' if I''ll be goin'' platinum in a year again?', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Don''t I got the shit the world wanna hear again? Don''t Michael Jordan still got his hoop earring in?', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Man, all of your flows bore me, paint dryin. And I don''t ever be trippin'' off of what ain''t mine', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And I be hearin'' the shit you say through the grapevine. But jealousy is just love and hate at the same time', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, it''s been that way from the beginnin''. I just been playin'', I ain''t even notice I was winnin''', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And this is the only sound you should fear. Man, these kids wear crowns over here and everything is alright', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('You say I''m old news, well, who the new star? ''Cause if I''m goin'' anywhere, it''s probably too far', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Just performed at a Bar Mitzvah over in the States. Used half of the money to beat my brother''s case', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Red wine over fed time. But shout out to the ni**as that''s doin'' dead time', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And shout out to the bitches there when it''s bedtime. And fuck you to the ni**as that think it''s their time', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, don''t make me take your life apart, boy. You and whoever the fuck gave you your start, boy', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Oh, you wanna be a motherfuckin'' funny guy? Don''t make me break your Kevin Hart, boy', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, it''s whatever. You know, feelin'' good, livin'' better', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('I think maybe I was numb to it last year. But you know I feel it now more than ever', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('My city love me like Mac Dre in the Bay. Second album, I''m back pavin'' the way', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('The backpackers are back on the bandwagon. Like this was my comeback season back, back in the day', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And I met your baby moms last night. We took a picture together, I hope she frames it', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('And I was drinkin'' at the Palms last night. And ended up losin'' everything that I came with', 'Over My Dead Body');
INSERT INTO bars(lyrics, song_name) VALUES('Yeah, feel like I''ve been here before, huh? I still got ten years to go, huh?', 'Over My Dead Body');

INSERT INTO bars(lyrics, song_name) VALUES('Somewhere between psychotic and iconic. Somewhere between I want it and I got it', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Somewhere between I''m sober and I''m lifted. Somewhere between a mistress and commitment', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('But I stay down, girl, I always stay down. Get down, never lay down', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Promise to break everybody off before I break down. Everyone just wait now', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('So much on my plate now. People I believed in, they don''t even show their face now', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('What they got to say now? Nothing they can say now', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Nothing really changed, but still they look at me a way now. What more can I say now?', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('You might feel like nothing was the same. I still been drinking on the low', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Mobbing on the low. Fucking on the low', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Smoking on the low. I still been plotting on the low', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Scheming on the low. The furthest thing from perfect, like everyone I know', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('And I hate that you don''t think that I belong to ya. Just too busy running shit to run home to you', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('You know that paper my passion. Bittersweet celebrations, I know I can''t change what happened', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('I can''t help it, I can''t help it. I was young and I was selfish', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('I made every woman feel like she was mine and no one else''s. Now you hate me, stop pretending', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Stop that fronting, I can''t take it. Girl, don''t treat me like a stranger', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Girl, you know I seen ya naked. Girl, you know that I remember, don''t be a pretender', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Getting high at the condo, that''s when it all comes together. You know I stay reminiscing and make-up sex is tradition', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('But you''ve been missing girl, and you might feel like nothing was the same. I still been drinking on the low', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Drinking, smoking, fucking, plotting, scheming. Plotting, scheming, getting money', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Uh, this the life for me. My mama told me this was right for me', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('I got ''em worried like make sure you save a slice for me. I should have spoons serve you up with a fork and knife for me', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('The actions make us doubt you. Your lack of effort got me rapping different', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('This the shit I wanna go out to. Play this shit at my funeral if they catch me slipping', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Naked women swimming that''s just how I''m living. Donate a million to some children that''s just how I''m feeling', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('A ni**a filling up arenas, who the fuck can see us? I had to Derrick Rose the knee up ''fore I got the re-up', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Yours truly, the boy. I just build and build more, y''all ni**as build and destroy', 'Furthest Thing');
INSERT INTO bars(lyrics, song_name) VALUES('Y''all ni**as party too much man, I just chill and record. No filler, you feel it now if you ain''t feel it before', 'Furthest Thing');
