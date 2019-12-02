Dec 1st, 2019

bweeeeee!

I refactored the entire thing into a seperate
golang-client and elixir server. As well as
rabbitmq in place of zeromq, postgresql in
place of mongodb and am working under the 
assumption that one would play the game with
multiple monitors. If you have Opinions on 
this, join the discord or slack and let your voice
be heard!

https://discord.gg/7gu2tkX

Cheers!
Weasel

![ouroboro](docs/ouroboro.png)

***November 16th, 2019***
I tossed the idea of profile photos in favour of

an equipment/inventory system. I'm thinking about how

to roll in a battle system and what kind I'd like it

to be. Currently the old implementation got broken along

the way and I'm not too excited about re-doing it as it

was mostly placeholder anyway. The one part I want to keep

is the roar intro to battle. That got me jazzed every time.

Especially when I surprised myself with it.



I'm thinking something odd for the battle system. For one

it will not be one kind of battle for everything. I want

to have different types of battle for different classes.

For instance, a dexterous hacker has a ddr-style keyboard

battle to produce battlespam. (this is going to be

a typing test game) While a harmonious enchanter will have

a keypad game of simon to produce the same.

Matching the patterns correctly will result in more "hits"

and more messups will correlate to "misses" and getting hit.




I'm unsure of multiclassing at this point. I wanted to be

able to multiclass, but that was when I was thinking this

thing would look more like a diku than anything else.

I think the biggest thing is going to be making sure there's

little to no penalty to trying out a class, but advancing in

a class will take real life time digestion of skill points.

Ie I want to advance in hacker to close the level gap

between me and my prey so the keyboard battle is less

insanely hard. I select my chosen path, the amount of

points to spend, and it gives me an amount of time until

that is complete. I agree and it starts digesting. I come

back in a few hours and lo and behold! A battle that was

once inhumanely hard is now manageable.


As for social spaces, the only way to "see" another player

right now is if they post to a public channel.



In other news I migrated the vast majority of sncn's backend

to a mongo atlas server, and increased the reactivity of the

game tenfold.


![sunshot](docs/Screenshot from 2019-11-07 21-40-15.png)

****November 8th, 2019****

I'm in the middle of adding the social aspects of the game.

Currently it will be an interface like the above, with
the left hand side holding your profile photo, as well
as your stats. Below in the empty space I will put
contents of the cels in the center top. So the short
messages will act more like a subject line if marked so.
If a "subject" cel has more content to it, it will show
up as different colour with a "...Read More" tag.

You will be able to switch through each messaage using
the arrow keys thanks to termbox-go. To start with each
message will hold a "core seed" which will allow a player
to generate a "corrupted coreboard" which is basically
a fun way of saying spawn a dungeon to be run.

Other functionality added would be the ability to either
capture a picture from your webcamera for use as a profile
picture, or using a pre-taken image and rasterizing it
to be used as a profile picture.

I think one thing I'm going to add right now is a "HELP"
option that spams all the available commands, as there
is getting to be a lot of them.

As well as generate some splash screens for the starting
coreboard set.

That's all for now!

**UPDATE**
Here's a list of the current commands and what they do

HELP

pewpew <num> = play a sound
broadside=<num>:<num> = show a broadcast at a position
bs=<num> = select broadcast numbered
show soc = show social broadcasts
soc = generate social broadcasts
hide soc = hide social broadcasts

capture profile picture = get a picture from the camera
load profile picture = use an existing photo

score = show player info
addclass = adds a class type from a template

SAVE ZONES = dumps the zone files to dat/zone.bson

view from <num> = shows the view at a certain vnum

craft mobile = UNFINISHED create a mobile

open map = does nothing right now

gen coreboard=<num>:<num> = create a coreboard dungeon of size numXnum
tc:<num>|<num> = enter said coreboard at position num, num

look = look at the current room

report = show current players classes

show chat = show the out of character chats
hide chat = hides said chat

show grape = shows the grapevine broadcasts
hide grape = hides the grapevine broadcasts

gvsub <channelname> = subscribe to a grapevine channel
gvunsub <channelname> = unsubscribe from a grapevine channel
g:<words> = compose a grapevine chat


update zonemap = copies update zone maps to rooms

merge <zonesource> <zonedest> = merges cross zone map positions

count keys = counts a number of characters

show channels = shows grapevine channels currently subscribed to

edit desc = edit the room Description

show zone info = get information about the current zone

heal = gain Rezz and Tech

dam rezz = damage your rezz

dam tech = damage your tech

show room vnum = shows the room virtual number

shutdown server = shut the server down

login <player> <pass> = login with the supplied credentials

create <player> <pass> = create a player with the supplied credentials
                          most useful when using the dorp:norp initilizer characters

blit = update the output with a clean copy of your view

ooc <words> = out of character chat

dig new = create rooms

save = save your character

quit = exit game
logout = another way to exit



![dead](docs/dead.png)
****November 6th, 2019****

So for now I've settled on a way to handle battle.

Basically there is another "mode" rather than straight text input.

It's accessible by generating a coreboard with ```gen coreboard```
and ```lock coreboard``` ```tc:1|1``` will plop you in to space 1,1

![coreboard](docs/coreboard.png)

As well as battle, I've settled on a way to start the game itself.

Once you have sncn-core on your computer, run ```sncn-core --wizinit```

then ```sncn-core --connect-core``` for the full experience, ```sncn-core --connect-core --safe-mode```
for the fallback.

It autogenerates a default persona to be used to create other characters.
Also to keep the entire engine from crapping out due to lack of playerfiles.

As for logins, today I'm implementing a hash pattern that allows for
more than one character to have a certain password. If that sounds
like the most ridiculous problem to have, you would be right. However it
is an easy fix. When I created the password hash, I only took in to
account the accounts password, which meant that if you used the same
password as someone else, you would be logged in to whomever showed
up first in the database. Which was not who it was expected to be
half the time.

But! The simple but elegant solution is to calculate the hash according
to the player's name AND password to create a unique pass/hash pair.

I thought about creating a UUID and locking it behind a password, but
I couldn't get past the fact I would have to save the password in plain
text somewhere.

So there! hash(pass) becomes hash(name+pass) and all of a sudden everyone
is unique!

The next slice of life improvement bit will be to make your login different
from your character, so one doesn't have to advertise half their
login credentials simply by logging in.

Till next time!

Entropy




![glitch](docs/glitch.png)
****November 1st, 2019****

Today was a productive day, and I managed to create a lynchpin part of this MUD. The interface for the roguelike section of snowcrash.network

![old](docs/old.png)

This will populate with randomized creatures and treasure, right now they show up as golden tiaras and rabid ferrets.

![coreboard](docs/coreboard.png)

You can clip through walls, but that'll be an easy fix.

The battle spam design will be tomorrow, and you can bet I'm going to do something fun for that. I'm thinking a health
bar that spans the top screen, giving an imperative feel to fighting. As well for skills and attacks, things are going
to be tied to maybe half a dozen different attacks at any one time. I want to work stances in as well, to give that
extra bit of "off" to help balance.


Also, I think I'm going to commit the heretical sin of making a MOO inside the MUD. The social aspect of this game will
rely on crafting for dungeon runs, and crafting rooms and spaces in a pve only environment.

Also I should mention, this game is going to be best played as WSAD(Q-to-win) with keypad, or straight up game pad.

Because face it, we all have at least some kind of game pad lying around. Mine is a knockoff gamecube controller.

![login](docs/login.png)


****October 30th, 2019****
**The wonders of the year 2019**


Hah, it felt futuristic to type that.


The topic of today's blog is networking and hash functions. This MUD's first quirk is that it is not going to be running over telnet, and will require a client either made by the community or by me. The MUD itself is going to be a public facing API, linking playerfiles to database data. The user authenticates in their client, the connection from the client to the auth server is encrypted, and the password is stored in a hash table, which is a table of values that are consistently mangled by something called a hash function.


The hash function can be anything you want it to be, as long as it's consistent. For example. I'm going to take a UUID created by a chunk of code I wrote for the original social network, iterate over the characters and perform an arbitrary mathematical function on them. It will produce a string of "garbage", which we will store. The client will send a mangled key of its own, however the server will do the mangling, and associate the user with the key send at that time. As with most APIs, there will be a scope for each key, stored as a prefix on the hash function.
You might ask why bother? The answer is paranoia. If the database were ever to be compromised, the attacker dumping our precious playerfiles into their pockets, they would get a useless hashed garbage string instead of the password you happen to share with a facebook account.

So that's my reasoning.


Currently I have a working communication that utilizes zeromq for networking. zmq is one of my desert island libraries, I always manage to build something useful or fun with it with minimal hassle. 

![zmq](docs/zmq.png)

and there we are! small asynchronous atomic communication! Encrypted by a set of keys generated at runtime even! No re-using old keys or storing them like so many acorns. However, this is the solution to a totally different problem than we were talking about. The hash function will look more like this

![hash](docs/hash.png)

Now for securities sake I'm not going to be keeping this formula, but it makes sense, no? mangle the string but do it in a way that's reproducible. That's what a hash function is all about.


So there we go! I'm working on refactoring the current client into a zmq utilizing API otherwise, so that's all for tonight.