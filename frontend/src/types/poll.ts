import { User } from "./user";

export type PollItem = {
    title: string;
};

export type PollItemPair = {
    left: PollItem | null;
    right: PollItem | null;
};

export type Poll = {
    title: string;
    id: string;
    author: User;
    likes: number;
    createdAt: string;
    updatedAt: string;
};
