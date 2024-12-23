import { User } from '../../models/user';
import jwt from 'jsonwebtoken'
import { Logger } from '../../loggers/logger'
import { JWT_SECRET } from '../../constants/constants';

export const generateJwt = async (user: User):  Promise<string | null> => {
    try{
        const payload = {
            username: user.username,
            email: user.email
        }
        const jwtToken = jwt.sign(payload, JWT_SECRET, {expiresIn: "1200m"})
        return jwtToken
    } catch(error){
        Logger.Error(error.toString())
        return null
    }
}

export const validateJwt = async (token: string): Promise<User | null> => {
    try {
        const decodedToken = jwt.verify(token, JWT_SECRET) as { username: string; email: string };
        return {
            username: decodedToken.username,
            email: decodedToken.email,
        };
    } catch (error) {
        if (error instanceof jwt.TokenExpiredError) {
            Logger.Error(error.toString())
        } else {
            Logger.Error(error.toString());
        }
        return null;
    }
};