const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(inputLines) {
  let t = +inputLines[0].split(/\s+/).slice(1).join('');
  let s = +inputLines[1].split(/\s+/).slice(1).join('');
  let ways = 0;

  for (let v = 1; v <= t - 1; v++) {
    if (v * (t - v) <= s) continue;
    ways++;
  }

  console.log(ways);
}
