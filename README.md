# AHOY, ME HARTIES

You are shipwrecked with a PC that only has Golang installed. You find a scroll with some instructions on it. The scroll consists of an array of instructions of only a single word, either TURN or WALK followed by a space and a number. 

For TURN instructions the number is in degrees so 90 would be a left turn and 270 for right. WALK is followed by a number of steps forward. Start facing north at the center of the island (by the obligatory palm tree). What is the list of instructions you can use to get to the end of the trail and how few can you do it in? As usual, points for brevity, verbosity, and performance and of course finding the treasure.

The instructions are as follows:

```
[
"WALK 14",
"TURN 90",
"WALK 13",
"TURN 270",
"WALK 12",
"TURN 270",
"WALK 8",
"TURN 270",
"WALK 7",
"TURN 270",
"WALK 11",
"TURN 270",
"WALK 9",
"TURN 90",
"WALK 12",
"TURN 90",
"WALK 15",
"TURN 270",
"WALK 15",
"TURN 270",
"WALK 10",
"TURN 90",
"WALK 9",
"TURN 90",
"WALK 14",
"TURN 90",
"WALK 6",
"TURN 270",
"WALK 7",
"TURN 90",
"WALK 6",
"TURN 90",
"WALK 10",
"TURN 270",
"WALK 15",
"TURN 270",
"WALK 15",
"TURN 270",
"WALK 14",
"TURN 90",
"WALK 6",
"TURN 90",
"WALK 6",
"TURN 270",
"WALK 6",
"TURN 270",
"WALK 10",
"TURN 270",
"WALK 15",
"TURN 270",
"WALK 13",
"TURN 270",
"WALK 13",
"TURN 270",
"WALK 6",
"TURN 90",
"WALK 13",
"TURN 270",
"WALK 10",
"TURN 90",
"WALK 14",
"TURN 270",
"WALK 12",
"TURN 270",
"WALK 10",
"TURN 270",
"WALK 6",
"TURN 270",
"WALK 8",
"TURN 90",
"WALK 11",
"TURN 90",
"WALK 10",
"TURN 270",
"WALK 14",
"TURN 90",
"WALK 15",
"TURN 270",
"WALK 12",
"TURN 90",
"WALK 14",
"TURN 90",
"WALK 12",
"TURN 90",
"WALK 12",
"TURN 90",
"WALK 9",
"TURN 90",
"WALK 14",
"TURN 90",
"WALK 7",
"TURN 270",
"WALK 10",
"TURN 90",
"WALK 13",
"TURN 270",
"WALK 8",
"TURN 270",
"WALK 14",
"TURN 270",
"WALK 9"
]
```