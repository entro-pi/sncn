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


Gotten from
https://5e.tools
CR	XP	Prof B	AC	HP	Atk+	Dam/Rn	SaveDC
0	0 or 10	2	13	1-6	3	0-1	13
1/8	25	2	13	7-35	3	2-3	13
1/4	50	2	13	36-49	3	4-5	13
1/2	100	2	13	50-70	3	6-8	13
1	200	2	13	71-85	3	9-14	13
2	450	2	13	86-100	3	15-20	13
3	700	2	13	101-115	4	21-26	13
4	1,100	2	14	116-130	5	27-32	14
5	1,800	3	15	131-145	6	33-38	15
6	2,300	3	15	146-160	6	39-44	15
7	2,900	3	15	161-175	6	45-50	15
8	3,900	3	16	176-190	7	51-56	16
9	5,000	4	16	191-205	7	57-62	16
10	5,900	4	17	206-220	7	63-68	16
11	7,200	4	17	221-235	8	69-74	17
12	8,400	4	17	236-250	8	75-80	18
13	10,000	5	18	251-265	8	81-86	18
14	11,500	5	18	266-280	8	87-92	18
15	13,000	5	18	281-295	8	93-98	18
16	15,000	5	18	296-310	9	99-104	18
17	18,000	6	19	311-325	10	105-110	19
18	20,000	6	19	326-340	10	111-116	19
19	22,000	6	19	341-355	10	117-122	19
20	25,000	6	19	356-400	10	123-140	19
21	30,000	7	19	401-445	11	141-158	20
22	41,000	7	19	446-490	11	159-176	20
23	50,000	7	19	491-535	11	177-194	20
24	62,000	7	19	536-580	11	195-212	21
25	75,000	8	19	581-625	12	213-230	21
26	90,000	8	19	626-670	12	231-248	21
27	105,000	8	19	671-715	13	249-266	22
28	120,000	8	19	716-760	13	267-284	22
29	135,000	9	19	760-805	13	285-302	22
30	155,000	9	19	805-850	14	303-320	23


