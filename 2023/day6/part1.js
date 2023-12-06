const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(inputLines) {
  let time = inputLines[0].split(/\s+/).slice(1).map(n => +n);
  let dist = inputLines[1].split(/\s+/).slice(1).map(n => +n);
  let mult = 1;

  for (let i in time) {
    let t = time[i];
    let s = dist[i];
    let ways = 0;

    for (let v = 1; v <= t - 1; v++) {
      if (v * (t - v) <= s) continue;
      ways++;
    }

    mult *= ways || 1;
  }

  console.log(mult);
}
