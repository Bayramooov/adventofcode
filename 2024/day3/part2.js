const { readFile, writeFileSync } = require('fs');
let outputText = '';
let cout = text => (outputText += text);
readFile('INPUT.TXT', 'utf8', (err, data) => {
  if (err) return console.error(err);
  const input = data;
  // .trim()
  // .split(/\r?\n/)
  // .map(line =>
  //   line
  //     .trim()
  //     .split(/\s+/)
  //     .map(x => (isNaN(+x) ? String(x) : +x))
  // );
  main(input);
  writeFileSync('OUTPUT.TXT', String(outputText));
});

function main(input2) {
  const input = input2
    .split('do()')
    .map(x => x.split("don't()")[0].trim())
    .join('');

  cout([...input.matchAll(/mul\((\d{1,3}),(\d{1,3})\)/gm)].reduce((sum, matched) => (sum += +matched[1] * +matched[2]), 0));
}
