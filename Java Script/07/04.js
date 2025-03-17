'use strict';

const fizzBuzz =i => i.map(n => n % 15 === 0 ? 'Fizz Buzz' : n % 3 === 0 ? 'Fizz' : n % 5 === 0 ? 'Buzz' : n);
let range = [];
for (let i = 1; i <= 100; i++) {
    range.push(i);
}
console.log(fizzBuzz(range));

