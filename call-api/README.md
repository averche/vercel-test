## Invoke Vercel API

### Try it out

```sh
$ export VERCEL_TOKEN="<token>"
$ go run main.go | jq
```

```json
{
  "type": "encrypted",
  "value": "X4cvr0SI6R91UJYWKReZr1PgQy4jhW/cMyr0V7wIKBg=",
  "target": [
    "preview",
    "development",
    "production"
  ],
  "configurationId": null,
  "id": "6UL0xeGrzoE0kPIJ",
  "key": "MY_NEW_ENV2",
  "createdAt": 1687302007973,
  "updatedAt": 1687302007973,
  "createdBy": "4271JZfdlbJECRd6aTVRfR52",
  "updatedBy": null
}
{
  "type": "encrypted",
  "key": "TEST_ENV",
  "value": "llyYZkmV0Mq7GcQN9p5NuAruWHsWJd2rh3fyaLnyNMY0pk9MU7zRK49Ae6lfLs1K",
  "target": [
    "production",
    "preview",
    "development"
  ],
  "configurationId": null,
  "createdAt": 1687276974186,
  "updatedAt": 1687302008179,
  "createdBy": "4271JZfdlbJECRd6aTVRfR52",
  "updatedBy": "4271JZfdlbJECRd6aTVRfR52",
  "id": "08w3OJ5CpocL9okK"

```

### Environment Variable Operations
1. [Create](https://vercel.com/docs/rest-api/endpoints#create-one-or-more-environment-variables) (`upsert=true` is not working at the moment)
2. [Edit](https://vercel.com/docs/rest-api/endpoints#edit-an-environment-variable)
3. [Remove](https://vercel.com/docs/rest-api/endpoints#remove-an-environment-variable)
