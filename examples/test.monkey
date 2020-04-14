let hello = "Hello World";
puts(hello);

let addFunction = fn(a) {
  return fn(b) {
    return a + b;
  };
};
let addTwo = addFunction(2);
puts(addTwo(3));

let array = [1, 1 * 2, 1 + 2];
let array = push(array, 2 * 2);
puts(last(array));

let hash = {"hello": "world", 1: "one"};
let hello = "hello";
puts(hash[hello]);
puts(hash[1]);

let unless = macro(cond, cons) {
  quote(if (!unquote(cond)) { unquote(cons); });
};
unless(3 > 5, puts("3 is less than 5"));
