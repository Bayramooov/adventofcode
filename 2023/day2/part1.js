const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(games) {
  let sum = 0;

  games.forEach((game, index) => {
    let game_number = index + 1;
    let is_valid = true;
    let sets = game.split(': ')[1].split('; ');

    for (let s of sets) {
      let set = {};
      s.split(', ').forEach(x => {
        let cube = x.split(' ');
        if (!set[cube[1]]) set[cube[1]] = 0;
        set[cube[1]] += +cube[0];
      });
      if (set?.red > 12 || set?.green > 13 || set?.blue > 14) {
        is_valid = false;
        break;
      };
    }
    if (is_valid) sum += game_number;
  });
  console.log(sum);
}
