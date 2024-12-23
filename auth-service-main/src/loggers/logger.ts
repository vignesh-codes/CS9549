import morgan from 'morgan';
import winston from 'winston';
import fs from 'fs';

const logsDir = 'logs';
if (!fs.existsSync(logsDir)) {
    fs.mkdirSync(logsDir);
}

const accessLogStream = fs.createWriteStream(`${logsDir}/access.log`, { flags: 'a' });
const infoLogStream = fs.createWriteStream(`${logsDir}/info.log`, { flags: 'a' });
const errorLogStream = fs.createWriteStream(`${logsDir}/error.log`, { flags: 'a' });
const infoLogger = winston.createLogger({
    level: 'info',
    format: winston.format.combine(
        winston.format.timestamp(),
        winston.format.json()
    ),
    transports: [
        new winston.transports.Console(),
        new winston.transports.Stream({ stream: infoLogStream })
    ],
});

const errorLogger = winston.createLogger({
    level: 'error',
    format: winston.format.combine(
        winston.format.timestamp(),
        winston.format.json()
    ),
    transports: [
        new winston.transports.Console(),
        new winston.transports.Stream({ stream: errorLogStream })
    ],
});

export const accessLogger = morgan('combined', { stream: accessLogStream });

export const Logger = {
    Info: (message: string) => infoLogger.info(message),
    Error: (message: string) => errorLogger.error(message),
};
