const net = require('net')

const conn = net.createConnection(4500, "localhost")

const mmm = {
    "action": "matchmake",
    "auth": "None, RIP",
    "userId": "0dde213a-a81b-4b02-a665-41ec6c037112",
    "armyId": "fabc3e24-5dca-47f4-86ba-b3e504de4ccb",
    // "userId": "0c79665d-0ff8-4df9-8d9e-fe44b4b36308",
    // "armyId": "abcd661c-18c4-4c7a-bd9d-e35ac06a48f5",
    "gameId":"640129d2-4277-47ce-aed1-fc50ede03c9e",
}

conn.on('connect', () => {
    console.log('I connected!')
    conn.write(JSON.stringify(mmm))
    // setTimeout(() => {
    //     console.log("Asking politely to cancel matchmaking")
    //     conn.write("cancel")
    // }, 1000);
})

conn.on("close", () => {
    console.log('Socket Killed')
})

conn.on("end", () => {
    console.log("???")
})

conn.on("data", (data) => {
    console.log(data.toString("utf-8"))
})
