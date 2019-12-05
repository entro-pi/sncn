***DESIGN DOC for SnowcrashNetwork, or sncn***

*The game loop will have a server-->client-->server loop*

This means that all actions will have to "Bounce" off
the server to get their results. Whether an attack hits,
a message gets sent, a magic missile makes it across the
void etc.

This gives the ability to make different client versions
for this, and I want to take full advantage of that fact.

***One interface will be a MUD-like setup, that is, text***
***based. Not extremely wordy, but then again not totally***
***unaware that words are what are needed to drive the story.***

***The second interface will be a pixellated version of the***
***first.***

***The third interface will be the social aspect of the game.***

using small packets of words to command your characters'
minions from outside the regular interface will hopefully
give just the right amount of "still connected" that we
as a society absolutely love.

For example, a Character named @loki has two hirelings
$noki and $nod @loki can get $noki to buy crafting supplies
for @loki by scrolling through a menu in the mobile app,
selecting $noki,
Then #purchase-crafting-->silk
and $ending the command.

I really want to get the feeling of getting something done
in the way that games like animal crossing seem to do so well.

The app will be android to start, then mac if I figure a way
to do that without owning an apple.

The app will not be a webapp in disguise. I'd like to have
some heft to this and double down that the mobile experience
will be just as enjoyable as sitting down mashing keys can be.

Some of the things that will make interactions between mobile
and tower will be if NPCs are running around doing quests and
harvesting for someone, they will have a banner of their lord
or lady and spout things about how wonderful blah is. The fun
part will be the things that NPC sees and interacts with will
make it back to the lady(or lord) whether the NPC completes
the mission or not.

Ok, THAT is a good idea. running with it.

WIP!


Solving the Kick Problem, basically making battling more exciting
than typing kick over and over. One solution I'm going to use
is to colour code your attacks, as well as have them trigger
via a Q-to-win setup that triggers when you're in a dungeon run.

Another thing is going to be voice chat and speech to text.

Fighting stances? maybe.

Dungeons will be jack-chips that can be mastered by running
them yourself, or getting a hireling to run them enough times
to "master" them.

Getting a jack-chip mastered will either reward with quest xp
or crafting xp. Depending on whether it's mastered by a character
or by a minion. Haven't decided which one it is yet.



"Digging new ground"

I want this to be a tool that will get used more
as it is needed, so I'm trying to get past the
obvious problems. The Snake problem is the
first I've seen. When one wants to turn back
on oneself, there's nothing telling the digger
that there is another room already "in" that spot

My solution is to have a grid laid out for each Digging
session and populated by "dig" movements. We'll deal
with making mazes later. For now building sanely is
the goal.

The tool will assume a grid of 50-50 with three layers
to begin with.



This will be fuuuuun!

I will have to add a timing function eventually, why not
make the first timed thing to be the blink of a new message!


map movement is going to be very atypical. Rather than running from
one place to another, there will be a quicktravel option open to
any and all players. The idea is maps that have a small description
and name, that can be selected and teleported to. **this will take some work to make it sensible, just getting the idea out of my brain**


As far as asymmetric and symmetric games go, this will be a 
start-even type game. You will basically start as a generic
user and change to the player's will as you go. This is to help
with balancing. Races and classes are usually unknown as you
start anyway, this will give you some time to come up with a
back story to your character anyways.
