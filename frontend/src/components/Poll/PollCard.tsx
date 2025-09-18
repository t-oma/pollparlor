"use client";

import { Poll } from "@/types/poll";
import Button from "../Button";
import { formatRelativeDate } from "@/utils/date";
import Divider from "../Divider";
import { memo } from "react";

function PollCard({
    poll,
    onClick,
}: Readonly<{
    poll: Poll;
    onClick: () => void;
}>) {
    const isUpdated =
        new Date(poll.updatedAt).getTime() > new Date(poll.createdAt).getTime();

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

            <Divider className="my-2" />
            <p className="text-sm">By: {poll.author.email}</p>
            <Divider className="my-2" />

            <p className="text-sm">
                {isUpdated
                    ? `upd: ${formatRelativeDate(poll.updatedAt)}`
                    : formatRelativeDate(poll.createdAt)}
            </p>
        </Button>
    );
}

export default memo(PollCard);
