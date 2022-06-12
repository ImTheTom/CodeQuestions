// Look at https://howtodoinjava.com/typescript/maps/

let myMap = new Map<string, number>();

myMap.set("Tim", 99);
myMap.set("Tam", 23);
myMap.set("Kangaroo", 22);

console.log(myMap.get("Tim"));

console.log(myMap.has("Tam"));
console.log(myMap.has("Tams"));

console.log(myMap.size);

console.log(myMap.has("Kangaroo"));
console.log(myMap.delete("Kangaroo"));
console.log(myMap.has("Kangaroo"));

for (let key of myMap.keys()) {
    console.log(key);
}

for (let value of myMap.values()) {
    console.log(value);
}

for (let entry of myMap.entries()) {
    console.log(entry[0], entry[1]);
}

for (let [key, value] of myMap) {
    console.log(key, value);
}

myMap.clear();