"use client";

import { Poll } from "@/types/poll";
import PollCard from "./PollCard";
import { useState } from "react";
import Modal from "../Modal";

export default function PollsGrid({ polls }: { polls: Poll[] }) {
    const [open, setOpen] = useState(false);
    const [selected, setSelected] = useState<Poll | null>(null);

    return (
        <div className="grid grid-cols-2 gap-3">
            {polls.map((poll) => (
                <PollCard
                    key={poll.id}
                    poll={poll}
                    onClick={() => {
                        setSelected(poll);
                        setOpen(true);
                    }}
                />
            ))}

            <Modal
                open={open}
                onClose={() => setOpen(false)}
                title={selected?.title}
                inertRootId="page-root"
            >
                {selected && (
                    <div className="space-y-2">
                        <p className="text-sm">By: {selected.author.email}</p>
                        <p className="text-xs">
                            Created:{" "}
                            {new Date(selected.createdAt).toLocaleString()}
                        </p>
                        <p className="text-xs">
                            Updated:{" "}
                            {new Date(selected.updatedAt).toLocaleString()}
                        </p>
                    </div>
                )}
            </Modal>
        </div>
    );
}
