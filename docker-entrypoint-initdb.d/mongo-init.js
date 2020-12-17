print(
	"Start #################################################################"
);

db = db.getSiblingDB("passnet_auth");
db.createUser({
	user: "passnet-auth-app",
	pwd: "tungduong98",
	roles: [
		{
			role: "readWrite",
			db: "passnet_auth",
		},
	],
});

db.createCollection("test");
db.test.insert({ key: "value" });

print("END #################################################################");
