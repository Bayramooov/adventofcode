const fs = require('fs');
fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) return console.error(err);
  solve(data.split(/\r?\n/));
});

// Solution here
let games = [];
let sum = 0;

function solve(inputLines) {
  for (let line of inputLines) {
    let hand = line
      .split(' ')[0]
      .split('')
      .map(n => {
        switch (n) {
          case 'A':
            return 14;
          case 'K':
            return 13;
          case 'Q':
            return 12;
          case 'J':
            return 11;
          case 'T':
            return 10;
          default:
            return +n;
        }
      });
    let bid = +line.split(' ')[1];
    let hand_type = handType(hand);
    games.push({ card: line.split(' ')[0], hand, hand_type, bid });
  }
  games = games
    .sort((g1, g2) => {
      if (g1.hand_type === g2.hand_type) {
        for (let j in g1.hand) {
          if (g1.hand[j] === g2.hand[j]) {
            continue;
          }
          return g1.hand[j] - g2.hand[j];
        }
      }
      return g1.hand_type - g2.hand_type;
    })
    .map((g, i) => {
      g.rank = i + 1;
      return g;
    });

  for (let g of games) {
    sum += g.bid * g.rank;
  }

  console.log(sum);
}

function handType(hand) {
  // 1 step
  let obj = {};
  for (let c of hand) {
    if (!obj[c]) obj[c] = 1;
    else obj[c]++;
  }
  // 2 step
  let arr = [];
  for (let x in obj) {
    arr.push(obj[x]);
  }
  // 3 step
  arr = arr.sort().reverse();
  // 4 step
  if (arr.length === 1) return 7;
  if (arr.length === 2 && arr[0] === 4) return 6;
  if (arr.length === 2 && arr[0] === 3) return 5;
  if (arr.length === 3 && arr[0] === 3) return 4;
  if (arr.length === 3 && arr[0] === 2) return 3;
  if (arr.length === 4) return 2;
  else return 1;
}
