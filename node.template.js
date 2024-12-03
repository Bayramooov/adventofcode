const { readFile, writeFileSync } = require('fs');
let outputText = '';
let cout = text => (outputText += text);
readFile('INPUT.TXT', 'utf8', (err, data) => {
  if (err) return console.error(err);
  const input = data
    .trim()
    .split(/\r?\n/)
    .map(line =>
      line
        .trim()
        .split(/\s+/)
        .map(x => (isNaN(+x) ? String(x) : +x))
    );
  main(input);
  writeFileSync('OUTPUT.TXT', String(outputText));
});

function main(input) {
  let result = '';
  cout(result);
}
