import mongoose from 'mongoose';
import { MONGOURL } from "../constants/constants"
import {Logger} from "../loggers/logger"

export class MongoDB {
    public static async initDb() {
        try {
            Logger.Info("Trying to connect to MongoDB")
            await mongoose.connect(MONGOURL);
            Logger.Info("Connected to MongoDB");
        } catch (error) {
            Logger.Error("Error connecting to MongoDB:" + error.toString());
        }
        this.connectToDb()
    }

    private static connectToDb() {
        mongoose.connection.on('error', (error: Error) => {
            Logger.Error("MongoDB connection error:"+ error.toString());
        });
    }
}