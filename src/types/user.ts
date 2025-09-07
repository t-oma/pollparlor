export type User = {
    id: string;
    email: string;
    password: string;
    createdAt: Date;
    passwordChangedAt: Date;
    verification: {
        valid: boolean;
    };
};
