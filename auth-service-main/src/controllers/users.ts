import express from 'express'

import { UserDatabase } from "../db/users";
import { Logger } from '../loggers/logger'
export class UserController {
    private static userDB = UserDatabase.getInstance();

    public static getAllUsers = async (req: express.Request, res: express.Response) => {
        try {
            const users = await UserController.userDB.getUsers();
            res.send(users)
            res.end()
        } catch (error) {
            Logger.Error(error.toString())
            return res.sendStatus(400)
        }
    }

    public static updateUser = async (req: express.Request, res: express.Response) => {
        try {
            const { username } = req.body
            const { id } = req.params
            if (!username) {
                return res.sendStatus(400)
            }
            const userToUpdate = await UserController.userDB.getUserById(id);
            userToUpdate.username = username;
            await userToUpdate.save();
            return res.status(200).json(userToUpdate).end()
        } catch (error) {
            Logger.Error(error.toString())
            res.sendStatus(400)
        }
    }

    public static deleteUser = async (req: express.Request, res: express.Response) => {
        try {
            const { id } = req.params;
            const deletedUser = await UserController.userDB.deleteUserById(id);
            return res.json(deletedUser)
        } catch (error) {
            Logger.Error(error.toString())
            res.sendStatus(400)
        }
    }
}