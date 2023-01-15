import { Level } from 'level';

const db = new Level("./mydata", {
    valueEncoding: "json"
});

const r = await db.put("name", {
    user: {
        name: "Mihai"
    }
});
console.log(r);

console.log("stored name = ", await db.get('name'));

