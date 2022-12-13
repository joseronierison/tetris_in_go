# Board Data Structure

## Status

This version was accepted for the first version 0.0.0

## Context

There is a need for a representation of a board that enables both to draw the board and compute it physics.

It should be easier to draw since it will happens in each game frame (tick) and easy to compute the game rules.

## Decision

For that need an array of array with boolean like this `[49][28]bool` where true means that the field is occupied and false that is clear.

That is a example, where 0 represents a false and 1 represents true.
```
|000000000|
|000001000|
|000001000|
|000001110|
|000000000|
```

## Consequences

However it was really easy to draw the gameboard, compute the game rules required a step more since the data structure is a columns of lines instead of lines of columns. Currently the code is reversing it to compute game rules.

```
the array goes
this direction
    v
---------                                ---------
000000000            it should be       [000000000]
000001000            better this was    [000001000]
000001000                               [000001000] < the array goes this direction
000001110                               [000001110]
000000000                               [000000000]
---------                                ---------
```

It would improve the way to compute score since the game rules are applied line by line from bottom to top.
