# Team Generator

Team Generator is a web application to help game organizers/drop-in admins divide players into teams in a game session. User can rate players to better reflect their abilities and let the application divide the team equally depending on skill levels. Players can also improve their ratings by winning games.

## How do you divide the teams

We use a special algorithm that we believe can divide the teams as equal as possible while still maintaining the "randomness". Therefore, some of the time, some teams can be very OP but when it happens, the user can always redraw the team. Here is the step-by-step:

1. First we shuffle the players in a randomly ordered list.
2. Calculate the ideal maximum ratings that each team can take. This is done by calculating total ratings of all players divided by the number of teams + half of average player ratings.
3. we will try to put the first player on the first team and see if the team becomes too OP (a.k.a going above ideal maximum).
4. If the team becomes too OP, then try to put the player on the next team and see if the players can fit in before meeting the team rating cap.
5. If for some reason the player cannot fit in any of the teams anymore, then add the player to the team that has the lowest current overall ratings that still needs players.
6. Repeat the procedure until team is fully assigned.

Database Schema:
```
## User
- id: int
- username: string
- name: string
- password: string

## Player
- id: int
- name: string
- ratings: int
- form: int
- community_id (FK): int

## Community
- id: int
- name: string,  unique
- location: string
- admin_id (FK): **NOT IMPLEMENTED YET**

##  Sessions
- id
- community_id (FK)
- State (Enum: Draft, Closed)
- player_id[] (FK)

## Ratings
- id: int
- rating_gain: int
- user_id (FK): int
- session_id (FK): int

## SessionPlayer
- id: int
- player_id (FK): int
- community_id (FK): : int
- rating: int
- form: int
- form_change: int
```
