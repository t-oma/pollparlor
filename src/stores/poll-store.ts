import { PollItem, PollItemPair } from "@/types/poll";
import { shuffle } from "@/utils/array";
import { createStore } from "zustand";

export type PollState = {
    queue: PollItem[];
    winners: PollItem[];
    active: PollItemPair;
    champion: PollItem | null;
};

export type PollActions = {
    start: (items: PollItem[]) => void;
    pick: (item: PollItem) => void;
};

export type PollStore = PollState & PollActions;

const defaultInitState: PollState = {
    queue: [],
    winners: [],
    active: { left: null, right: null },
    champion: null,
};

export const createPollStore = (initState: PollState = defaultInitState) => {
    return createStore<PollStore>()((set, get) => ({
        ...initState,
        start: (items) => {
            const q = shuffle(items);
            const [left, right, ...rest] = q;
            set({
                queue: rest,
                winners: [],
                active: { left, right },
                champion: null,
            });
        },
        pick: (item) => {
            const { queue, winners } = get();
            const newWinners = [...winners, item];

            if (queue.length >= 2) {
                const [left, right, ...rest] = queue;
                set({
                    queue: rest,
                    winners: newWinners,
                    active: { left, right },
                });
            } else if (queue.length === 1) {
                const auto = queue[0];
                const allWinners = [...newWinners, auto];

                if (allWinners.length >= 2) {
                    const q2 = shuffle(allWinners);
                    const [l2, r2, ...rest2] = q2;
                    set({
                        queue: rest2,
                        winners: [],
                        active: { left: l2, right: r2 },
                    });
                } else {
                    set({
                        queue: [],
                        winners: [],
                        active: { left: null, right: null },
                        champion: allWinners[0],
                    });
                }
            } else {
                if (newWinners.length >= 2) {
                    const q2 = shuffle(newWinners);
                    const [l2, r2, ...rest2] = q2;
                    set({
                        queue: rest2,
                        winners: [],
                        active: { left: l2, right: r2 },
                    });
                } else {
                    set({
                        queue: [],
                        winners: [],
                        active: { left: null, right: null },
                        champion: newWinners[0],
                    });
                }
            }
        },
    }));
};
