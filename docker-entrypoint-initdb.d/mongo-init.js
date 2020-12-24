print(
    "Start #################################################################"
);

db = db.getSiblingDB("passnet_auth");
db.dropUser("passnet-auth-app");
db.createUser({
    user: "passnet-auth-app",
    pwd: "tungduong98",
    roles: [
        {
            role: "readWrite",
            db: "passnet_auth",
        },
        {role: "readAnyDatabase", db: "admin"}
    ],
});

db.createCollection("test");
db.test.insert({key: "value"});

print("END #################################################################");
