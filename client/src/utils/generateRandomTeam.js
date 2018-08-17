import shuffle from 'shuffle-array';

const generateRandomTeam = (players, playersPerTeam) => {

  const shuffledPlayers = shuffle(players);

  const numberOfTeams = Math.ceil(shuffledPlayers.length / playersPerTeam);

  const averagePlayerRatings = shuffledPlayers.reduce((acc, p) => acc + p.ratings, 0) / shuffledPlayers.length;

  console.log('average player ratings: ', averagePlayerRatings);//

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
    console.log('--- adding ' + shuffledPlayers[i].name + ' (' + shuffledPlayers[i].ratings + ') ---');
    let teamToPutPlayerIndex = 0;

    console.log('+----- trying to put on team' + (teamToPutPlayerIndex + 1) + '. Players in this team: ' + teams[teamToPutPlayerIndex].length + '  (' + (teams[teamToPutPlayerIndex].length >= playersPerTeam ? 'team already full' : 'checked') + ')');
    console.log('+----- check if player rating can fit in the team   : ' + (teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings) + '/' + maxTeamRatings + '(' + (teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings > maxTeamRatings ? 'too full': 'checked') + ')');

    while (
      teams[teamToPutPlayerIndex] && 
      (teams[teamToPutPlayerIndex].length >= playersPerTeam ||
      teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings > (maxTeamRatings + averagePlayerRatings / 2))
    ) {
      teamToPutPlayerIndex++;
      if (teams[teamToPutPlayerIndex]) {
        console.log('+----- trying to put on team' + (teamToPutPlayerIndex + 1) + '. Players in this team: ' + teams[teamToPutPlayerIndex].length + '  (' + (teams[teamToPutPlayerIndex].length >= playersPerTeam ? 'team already full' : 'checked') + ')');
        console.log('+----- check if player rating can fit in the team   : ' + (teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings) + '/' + maxTeamRatings + '(' + (teams[teamToPutPlayerIndex].reduce((acc, p) => acc + p.ratings, 0) + shuffledPlayers[i].ratings > maxTeamRatings ? 'too full': 'checked') + ')');
      }
    }

    if(!teams[teamToPutPlayerIndex]) {
      console.log('!!!!!!!!!!! There is no team available to put this player on. Putting this player in the team with the lowest total stars that still needs players')
      const totalRatingsOfTeams = teams.map(team => team.reduce((acc, player) => acc + player.ratings, 0));
      console.log('---------- team ratings: ' + totalRatingsOfTeams.join(', '));
      console.log('---------- team need more player: ' + teams.map(team => team.length < playersPerTeam).join(', '));

      const sortedTeamsListByRatings = teams.slice().sort((t1, t2) => t1.reduce((acc, player) => acc + player.ratings, 0) >= t2.reduce((acc, player) => acc + player.ratings, 0));

      console.log(sortedTeamsListByRatings);

      const teamToAddPLayerTo = sortedTeamsListByRatings.find(team => team.length < playersPerTeam);

      teamToAddPLayerTo.push(shuffledPlayers[i]);

      // const teamWithLowestScoreIndex = totalRatingsOfTeams.indexOf(lowestScore);
      // console.log('---------- putting in team ' + (teamWithLowestScoreIndex + 1));
      // teams[teamWithLowestScoreIndex].push(shuffledPlayers[i]);
    } else {
      teams[teamToPutPlayerIndex].push(shuffledPlayers[i]);
      console.log('--vvvvvvvvvvv-- inserting ', shuffledPlayers[i].name, ' into team ', (teamToPutPlayerIndex + 1));``
    }
  }

  return teams;
}

export default generateRandomTeam;