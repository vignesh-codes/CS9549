import express from 'express';
import { get, merge } from 'lodash';
import { Logger } from '../../loggers/logger'
import { UserDatabase } from '../../db/users';

export class V1Middleware {
    private static userDB = UserDatabase.getInstance();

    public static isAuthenticated = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
        try {
            const sessionToken = req.cookies['VWS-AUTH'];

            if (!sessionToken) {
                return res.sendStatus(401);
            }

            const existingUser = await V1Middleware.userDB.getUserBySessionToken(sessionToken);
            if (!existingUser) {
                return res.sendStatus(403);
            }
            merge(req, { identity: existingUser });
            return next();

        } catch (error) {
            Logger.Error(error.toString());
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
            if (currentUserId.toString() !== id) {
                return res.sendStatus(400);
            }
            next();
        } catch (error) {
            Logger.Error(error.toString());
            res.sendStatus(400);
        }
    }
}
