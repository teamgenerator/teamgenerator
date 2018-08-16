import shuffle from 'shuffle-array';

const generateRandomTeam = (players, playersPerTeam) => {

  const shuffledPlayers = shuffle(players);

  const numberOfTeams = Math.ceil(shuffledPlayers.length / playersPerTeam);

  const teams = [];

  const totalRatings = players.reduce((acc, p) => acc + p.ratings, 0);

  const maxTeamRatings = totalRatings / numberOfTeams;

  console.log('total ratings: ', totalRatings);
  console.log('number of teams: ', numberOfTeams);
  console.log("max ratings per team: ", maxTeamRatings);

  for (let i = 0; i < numberOfTeams; i++) {
    teams[i] = [];
  }

  for (let i = 0; i < shuffledPlayers.length; i++) {
    let teamToPutPlayerIndex = 0;

    while (
      teams[teamToPutPlayerIndex].length >= playersPerTeam &&
      teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings > maxTeamRatings
    ) {
      teamToPutPlayerIndex++;
    }

    teams[teamToPutPlayerIndex].push(shuffledPlayers[i]);
  }

  return teams;
}

export default generateRandomTeam;