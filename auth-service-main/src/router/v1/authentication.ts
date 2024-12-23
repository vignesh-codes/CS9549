import express from 'express';

import { V1AuthController } from '../../controllers/v1/authentication';

export default (router: express.Router) => {
    router.post('/auth/register', V1AuthController.v1logSuccessMsg, V1AuthController.v1register);
    router.post('/auth/login', V1AuthController.v1login)
};
