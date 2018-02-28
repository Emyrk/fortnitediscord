// require the package
const Fortnite = require('fortnite-api');
const args = process.argv;

token = process.env.DISCORD_TOKEN
launcherkey = process.env.LAUNCHER_KEY
clientkey = process.env.CLIENT_KEY
user = process.env.EPIC_USERNAME
pass = process.env.EPIC_PASSWORD



let fortniteAPI = new Fortnite([user, pass, launcherkey, clientkey]);

fortniteAPI.login()
    .then(() => {
        name = args[2].replace(/['"]+/g, '')
        fortniteAPI.getStatsBR(name, "pc")
        // fortniteAPI.checkPlayer("Emyrks", "pc")
            .then((stats) => {
                console.log(JSON.stringify(stats));
                process.exit()
            })
            .catch((err) => {
                console.log(args[2], err);
                process.exit()
            });
    });

