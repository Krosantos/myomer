const net = require('net')

const conn = net.createConnection(4500, "localhost")

const mmm = {
    "type": "matchmake",
    "auth": "None, RIP",
    "userId": "0dde213a-a81b-4b02-a665-41ec6c037112",
    "armyId": "GOOD_ARMY",
}

conn.on('connect', () => {
    console.log('I connected!')
    conn.write(JSON.stringify(mmm))
})

conn.on("close", () => {
    console.log('I AM CLOSE')
})

conn.on("end", () => {
    console.log("???")
})

conn.on("data", (data) => {
    console.log(data.toString("utf-8"))
})
