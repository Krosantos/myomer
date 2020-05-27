const net = require('net')

// const conn = net.createConnection(3000, "localhost")
conn = {}
conn.fuckyou = function(string){
    this.write(string)
    this.write('\n')
}

const mmm = {
    "action": "matchmake",
    "auth": "None, RIP",
    "userId": "0dde213a-a81b-4b02-a665-41ec6c037112",
    "armyId": "fabc3e24-5dca-47f4-86ba-b3e504de4ccb",

    // "userId": "0c79665d-0ff8-4df9-8d9e-fe44b4b36308",
    // "armyId": "abcd661c-18c4-4c7a-bd9d-e35ac06a48f5",
    
    "gameId": "d2696731-0c23-4dec-ad5b-0e54eee70756"
}

const mmb = {
    "action": "matchmake",
    "auth": "None, RIP",
    // "userId": "0dde213a-a81b-4b02-a665-41ec6c037112",
    // "armyId": "fabc3e24-5dca-47f4-86ba-b3e504de4ccb",

    "userId": "0c79665d-0ff8-4df9-8d9e-fe44b4b36308",
    "armyId": "abcd661c-18c4-4c7a-bd9d-e35ac06a48f5",
    
    "gameId": "d2696731-0c23-4dec-ad5b-0e54eee70756"
}

const forfeit = {
    action:"FORFEIT",
    automatic: true
}
console.log(JSON.stringify(mmm))
console.log()
console.log(JSON.stringify(mmb))
console.log()
console.log(JSON.stringify(forfeit))

// conn.on('connect', () => {
//     console.log('I connected!')
//     conn.fuckyou(JSON.stringify(mmm))
//     // setTimeout(() => {
//     //     console.log("Asking politely to cancel matchmaking")
//     //     conn.fuckyou("cancel")
//     // }, 1000);
// })

// conn.on("close", () => {
//     console.log('Socket Killed')
// })

// conn.on("end", () => {
//     console.log("???")
// })

// conn.on("data", (data) => {
//     const message = data.toString("utf-8")
//     console.log(message)


//     // if (message.includes("reconnected") && mmm.action === 'reconnect') {
//     //     setTimeout(() => {
//     //         console.log('Asking to lose')
//     //         conn.fuckyou(JSON.stringify(forfeit))
//     //     }, 5000);
//     // }


//     if(message.includes("Successfully added to game")){
//         setTimeout(() => {
//             console.log('Asking to lose')
//             conn.fuckyou(JSON.stringify(forfeit))
//         }, 5000);
//     }


//     // setInterval(() => {
//     //     conn.fuckyou("Ping, motherfucker")
//     //     console.log("Pinging, what the fuck")
//     // }, 1500);
// })
