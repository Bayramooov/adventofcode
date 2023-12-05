const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(games) {
  let sum = 0;

  games.forEach(game => {
    let sets = game.split(': ')[1].split('; ');
    let max = {
      red: 0,
      green: 0,
      blue: 0
    };

    for (let s of sets) {
      let set = {};
      s.split(', ').forEach(x => {
        let cube = x.split(' ');
        if (!set[cube[1]]) set[cube[1]] = 0;
        set[cube[1]] += +cube[0];
      });
      if (set?.red > max.red) max.red = set.red;
      if (set?.green > max.green) max.green = set.green;
      if (set?.blue > max.blue) max.blue = set.blue;
    }
    sum += +max.red * +max.green * +max.blue;
  });
  console.log(sum);
}
