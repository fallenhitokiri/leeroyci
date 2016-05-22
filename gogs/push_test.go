package gogs

var testPush = `
{
  "secret": "asdf",
  "ref": "refs/heads/foo",
  "before": "0000000000000000000000000000000000000000",
  "after": "c9f89a3e1ed76867999b708c75017ab293dc7749",
  "compare_url": "http://localhost:3000/",
  "commits": [
    {
      "id": "c9f89a3e1ed76867999b708c75017ab293dc7749",
      "message": "var\n",
      "url": "http://localhost:3000/timo/test/commit/c9f89a3e1ed76867999b708c75017ab293dc7749",
      "author": {
        "name": "Timo Zimmermann",
        "email": "timo@screamingatmyscreen.com",
        "username": ""
      }
    },
    {
      "id": "2a73fa48186e965f8d6c86da791787e4c4f78782",
      "message": "foo\n",
      "url": "http://localhost:3000/timo/test/commit/2a73fa48186e965f8d6c86da791787e4c4f78782",
      "author": {
        "name": "Timo Zimmermann",
        "email": "timo@screamingatmyscreen.com",
        "username": ""
      }
    },
    {
      "id": "21c9636a8258549b95c49cd26d08e9fba25ba2f9",
      "message": "foo\n",
      "url": "http://localhost:3000/timo/test/commit/21c9636a8258549b95c49cd26d08e9fba25ba2f9",
      "author": {
        "name": "Timo Zimmermann",
        "email": "timo@screamingatmyscreen.com",
        "username": ""
      }
    },
    {
      "id": "269774b4c342c71f68b1055061bfbc93a2b2304b",
      "message": "initial\n",
      "url": "http://localhost:3000/timo/test/commit/269774b4c342c71f68b1055061bfbc93a2b2304b",
      "author": {
        "name": "Timo Zimmermann",
        "email": "timo@screamingatmyscreen.com",
        "username": ""
      }
    }
  ],
  "repository": {
    "id": 1,
    "name": "test",
    "url": "http://localhost:3000/timo/test",
    "ssh_url": "timo@localhost:timo/test.git",
    "clone_url": "http://localhost:3000/timo/test.git",
    "description": "",
    "website": "",
    "watchers": 1,
    "owner": {
      "name": "timo",
      "email": "admin@foo.tld",
      "username": "timo"
    },
    "private": false,
    "default_branch": "master"
  },
  "pusher": {
    "name": "timo",
    "email": "admin@foo.tld",
    "username": "timo"
  },
  "sender": {
    "login": "timo",
    "id": 1,
    "avatar_url": "https://secure.gravatar.com/avatar/983665f6fb4eae7c9e546da0c94eb71c"
  }
}
`
