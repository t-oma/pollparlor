"use client";

import { Poll } from "@/types/poll";
import Button from "../Button";

export default function PollCard({
    poll,
    onClick,
}: Readonly<{
    poll: Poll;
    onClick: () => void;
}>) {
    return (
        <Button
            onClick={onClick}
            className="min-w-[150px] flex-col gap-0 p-4 transition-all hover:scale-105 hover:rotate-2"
        >
            <div className="flex flex-1 flex-col items-center justify-center">
                <h3 className="text-center text-sm font-semibold">
                    {poll.title}
                </h3>
            </div>
            <hr className="border-border my-2 w-full" />
            <p className="text-sm">By: {poll.author.email}</p>
            <hr className="border-border my-2 w-full" />
            {new Date(poll.createdAt) >= new Date(poll.updatedAt) ? (
                <p className="text-sm">{poll.createdAt.toString()}</p>
            ) : (
                <p className="text-sm">u: {poll.updatedAt.toString()}</p>
            )}
        </Button>
    );
}
