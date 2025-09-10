export type User = {
    id: string;
    email: string;
    createdAt: Date;
    passwordChangedAt: Date;
    verification: {
        valid: boolean;
    };
};
