import express from 'express';

import { UserDatabase } from '../../db/users';
import { authentication, random } from '../../helpers'
import { Logger } from '../../loggers/logger'
export class V1AuthController {
    private static userDB = UserDatabase.getInstance();

    public static v1login = async (req: express.Request, res: express.Response) => {
        try {
            const { email, password} = req.body;
            if (!email || !password ) {
                return res.sendStatus(400)
            }
            const user = await V1AuthController.userDB.getUserByEmail(email).select("+authentication.salt +authentication.password");

            if (!user) {
                return res.sendStatus(400)
            }
            
            const expectedHash = authentication(user.authentication.salt, password);
            if (user.authentication.password != expectedHash) {
                return res.sendStatus(401)
            }

            const salt = random();
            user.authentication.sessionToken = authentication(salt, user._id.toString());
            await user.save()

            res.cookie("VWS-AUTH", user.authentication.sessionToken, {
                domain: 'localhost',
                path: "/"
            });
            return res.status(200).json(user).end()

        } catch (error) {
            Logger.Error(error.toString())
            return res.sendStatus(400)
        }
    }

    public static v1register = async (req: express.Request, res: express.Response) => {
        try {
            const { email, password, username } = req.body;

            if (!email || !password || !username) {
                return res.sendStatus(400);
            };

            const existingUser = await V1AuthController.userDB.getUserByEmail(email);
            if (existingUser) {
                return res.sendStatus(400);
            };

            const salt = random();
            const user = await V1AuthController.userDB.createUser({
                email,
                username,
                authentication: {
                    salt,
                    password: authentication(salt, password)
                },
            });

            return res.status(200).json(user).end()

        } catch (error) {
            Logger.Error(error.toString());
            return res.sendStatus(400);
        }
    }

    public static v1logSuccessMsg = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
        Logger.Info("successfully registered");
        next(); 
    }
}
