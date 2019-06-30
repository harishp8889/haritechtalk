// Usage
import { Singleton } from "./Singleton";

const singleInstance = Singleton.getInstance();
singleInstance.data = 10;
console.log(singleInstance.data);

//creating second instance should return same 
const singleInstance2 = Singleton.getInstance();
console.log(singleInstance2.data); //will return 10