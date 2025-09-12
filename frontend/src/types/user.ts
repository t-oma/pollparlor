export type User = {
    id: string;
    email: string;
    createdAt: string;
    passwordChangedAt: string;
    verification: {
        valid: boolean;
    };
};
