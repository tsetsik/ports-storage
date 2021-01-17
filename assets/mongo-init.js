db.auth('root', 'secure')

db = db.getSiblingDB('ports')

db.createUser({
  user: 'test-user',
  pwd: 'test-password',
  roles: [
    {
      role: 'readWrite',
      db: 'ports',
    },
  ],
});

db.createCollection('ports');