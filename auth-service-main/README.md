# Typescript-Node-Auth-Boilerplate
This is my personal boilerplate for ts api server. This repo has two endpoints v1 for cookie based auth and v2 for jwt based auth

```
src
├── constants
│   └── constants.ts
├── controllers
│   ├── users.ts
│   ├── v1
│   │   └── authentication.ts
│   └── v2
│       └── authentication.ts
├── db
│   ├── initdb.ts
│   └── users.ts
├── helpers
│   └── index.ts
├── index.ts
├── loggers
│   └── logger.ts
├── middlewares
│   ├── v1
│   │   └── auth.ts
│   └── v2
│       ├── auth.ts
│       └── jwt.ts
├── models
│   └── user.ts
└── router
    ├── index.ts
    ├── v1
    │   ├── authentication.ts
    │   └── users.ts
    └── v2
        ├── authentication.ts
        └── users.ts
```

## To run the application
- Update the mongo url in src/constants/constants.ts

```
npm start
```