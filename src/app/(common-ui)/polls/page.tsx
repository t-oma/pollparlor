"use client";

import PollCard from "@/components/PollCard";
import { Poll, PollItem } from "@/types/poll";
import { useRef, useState } from "react";

const initPolls: Poll[] = [
    {
        title: "Programming languages",
        uuid: "9e3d57d6-4005-4b14-81b9-6034f6740f43",
        author: {
            id: "2",
            email: "tets@tets.com",
            createdAt: new Date(),
            passwordChangedAt: new Date(),
            verification: { valid: true },
        },
        createdAt: "09/09/2025, 12:20:33" as unknown as Date,
        updatedAt: "09/09/2025, 15:20:33" as unknown as Date,
    },
    {
        title: "Web frameworks",
        uuid: "423ba1b8-e5bd-4193-8215-83329fcf1427",
        author: {
            id: "1",
            email: "test@test.com",
            createdAt: new Date(),
            passwordChangedAt: new Date(),
            verification: { valid: true },
        },
        createdAt: "09/09/2025, 12:20:33" as unknown as Date,
        updatedAt: "09/09/2025, 12:20:33" as unknown as Date,
    },
];

export default function Polls() {
    const polls = useRef<Poll[]>(initPolls);

    return (
        <main className="flex flex-1">
            <section className="flex flex-1 flex-col items-center space-y-4 p-2">
                <hr className="border-border w-full" />
                <h2 className="text-lg font-medium">Available polls</h2>
                <hr className="border-border w-full" />

                <div className="grid grid-cols-2 gap-3">
                    {polls.current.map((poll) => (
                        <PollCard
                            key={poll.uuid}
                            poll={poll}
                        />
                    ))}
                </div>
            </section>
        </main>
    );
}
