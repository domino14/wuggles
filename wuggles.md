base scores by WL:
                                    boggle
3s - 10 pts                         1
4s - 20 pts         (+10)           1
5s - 40 pts         (+20)           2
6s - 70 pts         (+30)           3
7s - 110 pts        (+40)           5
8s - 150 pts        (+40)           11

My additions:
9s - 200 pts
10 - 250
11 - 300
12 - 360
13 - 420
14 - 480
15 - 600

there is a Qu tile
missed words can be seen by scrolling down through the list.

round 1:  (2 minutes)       16 squares

. . . .
. . . .
. . . .
. . . .

round 2: starts in 10 secs (countdown)   24 squares

    ^ .
  . . . .
. . . . . .
. . . . . .
  . . . .
    . ^

hitting the ^ multiplies word score by 2.

round 3: after 10 seconds       28 squares

@ . . .
. . . .
. . . . . .
. . . . . .
    . . . .
    . . . @

hitting the @ multiplies word score by 3.

round 4: after 10 seconds

^ . . . . @         32 squares
. . . . . .
. .     . .
. .     . .
. . . . . .
@ . . . . ^


### Multiplayer

- Should be 8 players max
- Simple lobby
- Golang server backend. **Build a library that we can also use for my future crossword app**. Ala Yahoo Games.
- Websocket. Put some effort into a disconnect/reconnect flow.

DB Tables:

- Presence
- "Table" model vs something else (game ID?)
- Table model is more expandable but perhaps a bit antiquated. Multiple game IDs per table, every time you click Start.

Table MVP:
- game ID (changeable, like word list ID in Aerolith)
- Host
-

