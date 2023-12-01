const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
function solve(inputLines) {
  let sum = 0;
  inputLines.forEach(line => {
    let validLine = line
      .replaceAll('one', 'o1e')
      .replaceAll('two', 't2o')
      .replaceAll('three', 't3e')
      .replaceAll('four', 'f4r')
      .replaceAll('five', 'f5e')
      .replaceAll('six', 's6x')
      .replaceAll('seven', 's7n')
      .replaceAll('eight', 'e8t')
      .replaceAll('nine', 'n9e');

    let nums = validLine.replace(/[^0-9]/g, '');
    sum += parseInt(nums[0] + nums[nums.length - 1]);
  });

  console.log(sum);
}
