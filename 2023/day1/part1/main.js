const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(inputLines) {
  let sum = 0;
  inputLines.forEach(line => {
    let nums = line.replace(/[^0-9]/g, '');
    sum += parseInt(nums[0] + nums[nums.length - 1]);
  });
  console.log(sum);
}
