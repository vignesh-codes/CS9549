import express from 'express';

import { UserController } from '../../controllers/users';
import { V2Middleware } from '../../middlewares/v2/auth';

export default (router: express.Router) => {
    router.get('/users', V2Middleware.isAuthenticated, UserController.getAllUsers);
    router.delete('/users/:id', V2Middleware.isAuthenticated, V2Middleware.isOwner, UserController.deleteUser)
    router.patch('/users/:id', V2Middleware.isAuthenticated, V2Middleware.isOwner, UserController.updateUser)
}
