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

let hash = {"hello": "world"};
let hello = "hello";
puts(hash[hello]);
