const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/).map(line => line.split('.')));
});

// Solution here
function solve(inputMatrix) {
  let sum = 0;

  console.log(inputMatrix);

  for (let i in inputMatrix) {
    for (let j in inputMatrix[i]) {
      
    }
  }
  
  console.log(sum);
}
