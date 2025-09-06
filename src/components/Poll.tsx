"use client";

import { useEffect, useState } from "react";
import { shuffle } from "@/utils/array";
import { PollItem, PollItemPair } from "@/types/poll";

const items: PollItem[] = [
    {
        title: "JavaScript",
    },
    {
        title: "TypeScript",
    },
    {
        title: "Rust",
    },
    {
        title: "Go",
    },
];

function drawPair(queue: PollItem[]): { pair: PollItemPair; rest: PollItem[] } {
    if (queue.length >= 2) {
        const [left, right, ...rest] = queue;
        return { pair: { left, right }, rest };
    }
    return { pair: { left: null, right: null }, rest: queue.slice() };
}

function getRandomPollItem(items: PollItem[]): PollItem {
    return shuffle(items)[0];
}

export default function Poll() {
    const [pollItems, setPollItems] = useState<PollItem[]>(items);
    const [winners, setWinners] = useState<PollItem[]>([]);
    const [activePair, setActivePair] = useState<PollItemPair | null>(null);

    function pullItem(): PollItem {
        const item = getRandomPollItem(pollItems);
        setPollItems((prev) => prev.filter((i) => i !== item));
        return item;
    }

    function onSelect(item: PollItem | null) {
        if (!item) return;

        setWinners((prevWinners) => {
            const newWinners = [...prevWinners, item];

            setPollItems((prevItems) => {
                const remaining = prevItems.filter((i) => i !== item);

                if (remaining.length >= 2) {
                    setActivePair({
                        left: getRandomPollItem(remaining),
                        right: getRandomPollItem(
                            remaining.filter((x) => x !== remaining[0])
                        ),
                    });
                    return remaining;
                }

                if (prevWinners.length + 1 >= 2) {
                    const restored = newWinners;
                    const newPair = {
                        left: getRandomPollItem(restored),
                        right: getRandomPollItem(
                            restored.filter((x) => x !== restored[0])
                        ),
                    };
                    setActivePair(newPair);
                    return restored;
                }

                return remaining;
            });

            return newWinners;
        });
    }

    useEffect(() => {
        const left = pullItem();
        const right = pullItem();

        setActivePair({ left, right });
    }, []);

    useEffect(() => {
        console.log("winners:", winners);
    }, [winners]);

    useEffect(() => {
        console.log("items:", pollItems);
    }, [pollItems]);

    return (
        <div className="flex w-full flex-1 flex-col lg:flex-row">
            <PollItem
                onSelect={() => onSelect(activePair?.left ?? null)}
                item={activePair?.left ?? null}
            />
            <div className="h-1 w-full lg:h-auto lg:w-1 dark:bg-zinc-500"></div>
            <PollItem
                onSelect={() => onSelect(activePair?.right ?? null)}
                item={activePair?.right ?? null}
            />
        </div>
    );
}
