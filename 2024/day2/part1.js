const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(lines) {
  let saved_count = 0;

  lines.forEach(line => {
    let numbers = line.split(' ');
    let is_saved = true;
    let is_increasing = true;
    let prev_number = numbers[0];

    for (let i = 1; i < numbers.length; i++) {
      if (prev_number === numbers[i] || Math.abs(prev_number - numbers[i]) > 3) {
        is_saved = false;
        break;
      }

      let this_increased = prev_number - numbers[i] < 0;
      if (i === 1) is_increasing = this_increased;

      if (this_increased !== is_increasing) {
        is_saved = false;
        break;
      }

      prev_number = numbers[i];
    }

    if (is_saved) saved_count++;
  });

  console.log(saved_count);
}
