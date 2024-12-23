import express from 'express';

import { V2AuthController } from '../../controllers/v2/authentication';

export default (router: express.Router) => {
    router.post('/auth/register', V2AuthController.v2register);
    router.post('/auth/login', V2AuthController.v2login)
};