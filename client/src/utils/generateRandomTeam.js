import shuffle from 'shuffle-array';

const generateRandomTeam = (players, playersPerTeam) => {

  const shuffledPlayers = shuffle(players);

  const numberOfTeams = Math.ceil(shuffledPlayers.length / playersPerTeam);

  const teams = [];

  for (let i = 0; i < numberOfTeams; i++) {
    teams[i] = [];
  }

  for (let i = 0; i < shuffledPlayers.length; i++) {
    let teamToPutPlayerIndex = 0;

    while (teams[teamToPutPlayerIndex].length >= playersPerTeam) {
      teamToPutPlayerIndex++;
    }

    teams[teamToPutPlayerIndex].push(shuffledPlayers[i]);
  }

  return teams;
}

export default generateRandomTeam;