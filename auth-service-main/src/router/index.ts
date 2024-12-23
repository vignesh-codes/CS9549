import express from 'express';
import v1authentication from './v1/authentication';
import v2authentication from './v2/authentication';

import v1users from './v1/users';
import v2users from './v2/users';
const v1router = express.Router();

const v2router = express.Router();


export default (): { v1router: express.Router, v2router: express.Router } => {
    v1authentication(v1router);
    v1users(v1router);

    v2authentication(v2router);
    v2users(v2router);
    return { v1router, v2router };
}

