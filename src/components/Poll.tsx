"use client";

import { useEffect, useState } from "react";
import { usePollStore } from "@/providers/poll-store-provider";
import { PollItem } from "@/types/poll";
import clsx from "clsx";
import Button from "./Button";

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
                            "flex transition-all duration-500",
                            selected === "left" && "flex-1",
                            selected === "right" && "flex-[0] overflow-hidden",
                            !selected && "flex-1/2"
                        )}
                    >
                        <Button
                            disabled={animating}
                            onClick={() => onSelect("left")}
                            className="flex-1"
                        >
                            <span className="text-3xl">
                                {active.left ? active.left.title : "-"}
                            </span>
                        </Button>
                    </div>
                    <div className="bg-border h-1 w-full rounded-md lg:h-auto lg:w-1"></div>
                    <div
                        className={clsx(
                            "flex transition-all duration-500",
                            selected === "right" && "flex-1",
                            selected === "left" && "flex-[0] overflow-hidden",
                            !selected && "flex-1/2"
                        )}
                    >
                        <Button
                            disabled={animating}
                            onClick={() => onSelect("right")}
                            className="flex-1"
                        >
                            <span className="text-3xl">
                                {active.right ? active.right.title : "-"}
                            </span>
                        </Button>
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
