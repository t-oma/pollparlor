"use client";

import { useEffect, useState } from "react";
import { usePollStore } from "@/providers/poll-store-provider";
import { PollItem } from "@/types/poll";
import clsx from "clsx";

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

export default function Poll() {
    const { queue, winners, active, champion, start, pick } = usePollStore(
        (store) => store
    );
    const [animating, setAnimating] = useState(false);
    const [selected, setSelected] = useState<"left" | "right" | null>(null);

    useEffect(() => {
        start(items);
    }, [start]);

    const onSelect = (side: "left" | "right") => {
        if (animating) return;
        const item = side === "left" ? active.left : active.right;
        if (!item) return;

        setSelected(side);
        setAnimating(true);

        setTimeout(() => {
            pick(item);
            setAnimating(false);
            setSelected(null);
        }, 650);
    };

    return (
        <div className="flex w-full flex-1 flex-col lg:flex-row">
            {champion === null ? (
                <>
                    <div
                        className={clsx(
                            "flex border transition-all duration-500 dark:border-zinc-500",
                            selected === "left" && "flex-1",
                            selected === "right" && "flex-[0] overflow-hidden",
                            !selected && "flex-1/2"
                        )}
                    >
                        <button
                            type="button"
                            className="flex-1 cursor-pointer p-4 transition-colors disabled:cursor-default dark:bg-zinc-900 dark:hover:bg-zinc-800"
                            disabled={animating}
                            onClick={() => onSelect("left")}
                        >
                            <span className="text-3xl">
                                {active.left ? active.left.title : "-"}
                            </span>
                        </button>
                    </div>
                    <div className="h-1 w-full lg:h-auto lg:w-1 dark:bg-zinc-500"></div>
                    <div
                        className={clsx(
                            "flex border transition-all duration-500 dark:border-zinc-500",
                            selected === "right" && "flex-1",
                            selected === "left" && "flex-[0] overflow-hidden",
                            !selected && "flex-1/2"
                        )}
                    >
                        <button
                            type="button"
                            className="flex-1 cursor-pointer p-4 transition-colors disabled:cursor-default dark:bg-zinc-900 dark:hover:bg-zinc-800"
                            disabled={animating}
                            onClick={() => onSelect("right")}
                        >
                            <span className="text-3xl">
                                {active.right ? active.right.title : "-"}
                            </span>
                        </button>
                    </div>
                </>
            ) : (
                <div className="flex flex-1 items-center justify-center">
                    <p className="text-center text-4xl font-bold">
                        champion: {champion.title}
                    </p>
                </div>
            )}
        </div>
    );
}
