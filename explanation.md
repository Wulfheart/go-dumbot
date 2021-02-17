# How it works.

DumbBot works in two stages. In the first stage it calculates a value for
each province and each coast. Then in the second stage, it works out an
order for each unit, based on those values.

For each province, it calculates the following :

- If it is our supply centre, the size of the largest adjacent power
- If it is not our supply centre, the size of the owning power
- If it is not a supply centre, zero.

Then

Proximity[0] for each coast, is the above value, multiplied by a weighting.

Then

Proximity[n] for each coast, = ( sum( proximity[ n-1 ] for all adjacent coasts
                               * proximity[ n-1 ] for this coast )
                             / 5

Also for each province, it works out the strength of attack we have on that
province - the number of adjacent units we have, and the competition for
that province - the number of adjacent units any one other power has.

Finally it works out an overall value for each coast, based on all the proximity
values for that coast, and the strength and competition for that province, each
of which is multiplied by a weighting.

Then for move orders, it tries to move to the best adjacent coast, with a random
chance of it moving to the second best, third best, etc (with that chance varying
depending how close to first the second place is, etc). If the best place is where
it already is, it holds. If the best place already has an occupying unit, or already
has a unit moving there, it either supports that unit, or moves elsewhere,
depending on whether the other unit is guaranteed to succeed or not.

Retreats are the same, except there are no supports...

Builds and Disbands are also much the same - build in the highest value home centre
disband the unit in the lowest value location.