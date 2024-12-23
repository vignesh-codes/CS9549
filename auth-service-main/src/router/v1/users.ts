import express from 'express';

import { UserController } from '../../controllers/users';
import { V1Middleware } from '../../middlewares/v1/auth';

export default (router: express.Router) => {
    router.get('/users', V1Middleware.isAuthenticated, UserController.getAllUsers);
    router.delete('/users/:id', V1Middleware.isAuthenticated, V1Middleware.isOwner, UserController.deleteUser)
    router.patch('/users/:id', V1Middleware.isAuthenticated, V1Middleware.isOwner, UserController.updateUser)
}
