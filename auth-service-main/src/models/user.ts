
export interface User {
    username: string;
    email: string;
    authentication?: {
        password: string;
        salt: string;
        sessionToken?: string;
        jwtToken?: string;
    };
}
