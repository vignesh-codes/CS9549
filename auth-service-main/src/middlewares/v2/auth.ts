import express from 'express';
import { get, merge } from 'lodash';
import { Logger } from '../../loggers/logger'
import { UserDatabase } from '../../db/users';
import { validateJwt } from './jwt';

export class V2Middleware {
    private static userDB = UserDatabase.getInstance();

    public static isAuthenticated = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
        try {
            const jwt = req.headers["jwt"] as string;
            if (!jwt) {
                return res.sendStatus(401);
            }

            const decodedToken = await validateJwt(jwt)
            if (!decodedToken) {
                return res.sendStatus(400);
            }
            const existingUser = await V2Middleware.userDB.getUserByEmail(decodedToken.email);
            if (!existingUser) {
                return res.sendStatus(403);
            }
            merge(req, { identity: existingUser });
            return next();

        } catch (error) {
            Logger.Error(error.toString())
            return res.sendStatus(400);
        }
    }

    public static isOwner = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
        try {
            const { id } = req.params;
            const currentUserId = get(req, 'identity._id') as string;
            if (!currentUserId) {
                return res.sendStatus(400);
            }
            if (currentUserId.toString() != id) {
                return res.sendStatus(400)
            }
            next();
        } catch (error) {
            Logger.Error(error.toString())
            res.sendStatus(400)
        }
    }
}
