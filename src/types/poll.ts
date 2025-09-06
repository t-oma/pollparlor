export type PollItem = {
    title: string;
};

export type PollItemPair = {
    left: PollItem | null;
    right: PollItem | null;
};
