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
