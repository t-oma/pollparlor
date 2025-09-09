"use client";

import { Poll } from "@/types/poll";
import Link from "next/link";

export default function PollCard({ poll }: { poll: Poll }) {
    return (
        <Link
            href={`/polls/${poll.uuid}`}
            className="border-border flex flex-col gap-2 rounded-md border transition-all hover:scale-105 hover:rotate-2"
        >
            <div className="flex flex-col px-2 pt-2">
                <h3 className="text-center text-sm font-semibold">
                    {poll.title}
                </h3>
                <hr className="border-border my-2" />
                <p className="text-sm">By: {poll.author.email}</p>
            </div>
            <div className="border-border flex flex-1 flex-col items-center justify-center border-t p-1">
                {new Date(poll.createdAt) >= new Date(poll.updatedAt) ? (
                    <p className="text-sm">{poll.createdAt.toLocaleString()}</p>
                ) : (
                    <p className="text-xs">
                        Update: {poll.updatedAt.toLocaleString()}
                    </p>
                )}
            </div>
        </Link>
    );
}
