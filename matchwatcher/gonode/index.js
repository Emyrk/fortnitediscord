// require the package
const Fortnite = require('fortnite-api');
const args = process.argv;

token = process.env.DISCORD_TOKEN
launcherkey = process.env.LAUNCHER_KEY
clientkey = process.env.CLIENT_KEY
user = process.env.EPIC_USERNAME
pass = process.env.EPIC_PASSWORD


console.log(args[1])

let fortniteAPI = new Fortnite(EPIC_USERNAME, EPIC_PASSWORD, n1, h1]);

fortniteAPI.login()
    .then(() => {
        fortniteAPI.getStatsBR(args[1], "pc")
        // fortniteAPI.checkPlayer("Emyrks", "pc")
            .then((stats) => {
                console.log(JSON.stringify(stats));
                process.exit()
            })
            .catch((err) => {

                console.log(err);
            });
    });

